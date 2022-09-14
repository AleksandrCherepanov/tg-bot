package template

import (
	"bytes"
	"text/template"
)

type HelpTemplate struct {
	Commands map[string]string
}

func NewHelpTemplate() *HelpTemplate {
	t := &HelpTemplate{}
	t.Commands = map[string]string{
		"/start":       "start to work",
		"/help":        "help",
		"/info":        "get current user summary",
		"/l":           "get lists of user",
		`/lc \{name\}`: `create list for user, where \{name\} is a name for you list`,
		"/ls":          "set list as a current one",
		"/lg":          "get list",
		`/ld \{id\}`:   `delete list, where \{id\} is list id`,
		"/lda":         "delete all lists",
		"/t":           "get tasks of current list",
		"/tg":          "get task of current list",
		"/tc":          "create task for current list",
		"/td":          "delete task for current list",
		"/tda":         "delete all tasks for current list",
		"/tm":          "mark task as done/undone for current list",
		"/tma":         "mark all task as done/undone for current list",
	}

	return t
}

func (ht *HelpTemplate) GetText() (string, error) {
	tmpl := template.New("help")
	text := ` 
*Remember* working with tasks is possible after setting some of your lists as a current one\.

*Command list:*
{{range $command, $description := .Commands}}{{$command}} \- {{$description}}
{{end}}
`
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
