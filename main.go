package main

import (
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
	u "github.com/rostrovsky/sourceprompt/pkg/utils"
)

const DEFAULT_PROMPT = `You will be provided with a markdown text (under the "---" separator) containing the contents of a codebase. Each code snippet will be enclosed in code fences, along with the corresponding file name. Your task is to analyze the codebase and gain a comprehensive understanding of its structure, functionality, and key features.

Please follow these steps:

1. Read through the entire codebase carefully, paying attention to the file names and the code within each code fence.
2. Identify the main components, modules, or classes of the codebase and their responsibilities. Summarize the purpose and functionality of each significant component.
3. Analyze the relationships and dependencies between different parts of the codebase. Identify any important interactions, data flow, or control flow between the components.
4. Extract the most important features and functionalities implemented in the codebase. Highlight any critical algorithms, data structures, or design patterns used.
5. Consider the overall architecture and design of the codebase. Identify any architectural patterns or principles followed, such as MVC, MVVM, or microservices.
6. Evaluate the code quality, readability, and maintainability. Note any areas that could be improved or any potential issues or vulnerabilities.
7. Provide a summary of your analysis, including the key insights, strengths, and weaknesses of the codebase. Offer suggestions for improvements or optimizations, if applicable.
8. Based on your understanding of the codebase, provide guidance on how AI agents can effectively operate across the entire codebase. Identify the entry points, important functions, or APIs that the agents should focus on for interaction and manipulation.
9. Discuss any specific considerations or challenges that AI agents may face when working with this codebase, such as dependencies, external libraries, or platform-specific requirements.
10. Conclude your analysis by providing a high-level overview of the codebase's functionality, architecture, and potential use cases. Highlight any notable features or aspects that make this codebase unique or valuable.

Your analysis should be thorough, insightful, and aimed at enabling AI agents to effectively understand and operate within the given codebase. Provide clear explanations and examples to support your findings and recommendations.

---

`

const (
	version = "1.0.2"
)

var (
	rFlag bool
	vFlag bool
	eFlag string
	iFlag string
	oFlag string
	pFlag string
)

var rootCmd = &cobra.Command{
	Use:   "sourceprompt [path]",
	Short: "Converts your codebase into LLM prompt.\nAccepts local directory path or git repo URL as an argument.",
	Args:  cobra.ExactArgs(1),
	Run:   run,
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the version of sourceprompt",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("sourceprompt version %s\n", version)
	},
}

var programLevel *slog.LevelVar

func init() {
	rootCmd.Flags().BoolVarP(&rFlag, "raw", "r", false, "Return just file contents without LLM prompt")
	rootCmd.Flags().BoolVarP(&vFlag, "verbose", "v", false, "Enable verbose output")
	rootCmd.Flags().StringVarP(&oFlag, "output", "o", "", "Output file path")
	rootCmd.Flags().StringVarP(&pFlag, "prompt", "p", "", "Prompt file path or URL")
	rootCmd.Flags().StringVarP(&eFlag, "exclude", "e", "", "Regular expression of filename patterns to exclude")
	rootCmd.Flags().StringVarP(&iFlag, "include", "i", "", "Regular expression of filename patterns to include")

	rootCmd.AddCommand(versionCmd)

	programLevel = new(slog.LevelVar)
	h := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: programLevel})
	slog.SetDefault(slog.New(h))
}

func run(cmd *cobra.Command, args []string) {
	if vFlag {
		programLevel.Set(slog.LevelDebug)
	}

	path := args[0]

	if !u.IsURL(path) && !u.IsFilePath(path) {
		u.LogErrAndExit(fmt.Errorf("argument must be a valid git URL or file path"))
	}

	sb := strings.Builder{}

	if rFlag {
		slog.Debug("Raw mode - skipping LLM prompt")
	} else {
		if pFlag != "" {
			slog.Debug("Using custom prompt")
			promptContent, err := u.GetCustomPromptContent(pFlag)
			if err != nil {
				u.LogErrAndExit(err)
			}
			sb.Write(promptContent)
			sb.WriteString("\n\n")
		} else {
			sb.WriteString(DEFAULT_PROMPT + "\n\n")
		}
	}

	prefixToRemove := path

	if u.IsURL(path) {
		slog.Debug("Cloning using git", "url", path)

		tempDir, err := os.MkdirTemp("", "sourceprompt-git-clone-")
		if err != nil {
			u.LogErrAndExit(fmt.Errorf("failed to create temporary directory: %v", err))
		}
		defer func() {
			os.RemoveAll(tempDir)
			slog.Debug("Temporary directory removed", "tempDir", tempDir)
		}()

		cmd := exec.Command("git", "clone", path, tempDir)
		err = cmd.Run()
		if err != nil {
			u.LogErrAndExit(fmt.Errorf("failed to clone Git repository: %v", err))
		}

		slog.Debug("Repository cloned succesfully", "tempDir", tempDir)
		path = tempDir
		prefixToRemove = tempDir
	}

	var includeRe *regexp.Regexp
	var excludeRe *regexp.Regexp

	if iFlag != "" {
		re, err := regexp.Compile(iFlag)
		if err != nil {
			u.LogErrAndExit(err)
		}
		includeRe = re
	}

	if eFlag != "" {
		re, err := regexp.Compile(eFlag)
		if err != nil {
			u.LogErrAndExit(err)
		}
		excludeRe = re
	}

	err := u.ProcessPath(path, prefixToRemove, includeRe, excludeRe, &sb)
	if err != nil {
		u.LogErrAndExit(err)
	}

	slog.Debug("Processing done")

	if oFlag != "" {
		slog.Debug("Saving output", "file", oFlag)
		err := u.WriteToFile(oFlag, []byte(sb.String()))
		if err != nil {
			u.LogErrAndExit(err)
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
