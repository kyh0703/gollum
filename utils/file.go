package util

import (
	"fmt"
	"html/template"
	"os"
)

type TemplateFile struct{}

func (tf *TemplateFile) MakeDir(dirPath string) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	// make service dir
	if err := os.MkdirAll(fmt.Sprintf("%s/%s", wd, dirPath), 0o751); err != nil {
		return err
	}
	return nil
}

func (tf *TemplateFile) MakeFile(path, fileName string) (*os.File, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	file, err := os.Create(wd + "/" + path + "/" + fileName)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (tf *TemplateFile) WriteFileWithTemplate(file *os.File, name, content string, data any) error {
	tmpl := template.Must(template.New(name).Parse(content))
	return tmpl.Execute(file, data)
}
