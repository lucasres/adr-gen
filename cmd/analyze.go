package cmd

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/lucasres/adr-gen/internal/engine"
	"github.com/lucasres/adr-gen/internal/file"
	"github.com/lucasres/adr-gen/pkg/helpers"
	"github.com/spf13/cobra"
)

const (
	analyzeFlagTimeout = "timeout"
	analyzeFlagPath    = "path"
)

func NewAnalyzeCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "analyze",
		Run: runAnalyze,
	}

	// Add as flags que o comando tem
	cmd.Flags().IntP(analyzeFlagTimeout, "t", 30, "Set timeout of process")
	cmd.Flags().StringP(analyzeFlagPath, "p", "", "Specify the path to be analyzed")

	return cmd
}

func runAnalyze(cmd *cobra.Command, args []string) {
	timeout, err := cmd.Flags().GetInt(analyzeFlagTimeout)
	helpers.PrintAndExitIfGetFlagReturnError(analyzeFlagTimeout, err)

	path, err := cmd.Flags().GetString(analyzeFlagPath)
	helpers.PrintAndExitIfGetFlagReturnError(analyzeFlagPath, err)

	if len(path) < 1 {
		helpers.PrintErrorAndExit(
			fmt.Errorf("the path to be analyzed is invalid, please specify a valid path with \"--path\" flag"),
		)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	errChanel := make(chan error)
	endChannel := make(chan interface{})

	w, err := getAnalyzeWalker()
	if err != nil {
		helpers.PrintErrorAndExit(err)
	}

	r := getAnalyzeReader()

	go func() {
		if err := w.Walk(ctx, path); err != nil {
			errChanel <- err
		}
	}()

	go func() {
		if err := r.Read(ctx, w); err != nil {
			errChanel <- err
		}
	}()

	var wg sync.WaitGroup

	go func() {
		for content := range r.Out() {
			wg.Add(1)

			go func(fileContents []byte) {
				defer wg.Done()
				e := getAnalyzeEngine()
				if err := e.Run(fileContents); err != nil {
					errChanel <- err
				}
			}(content)
		}

		// Wait all engine goroutine finish
		wg.Wait()
		endChannel <- nil
	}()

	select {
	case <-endChannel:
		fmt.Println("finished")
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
	// @todo: Construct Reader based in some configuration
	return file.NewLocalReader(5)
}

func getAnalyzeEngine() engine.Engine {
	// @todo: Create Engine based in some configuration
	return engine.NewSengine()
}
