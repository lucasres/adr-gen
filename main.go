package main

import (
	"fmt"
	"os"

	"github.com/lucasres/adr-gen/cmd"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "adr-gen",
		Short: "ADR Generator",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	rootCmd.AddCommand(cmd.NewVersionCommand())
	rootCmd.AddCommand(cmd.NewAnalyzeCommand())

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}