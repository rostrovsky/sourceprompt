package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"unicode/utf8"

	"github.com/spf13/cobra"
)

var (
	rFlag bool
	vFlag bool
	oFlag string
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

	sb := strings.Builder{}

	if rFlag {
		slog.Debug("Raw mode - skipping LLM prompt")
	} else {
		sb.WriteString("This is the LLM prompt blablabla\n\n")
	}

	err := processPath(path, &sb)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	slog.Debug("Processing done")

	if oFlag != "" {
		slog.Debug("Saving", "output file", oFlag)
		err := writeToFile(oFlag, []byte(sb.String()))
		if err != nil {
			slog.Error(err.Error())
			os.Exit(1)
		}
		slog.Debug("File saved sucessfully")
	} else {
		fmt.Println(sb.String())
	}
}

func isBinary(filename string) (bool, error) {
	f, err := os.Open(filename)
	if err != nil {
		return false, err
	}
	defer f.Close()

	buf := make([]byte, 1024)
	_, err = f.Read(buf)
	if err != nil {
		return false, err
	}

	return !utf8.ValidString(string(buf)), nil
}

func processPath(path string, stringBuilder *strings.Builder) error {
	return filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// skip directories and hidden files
		if info.IsDir() || strings.HasPrefix(filePath, ".") {
			return nil
		}

		// skip binary files
		isBinary, err := isBinary(filePath)
		if err != nil {
			return err
		}

		if isBinary {
			return nil
		}

		slog.Debug("Processing", "path", filePath)

		content, err := os.ReadFile(filePath)
		if err != nil {
			return err
		}

		stringBuilder.WriteString("#### " + filePath + "\n\n")
		stringBuilder.WriteString("```" + "\n")
		stringBuilder.Write(content)
		stringBuilder.WriteString("```" + "\n\n")

		return nil
	})
}

func writeToFile(filePath string, content []byte) error {
	dir := filepath.Dir(filePath)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.Write(content)
	if err != nil {
		return err
	}

	err = writer.Flush()
	if err != nil {
		return err
	}

	return nil
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
