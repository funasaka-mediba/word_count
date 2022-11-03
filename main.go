package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"unicode"
)

func main() {
	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("cannot file open: %v, %s", filename, err)
	}
	defer f.Close()

	words, err := wordCount(f)
	if err != nil {
		log.Fatalf("read failed: %s", err)
	}

	fmt.Printf("%q: %d words\n", filename, words)
}

// wordCount 引数をio.ReaderにすることでファイルでもネットでもHTTPレスポンスボディでも色々扱えるようになるぞ。
func wordCount(r io.Reader) (int, error) {
	words, inword := 0, false
	b := bufio.NewReader(r)

	for {
		rune, _, err := b.ReadRune()
		if unicode.IsSpace(rune) {
			if inword {
				words++
			}
			inword = false
		} else {
			inword = true
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return 0, err
		}
	}
	return words, nil
}
