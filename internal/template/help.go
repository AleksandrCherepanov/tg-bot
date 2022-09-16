package template

import (
	"bytes"
	"text/template"
)

type HelpTemplate struct {
	Commands []string
}

func NewHelpTemplate() *HelpTemplate {
	t := &HelpTemplate{}
	t.Commands = []string{
		"`/start` \\- start to work",
		"`/help` \\- help",
		"`/l` \\- get lists of user",
		"`/lc name` \\- create list for user, where `name` is a name for you list",
		"`/ls id` \\- set list as a current one, where `id` is identifier of your list",
		"`/ld id` \\- delete list, where `id` is identifier of your list",
		"`/lda` \\- delete all lists",
		"`/t` \\- get tasks of current list",
		"`/tc name` \\- create task for current list, where `name` is a name of your task",
		"`/td id` \\- delete task for current list, where `id` is identifier of your task",
		"`/tda` \\- delete all tasks for current list",
		"`/tm id` \\- mark task as done/undone for current list, where `id` is identifier of your task",
		"`/tma flag` \\- mark all task as done/undone for current list, where flag is 0 or 1, 0 \\- undone, 1 \\- done",
	}

	return t
}

func (ht *HelpTemplate) GetText() (string, error) {
	tmpl := template.New("help")

	text := "*Remember* working with tasks is possible after setting some of your lists as a current one\\.\n\n"
	text += "*Command list:*\n"
	text += "{{range $i, $command := .Commands}}{{$command}}\n"
	text += "{{end}}"

	tmpl, err := tmpl.Parse(text)
	if err != nil {
		return "", err
	}

	var result bytes.Buffer
	err = tmpl.Execute(&result, ht)
	if err != nil {
		return "", err
	}

	return result.String(), nil
}
