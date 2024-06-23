package service

import (
	"web-server-template/internal/cmd/service/run"

	"github.com/spf13/cobra"
)

var ServiceCmd = &cobra.Command{
	Use:   "service",
	Short: "The main service command",
	Long:  ``,
	// Args:  cobra.MinimumNArgs(1),
}

func init() {
	ServiceCmd.AddCommand(run.RunCmd)
}
