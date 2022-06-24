package html

import (
	"bytes"
	"html/template"
)

// Format 发送邮件的html格式化
func Format(htmlText string, sendData interface{}) (string, error) {
	tmpl, err := template.New("emailTemp").Parse(htmlText)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, sendData)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
