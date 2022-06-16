package main

import (
	"bufio"
	"log"
	"os"

	"github.com/pborman/getopt"
	"github.com/uladzislaudd/wordle/pkg/dictionary"
)

var (
	words  = "words.txt"
	output = "output.json"
)

func init() {
	getopt.StringVarLong(&words, "words", 'w', "filepath to json with words, one line for each word", "/path/to/words.txt")
	getopt.StringVarLong(&output, "output", 'o', "filepath to save", "/path/to/output.json")
	if err := getopt.Getopt(nil); err != nil {
		panic(err)
	}
}

func main() {
	d := dictionary.Dictionary{}

	file, err := os.Open(words)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	i := 1
	rs := map[rune]any{}
	for scanner.Scan() {
		word := scanner.Text()
		d.Words = append(d.Words, word)
		if d.WordSize == 0 {
			d.WordSize = uint8(len(word))
			continue
		}

		if uint8(len(word)) != d.WordSize {
			log.Panicf("word \"%s\" at line %d have invalid WordSize (%d != %d)", word, i, len(word), d.WordSize)
		}

		for _, r := range word {
			rs[r] = 1
		}

		i++
	}

	for r := range rs {
		d.Alphabet = append(d.Alphabet, r)
	}

	if err = d.SaveToJSON(output); err != nil {
		log.Panic(err)
	}
}
