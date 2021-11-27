package main

import (
	"github.com/lucasres/adr-gen/cmd"
	"github.com/lucasres/adr-gen/pkg/helpers"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "adrgen",
		Short: "ADR Generator",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	rootCmd.AddCommand(cmd.NewVersionCommand())
	rootCmd.AddCommand(cmd.NewAnalyzeCommand())

	if err := rootCmd.Execute(); err != nil {
		helpers.PrintErrorAndExit(err)
	}
}
