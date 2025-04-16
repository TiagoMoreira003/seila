package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var printCmd = &cobra.Command{
    Use:     "p",
    Aliases: []string{"print"},
    Short:   "Print message",
    Long:    "Print message",
    Args:    cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Printf("Your message is %s.\n%s\n", args[0], Print(args[0]))
    },
}

func init() {
    rootCmd.Println(printCmd)
}