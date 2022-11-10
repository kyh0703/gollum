package cmd

import (
	"github.com/spf13/cobra"
)

// command instance
var Init *InitCommand

// initialize command
func init() {
	Init = NewInitCommand()
	rootCmd.AddCommand(Init.Command)
}

// initCmd represents the init command
type InitCommand struct {
	ModName string
	Command *cobra.Command
}

func NewInitCommand() *InitCommand {
	init := new(InitCommand)
	init.Command = &cobra.Command{
		Use:   "init [path]",
		Short: "Initialize a gollum application",
		Run:   init.Run,
	}
	return init
}

func (cmd *InitCommand) Run(command *cobra.Command, args []string) {
}

func (cmd *InitCommand) parseMod() string {
	return ""
}

// func parseModInfo() (Mod, CurDir) {
// 	var (
// 		mod Mod
// 		dir CurDir
// 	)
// 	m := modInfoJSON("-m")
// 	cobra.CheckErr(json.Unmarshal(m, &mod))
// 	if mod.Path == "command-line-arguments" {
// 		cobra.CheckErr("please run `go mod init <MODNAME>` before `gollum init`")
// 	}
// 	e := modInfoJSON("-e")
// 	cobra.CheckErr(json.Unmarshal(e, &dir))
// 	return mod, dir
// }
