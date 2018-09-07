package cmd

import (
	"fmt"

	"github.com/pathfinder-cm/pfi/global"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(rescheduleCmd)
	rescheduleCmd.AddCommand(rescheduleContainerCmd)
}

var rescheduleCmd = &cobra.Command{
	Use:   "reschedule",
	Short: "Specify things to reschedule",
	Long:  `Specify things to reschedule`,
}

var rescheduleContainerCmd = &cobra.Command{
	Use:   "container",
	Short: "Reschedule a new container",
	Long:  `Reschedule a new container`,
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
		container, _ := client.RescheduleContainer(args[0])
		fmt.Println(fmt.Sprintf("Container %s rescheduled!", container.Hostname))
	},
}
