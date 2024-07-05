package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/vivekprm/gophercises/cli-task-manager/db"
)

var addCommand = &cobra.Command{
	Use: "add",
	Short: "Adds a task to task list.",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		_, err := db.CreateTask(task)
		if err != nil {
			fmt.Println("Something went wrong:", err.Error())
			os.Exit(1)
		}
		fmt.Printf("Added \"%s\" to your task list\n", task)
	},
}

// It can run before the main function.
// https://go.dev/doc/effective_go
func init() {
	RootCmd.AddCommand(addCommand)
}