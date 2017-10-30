package template

import (
	"fmt"
	"html/template"
	"path/filepath"
	"time"
)

var templatePath = "./template"

var funcMap = template.FuncMap{
	"dateFormat": dateFormat,
}

func dateFormat(t time.Time) string {
	return t.Format("2006/01/02 15:04:05")
}

func SetTemplatePath(path string) {
	templatePath = path
}
func LoadTemplate(templateName string) (*template.Template, error) {
	fmt.Println(filepath.Abs(templatePath))

	return template.New(templateName).Funcs(funcMap).ParseFiles(templatePath + "/" + templateName)
}
