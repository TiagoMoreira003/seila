package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "seila",
	Short: "seila is a CLI to create a infrastructure in AWS",
	Long: "seila is a CLI to create a infrastructure in AWS",
	Run: func (cmd *cobra.Command, args []string)  {
		fmt.Println("funcionou")
	},
}

func Execute(){
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error while using seila", err)
		os.Exit(1)
	}
}