package template

import (
	"html/template"
	"time"
)

var templatePath = "./template"

var funcMap = template.FuncMap{
	"dateFormat": dateFormat,
}

func dateFormat(t time.Time) string {
	return t.Format("2006/01/02 15:04:05")
}

//SetTemplatePath is helper funcs setup custom template path
func SetTemplatePath(path string) {
	templatePath = path
}

//GetTemplatePath return templatepath
func GetTemplatePath() string {
	return templatePath
}

//LoadTemplate is helper funcs load templatepath
func LoadTemplate(templateName string) (*template.Template, error) {
	//fmt.Println(filepath.Abs(templatePath))
	return template.New(templateName).Funcs(funcMap).ParseFiles(templatePath + "/" + templateName)
}
