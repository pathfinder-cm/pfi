package cmd

import (
	"fmt"

	. "github.com/pathfinder-cm/pfi/global"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of pfi",
	Long:  `Print the version number of pfi`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("(P)ath(f)inder (I)nterface --", Version)
	},
}
