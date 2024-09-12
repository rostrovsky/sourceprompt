package utils

import (
	"bufio"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"unicode/utf8"
)

func LogErrAndExit(err error) {
	slog.Error(err.Error())
	os.Exit(1)
}

func IsURL(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func IsFilePath(str string) bool {
	x, err := filepath.Abs(str)
	slog.Debug(x)
	return err == nil
}

func IsBinary(filename string) (bool, error) {
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

func IsMarkdown(filename string) bool {
	return strings.HasSuffix(strings.ToLower(filename), ".md")
}

func DetectLanguage(path string) string {
	filename := filepath.Base(path)
	ext := strings.ToLower(filepath.Ext(filename))

	switch ext {
	case ".go", ".templ":
		return "go"
	case ".py":
		return "python"
	case ".js":
		return "javascript"
	case ".ts":
		return "typescript"
	case ".java":
		return "java"
	case ".c", ".h":
		return "c"
	case ".cpp", ".cxx", ".cc", ".hpp":
		return "cpp"
	case ".cs":
		return "csharp"
	case ".rb":
		return "ruby"
	case ".php":
		return "php"
	case ".swift":
		return "swift"
	case ".kt", ".kts":
		return "kotlin"
	case ".rs":
		return "rust"
	case ".html", ".htm", ".gohtml":
		return "html"
	case ".css":
		return "css"
	case ".sql":
		return "sql"
	case ".sh":
		return "bash"
	case ".pl":
		return "perl"
	case ".r":
		return "r"
	case ".m":
		return "objectivec" // This could also be MATLAB
	case ".vb":
		return "vbnet"
	case ".scala":
		return "scala"
	case ".lua":
		return "lua"
	case ".groovy":
		return "groovy"
	case ".dart":
		return "dart"
	case ".md", ".markdown":
		return "markdown"
	case ".json":
		return "json"
	case ".xml":
		return "xml"
	case ".yaml", ".yml":
		return "yaml"
	case ".tex":
		return "tex"
	case ".dockerfile", ".df":
		return "dockerfile"
	case ".ps1":
		return "powershell"
	case ".scss":
		return "scss"
	case ".toml":
		return "toml"
	case ".zig":
		return "zig"
	case ".nim":
		return "nim"
	case ".hs":
		return "haskell"
	default:
		return ""
	}
}

func ProcessPath(path string, prefixToRemove string, includes []*regexp.Regexp, excludes []*regexp.Regexp, stringBuilder *strings.Builder) error {
	return filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
		slog.Debug("Processing", "path", filePath)

		if err != nil {
			return err
		}

		// skip directories and hidden files
		trimmedPath := filePath
		if prefixToRemove != "." {
			trimmedPath = strings.TrimPrefix(filePath, prefixToRemove)
		}
		trimmedPath = strings.TrimLeft(trimmedPath, "/")
		trimmedPath = strings.TrimLeft(trimmedPath, "\\")
		if info.IsDir() || strings.HasPrefix(trimmedPath, ".") {
			slog.Debug("Skipped: path is dir or starts with dot", "path", filePath, "trimmedPath", trimmedPath)
			return nil
		}

		// skip files that don't match the include pattern
		matchesInclude := len(includes) == 0 // If no include patterns, match all
		for _, include := range includes {
			if include.MatchString(trimmedPath) {
				matchesInclude = true
				break
			}
		}
		if !matchesInclude {
			slog.Debug("Skipped: doesn't match any include pattern", "path", filePath, "trimmedPath", trimmedPath)
			return nil
		}

		// skip files that match the exclude pattern
		for _, exclude := range excludes {
			if exclude.MatchString(trimmedPath) {
				slog.Debug("Skipped: matches exclude pattern", "path", filePath, "trimmedPath", trimmedPath)
				return nil
			}
		}

		// skip binary files
		isBinary, err := IsBinary(filePath)
		if err != nil {
			slog.Debug("Error: can't open file", "path", filePath, "trimmedPath", trimmedPath)
			return err
		}

		if isBinary {
			slog.Debug("Skipped: binary file", "path", filePath, "trimmedPath", trimmedPath)
			return nil
		}

		content, err := os.ReadFile(filePath)
		if err != nil {
			slog.Debug("Error: can't read file", "path", filePath, "trimmedPath", trimmedPath)
			return err
		}

		if prefixToRemove != "" {
			filePath = trimmedPath
		}

		fences := "```"

		if IsMarkdown(filePath) {
			fences = "````"
		}

		stringBuilder.WriteString("`" + filePath + "`\n\n")
		stringBuilder.WriteString(fences + DetectLanguage(filePath) + "\n")
		stringBuilder.Write(content)
		if content[len(content)-1] != '\n' {
			stringBuilder.WriteString("\n")
		}
		stringBuilder.WriteString(fences + "\n\n")

		return nil
	})
}

func WriteToFile(filePath string, content []byte) error {
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

func GetCustomPromptContent(promptFilepathOrUrl string) ([]byte, error) {
	if IsURL(promptFilepathOrUrl) {
		slog.Debug("Downloading prompt file", "url", promptFilepathOrUrl)
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

	slog.Debug("Reading prompt file", "path", promptFilepathOrUrl)
	content, err := os.ReadFile(promptFilepathOrUrl)
	if err != nil {
		return nil, err
	}

	return content, nil
}
