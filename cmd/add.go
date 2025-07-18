package cmd

import (
	"aiden/taskmaker/services"
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add \"TASK_DESCRIPTION\"",
	Short: "Add a new task to the task list",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("invalid args: task description is required")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		id, err := services.AddTask(args[0])
		if err != nil {
			fmt.Println("Error adding task:", err)
			return
		}

		fmt.Printf("Task %v added successfully\n", id)
	},
}

// func runAddCmd(args []string) error {

// }
