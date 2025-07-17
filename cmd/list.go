package cmd

import (
	"aiden/taskmaker/services"
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Long: `List all tasks. You can filter tasks by status:

	list todo
	list in-progress
	list done
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			err := services.ListTasks(args[0])
			if err != nil {
				fmt.Println("Error listing tasks:", err)
				return
			}
		}

		services.ListTasks("all")
	},
}
