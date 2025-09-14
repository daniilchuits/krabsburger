package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
	"unicode"
)

type Res struct {
	Sentence     string
	LengtOfWords int
}

type Result struct {
	Sentence   string
	Words      int
	Chars      int
	Vowels     int
	Consonants int
	WorkerID   int
}

func Producer(jobs chan Res, sentences []string, pwg *sync.WaitGroup, id int, numProducers int) {
	defer pwg.Done()
	for j := id; j < len(sentences); j += numProducers {
		sent := strings.Trim(sentences[j], ".")
		words := len(strings.Fields(sent))
		jobs <- Res{
			Sentence:     sent,
			LengtOfWords: words,
		}
	}
}

func Consumer(jobs chan Res, results chan Result, cwg *sync.WaitGroup, glas []rune, id int) {
	defer cwg.Done()
	for res := range jobs {
		runes := []rune(res.Sentence)
		vowels := 0
		Letters := 0

		for _, r := range runes {
			if unicode.IsLetter(r) {
				Letters++
				for _, gl := range glas {
					if r == gl {
						vowels++
					}
				}
			}
		}
		con := Letters - vowels

		results <- Result{
			Sentence:   res.Sentence,
			Words:      res.LengtOfWords,
			Chars:      len(runes),
			Vowels:     vowels,
			Consonants: con,
			WorkerID:   id,
		}
	}
}

func main() {
	var sentences = []string{
		"Go is a compiled language.",
		"Concurrency is not parallelism.",
		"Worker pools help manage goroutines efficiently.",
		"Channels provide a way for goroutines to communicate.",
		"Fan in and fan out are common concurrency patterns.",
		"Synchronization is important in concurrent programming.",
		"Mutex is used to protect shared resources.",
		"Buffered channels can reduce blocking between goroutines.",
		"Select statement allows waiting on multiple channels.",
		"Deadlocks happen when goroutines wait forever.",
		"Starvation occurs if some goroutines never get CPU time.",
		"Goroutines are lightweight threads managed by Go runtime.",
		"Parallelism depends on multiple CPU cores.",
		"Context helps with cancellation and timeouts.",
		"Race conditions are tricky to debug.",
		"Immutability can simplify concurrent programs.",
		"Go scheduler decides which goroutine runs next.",
		"Pipelines can process streams of data step by step.",
		"Error handling is an important part of concurrency.",
		"Graceful shutdown ensures resources are cleaned up.",
	}
	glas := []rune{'e', 'u', 'i', 'o', 'a'}
	jobs := make(chan Res, len(sentences)*10)
	results := make(chan Result, len(sentences)*10)
	var cwg sync.WaitGroup
	var pwg sync.WaitGroup
	rand.Seed(time.Now().UnixNano())

	var numProducers = rand.Intn(2) + 2
	var numConsumers = 3

	for i := 1; i <= numConsumers; i++ {
		cwg.Add(1)
		go Consumer(jobs, results, &cwg, glas, i)
	}

	for i := 0; i < numProducers; i++ {
		pwg.Add(1)
		go Producer(jobs, sentences, &pwg, i, numProducers)
	}

	go func() {
		pwg.Wait()
		close(jobs)
	}()

	go func() {
		cwg.Wait()
		close(results)
	}()

	for r := range results {
		fmt.Printf("Sentence: %-40s | Words: %d | Chars: %d | Vowels: %d | Consonants: %d | WorkerID: %d\n",
			r.Sentence, r.Words, r.Chars, r.Vowels, r.Consonants, r.WorkerID)
	}

}
