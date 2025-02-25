# Tree View:
```
.
├─LICENSE
├─README.md
├─build_and_run.ps1
├─go.mod
├─goreleaser.yaml
├─main.go
└─test_root
  ├─File at root A.txt
  ├─File at root B.md
  ├─folder 01
  │ ├─File at folder 01 I.txt
  │ ├─File at folder 01 II.md
  │ └─File at folder 01 III.csv
  └─folder 02
    ├─File at folder 02 I.txt
    ├─File at folder 02 II.md
    ├─File at folder 02 III.csv
    └─folder 02 01
      ├─File at folder 02 01 I.txt
      ├─File at folder 02 01 II.md
      └─File at folder 02 01 III.csv
```

# Content:

## LICENSE
```
MIT License

Copyright (c) 2024 Carlos Tarjano

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

```

## README.md
```md



# CodeWeaver: Generate a Markdown Document of Your Codebase Structure and Content

CodeWeaver is a command-line tool designed to weave your codebase into a single, easy-to-navigate Markdown document. It recursively scans a directory, generating a structured representation of your project's file hierarchy and embedding the content of each file within code blocks. This tool simplifies codebase sharing, documentation, and integration with AI/ML code analysis tools by providing a consolidated and readable Markdown output.
The output for the current repository can be found [here](https://github.com/tesserato/CodeWeaver/blob/main/codebase.md).

# Key Features

* **Comprehensive Codebase Documentation:** Generates a Markdown file that meticulously outlines your project's directory and file structure in a clear, tree-like format.
* **Code Content Inclusion:** Embeds the complete content of each file directly within the Markdown document, enclosed in syntax-highlighted code blocks based on file extensions.
* **Flexible Path Filtering:**  Utilize regular expressions to define ignore patterns, allowing you to exclude specific files and directories from the generated documentation (e.g., `.git`, build artifacts, specific file types).
* **Optional Path Logging:** Choose to save lists of included and excluded file paths to separate files for detailed tracking and debugging of your ignore rules.
* **Simple Command-Line Interface:**  Offers an intuitive command-line interface with straightforward options for customization.

# Installation

If you have Go installed, run `go install github.com/tesserato/CodeWeaver@latest`to install the latest version of CodeWeaver or `go install github.com/tesserato/CodeWeaver@vX.Y.Z` to install a specific version.

Alternatively, download the appropriate pre built executable from the [releases page](https://github.com/tesserato/CodeWeaver/releases).

If necessary, make the `codeweaver` executable by using the `chmod` command:

```bash
chmod +x codeweaver
```

# Usage

## For help, run
```bash
codeweaver -h
```

## For actual usage, run
```bash
codeweaver [options]
```

**Options:**

| Option                            | Description                                                               | Default Value           |
| --------------------------------- | ------------------------------------------------------------------------- | ----------------------- |
| `-dir <directory>`                | The root directory to scan and document.                                  | Current directory (`.`) |
| `-output <filename>`              | The name of the output Markdown file.                                     | `codebase.md`           |
| `-ignore "<regex patterns>"`      | Comma-separated list of regular expression patterns for paths to exclude. | `\.git.*`               |
| `-included-paths-file <filename>` | File to save the list of paths that were included in the documentation.   | None                    |
| `-excluded-paths-file <filename>` | File to save the list of paths that were excluded from the documentation. | None                    |
| `-help`                           | Display this help message and exit.                                       |                         |

# Examples

## **Generate documentation for the current directory:**

   ```bash
   ./codeweaver
   ```
   This will create a file named `codebase.md` in the current directory, documenting the structure and content of the current directory and its subdirectories (excluding paths matching the default ignore pattern `\.git.*`).

## **Specify a different input directory and output file:**

   ```bash
   ./codeweaver -dir=my_project -output=project_docs.md
   ```
   This command will process the `my_project` directory and save the documentation to `project_docs.md`.

## **Ignore specific file types and directories:**

   ```bash
   ./codeweaver -ignore="\.log,temp,build" -output=detailed_docs.md
   ```
   This example will generate `detailed_docs.md`, excluding any files or directories with names containing `.log`, `temp`, or `build`. Regular expression patterns are comma-separated.

## **Save lists of included and excluded paths:**

   ```bash
   ./codeweaver -ignore="node_modules" -included-paths-file=included.txt -excluded-paths-file=excluded.txt -output=code_overview.md
   ```
   This command will create `code_overview.md` while also saving the list of included paths to `included.txt` and the list of excluded paths (due to the `node_modules` ignore pattern) to `excluded.txt`.

# Contributing

Contributions are welcome! If you encounter any issues, have suggestions for new features, or want to improve CodeWeaver, please feel free to open an issue or submit a pull request on the project's GitHub repository.

# License

CodeWeaver is released under the [MIT License](LICENSE). See the `LICENSE` file for complete license details.

# Star History

[![Star History Chart](https://api.star-history.com/svg?repos=tesserato/CodeWeaver&type=Date)](https://star-history.com/#tesserato/CodeWeaver&Date)

# Alternatives

## GitHub Repositories

- **ai-context** - [https://github.com/tanq16/ai-context](https://github.com/tanq16/ai-context) [![GitHub stars](https://img.shields.io/github/stars/tanq16/ai-context?style=social)](https://github.com/tanq16/ai-context)
- **bundle-codebases** - [https://github.com/manfrin/bundle-codebases](https://github.com/manfrin/bundle-codebases) [![GitHub stars](https://img.shields.io/github/stars/manfrin/bundle-codebases?style=social)](https://github.com/manfrin/bundle-codebases)
- **code2prompt** - [https://github.com/mufeedvh/code2prompt](https://github.com/mufeedvh/code2prompt) [![GitHub stars](https://img.shields.io/github/stars/mufeedvh/code2prompt?style=social)](https://github.com/mufeedvh/code2prompt)
- **code2text** - [https://github.com/forrest321/code2text](https://github.com/forrest321/code2text) [![GitHub stars](https://img.shields.io/github/stars/forrest321/code2text?style=social)](https://github.com/forrest321/code2text)
- **codefetch** - [https://github.com/regenrek/codefetch](https://github.com/regenrek/codefetch) [![GitHub stars](https://img.shields.io/github/stars/regenrek/codefetch?style=social)](https://github.com/regenrek/codefetch)
- **copcon** - [https://github.com/kasperjunge/copcon](https://github.com/kasperjunge/copcon) [![GitHub stars](https://img.shields.io/github/stars/kasperjunge/copcon?style=social)](https://github.com/kasperjunge/copcon)
- **describe** - [https://github.com/rodlaf/describe](https://github.com/rodlaf/describe) [![GitHub stars](https://img.shields.io/github/stars/rodlaf/describe?style=social)](https://github.com/rodlaf/describe)
- **feed-llm** - [https://github.com/nahco314/feed-llm](https://github.com/nahco314/feed-llm) [![GitHub stars](https://img.shields.io/github/stars/nahco314/feed-llm?style=social)](https://github.com/nahco314/feed-llm)
- **files-to-prompt** - [https://github.com/simonw/files-to-prompt](https://github.com/simonw/files-to-prompt) [![GitHub stars](https://img.shields.io/github/stars/simonw/files-to-prompt?style=social)](https://github.com/simonw/files-to-prompt)
- **ggrab** - [https://github.com/keizo/ggrab](https://github.com/keizo/ggrab) [![GitHub stars](https://img.shields.io/github/stars/keizo/ggrab?style=social)](https://github.com/keizo/ggrab)
- **gitingest** - [https://gitingest.com/](https://gitingest.com/) [![GitHub stars](https://img.shields.io/github/stars/cyclotruc/gitingest?style=social)](https://github.com/cyclotruc/gitingest)
- **gitpodcast** - [https://gitpodcast.com](https://gitpodcast.com) [![GitHub stars](https://img.shields.io/github/stars/BandarLabs/gitpodcast?style=social)](https://github.com/BandarLabs/gitpodcast)
- **globcat.sh** - [https://github.com/jzombie/globcat.sh](https://github.com/jzombie/globcat.sh) [![GitHub stars](https://img.shields.io/github/stars/jzombie/globcat.sh?style=social)](https://github.com/jzombie/globcat.sh)
- **grimoire** - [https://github.com/foresturquhart/grimoire](https://github.com/foresturquhart/grimoire) [![GitHub stars](https://img.shields.io/github/stars/foresturquhart/grimoire?style=social)](https://github.com/foresturquhart/grimoire)
- **llmcat** - [https://github.com/azer/llmcat](https://github.com/azer/llmcat) [![GitHub stars](https://img.shields.io/github/stars/azer/llmcat?style=social)](https://github.com/azer/llmcat)
- **RepoMix** - [https://github.com/yamadashy/repomix](https://github.com/yamadashy/repomix) [![GitHub stars](https://img.shields.io/github/stars/yamadashy/repomix?style=social)](https://github.com/yamadashy/repomix)
- **techdocs** - [https://github.com/thesurlydev/techdocs](https://github.com/thesurlydev/techdocs) [![GitHub stars](https://img.shields.io/github/stars/thesurlydev/techdocs?style=social)](https://github.com/thesurlydev/techdocs)
- **thisismy** - [https://github.com/franzenzenhofer/thisismy](https://github.com/franzenzenhofer/thisismy) [![GitHub stars](https://img.shields.io/github/stars/franzenzenhofer/thisismy?style=social)](https://github.com/franzenzenhofer/thisismy)
- **yek** - [https://github.com/bodo-run/yek](https://github.com/bodo-run/yek) [![GitHub stars](https://img.shields.io/github/stars/bodo-run/yek?style=social)](https://github.com/bodo-run/yek)
- **your-source-to-prompt** - [https://github.com/Dicklesworthstone/your-source-to-prompt.html](https://github.com/Dicklesworthstone/your-source-to-prompt.html) [![GitHub stars](https://img.shields.io/github/stars/Dicklesworthstone/your-source-to-prompt.html?style=social)](https://github.com/Dicklesworthstone/your-source-to-prompt)
- **ingest** - [https://github.com/sammcj/ingest](https://github.com/sammcj/ingest) [![GitHub stars](https://img.shields.io/github/stars/sammcj/ingest?style=social)](https://github.com/sammcj/ingest)
- **onefilellm** - [https://github.com/jimmc414/onefilellm](https://github.com/jimmc414/onefilellm) [![GitHub stars](https://img.shields.io/github/stars/jimmc414/onefilellm?style=social)](https://github.com/jimmc414/onefilellm)
- **repo2file** - [https://github.com/artkulak/repo2file](https://github.com/artkulak/repo2file) [![GitHub stars](https://img.shields.io/github/stars/artkulak/repo2file?style=social)](https://github.com/artkulak/repo2file)
- **code2prompt** - [https://github.com/mufeedvh/code2prompt](https://github.com/mufeedvh/code2prompt) [![GitHub stars](https://img.shields.io/github/stars/mufeedvh/code2prompt?style=social)](https://github.com/mufeedvh/code2prompt)
- **clipsource** - [https://github.com/strizzo/clipsource](https://github.com/strizzo/clipsource) [![GitHub stars](https://img.shields.io/github/stars/strizzo/clipsource?style=social)](https://github.com/strizzo/clipsource)

## Other

- **r2md** - [https://crates.io/crates/r2md](https://crates.io/crates/r2md)
- **repo2txt** - [https://chathub.gg/repo2txt](https://chathub.gg/repo2txt)
- **repo2txt** - [https://repo2txt.simplebasedomain.com/local.html](https://repo2txt.simplebasedomain.com/local.html)
- **repoprompt** - [https://www.repoprompt.com](https://www.repoprompt.com)


## VSCode Extensions

- **Codebase to Markdown** - [https://marketplace.visualstudio.com/items?itemName=DVYIO.combine-open-files](https://marketplace.visualstudio.com/items?itemName=DVYIO.combine-open-files)

```

## build_and_run.ps1
```ps1
go build .

git describe --tags --abbrev=0

./CodeWeaver -version

./CodeWeaver -ignore="\.git.*,.+\.exe,codebase.md,excluded_paths.txt" -excluded-paths-file="excluded_paths.txt"
```

## go.mod
```mod
module github.com/tesserato/CodeWeaver

go 1.23.0

```

## goreleaser.yaml
```yaml
# vim: set ts=2 sw=2 tw=0 fo=cnqoj

version: 2

builds:
  - env:
      - CGO_ENABLED=0
    goos:
      - linux
      - windows
      - darwin
    ldflags:
      - -s -w
      - -X main.version={{.Version}}
      - -X main.commit={{.Commit}}
      - -X main.date={{.Date}}

archives:
  - formats: [tar.gz]
    # this name template makes the OS and Arch compatible with the results of `uname`.
    name_template: >-
      {{ .ProjectName }}_
      {{- title .Os }}_
      {{- if eq .Arch "amd64" }}x86_64
      {{- else if eq .Arch "386" }}i386
      {{- else }}{{ .Arch }}{{ end }}
      {{- if .Arm }}v{{ .Arm }}{{ end }}
    # use zip for windows archives
    format_overrides:
      - goos: windows
        formats: [ 'zip' ]

changelog:
  sort: asc
  filters:
    exclude:
      - "^docs:"
      - "^test:"

release:
  footer: >-

    ---

    Released by [GoReleaser](https://github.com/goreleaser/goreleaser).
```

## main.go
```go
package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var version = "v0.0.8"

func main() {
	// Define command line flags
	dirPath := flag.String("input", ".", "Directory to scan")
	outputFileName := flag.String("output", "codebase.md", "Output file name")
	ignorePatterns := flag.String("ignore", `\.git.*`, "Comma-separated list of regular expression patterns that match the paths to be ignored")
	includedPathsFile := flag.String("included-paths-file", "", "File to save included paths (optional). If provided, the included paths will be saved to the file and not printed to the console.")
	excludedPathsFile := flag.String("excluded-paths-file", "", "File to save excluded paths (optional). If provided, the excluded paths will be saved to the file and not printed to the console.")
	showVersion := flag.Bool("version", false, "Show version and exit")
	showHelp := flag.Bool("help", false, "Show help message and exit")

	flag.Parse()

	if *showVersion {
		fmt.Println(version)
		return
	}

	// Check if help flag is set or no arguments are provided
	if *showHelp || len(os.Args) == 1 {
		printHelp()
		return
	}

	// Split ignore patterns string into a slice
	ignoreListString := strings.Split(*ignorePatterns, ",")
	ignoreList := make([]*regexp.Regexp, len(ignoreListString))

	fmt.Println("Patterns to ignore:")
	for i, pattern := range ignoreListString {
		fmt.Println(ignoreListString[i])
		ignoreList[i] = regexp.MustCompile(strings.TrimSpace(pattern))

	}

	// Create the output file
	outputFile, err := os.Create(*outputFileName)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	// Write the codebase tree to the output file
	fmt.Fprintln(outputFile, "# Tree View:\n```")
	fmt.Fprintf(outputFile, "%s\n", *dirPath)

	var depthOpen map[int]bool
	depthOpen = make(map[int]bool)
	err = printTree(*dirPath, 0, depthOpen, ignoreList, outputFile)
	if err != nil {
		fmt.Println("Error printing codebase tree:", err)
		return
	}
	fmt.Fprintln(outputFile, "```")

	// Write the code content to the output file
	fmt.Fprintln(outputFile, "\n# Content:\n")
	err = writeCodeContent(*dirPath, ignoreList, outputFile, *includedPathsFile, *excludedPathsFile)
	if err != nil {
		fmt.Println("Error writing code content:", err)
		return
	}

	fmt.Println("Codebase documentation generated successfully!")
}

// printTree recursively walks the directory tree and prints the structure to the output file
func printTree(dirPath string, depth int, depthOpen map[int]bool, ignoreList []*regexp.Regexp, outputFile *os.File) error {
	files, err := os.ReadDir(dirPath)
	if err != nil {
		return err
	}

	// Filter out ignored files and directories
	var filteredFiles []fs.DirEntry
	for _, file := range files {
		filePath := filepath.Join(dirPath, file.Name())
		relPath, _ := filepath.Rel(".", filePath)
		if !shouldIgnore(relPath, ignoreList) {
			filteredFiles = append(filteredFiles, file)
		}
	}

	for i, file := range filteredFiles { // Iterate over filtered files
		filePath := filepath.Join(dirPath, file.Name())

		var pipe string = "├─"
		depthOpen[depth] = true
		if i == len(filteredFiles)-1 { // Use filteredFiles length
			pipe = "└─"
			depthOpen[depth] = false
		}

		var indent = []rune("")
		if depth > 0 {
			indent = []rune(strings.Repeat("  ", depth))
			for j := 0; j < depth; j++ {
				if depthOpen[j] {
					indent[j*2] = '│'
				}
			}
		}

		if file.IsDir() {
			fmt.Fprintf(outputFile, "%s%s%s\n", string(indent), pipe, file.Name())
			printTree(filePath, depth+1, depthOpen, ignoreList, outputFile)
			depthOpen[depth] = false
		} else {
			fmt.Fprintf(outputFile, "%s%s%s\n", string(indent), pipe, file.Name())
		}
	}

	return nil
}

// writeCodeContent reads the content of each file and writes it to the output file within a code block
func writeCodeContent(dirPath string, ignoreList []*regexp.Regexp, outputFile *os.File, includedPathsFile, excludedPathsFile string) error {
	var Red = "\033[31m"
	var Green = "\033[32m"
	var Reset = "\033[0m"
	var includedPaths []string
	var excludedPaths []string

	err := filepath.WalkDir(dirPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Check if the file should be ignored
		relPath, _ := filepath.Rel(".", path)
		if shouldIgnore(relPath, ignoreList) {
			if excludedPathsFile == "" {
				fmt.Println(Red + "- " + path + Reset)
			} else {
				excludedPaths = append(excludedPaths, path)
			}
			return nil
		}

		if includedPathsFile == "" {
			fmt.Println(Green + "+ " + path + Reset)
		} else {
			includedPaths = append(includedPaths, path)
		}

		if !d.IsDir() {
			content, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			extension := filepath.Ext(path)
			extension = strings.ToLower(extension)
			extension = strings.TrimPrefix(extension, ".")
			fmt.Fprintf(outputFile, "## %s\n", path)
			fmt.Fprintf(outputFile, "```%s\n%s\n```\n\n", extension, content)
		}

		return nil
	})

	// Save included paths to file (if filename provided)
	if includedPathsFile != "" {
		err = savePathsToFile(includedPathsFile, includedPaths)
		if err != nil {
			return fmt.Errorf("error saving included paths to file: %w", err)
		}
	}

	// Save excluded paths to file (if filename provided)
	if excludedPathsFile != "" {
		err = savePathsToFile(excludedPathsFile, excludedPaths)
		if err != nil {
			return fmt.Errorf("error saving excluded paths to file: %w", err)
		}
	}

	return err
}

// shouldIgnore checks if a given path should be ignored based on the ignore patterns
func shouldIgnore(path string, ignoreList []*regexp.Regexp) bool {
	if path == "." {
		return true
	}
	for _, pattern := range ignoreList {
		if pattern.MatchString(path) {
			return true
		}
	}
	return false
}

// printHelp prints the help message
func printHelp() {
	fmt.Println("Usage: go run codemerge.go [options]")
	fmt.Println("\nOptions:")
	flag.PrintDefaults()
}

// savePathsToFile saves a list of paths to a file, one per line
func savePathsToFile(filename string, paths []string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, path := range paths {
		_, err := fmt.Fprintln(file, path)
		if err != nil {
			return err
		}
	}

	return nil
}

```

## test_root\File at root A.txt
```txt

```

## test_root\File at root B.md
```md

```

## test_root\folder 01\File at folder 01 I.txt
```txt

```

## test_root\folder 01\File at folder 01 II.md
```md

```

## test_root\folder 01\File at folder 01 III.csv
```csv

```

## test_root\folder 02\File at folder 02 I.txt
```txt

```

## test_root\folder 02\File at folder 02 II.md
```md

```

## test_root\folder 02\File at folder 02 III.csv
```csv

```

## test_root\folder 02\folder 02 01\File at folder 02 01 I.txt
```txt

```

## test_root\folder 02\folder 02 01\File at folder 02 01 II.md
```md

```

## test_root\folder 02\folder 02 01\File at folder 02 01 III.csv
```csv

```

