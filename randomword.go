// Public domain

// randomword command prints random words from a list of 4096 words
// specifically selected to be easy to memorize when used in passphrases.
package main

import (
	"crypto/rand"
	"encoding/binary"
	"flag"
	"fmt"
	"io"
	"strings"
)

func init() {
	// Safety checks.
	if len(words)&(len(words)-1) != 0 {
		// Number of words must be a power of two to simplify unbiased
		// generation of random words.
		panic("number of words is not a power of two")
	}
	if len(words) > 1<<31-1 {
		panic("too many words")
	}
}

func getRandomWord() string {
	var b [4]byte
	if _, err := io.ReadFull(rand.Reader, b[:]); err != nil {
		panic("failed to read random bytes: " + err.Error())
	}
	return words[binary.LittleEndian.Uint32(b[:])%uint32(len(words))]
}

var (
	fNumberOfWords = flag.Int("c", 1, "number of words to generate")
	fNoNewLine     = flag.Bool("n", false, "do not print new line at the end")
	fLines         = flag.Bool("l", false, "print each word on its own line")
)

func main() {
	flag.Parse()
	w := make([]string, *fNumberOfWords)
	for i := range w {
		w[i] = getRandomWord()
	}
	sep := " "
	if *fLines {
		sep = "\n"
	}
	fmt.Print(strings.Join(w, sep))
	if !*fNoNewLine {
		fmt.Println("")
	}
}
