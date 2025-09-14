package main

import (
	"fmt"
	"regexp"
	"strings"
	"sync"
	"unicode"
)

type Step1Result struct {
	Sentence string
	Words    int
	Chars    int
}

type FinalResult struct {
	Sentence   string
	Words      int
	Chars      int
	Vowels     int
	Consonants int
}

func SplitIntoSentences(text *[]string) []string {
	re := regexp.MustCompile(`[.!?]`)
	var sentences []string
	for _, sent := range *text {
		parts := re.Split(sent, -1)

		for _, s := range parts {
			s = strings.TrimSpace(s)
			if s != "" {
				sentences = append(sentences, s)
			}
		}
	}
	return sentences
}

func Worker1(jobs chan string, result1 chan Step1Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for sent := range jobs {
		words := strings.Fields(sent)
		wordsCount := len(words)
		chars := []rune(sent)
		charCount := len(chars)
		result1 <- Step1Result{
			Sentence: sent,
			Words:    wordsCount,
			Chars:    charCount,
		}
	}
}

func Worker2(result1 chan Step1Result, result2 chan FinalResult, gl []rune, wg *sync.WaitGroup) {
	defer wg.Done()
	for r := range result1 {
		vowelsCount := 0
		totalLetter := 0
		for _, symbol := range strings.ToLower(r.Sentence) {
			for _, gls := range gl {
				if gls == symbol {
					vowelsCount++
				}
			}
			if unicode.IsLetter(symbol) {
				totalLetter++
			}
		}
		con := totalLetter - vowelsCount
		result2 <- FinalResult{
			Sentence:   r.Sentence,
			Words:      r.Words,
			Chars:      r.Chars,
			Vowels:     vowelsCount,
			Consonants: con,
		}
	}
}

func main() {
	sentences := []string{
		"Golang is fun and powerful.",
		"Concurrency makes programs efficient!",
		"Fan-in and fan-out are useful patterns?",
		"Channels provide safe communication between goroutines.",
		"WaitGroups help synchronize concurrent tasks.",
	}
	jobs := make(chan string, len(sentences))
	result1 := make(chan Step1Result, len(sentences))
	result2 := make(chan FinalResult, len(sentences))
	gl := []rune{'e', 'u', 'o', 'a', 'i'}
	var wg sync.WaitGroup

	sentences = SplitIntoSentences(&sentences)

	for _, sent := range sentences {
		jobs <- sent
	}
	close(jobs)

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go Worker1(jobs, result1, &wg)
	}

	go func() {
		wg.Wait()
		close(result1)
	}()

	wg1 := sync.WaitGroup{}
	for i := 0; i < 3; i++ {
		wg1.Add(1)
		go Worker2(result1, result2, gl, &wg1)
	}

	go func() {
		wg1.Wait()
		close(result2)
	}()

	for r := range result2 {
		fmt.Printf("Sentence: %s | Words: %d | Chars %d | Vowel: %d | Consonants: %d\n", r.Sentence, r.Words, r.Chars, r.Vowels, r.Consonants)
	}
}
