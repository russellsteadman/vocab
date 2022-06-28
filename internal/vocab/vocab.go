package vocab

import (
	"bytes"
	"errors"
	"fmt"
	"regexp"
	"sort"
)

var wordRegex = regexp.MustCompile("(?i)[^-'0-9a-zÀ-ÿ`]")
var sepRegex = regexp.MustCompile("[\n\r\t—]")

type Word struct {
	Name  []byte
	Count int
}

type VocabInput struct {
	File      []byte `json:"file"`
	Threshold int    `json:"threshold"`
	Count     int    `json:"count"`
	Simple    bool   `json:"simple"`
	Group     int    `json:"group"`
}

func CreateSet(input *VocabInput) ([]byte, error) {
	if input.File == nil || len(input.File) == 0 {
		return nil, errors.New("missing or empty input file")
	} else if input.Threshold > 0 && input.Count > 0 {
		return nil, errors.New("cannot use both threshold and count")
	}

	fileReformatted := sepRegex.ReplaceAll(input.File, []byte(" "))
	fileWords := bytes.Split(fileReformatted, []byte(" "))
	fileWordsFormatted := make([][]byte, 0, len(fileWords))

	for _, word := range fileWords {
		word = wordRegex.ReplaceAll(word, []byte{})
		word = bytes.ToLower(word)
		if len(word) > 0 {
			fileWordsFormatted = append(fileWordsFormatted, word)
		}
	}

	if input.Group > 1 {
		for i := 0; i <= len(fileWordsFormatted)-input.Group; i++ {
			fileWordsFormatted[i] = bytes.Join(fileWordsFormatted[i:i+input.Group], []byte(" "))
		}
		fileWordsFormatted = fileWordsFormatted[:len(fileWordsFormatted)-input.Group+1]
	}

	wordCount := make(map[string]int, len(fileWordsFormatted))

	for _, word := range fileWordsFormatted {
		if _, ok := wordCount[string(word)]; !ok {
			wordCount[string(word)] = 1
		} else {
			wordCount[string(word)]++
		}
	}

	words := make([]Word, 0, len(wordCount))

	for word, count := range wordCount {
		words = append(words, Word{[]byte(word), count})
	}

	sort.Slice(words, func(i int, j int) bool {
		return words[i].Count > words[j].Count
	})

	max := len(words)
	if input.Count > 0 {
		max = input.Count
	}

	thresh := 0
	if input.Threshold > 0 {
		thresh = input.Threshold
	}

	output := []byte{}

	for i := 0; i < max; i++ {
		if words[i].Count >= thresh {
			if input.Simple {
				output = append(output, words[i].Name...)
				output = append(output, []byte("\n")...)
			} else {
				output = append(output, []byte(fmt.Sprintf("#%-5d - %5d\t%s\n", i+1, words[i].Count, words[i].Name))...)
			}
		}
	}

	return output, nil
}
