package cmd

import (
	"encoding/json"
	"os/exec"

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
		Use:   "init",
		Short: "Initialize a Gollum Application",
		Run:   init.Run,
	}
	return init
}

func (cmd *InitCommand) Run(command *cobra.Command, args []string) {
	if len(args) < 1 {
		command.Help()
		return
	}
	// Check go mod file
}

func parseModInfo() (Mod, CurDir) {
	var (
		mod Mod
		dir CurDir
	)
	m := modInfoJSON("-m")
	cobra.CheckErr(json.Unmarshal(m, &mod))
	if mod.Path == "command-line-arguments" {
		cobra.CheckErr("please run `go mod init <MODNAME>` before `gollum init`")
	}
	e := modInfoJSON("-e")
	cobra.CheckErr(json.Unmarshal(e, &dir))
	return mod, dir
}

func modInfoJSON(args ...string) []byte {
	cmdArgs := append([]string{"list", "-json"}, args...)
	out, err := exec.Command("go", cmdArgs...).Output()
	cobra.CheckErr(err)
	return out
}

func goGet(mod string) error {
	return exec.Command("go", "get", mod).Run()
}
