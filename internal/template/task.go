package template

import (
	"bytes"
	"text/template"
	"tg-bot/internal/task"
)

type AllTaskTemplate struct {
	Tasks []task.Task
}

func NewAllTaskTemplate(tasks []task.Task) *AllTaskTemplate {
	return &AllTaskTemplate{
		Tasks: tasks,
	}
}

func (t *AllTaskTemplate) GetText() (string, error) {
	tmpl := template.New("all_tasks")
	text := "Your tasks:\n"
	text += "{{range $index, $task := .Tasks}}"
	text += "{{if .IsDone}}~{{$task.Id}} \\- {{$task.Text}}~{{- else}}{{$task.Id}} \\- {{$task.Text}}{{- end}}\n"
	text += "{{end}}"

	tmpl, err := tmpl.Parse(text)
	if err != nil {
		return "", err
	}

	var result bytes.Buffer
	err = tmpl.Execute(&result, t)
	if err != nil {
		return "", err
	}

	return result.String(), nil
}
