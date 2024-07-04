package inlama

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type Cli struct {
	Stream       bool   // Stream input to model
	SystemPrompt string // System prompt for model
	BufferTime   int    // Buffer time for streaming input (in seconds)
	Url          string // Url for model
	Model        string // Model to use
}

func Init() Cli {
	var cli Cli

	var rootCmd = &cobra.Command{
		Use:   "inlama",
		Short: "Inlama is a CLI tool for using Stdin with LLM models",
		Long:  "Inlama allows you to pass stdin to LLMs to generate statistics, reports and more... \nThe behaviour can be tweaked with flags allowing for streaming input as well as configuring different models and system prompts for precise control over the output.",

		Run: func(cmd *cobra.Command, args []string) {
			//
		},
	}

	var defaults = Cli{
		Stream:       false,
		SystemPrompt: "generate a report of the following text. This summary should include the following: 1 line of summary, 1 line of insights, 1 line of questions, 1 line of gaps, 1 line of recommendations. This report should be structured in a yaml format.",
		BufferTime:   1,
		Url:          "http://localhost:11434",
		Model:        "llama3",
	}

	rootCmd.Flags().BoolVarP(&cli.Stream, "follow", "f", defaults.Stream, "Stream input to model")
	rootCmd.Flags().StringVarP(&cli.SystemPrompt, "prompt", "p", defaults.SystemPrompt, "System prompt for model")
	rootCmd.Flags().IntVarP(&cli.BufferTime, "buffer-time", "b", defaults.BufferTime, "Buffer time for streaming input (in seconds)")
	rootCmd.Flags().StringVarP(&cli.Url, "url", "u", defaults.Url, "Url for model")
	rootCmd.Flags().StringVarP(&cli.Model, "model", "m", defaults.Model, "Model to use")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return cli
}
