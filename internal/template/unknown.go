package template

import (
	"bytes"
	"text/template"
)

type UnknownTemplate struct {
	Command string
}

func NewUnknownTemplate(command string) *UnknownTemplate {
	t := &UnknownTemplate{}
	t.Command = command
	return t
}

func (ut *UnknownTemplate) GetText() (string, error) {
	tmpl := template.New("unknown")
	text := ` 
Command {{.Command}} is unknown or is not supported yet\.
`
	tmpl, err := tmpl.Parse(text)
	if err != nil {
		return "", err
	}

	var result bytes.Buffer
	err = tmpl.Execute(&result, ut)
	if err != nil {
		return "", err
	}

	return result.String(), nil
}
