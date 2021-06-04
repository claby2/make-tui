package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/alecthomas/chroma/styles"
	"github.com/spf13/pflag"
)

// Check is a helper function to error check functions which can be used to error check deferred functions
func Check(f func() error) {
	if err := f(); err != nil {
		log.Fatal(err)
	}
}

func getFileContent(filePath string) []string {
	var fileContent []string
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer Check(file.Close)
	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	for scanner.Scan() {
		fileContent = append(fileContent, scanner.Text())
	}
	return fileContent
}

func main() {
	var filePath string
	var theme string

	pflag.StringVarP(&filePath, "file-name", "f", "", "Parse given file as Makefile")
	help := pflag.BoolP("help", "h", false, "Print this message and exit")
	all := pflag.BoolP("all", "a", false, "Display all targets including special built-in targets")
	list := pflag.Bool("list-themes", false, "List built-in syntax highlighting themes")
	pflag.StringVar(&theme, "theme", "vim", "Use a built-in syntax highlighting theme")
	pflag.Parse()

	if *help {
		// Print help message
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		pflag.PrintDefaults()
		os.Exit(0)
	} else if *list {
		fmt.Println("Built-in themes:")
		for theme := range styles.Registry {
			fmt.Println("\t" + theme)
		}
		os.Exit(0)
	} else if filePath == "" {
		// Attempt to find makefile in current directory
		defaultMakefileNames := []string{"GNUmakefile", "makefile", "Makefile"}
		foundFile := false
		for _, name := range defaultMakefileNames {
			if _, err := os.Stat(name); os.IsNotExist(err) == false {
				// File exists
				filePath = name
				foundFile = true
				break
			}
		}
		if !foundFile {
			log.Fatalln("No Makefile found.")
		}
	}

	if theme != "" {
		// Ensure given theme exists in registry.
		if _, ok := styles.Registry[theme]; !ok {
			log.Fatalln("\"" + theme + "\" is not a built-in theme.")
		}
	}

	Check(LoadConfig)

	content := NewParsedContent(filePath, getFileContent(filePath))
	if *all {
		content.SetIncludeSpecialTargets(true)
	}
	content.Parse()

	Render(content, theme)
}
