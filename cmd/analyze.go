package cmd

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/lucasres/adr-generator/internal/file"
	"github.com/spf13/cobra"
)

func NewAnalyzeCommand() *cobra.Command {
	return &cobra.Command{
		Use: "analyze",
		Run: runAnalyze,
	}
}

func runAnalyze(cmd *cobra.Command, args []string) {
	w, err := file.NewLocalWalk(10, nil)
	if err != nil {
		// @todo: Define helper for handle errors
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
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
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	case <-ctx.Done():
		if err := ctx.Err(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}
