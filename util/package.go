package util

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

type GoMod struct {
	Path      string `json:"Path"`
	Main      bool   `json:"Main"`
	Dir       string `json:"Dir"`
	GoMod     string `json:"GoMod"`
	GoVersion string `json:"GoVersion"`
}

type Package struct{}

func (pkg *Package) GetGoModFromExec() (*GoMod, error) {
	out, err := exec.Command("go", "list", "-json", "-m").Output()
	if err != nil {
		return nil, err
	}

	var modFile GoMod
	if err := json.Unmarshal(out, &modFile); err != nil {
		return nil, err
	}

	if modFile.Path == "command-line-arguments" {
		return nil, fmt.Errorf(`Please run 'go mod init <MODNAME>' before 'gollum init'`)
	}

	return &modFile, nil
}

func (pkg *Package) ExecGoGet(mod string) error {
	return exec.Command("go", "get", mod).Run()
}

func (pkg *Package) ExecSwagger() error {
	return nil
}
