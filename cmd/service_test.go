package cmd

import (
	"testing"

	"github.com/spf13/cobra"
)

func TestServiceCommand(t *testing.T) {
	cmd := &cobra.Command{}
	svc := NewServiceCommand()
	svc.Run(cmd, []string{"auth"})
}
