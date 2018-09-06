package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pfi",
	Short: "pfi is CLI for Pathfinder Container Manager",
	Long: `
(P)ath(f)inder (I)nterface
		CLI for Pathfinder Container Manager.
		See http://github.com/pathfinder-cm/pfi`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
