package main

import (
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

const DEFAULT_PROMPT = `# LLM Prompt`

var (
	rFlag bool
	vFlag bool
	oFlag string
	pFlag string
)

var rootCmd = &cobra.Command{
	Use:   "sourceprompt [path]",
	Short: "Converts your codebase into LLM prompt ",
	Args:  cobra.ExactArgs(1),
	Run:   run,
}

var programLevel *slog.LevelVar

func init() {
	rootCmd.Flags().BoolVarP(&rFlag, "raw", "r", false, "Return just file contents without LLM prompt")
	rootCmd.Flags().BoolVarP(&vFlag, "verbose", "v", false, "Enable verbose output")
	rootCmd.Flags().StringVarP(&oFlag, "output", "o", "", "Output file path")
	rootCmd.Flags().StringVarP(&pFlag, "prompt", "p", "", "Prompt file path")

	programLevel = new(slog.LevelVar)
	h := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: programLevel})
	slog.SetDefault(slog.New(h))
}

func run(cmd *cobra.Command, args []string) {
	if vFlag {
		programLevel.Set(slog.LevelDebug)
	}

	path := args[0]

	slog.Debug("Processing", "path", path)

	if !isGitURL(path) && !isFilePath(path) {
		logErrAndExit(fmt.Errorf("argument must be a valid git URL or file path"))
	}

	sb := strings.Builder{}

	if rFlag {
		slog.Debug("Raw mode - skipping LLM prompt")
	} else {
		if pFlag != "" {
			slog.Debug("Using custom prompt")
			promptContent, err := getCustomPromptContent(pFlag)
			if err != nil {
				logErrAndExit(err)
			}
			sb.Write(promptContent)
			sb.WriteString("\n\n")
		} else {
			sb.WriteString(DEFAULT_PROMPT + "\n\n")
		}
	}

	err := processPath(path, &sb)
	if err != nil {
		logErrAndExit(err)
	}

	slog.Debug("Processing done")

	if oFlag != "" {
		slog.Debug("Saving", "output file", oFlag)
		err := writeToFile(oFlag, []byte(sb.String()))
		if err != nil {
			logErrAndExit(err)
		}
		slog.Debug("File saved sucessfully")
	} else {
		fmt.Println(sb.String())
	}
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
