package cmd

import (
	"github.com/kyh0703/gollum/templates"
	"github.com/kyh0703/gollum/util"
	"github.com/spf13/cobra"
)

// command instance
var Controller *ControllerCommand

// initialize command
func init() {
	Controller = NewControllerCommand()
	rootCmd.AddCommand(Controller.Command)
}

// controllerCommand represents the controller command
type ControllerCommand struct {
	util.Selector
	util.TemplateFile
	Name    string
	Command *cobra.Command
}

func NewControllerCommand() *ControllerCommand {
	controller := new(ControllerCommand)
	controller.Command = &cobra.Command{
		Use:     "controller [name]",
		Aliases: []string{"co"},
		Short:   "Generate a controller declaration",
		Example: "controller auth",
		Run:     controller.Run,
	}
	return controller
}

func (cmd *ControllerCommand) Run(command *cobra.Command, args []string) {
	if len(args) < 1 {
		command.Help()
		return
	}

	cmd.Name = args[0]

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

	if err := cmd.WriteFileWithTemplate(
		file,
		"controller",
		templates.Controller,
		cmd,
	); err != nil {
		cobra.CheckErr(err)
	}
}
