/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/kyh0703/gollum/templates"
	util "github.com/kyh0703/gollum/utils"
	"github.com/spf13/cobra"
)

// command instance
var Controller *ControllerCommand

// initialize command
func init() {
	Controller = NewControllerCommand()
	rootCmd.AddCommand(Controller.Command)
}

// controllerCommand represents the service command
type ControllerCommand struct {
	Name    string
	Command *cobra.Command
	util.TemplateFile
}

func NewControllerCommand() *ControllerCommand {
	controller := new(ControllerCommand)
	controller.Command = &cobra.Command{
		Use:     "controller [name]",
		Short:   "Generate a controller declaration",
		Example: "controller auth",
		Run:     controller.Run,
	}
	return controller
}

func (cmd *ControllerCommand) Run(command *cobra.Command, args []string) {
	if len(args) < 1 {
		cobra.CheckErr(fmt.Errorf("add need a name for the command"))
	}
	var (
		path     = "routes/" + cmd.Name
		fileName = cmd.Name + "_controller.go"
	)
	if err := cmd.MakeDir(path); err != nil {
		cobra.CheckErr(err)
	}
	file, err := cmd.MakeFile(path, fileName)
	if err != nil {
		cobra.CheckErr(err)
	}
	defer file.Close()
	cmd.WriteFileWithTemplate(file, "controller", string(templates.ServiceTemplate), data any)
	cmd.WriteFileWithTemplate(file, name string, content string, data any)
}
