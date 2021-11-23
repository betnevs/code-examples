package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

type ErrorHandling int

const (
	ContinueOnParseError ErrorHandling = iota + 1
	ExitOnParseError
	PanicOnParseError
	ReturnOnDividedByZero
	PanicOnDividedByZero
)

type OpType int

const (
	ADD OpType = iota + 1
	MINUS
	MULTIPLY
	DIVIDE
)

var (
	parseHandling int
)

var rootCmd = &cobra.Command{
	Use:   "math",
	Short: "Math calc result.",
	Run: func(cmd *cobra.Command, args []string) {
		Error(cmd, args, errors.New("unrecognized subcommand"))
	},
}

func init() {
	rootCmd.PersistentFlags().IntVarP(&parseHandling, "parse_error", "p", int(ContinueOnParseError), "parse case")
}

func Execute() {
	rootCmd.Execute()
}
