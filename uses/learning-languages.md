# Reading Foreign-Language Texts

My main motivation for making this tool was to get high-frequency words in a Spanish text I wanted to read. It is relatively easy to create a customized dictionary.

## Getting an .txt file

You will need an ebook format other than `.pdf` to create the `.txt` file. Use an online converter to convert the ebook to a text file.

## Getting word and phrase lists

Install the `vocab` tool as per the [README](https://pkg.go.dev/github.com/russellsteadman/vocab) and run the following commands in your `.txt` file's directory:

```sh
# Adjust the threshold to a reasonable number of recurrances for the text
# You can also remove the threshold entirely to get all words
vocab -i <filename>.txt -o single-word.txt --thresh 5 -s

# Adjust the word grouping as you see fit -- three is generally good
# Adjust the threshold as well -- it may be different than for single words
vocab -i <filename>.txt -o three-words.txt --thresh 5 -s --group 3
```

This will make a single-word frequency list and a three-word phrase frequency list.

## Translating the lists

Create a new Google Sheets workbook and paste the contents of `single-word.txt` and `three-words.txt` into separate sheets. Then, in the next column of each sheet, add `=GOOGLETRANSLATE(<FIRST CELL>, <FROM LANGUAGE>, <TO LANGUAGE>)`. For example, `=GOOGLETRANSLATE(A1, "es", "en")` would translate data in cell `A1` from Spanish (Español) to English.

## Learning the vocabulary

The Google Sheets information can then be added to Quizlet or other flashcard services.
