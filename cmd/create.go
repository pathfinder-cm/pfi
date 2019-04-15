package cmd

import (
	"fmt"

	"github.com/pathfinder-cm/pfi/global"
	"github.com/spf13/cobra"
)

var Image string

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.AddCommand(createContainerCmd)
	createContainerCmd.Flags().StringVarP(&Image, "image", "i", "18.04", "Image that pathfinder will use to create the container")
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Specify things to create",
	Long:  `Specify things to create`,
}

var createContainerCmd = &cobra.Command{
	Use:   "container",
	Short: "Create a new container",
	Long:  `Create a new container`,
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
		container, _ := client.CreateContainer(args[0], Image)
		fmt.Println(fmt.Sprintf("Container %s created!", container.Hostname))
	},
}
