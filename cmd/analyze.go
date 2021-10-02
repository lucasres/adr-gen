package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/lucasres/adr-gen/internal/file"
	"github.com/lucasres/adr-gen/pkg/helpers"
	"github.com/spf13/cobra"
)

func NewAnalyzeCommand() *cobra.Command {
	return &cobra.Command{
		Use: "analyze",
		Run: runAnalyze,
	}
}

func runAnalyze(cmd *cobra.Command, args []string) {
	w, err := getAnalyzeWalker()
	if err != nil {
		helpers.PrintErrorAndExit(err)
	}

	// @todo: Define timeout using user input
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	errChanel := make(chan error)
	endChannel := make(chan interface{})

	go func() {
		if err := w.Walk(ctx, "./internal"); err != nil {
			errChanel <- err
		}
	}()

	// @todo: Create file reader
	go func() {
		for path := range w.Out() {
			fmt.Println("path:", path)
		}

		endChannel <- nil
	}()

	select {
	case <-endChannel:
		fmt.Println("finish")
	case err := <-errChanel:
		helpers.PrintErrorAndExit(err)
	case <-ctx.Done():
		if err := ctx.Err(); err != nil {
			helpers.PrintErrorAndExit(err)
		}
	}
}

func getAnalyzeWalker() (file.Walker, error) {
	// @todo: Construct Walker based in some configuration
	return file.NewLocalWalk(10, nil)
}
