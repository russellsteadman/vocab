# Vocab Tool

Vocab is a simple CLI utility that takes a text file and finds the most frequent words.

## Motivation

I wanted an easy way to extract key vocabulary terms from a story in Spanish so I could learn the most common terms.

## Usage

If you have an `.mobi`, `.epub`, or other format, convert the file into a `.txt` text file. There are many online converters available to do so.

```sh
$ vocab --help
```

```txt
Usage: vocab 0.1.0
Vocab is a tool for generating a vocabulary file from a text file.

Options:
  -h, --help                    Show this help message and exit
  -v, --version                 Show the version number and exit
  -o, --output                  The output file to write to (required)
  -i, --input                   The input file to read from (required)
  -t, --thresh                  The minimum word count to include in the output
  -c, --count                   The maximum number of words to include in the output
  -s, --simple                  Only emit the words, not the counts
```

## Examples

Get all words in order of usage:

```sh
$ vocab -i book.txt -o book-vocab.txt
```

Get top 100 words:

```sh
$ vocab -i book.txt -o book-vocab.txt -c 100
```

Get words with 10 or more uses:

```sh
$ vocab -i charlie.txt -o charlie-vocab.txt -t 10
```

Get all words in order of count without additional markup:

```sh
$ vocab -i charlie.txt -o charlie-vocab.txt -s
```

## License

Open source. Released under an MIT License, see the `LICENSE` file for details.
