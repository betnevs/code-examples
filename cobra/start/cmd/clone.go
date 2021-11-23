package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var cloneCmd = &cobra.Command{
	Use:   "clone url",
	Short: "Clone a repository",
	Run: func(cmd *cobra.Command, args []string) {
		output, err := ExecuteCommand("git", "clone", args...)
		if err != nil {
			Error(cmd, args, err)
		}
		fmt.Fprintln(os.Stdout, output)
	},
}

func init() {
	rootCmd.AddCommand(cloneCmd)
}
