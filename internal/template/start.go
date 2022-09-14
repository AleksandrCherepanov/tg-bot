package template

import (
	"bytes"
	"text/template"
)

type StartTemplate struct {
	Steps []string
}

func NewStartTemplate() *StartTemplate {
	t := &StartTemplate{}
	t.Steps = []string{
		"`/lc name` \\- create your first list of tasks, where `name` is a name of your list;",
		"`/ls id` \\- make created list as a current one, where `id` is your list identifier;",
		"`/tc name` \\- create your first task, where `name` is a name of your task;",
	}
	return t
}

func (st *StartTemplate) GetText() (string, error) {
	tmpl := template.New("start")
	text := "Welcome\\!\n"
	text += "This bot allows you to manage your tasks\\.\n"
	text += "There are several steps to start:\n"
	text += "{{range $i, $description := .Steps}}{{$description}}\n"
	text += "{{end}}\n"
	text += "To get full list of commands use `/help`\\.\n\n"
	text += "*Remember* working with tasks is possible after setting some of your lists as a current one\\.\n"
	tmpl, err := tmpl.Parse(text)

	if err != nil {
		return "", err
	}

	var result bytes.Buffer
	err = tmpl.Execute(&result, st)
	if err != nil {
		return "", err
	}

	return result.String(), nil
}
