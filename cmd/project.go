package cmd

import (
	"fmt"
	"html/template"
	"os"

	"github.com/kyh0703/gollum/templates"
)

type Command struct {
	CmdName   string
	CmdParent string
	*Project
}

type Project struct {
	PackageName  string
	Copyright    string
	AbsolutePath string
	AppName      string
}

func (p *Project) Create() error {
	// check if AbsolutePath exists
	if _, err := os.Stat(p.AbsolutePath); os.IsNotExist(err) {
		// create directory
		if err := os.Mkdir(p.AbsolutePath, 0o754); err != nil {
			return err
		}
	}
	// create main.go
	mainFile, err := os.Create(fmt.Sprintf("%s/main.go", p.AbsolutePath))
	if err != nil {
		return err
	}
	defer mainFile.Close()
	mainTemplate := template.Must(template.New("main").Parse(string(templates.MainTemplate)))
	if err := mainTemplate.Execute(mainFile, p); err != nil {
		return err
	}
	if err := p.makeDir("service"); err != nil {
		return err
	}
	if err := p.makeDir("config"); err != nil {
		return err
	}

	return nil
}

func (p *Project) makeDir(dirName string) error {
	// make config dir
	configDir := fmt.Sprintf("%s/%s", p.AbsolutePath, dirName)
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		if err := os.Mkdir(configDir, 0o751); err != nil {
			return err
		}
	}
	return nil
}
