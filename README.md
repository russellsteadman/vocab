# Vocab Tool

Vocab is a simple CLI utility that takes a text file and finds the most frequent words.

## Motivation

I wanted an easy way to extract key vocabulary terms from a story in Spanish so I could learn the most common terms.

## Online Version

If you want to try out this tool or use it without installation, try the [experimental web version](https://vocab.docs.russellsteadman.com/experiments/web/). You can use all of the CLI features in your browser without any information leaving your device.

## Installation

### Using go

This requires go `1.18+`. Replace `<version>` with the latest release.

```sh
go install github.com/russellsteadman/vocab@<version>
```

If the `vocab` term is not available in your terminal, make sure your `$GO_PATH/bin` is in your `$PATH`.

### Via releases

Download the executable for your operating system and architecture from releases. You can then move the binary to a location in the path, or just use it locally.

```sh
# For MacOS (darwin)/Linux platforms
chmod +x ./vocab
./vocab --help
```

```bat
:: For Windows platforms
.\vocab --help
```

Note that these binaries are not signed and will raise "unidentified developer" errors. Code signing may be added in the future with enough usage.

## Usage (Command-line)

If you have an `.mobi`, `.epub`, or other format, convert the file into a `.txt` text file. There are many online converters available to do so.

```sh
vocab --help
```

```txt
Usage: vocab v0.2.0
Vocab is a tool for generating a vocabulary file from a text file.

Options:
  -h, --help                    Show this help message and exit
  -v, --version                 Show the version number and exit
  -o, --output                  The output file to write to (required)
  -i, --input                   The input file to read from (required)
  -t, --thresh                  The minimum word count to include in the output
  -c, --count                   The maximum number of words to include in the output
  -s, --simple                  Only emit the words, not the counts
  -g, --group                   Group words into # word groups
```

## Examples

Get all words in order of usage:

```sh
vocab -i book.txt -o book-vocab.txt
```

Get top 100 words:

```sh
vocab -i book.txt -o book-vocab.txt -c 100
```

Get words with 10 or more uses:

```sh
vocab -i charlie.txt -o charlie-vocab.txt -t 10
```

Get all words in order of count without additional markup:

```sh
vocab -i charlie.txt -o charlie-vocab.txt -s
```

Get groups of 3 words sorted by usage:

```sh
vocab -i charlie.txt -o charlie-vocab.txt -g 3
```

## License

Open source. Released under an MIT License, see the `LICENSE` file for details.
