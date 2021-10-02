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
	// @todo: Define timeout using user input
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	errChanel := make(chan error)
	endChannel := make(chan interface{})

	w, err := getAnalyzeWalker()
	if err != nil {
		helpers.PrintErrorAndExit(err)
	}

	r := getAnalyzeReader()

	go func() {
		if err := w.Walk(ctx, "./internal"); err != nil {
			errChanel <- err
		}
	}()

	go func() {
		if err := r.Read(ctx, w); err != nil {
			errChanel <- err
		}
	}()

	go func() {
		for content := range r.Out() {
			fmt.Printf("file Content:\n%s\n\n", content)
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

func getAnalyzeReader() file.Reader {
	// @todo: Construct Reader base in some configuration
	return file.NewLocalReader(5)
}
