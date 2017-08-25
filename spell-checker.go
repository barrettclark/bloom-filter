package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type SpellChecker struct {
	bloomFilter BloomFilter
}

func (checker SpellChecker) Load(file *os.File) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		checker.bloomFilter.Add(strings.ToLower(scanner.Text()))
	}
}

func (checker SpellChecker) CheckWord(str string) bool {
	return checker.bloomFilter.Contains(strings.ToLower(str))
}

func (checker SpellChecker) CheckDocument(file *os.File) {
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	var candidate string
	for scanner.Scan() {
		candidate = scanner.Text()
		fmt.Printf("Checking %s\n", candidate)
		if checker.CheckWord(candidate) == false {
			fmt.Println(candidate)
			return
		}
	}
}

func main() {
	h1 := HashFunction(HashSum)
	h2 := HashFunction(HashProduct)
	h3 := HashFunction(HashHash)

	bf := BloomFilter{
		HashFunctions: []HashFunction{h1, h2, h3},
		ByteArray:     make([]byte, 1000000),
	}

	checker := SpellChecker{
		bloomFilter: bf,
	}

	file, _ := os.Open("/usr/share/dict/words")
	checker.Load(file)
	fmt.Println(checker.CheckWord("HELLO"))

	var filename string
	if len(os.Args) == 1 {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter a filename to check: ")
		text, _ := reader.ReadString('\n')
		filename = strings.TrimSpace(text)
	} else {
		filename = os.Args[1]
	}
	document, _ := os.Open(filename)
	checker.CheckDocument(document)
}
