package cmd

import (
	"fmt"

	"github.com/pathfinder-cm/pfi/global"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.AddCommand(deleteContainerCmd)
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Specify things to delete",
	Long:  `Specify things to delete`,
}

var deleteContainerCmd = &cobra.Command{
	Use:   "container",
	Short: "Delete a new container",
	Long:  `Delete a new container`,
	Run: func(cmd *cobra.Command, args []string) {
		client, err := global.NewClient()
		if err != nil {
			fmt.Println(err)
			return
		}

		if len(args) < 1 {
			fmt.Println("You must specify container name")
			return
		}
		container, _ := client.DeleteContainer(args[0])
		fmt.Println(fmt.Sprintf("Container %s deleted!", container.Hostname))
	},
}
