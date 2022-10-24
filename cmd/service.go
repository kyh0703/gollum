/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	util "github.com/kyh0703/gollum/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra-cli/cmd"
)

// command instance
var Service *ServiceCommand

// initialize command
func init() {
	Service = NewServiceCommand()
	rootCmd.AddCommand(cmd.Command)
}

// serviceCommand represents the service command
type ServiceCommand struct {
	Name    string
	Command *cobra.Command
	util.TemplateFile
}

func NewServiceCommand() *ServiceCommand {
	service := new(ServiceCommand)
	service.Command = &cobra.Command{
		Use:     "service [name]",
		Aliases: []string{"s"},
		Short:   "Generate a service declaration",
		Example: "service auth",
		Run:     service.Run,
	}
	return service
}

func (cmd *ServiceCommand) Run(command *cobra.Command, args []string) {
	if len(args) < 1 {
		cobra.CheckErr(fmt.Errorf("add need a name for the command"))
	}
	var (
		path     = "routes/" + cmd.Name
		fileName = cmd.Name + "_service.go"
	)
	if err := cmd.MakeDir(path); err != nil {
		cobra.CheckErr(err)
	}
	file, err := cmd.MakeFile(path, fileName)
	if err != nil {
		cobra.CheckErr(err)
	}
	defer file.Close()
}
