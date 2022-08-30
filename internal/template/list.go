package template

import (
	"bytes"
	"text/template"
	"tg-bot/internal/list"
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
	You lists:
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

type CreateListTemplate struct {
	Id   int64
	Name string
}

func NewCreateListTemplate(id int64, name string) *CreateListTemplate {
	return &CreateListTemplate{
		Id:   id,
		Name: name,
	}
}

func (lt *CreateListTemplate) GetText() (string, error) {
	tmpl := template.New("create_list")
	text := ` 
	List was successfully created\.
	Id: {{.Id}} 
	Name: {{.Name}}
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
