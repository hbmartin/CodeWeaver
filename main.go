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
	includePatterns := flag.String("include", ``, "Comma-separated list of regular expression patterns that match the paths to be included")
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

	var ignoreList []*regexp.Regexp
	var includeList []*regexp.Regexp

	// Process ignore patterns if provided
	if *ignorePatterns != "" {
		ignoreListString := strings.Split(*ignorePatterns, ",")
		ignoreList = make([]*regexp.Regexp, len(ignoreListString))

		for i, pattern := range ignoreListString {
			fmt.Println(ignoreListString[i])
			ignoreList[i] = regexp.MustCompile(strings.TrimSpace(pattern))
		}
	}

	// Process include patterns if provided
	if *includePatterns != "" {
		includeListString := strings.Split(*includePatterns, ",")
		includeList = make([]*regexp.Regexp, len(includeListString))

		for i, pattern := range includeListString {
			fmt.Println(includeListString[i])
			includeList[i] = regexp.MustCompile(strings.TrimSpace(pattern))
		}
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

	depthOpen := make(map[int]bool)
	err = printTree(*dirPath, 0, depthOpen, ignoreList, includeList, outputFile)
	if err != nil {
		fmt.Println("Error printing codebase tree:", err)
		return
	}
	fmt.Fprintln(outputFile, "```")

	// Write the code content to the output file
	fmt.Fprintln(outputFile, "\n# Content:\n")
	err = writeCodeContent(*dirPath, ignoreList, includeList, outputFile, *includedPathsFile, *excludedPathsFile)
	if err != nil {
		fmt.Println("Error writing code content:", err)
		return
	}

	fmt.Println("Codebase documentation generated successfully!")
}

// printTree recursively walks the directory tree and prints the structure to the output file
func printTree(dirPath string, depth int, depthOpen map[int]bool, ignoreList, includeList []*regexp.Regexp, outputFile *os.File) error {
	files, err := os.ReadDir(dirPath)
	if err != nil {
		return err
	}

	// Filter files based on ignore/include patterns
	var filteredFiles []fs.DirEntry
	for _, file := range files {
		filePath := filepath.Join(dirPath, file.Name())
		relPath, _ := filepath.Rel(".", filePath)
		if shouldProcess(relPath, ignoreList, includeList) {
			filteredFiles = append(filteredFiles, file)
		}
	}

	for i, file := range filteredFiles {
		filePath := filepath.Join(dirPath, file.Name())

		var pipe string = "├─"
		depthOpen[depth] = true
		if i == len(filteredFiles)-1 { // Use filteredFiles length
			pipe = "└─"
			depthOpen[depth] = false
		}

		indent := []rune("")
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
			printTree(filePath, depth+1, depthOpen, ignoreList, includeList, outputFile)
			depthOpen[depth] = false
		} else {
			fmt.Fprintf(outputFile, "%s%s%s\n", string(indent), pipe, file.Name())
		}
	}

	return nil
}

// writeCodeContent reads the content of each file and writes it to the output file within a code block
func writeCodeContent(dirPath string, ignoreList, includeList []*regexp.Regexp, outputFile *os.File, includedPathsFile, excludedPathsFile string) error {
	Red := "\033[31m"
	Green := "\033[32m"
	Reset := "\033[0m"
	var includedPaths []string
	var excludedPaths []string

	err := filepath.WalkDir(dirPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Check if the file should be processed
		relPath, _ := filepath.Rel(".", path)
		if !shouldProcess(relPath, ignoreList, includeList) {
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

// shouldProcess determines if a file should be processed based on include and ignore patterns
func shouldProcess(path string, ignoreList, includeList []*regexp.Regexp) bool {
	if path == "." {
		return false
	}

	// If include patterns are specified, the path must match at least one
	if len(includeList) > 0 {
		matched := false
		for _, pattern := range includeList {
			if pattern.MatchString(path) {
				matched = true
				break
			}
		}

		// If it doesn't match any include pattern, exclude it
		if !matched {
			return false
		}
	}

	// Check if the path matches any ignore pattern
	for _, pattern := range ignoreList {
		if pattern.MatchString(path) {
			return false // Exclude if it matches any ignore pattern
		}
	}

	// Include the file if it passes both filters
	return true
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
