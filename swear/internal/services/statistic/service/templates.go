package service

import (
	"text/template"

	"pkg/errors"
)

type Templates struct {
	dailyStatistic *template.Template
}

func NewTemplates() (_ *Templates, err error) {

	templates := new(Templates)

	templates.dailyStatistic, err = template.New("dailyStatistic").Parse(dailyStatisticTemplate)
	if err != nil {
		return nil, errors.InternalServer.Wrap(err)
	}

	return templates, nil
}

const dailyStatisticTemplate = `
<b>Статистика за {{ .Date.Format "2006 January 02" }}</b>

{{ range .Statistics }}UserID: {{ .UserID }}
    {{ range .SwearStatistics }}{{ .Swear }}: {{ .Count }}
    {{ end }}
{{ end }}
`
