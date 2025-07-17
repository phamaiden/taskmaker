package cmd

import (
	"aiden/taskmaker/services"
	"errors"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(deleteCmd)
}

var deleteCmd = &cobra.Command{
	Use:   "delete TASK_ID ",
	Short: "Delete a task",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("invalid args: task id required")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid ID")
			return
		}

		err = services.DeleteTask(id)
		if err != nil {
			fmt.Println("Error deleting task:", err)
			return
		}

		fmt.Printf("Task ID: %v deleted successfully\n", id)
	},
}
