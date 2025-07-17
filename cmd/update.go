package cmd

import (
	"aiden/taskmaker/services"
	"errors"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(updateCmd)
}

var updateCmd = &cobra.Command{
	Use:   "update TASK_ID \"TASK_DESCRIPTION\"",
	Short: "Update a task with a new description",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			return errors.New("invalid args: task id and description required")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid ID")
			return
		}

		desc := args[1]
		err = services.UpdateTask(id, desc)
		if err != nil {
			fmt.Println("Error updating task:", err)
			return
		}

		fmt.Printf("Task %v updated successfully\n", id)
	},
}
