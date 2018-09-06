package cmd

import (
	"fmt"

	"github.com/pathfinder-cm/pfi/global"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.AddCommand(getNodesCmd)
	getCmd.AddCommand(getContainersCmd)
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Specify things to get",
	Long:  `Specify things to get`,
}

var getNodesCmd = &cobra.Command{
	Use:   "nodes",
	Short: "List all available nodes",
	Long:  `List all available nodes`,
	Run: func(cmd *cobra.Command, args []string) {
		client, err := global.NewClient()
		if err != nil {
			fmt.Println(err)
			return
		}
		nodes, _ := client.GetNodes()

		table := global.NewTable()
		table.SetHeader([]string{"Hostname", "Ipaddress", "Created At"})
		for _, node := range *nodes {
			table.Append([]string{node.Hostname, node.Ipaddress, node.CreatedAt})
		}
		fmt.Println()
		table.Render()
		fmt.Println()
	},
}

var getContainersCmd = &cobra.Command{
	Use:   "containers",
	Short: "List all available containers",
	Long:  `List all available containers`,
	Run: func(cmd *cobra.Command, args []string) {
		client, err := global.NewClient()
		if err != nil {
			fmt.Println(err)
			return
		}
		containers, _ := client.GetContainers()

		table := global.NewTable()
		table.SetHeader([]string{"Hostname", "Ipaddress", "Image", "Status"})
		for _, container := range *containers {
			table.Append([]string{container.Hostname, container.Ipaddress, container.Image, container.Status})
		}
		fmt.Println()
		table.Render()
		fmt.Println()
	},
}
