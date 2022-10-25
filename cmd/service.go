package cmd

import (
	"fmt"

	"github.com/kyh0703/gollum/util"
	"github.com/spf13/cobra"
)

// command instance
var Service *ServiceCommand

// initialize command
func init() {
	Service = NewServiceCommand()
	rootCmd.AddCommand(Service.Command)
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
