/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init [path]",
	Short: "Initialize a Gollum Application",
	Run: func(cmd *cobra.Command, args []string) {
		_, err := InitializeProject(args)
		cobra.CheckErr(err)
	},
}

func InitializeProject(args []string) (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	if len(args) > 0 {
		if args[0] != "." {
			wd = fmt.Sprintf("%s/%s", wd, args[0])
		}
	}

	return "", nil
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

func init() {
	rootCmd.AddCommand(initCmd)
}
