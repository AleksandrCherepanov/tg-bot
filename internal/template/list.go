package template

import (
	"bytes"
	"text/template"

	"github.com/AleksandrCherepanov/tg-bot/internal/list"
)

type AllListTemplate struct {
	Lists []list.TaskList
}

func NewAllListTemplate(lists []list.TaskList) *AllListTemplate {
	return &AllListTemplate{
		Lists: lists,
	}
}

func (lt *AllListTemplate) GetText() (string, error) {
	tmpl := template.New("all_lists")
	text := ` 
	Your lists:
{{range $index, $list := .Lists}}{{$list.Id}} \- {{$list.Name}}
{{end}}
`
	tmpl, err := tmpl.Parse(text)
	if err != nil {
		return "", err
	}

	var result bytes.Buffer
	err = tmpl.Execute(&result, lt)
	if err != nil {
		return "", err
	}

	return result.String(), nil
}
