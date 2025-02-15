package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/shopspring/decimal"
	httpSwagger "github.com/swaggo/http-swagger"
	"golang.org/x/sync/errgroup"
	"gopkg.in/telebot.v3"

	"pkg/database/postgresql"
	"pkg/errors"
	"pkg/http/router"
	"pkg/http/server"
	"pkg/log"
	"pkg/log/model"
	"pkg/migrator"
	"pkg/panicRecover"
	"pkg/stackTrace"
	"pkg/trace"
	"swearBot/internal/config"
	_ "swearBot/internal/docs"
	checkerService "swearBot/internal/services/checker/service"
	saverRepository "swearBot/internal/services/saver/repository"
	saverService "swearBot/internal/services/saver/service"
	shedulerService "swearBot/internal/services/scheduler"
	statisticEndpoint "swearBot/internal/services/statistic/endpoint"
	statisticRepository "swearBot/internal/services/statistic/repository"
	statisticService "swearBot/internal/services/statistic/service"
	swearEndpoint "swearBot/internal/services/swear/endpoint"
	swearRepository "swearBot/internal/services/swear/repository"
	swearService "swearBot/internal/services/swear/service"
	"swearBot/migrations"
)

// @title Swear word counting bot API
// @version @{version} (build @{build})
// @description API Documentation for Coin
// @contact.name Ilia Ivanov
// @contact.email bonavii@icloud.com
// @contact.url

// @securityDefinitions.apikey AuthJWT
// @in header
// @name Authorization
// @description JWT-токен авторизации

//go:generate go install github.com/swaggo/swag/cmd/swag@v1.8.2
//go:generate go mod download
//go:generate swag init -o docs --parseInternal --parseDependency

const version = "@{version}"
const build = "@{build}"

func main() {
	if err := run(); err != nil {
		log.Fatal(context.Background(), err)
	}
}

func run() error {

	// Основной контекст приложения
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	// Перехватываем возможную панику
	defer panicRecover.PanicRecover(func(err error) {
		log.Fatal(ctx, err)
	})

	// Парсим флаги
	logFormat := flag.String("log-format", string(log.JSONFormat), "text - Human readable string\njson - JSON format")
	envMode := flag.String("env-mode", "local", "Environment mode for log label: test, prod")
	flag.Parse()

	var logHandlers []log.Handler
	switch *logFormat {
	case "text":
		logHandlers = append(logHandlers, log.NewConsoleHandler(os.Stdout, log.LevelDebug))
	case "json":
		logHandlers = append(logHandlers, log.NewJSONHandler(os.Stdout, log.LevelDebug))
	}

	// Получаем имя хоста
	hostname, err := os.Hostname()
	if err != nil {
		return err
	}

	// Инициализируем логгер
	log.Init(
		model.SystemInfo{
			Hostname: hostname,
			Version:  version,
			Build:    build,
			Env:      *envMode,
		},
		logHandlers...,
	)

	// Получаем конфиг
	log.Info(ctx, "Получаем конфиг")
	cfg, err := config.GetConfig()
	if err != nil {
		return err
	}

	// Инициализируем все синглтоны
	log.Info(ctx, "Инициализируем синглтоны")
	initSingletons(cfg)

	log.Info(ctx, "Инициализируем трейсер")
	if err = trace.StartTracing(cfg.Tracer, cfg.ServiceName); err != nil {
		return err
	}

	// Подключаемся к базе данных
	log.Info(ctx, "Подключаемся к БД")
	pgsql, err := postgresql.NewClientSQL(cfg.Repository, cfg.DBName)
	if err != nil {
		return err
	}
	defer pgsql.Close()

	// Запускаем миграции в базе данных
	// TODO: Подумать, как откатывать миграции при ошибках
	log.Info(ctx, "Запускаем миграции")
	postgreSQLMigrator := migrator.NewMigrator(
		pgsql,
		migrator.MigratorConfig{
			EmbedMigrations: migrations.EmbedMigrationsPostgreSQL,
			Dir:             "pgsql",
		},
	)
	if err = postgreSQLMigrator.Up(ctx); err != nil {
		return err
	}

	log.Info(ctx, "Инициализируем Telegram-бота")
	tgBot, err := telebot.NewBot(telebot.Settings{
		URL:     "",
		Token:   cfg.Telegram.Token,
		Updates: 0,
		Poller: &telebot.LongPoller{
			Limit:          0,
			Timeout:        10 * time.Second,
			LastUpdateID:   0,
			AllowedUpdates: nil,
		},
		Synchronous: false,
		Verbose:     false,
		ParseMode:   telebot.ModeHTML,
		OnError: func(err error, c telebot.Context) {
			log.Error(ctx, err)
		},
		Client:  nil,
		Offline: false,
	})
	if err != nil {
		return errors.InternalServer.Wrap(err)
	}
	defer tgBot.Close()

	// Регистрируем репозитории
	saverRepository := saverRepository.NewSaverRepository(pgsql)
	statisticRepository := statisticRepository.NewStatisticRepository(pgsql)
	swearRepository := swearRepository.NewSwearRepository(pgsql)

	// Регистрируем сервисы
	saverService := saverService.NewSaverService(saverRepository)
	statisticService, err := statisticService.NewStatisticService(statisticRepository)
	if err != nil {
		return err
	}
	swearService := swearService.NewSwearService(swearRepository)
	checkerService.NewCheckerService(statisticService, swearService, saverService, tgBot)

	log.Info(ctx, "Запускаем планировщик")
	if err = shedulerService.NewScheduler().Start(); err != nil {
		return err
	}

	// Регистрируем HTTP-эндпоинты
	r := router.NewRouter()
	r.Mount("/swagger", httpSwagger.WrapHandler)

	// Регистрируем Телеграм-эндпоинты
	swearEndpoint.NewSwearEndpoint(tgBot, swearService)
	statisticEndpoint.NewStatisticEndpoint(tgBot, statisticService)
	tgBot.Handle("/start", func(c telebot.Context) error {
		if err := c.Send("Ну че ебана рот, погнали нахуй"); err != nil {
			return errors.InternalServer.Wrap(err)
		}
		return nil
	})

	server, err := server.GetDefaultServer(cfg.HTTP, r)
	if err != nil {
		return err
	}

	// Создаем wait группу
	eg, ctx := errgroup.WithContext(ctx)

	// Запускаем HTTP-сервер
	eg.Go(func() error { return server.Serve(ctx) })

	eg.Go(func() error {
		log.Info(ctx, "Telegram bot is listening")
		tgBot.Start()
		return nil
	})

	// Запускаем горутину, ожидающую завершение контекста
	eg.Go(func() error {

		// Если контекст завершился, значит процесс убили
		<-ctx.Done()

		// Плавно завершаем работу сервера
		server.Shutdown(ctx)

		return nil
	})

	// Ждем завершения контекста или ошибок в горутинах
	return eg.Wait()
}

func initSingletons(cfg config.Config) {

	stackTrace.Init(cfg.ServiceName)

	// Конфигурируем decimal, чтобы в JSON не было кавычек
	decimal.MarshalJSONWithoutQuotes = true

}
