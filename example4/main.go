package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	raw, err := os.ReadFile("note.txt")
	if err != nil {
		log.Fatal(err)
	}

	bWord := bytes.Fields(raw) // bWord is a slice of []byte

	wordCounts := make(map[string]int)
	for _, word := range bWord {
		word := string(word)
		fmt.Println(word)
		wordCounts[word]++
	}
	fmt.Println(wordCounts)
	for word, count := range wordCounts {
		fmt.Printf("%s %d\n", word, count)
	}
}
