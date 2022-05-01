package errors

import (
	"bytes"
	"html/template"
)

func makeTemplate(msg string, data any) (string, error) {
	t := template.Must(template.New("").Parse(msg))

	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return msg, err
	}

	return buf.String(), nil
}
