package service

import (
	"context"

	"gopkg.in/telebot.v3"

	"pkg/log"
	saverModel "swearBot/internal/services/saver/model"
)

func (s *CheckerService) CheckMessage(c telebot.Context) error {

	// Формируем базовый контекст
	ctx := context.Background()

	ctx, span := tracer.Start(ctx, "CheckMessage")
	defer span.End()

	go func() {
		if err := s.SaveAllData(ctx, c); err != nil {
			log.Error(ctx, err)
		}
	}()

	// Получаем данные
	message := c.Message()
	user := c.Sender()

	if err := s.swearHandler(ctx, message, user); err != nil {
		return err
	}

	return nil
}

func (s *CheckerService) SaveAllData(ctx context.Context, c telebot.Context) error {

	// Получаем данные
	message := c.Message()
	user := c.Sender()
	chat := c.Chat()

	// Сохраняем юзера
	if err := s.saverService.SaveUser(ctx, saverModel.SaveUserReq{
		ID:                user.ID,
		FirstName:         user.FirstName,
		LastName:          user.LastName,
		IsForum:           user.IsForum,
		Username:          user.Username,
		LanguageCode:      user.LanguageCode,
		IsBot:             user.IsBot,
		IsPremium:         user.IsPremium,
		AddedToMenu:       user.AddedToMenu,
		Usernames:         user.Usernames,
		CustomEmojiStatus: user.CustomEmojiStatus,
	}); err != nil {
		return err
	}

	// Сохраняем чат
	if err := s.saverService.SaveChat(ctx, saverModel.SaveChatReq{
		ID:                       chat.ID,
		Type:                     string(chat.Type),
		Title:                    chat.Title,
		FirstName:                chat.FirstName,
		LastName:                 chat.LastName,
		Username:                 chat.Username,
		Bio:                      chat.Bio,
		Description:              chat.Description,
		InviteLink:               chat.InviteLink,
		SlowMode:                 chat.SlowMode,
		StickerSet:               chat.StickerSet,
		CanSetStickerSet:         chat.CanSetStickerSet,
		CustomEmojiSetName:       chat.CustomEmojiSetName,
		LinkedChatID:             chat.LinkedChatID,
		Private:                  chat.Private,
		Protected:                chat.Protected,
		NoVoiceAndVideo:          chat.NoVoiceAndVideo,
		HasHiddenMembers:         chat.HasHiddenMembers,
		AggressiveAntiSpam:       chat.AggressiveAntiSpam,
		CustomEmojiID:            chat.CustomEmojiID,
		EmojiExpirationUnixtime:  chat.EmojiExpirationUnixtime,
		BackgroundEmojiID:        chat.BackgroundEmojiID,
		AccentColorID:            chat.AccentColorID,
		ProfileAccentColorID:     chat.ProfileAccentColorID,
		ProfileBackgroundEmojiID: chat.ProfileBackgroundEmojiID,
		HasVisibleHistory:        chat.HasVisibleHistory,
		UnrestrictBoosts:         chat.UnrestrictBoosts,
	}); err != nil {
		return err
	}

	// Сохраняем сообщение
	if err := s.saverService.SaveMessage(ctx, saverModel.SaveMessageReq{
		ID:                 message.ID,
		ThreadID:           message.ThreadID,
		Unixtime:           message.Unixtime,
		OriginalMessageID:  message.OriginalMessageID,
		OriginalSignature:  message.OriginalSignature,
		OriginalSenderName: message.OriginalSenderName,
		OriginalUnixtime:   message.OriginalUnixtime,
		AutomaticForward:   message.AutomaticForward,
		LastEdit:           message.LastEdit,
		TopicMessage:       message.TopicMessage,
		Protected:          message.Protected,
		AlbumID:            message.AlbumID,
		Signature:          message.Signature,
		Text:               message.Text,
		Payload:            message.Payload,
		Caption:            message.Caption,
		NewGroupTitle:      message.NewGroupTitle,
		GroupPhotoDeleted:  message.GroupPhotoDeleted,
		GroupCreated:       message.GroupCreated,
		SuperGroupCreated:  message.SuperGroupCreated,
		ChannelCreated:     message.ChannelCreated,
		MigrateTo:          message.MigrateTo,
		MigrateFrom:        message.MigrateFrom,
		ConnectedWebsite:   message.ConnectedWebsite,
		SenderBoostCount:   message.SenderBoostCount,
		HasMediaSpoiler:    message.HasMediaSpoiler,
	}); err != nil {
		return err
	}

	return nil
}
