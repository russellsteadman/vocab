package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

const version = "v0.2.0"

func printHelp() {
	fmt.Printf("Usage: vocab %s\n", version)
	fmt.Print("Vocab is a tool for generating a vocabulary file from a text file.\n\n")

	fmt.Println("Options:")
	fmt.Println("  -h, --help\t\t\tShow this help message and exit")
	fmt.Println("  -v, --version\t\t\tShow the version number and exit")
	fmt.Println("  -o, --output\t\t\tThe output file to write to (required)")
	fmt.Println("  -i, --input\t\t\tThe input file to read from (required)")
	fmt.Println("  -t, --thresh\t\t\tThe minimum word count to include in the output")
	fmt.Println("  -c, --count\t\t\tThe maximum number of words to include in the output")
	fmt.Println("  -s, --simple\t\t\tOnly emit the words, not the counts")
	fmt.Println("  -g, --group\t\t\tGroup words into # word groups")
}

var wordRegex = regexp.MustCompile("(?i)[^-'0-9a-zÀ-ÿ`]")
var sepRegex = regexp.MustCompile("[\n\r\t—]")

type Word struct {
	name  string
	count int
}

func main() {
	fmt.Print("\033[H\033[2J")

	args := os.Args
	vars := make(map[string]string)

	if len(args) < 2 {
		fmt.Print("Missing any arguments\n\n")
		printHelp()
		return
	}

	simple := false
	group := 1
	var err error

	for i, arg := range args {
		if arg == "-h" || arg == "--help" {
			printHelp()
			return
		} else if arg == "-v" || arg == "--version" {
			fmt.Printf("vocab version %s\n", version)
			return
		} else if arg == "-o" || arg == "--output" {
			if i+1 < len(args) {
				vars["output"] = args[i+1]
			}
		} else if arg == "-i" || arg == "--input" {
			if i+1 < len(args) {
				vars["input"] = args[i+1]
			}
		} else if arg == "-t" || arg == "--thresh" {
			if i+1 < len(args) {
				vars["threshold"] = args[i+1]
			}
		} else if arg == "-c" || arg == "--count" {
			if i+1 < len(args) {
				vars["count"] = args[i+1]
			}
		} else if arg == "-s" || arg == "--simple" {
			simple = true
		} else if arg == "-g" || arg == "--group" {
			if i+1 < len(args) {
				group, err = strconv.Atoi(args[i+1])
				if err != nil {
					fmt.Print("Missing group number (e.g. 2)\n\n")
					printHelp()
					return
				}
			}
		}
	}

	if (vars["input"] == "") || (vars["output"] == "") {
		fmt.Print("Missing input or output file\n\n")
		printHelp()
		return
	} else if vars["threshold"] != "" && vars["count"] != "" {
		fmt.Print("Cannot use both threshold and count\n\n")
		printHelp()
		return
	} else if vars["input"] == vars["output"] {
		fmt.Print("Input and output files cannot be the same\n\n")
		printHelp()
		return
	}

	file, err := os.Open(vars["input"])
	if err != nil {
		fmt.Printf("Error opening input file: %s\n\n", vars["input"])
		printHelp()
		return
	}
	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Print("Error reading input file, is it in use elsewhere?\n\n")
		printHelp()
		return
	}

	fileString := sepRegex.ReplaceAllString(string(fileBytes), " ")
	fileWords := strings.Split(fileString, " ")
	fileWordsFormatted := make([]string, 0, len(fileWords))

	for _, word := range fileWords {
		word = wordRegex.ReplaceAllString(word, "")
		word = strings.ToLower(word)
		if word != "" {
			fileWordsFormatted = append(fileWordsFormatted, word)
		}
	}

	if group > 1 {
		for i := 0; i <= len(fileWordsFormatted)-group; i++ {
			fileWordsFormatted[i] = strings.Join(fileWordsFormatted[i:i+group], " ")
		}
		fileWordsFormatted = fileWordsFormatted[:len(fileWordsFormatted)-group+1]
	}

	wordCount := make(map[string]int, len(fileWordsFormatted))

	for _, word := range fileWordsFormatted {
		if _, ok := wordCount[word]; !ok {
			wordCount[word] = 1
		} else {
			wordCount[word]++
		}
	}

	words := make([]Word, 0, len(wordCount))

	for word, count := range wordCount {
		words = append(words, Word{word, count})
	}

	sort.Slice(words, func(i int, j int) bool {
		return words[i].count > words[j].count
	})

	file, err = os.Create(vars["output"])
	if err != nil {
		fmt.Printf("Error opening output file: %s\n\n", vars["output"])
		printHelp()
		return
	}
	defer file.Close()

	max := len(words)
	if vars["count"] != "" {
		max, err = strconv.Atoi(vars["count"])
		if err != nil {
			fmt.Printf("Error parsing count: %s\n\n", vars["count"])
			printHelp()
			return
		}
	}

	thresh := 0
	if vars["threshold"] != "" {
		thresh, err = strconv.Atoi(vars["threshold"])
		if err != nil {
			fmt.Printf("Error parsing threshold: %s\n\n", vars["threshold"])
			printHelp()
			return
		}
	}

	for i := 0; i < max; i++ {
		if words[i].count >= thresh {
			if simple {
				file.WriteString(words[i].name + "\n")
			} else {
				file.WriteString(fmt.Sprintf("#%-5d - %5d\t%s\n", i+1, words[i].count, words[i].name))
			}
		}
	}

	fmt.Println("Success!")
	fmt.Print("Open the file in your editor to see the results:\n\n")
	abs, _ := filepath.Abs(vars["output"])
	fmt.Print(abs + "\n\n")
}
