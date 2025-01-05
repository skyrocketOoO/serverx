package run

import (
	"github.com/skyrocketOoO/web-server-template/internal"
	"github.com/spf13/cobra"
)

var RunCmd = &cobra.Command{
	Use:   "run",
	Short: "Run service",
	Long:  ``,
	Run:   internal.RunServer,
}

func init() {
	RunCmd.Flags().StringP("port", "p", "8080", "port")
	RunCmd.Flags().
		StringP("database", "d", "pg", `database enum. allowed: "pg", "sqlite"`)
}
