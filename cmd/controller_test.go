package cmd

import (
	"testing"

	"github.com/spf13/cobra"
)

func TestControllerCommand(t *testing.T) {
	cmd := &cobra.Command{}
	ctrl := NewControllerCommand()
	ctrl.Run(cmd, []string{"auth"})
}
