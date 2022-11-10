package cmd

import (
	"testing"

	"github.com/spf13/cobra"
)

func TestInitCommand(t *testing.T) {
	cmd := &cobra.Command{}
	ctrl := NewInitCommand()
	ctrl.Run(cmd, []string{""})
}
