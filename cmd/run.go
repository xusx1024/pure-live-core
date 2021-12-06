package cmd

import (
	"github.com/iyear/pure-live/server"
	"github.com/spf13/cobra"
)

var cfgFile string

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run the server",
	Long:  `run the server locally`,
	Run: func(cmd *cobra.Command, args []string) {
		server.Run(cfgFile)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "config.yaml", "config file (default is ./config.yaml)")

}