package main

import (
	"bufio"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"unicode/utf8"
)

func logErrAndExit(err error) {
	slog.Error(err.Error())
	os.Exit(1)
}

func isURL(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func isGitURL(url string) bool {
	if !isURL(url) {
		return false
	}

	resp, err := http.Get(url)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false
	}

	contentType := resp.Header.Get("Content-Type")
	return contentType == "application/x-git-upload-pack-advertisement"
}

func isFilePath(str string) bool {
	_, err := filepath.Abs(str)
	return err == nil
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

		stringBuilder.WriteString("`" + filePath + "`\n\n")
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

func getCustomPromptContent(promptFilepathOrUrl string) ([]byte, error) {
	if isURL(promptFilepathOrUrl) {
		resp, err := http.Get(promptFilepathOrUrl)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return nil, err
		}

		content, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return content, nil
	}

	content, err := os.ReadFile(promptFilepathOrUrl)
	if err != nil {
		return nil, err
	}

	return content, nil
}
