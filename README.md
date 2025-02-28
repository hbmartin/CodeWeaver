


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
| `-input <directory>`              | The root directory to scan and document.                                  | Current directory (`.`) |
| `-output <filename>`              | The name of the output Markdown file.                                     | `codebase.md`           |
| `-ignore "<regex patterns>"`      | Comma-separated list of regular expression patterns for paths to exclude. | `\.git.*`               |
| `-included-paths-file <filename>` | File to save the list of paths that were included in the documentation.   | None                    |
| `-excluded-paths-file <filename>` | File to save the list of paths that were excluded from the documentation. | None                    |
| `-version`                        | Display the version and exit.                                             |                         |
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
