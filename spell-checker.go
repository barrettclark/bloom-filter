package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
    "math/rand"
    "time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyz")

func randSeq(n int) string {
    b := make([]rune, n)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}

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
    rand.Seed(time.Now().UnixNano())

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

    dict := make(map[string]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		dict[strings.ToLower(scanner.Text())] = 1
	}

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

    attempts := 0
    falsePositives := 0
    for i := 1; i <= 10000; i++ {
        candidate := randSeq(5)
        if checker.CheckWord(candidate) == true && dict[candidate] == 0 {
            falsePositives += 1
        }
        attempts += 1
    }
    fmt.Printf("Attempts: %d\nFalse positives: %d\n", attempts, falsePositives)
}
