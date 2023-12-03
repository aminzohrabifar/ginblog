package cmd

import (
	"fmt"
	"os"
)
import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "help",
	Short: "help command",
	Long:  `display help command`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(1)
	}
}
