# sourceprompt

Converts your codebase into prompt which you can feed into LLM.

In short: it scans all the text files in specified directory / git repostiory and puts their content (along with their path) into structured markdown document.

By default it provides predefined prompt which can be replaced using `-p` option.

## Installation

Download binary from releases or install using:

```bash
go install github.com/rostrovsky/sourceprompt@latest
```

## Ignored files / paths

* Files/directories starting with `.`
* Binary files

## Flags

* `-i, --include` - Regex pattern of paths that should be included.
* `-e, --exclude` - Regex pattern of paths that should be excluded.
* `-o, --output` - Output file path. When not specified, output will be printed to `stdout`.
* `-p, --prompt` - Prompt file path or URL. Allows specifying custom prompt which will be put at the beginning of the output. **If not specified, [default prompt](#default-prompt) will be used.**
* `-r, --raw` - Removes prompt from the output.
* `-v, --verbose` - Prints debug info when enabled.

## Examples

### Output to stdout

```bash
# output to stdout
sourceprompt /path/to/dir
sourceprompt /path/to/dir -v # debug
sourceprompt https://github.com/some/repo

# use remote git repo for codebase and remote prompt
sourceprompt https://github.com/some/repo -p https://raw.githubusercontent.com/another/repo/prompt.md -o out.md

# output to stdout without default prompt
sourceprompt /path/to/dir -r

# output to file
sourceprompt /path/to/dir -o out.md

# output to file with custom prompt
sourceprompt /path/to/dir -o out.md -p my_prompt.txt

# include only src/ files
sourceprompt /path/to/dir -o out.md -i '^src'

# exclude markdown files
sourceprompt /path/to/dir -o out.md -e '\.md$'
```

## Default prompt

```
You will be provided with a markdown text (under the "---" separator) containing the contents of a codebase. Each code snippet will be enclosed in code fences, along with the corresponding file name. Your task is to analyze the codebase and gain a comprehensive understanding of its structure, functionality, and key features.

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

```
