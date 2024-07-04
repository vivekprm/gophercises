package cmd

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command {
	Use: "task",
	Short: "task is a CLI task manager app.",
	Long: `A fast and flexible static task manager built with
	love by spf13.
	`,
} 