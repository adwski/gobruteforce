package wordlist

import (
	"bufio"
	"github.com/adwski/gobruteforce/generators"
	"os"
)

// WordList is a generator from word list file
type WordList struct {
	FilePath string
}

// Gen takes suggestions from word list file
// each line is considered a suggestion
func (wl WordList) Gen(cfg generators.GeneratorConfig, suggestions chan<- string, done chan<- struct{}) {
	file, err := os.Open(wl.FilePath)
	if err != nil {
		cfg.Log.Fatal(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		suggestions <- scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		cfg.Log.Fatal(err)
	}
}
