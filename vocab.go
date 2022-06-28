//go:build !js
// +build !js

package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"

	"github.com/russellsteadman/vocab/internal/vocab"
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

func main() {
	fmt.Print("\033[H\033[2J")

	args := os.Args

	if len(args) < 2 {
		fmt.Print("Missing any arguments\n\n")
		printHelp()
		return
	}

	vocabInput := &vocab.VocabInput{}
	inputFile := ""
	outputFile := ""

	for i, arg := range args {
		if arg == "-h" || arg == "--help" {
			printHelp()
			return
		} else if arg == "-v" || arg == "--version" {
			fmt.Printf("vocab version %s\n", version)
			return
		} else if arg == "-o" || arg == "--output" {
			if i+1 < len(args) {
				outputFile = args[i+1]
			}
		} else if arg == "-i" || arg == "--input" {
			if i+1 < len(args) {
				inputFile = args[i+1]
			}
		} else if arg == "-t" || arg == "--thresh" {
			if i+1 < len(args) {
				thresh, err := strconv.Atoi(args[i+1])
				if err != nil {
					fmt.Print("Missing threshold number (e.g. 2)\n\n")
					printHelp()
					return
				}

				vocabInput.Threshold = thresh
			}
		} else if arg == "-c" || arg == "--count" {
			if i+1 < len(args) {
				count, err := strconv.Atoi(args[i+1])
				if err != nil {
					fmt.Print("Missing word count number (e.g. 2)\n\n")
					printHelp()
					return
				}

				vocabInput.Count = count
			}
		} else if arg == "-s" || arg == "--simple" {
			vocabInput.Simple = true
		} else if arg == "-g" || arg == "--group" {
			if i+1 < len(args) {
				group, err := strconv.Atoi(args[i+1])
				if err != nil {
					fmt.Print("Missing group number (e.g. 2)\n\n")
					printHelp()
					return
				}

				vocabInput.Group = group
			}
		}
	}

	if (inputFile == "") || (outputFile == "") {
		fmt.Print("Missing input or output file\n\n")
		printHelp()
		return
	} else if inputFile == outputFile {
		fmt.Print("Input and output files cannot be the same\n\n")
		printHelp()
		return
	}

	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("Error opening input file: %s\n\n", inputFile)
		printHelp()
		return
	}

	fileBytes, err := io.ReadAll(file)
	file.Close()
	if err != nil {
		fmt.Print("Error reading input file, is it in use elsewhere?\n\n")
		printHelp()
		return
	}

	vocabInput.File = fileBytes

	vocabOutput, err := vocab.CreateSet(vocabInput)
	if err != nil {
		fmt.Print(err.Error() + "\n\n")
		printHelp()
		return
	}

	file, err = os.Create(outputFile)
	if err != nil {
		fmt.Printf("Error opening output file: %s\n\n", outputFile)
		printHelp()
		return
	}

	file.Write(vocabOutput)
	file.Close()

	fmt.Println("Success!")
	fmt.Print("Open the file in your editor to see the results:\n\n")
	abs, _ := filepath.Abs(outputFile)
	fmt.Print(abs + "\n\n")
}
