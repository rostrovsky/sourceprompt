# sourceprompt

Converts your codebase into prompt which you can feed into LLM.

In short: it scans all the text files in specified directory / git repostiory and puts their contents (along with their path) into structured markdown document.

By default provides predefined prompt which can be replaced using `-p` option.

## Ignored files / paths

* Files/directories starting with `.`
* Binary files

## Flags

* `-o, --output` - Output file path. When not specified, output will be printed to `stdout`.
* `-p, --prompt` - Prompt file path or URL. Allows specifying custom prompt which will be put at the beginning of the output. If not specified, default prompt will be used.
* `-r, --raw` - Removes prompt from the output.
* `-v, --verbose` - Prints debug info when enabled.

## Examples

### Output to stdout

```bash
# output to stdout
sourceprompt /path/to/dir
sourceprompt /path/to/dir -v # debug
sourceprompt https://github.com/some/repo


# output to stdout without default prompt
sourceprompt /path/to/dir -r

# output to file
sourceprompt /path/to/dir -o out.md

# output to file with custom prompt
sourceprompt /path/to/dir -o out.md -p my_prompt.txt

# use remote git repo for codebase and remote file  forprompt
sourceprompt https://github.com/some/repo -p https://raw.githubusercontent.com/another/repo/prompt.md -o out.md
```
