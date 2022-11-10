package cmd

import (
	"github.com/kyh0703/gollum/templates"
	"github.com/kyh0703/gollum/util"
	"github.com/spf13/cobra"
)

// command instance
var Repository *RepositoryCommand

// initialize command
func init() {
	Repository = NewRepositoryCommand()
	rootCmd.AddCommand(Repository.Command)
}

// RepositoryCommand represents the Repository command
type RepositoryCommand struct {
	util.Selector
	util.TemplateFile
	Name    string
	Command *cobra.Command
}

func NewRepositoryCommand() *RepositoryCommand {
	Repository := new(RepositoryCommand)
	Repository.Command = &cobra.Command{
		Use:     "Repository [name]",
		Aliases: []string{"s"},
		Short:   "Generate a Repository declaration",
		Example: "Repository auth",
		Run:     Repository.Run,
	}
	return Repository
}

func (cmd *RepositoryCommand) Run(command *cobra.Command, args []string) {
	if len(args) < 1 {
		command.Help()
		return
	}

	cmd.Name = args[0]

	var (
		path     = "routes/" + cmd.Name
		fileName = cmd.Name + "_repository.go"
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
		"repository",
		templates.Repository,
		cmd,
	); err != nil {
		cobra.CheckErr(err)
	}
}
