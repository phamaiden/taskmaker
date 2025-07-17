package cmd

import (
	"aiden/taskmaker/services"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)

func init() {
	rootCmd.AddCommand(markInProgCmd)
}

var markInProgCmd = &cobra.Command{
	Use:   "mip TASK_ID ",
	Short: "Mark the status of a task to in-progress",
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

		err = services.UpdateTaskStatus(id, services.TASK_STATUS_IN_PROGRESS)
		if err != nil {
			fmt.Println("Error updating task:", err)
			return
		}

		fmt.Printf("Task %v updated successfully\n", id)
	},
}
