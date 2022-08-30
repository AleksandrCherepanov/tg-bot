package template

import (
	"bytes"
	"text/template"
)

type StartTemplate struct {
	Steps map[string]string
}

func NewStartTemplate() *StartTemplate {
	t := &StartTemplate{}
	t.Steps = map[string]string{
		"/lc": "create your first list of tasks",
		"/ls": "make created list as a current one",
		"/tc": "create your first task",
	}
	return t
}

func (st *StartTemplate) GetText() (string, error) {
	tmpl := template.New("start")
	text := ` 
Welcome\!
This bot allows you to manage your tasks\. 
There are several steps to start:
{{range $step, $description := .Steps}}{{$step}} \- {{$description}}
{{end}}
To get full list of commands use /help\.
*Remember* working with tasks is possible after setting some of your lists as a current one\.
`
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
