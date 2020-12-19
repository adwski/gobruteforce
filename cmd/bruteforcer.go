package main

import (
	"flag"
	"github.com/adwski/gobruteforce"
	"github.com/adwski/gobruteforce/generators"
	"github.com/adwski/gobruteforce/generators/combinations"
	"github.com/adwski/gobruteforce/generators/wordlist"
	"github.com/adwski/gobruteforce/tryers"
	"github.com/adwski/gobruteforce/tryers/privatekey"
	"log"
	"os"
)

// Supported modes
const (
	modePKPassPhrase string = "PKPassPhrase"
)

// Supported sources (generators)
const (
	sourceWordList     string = "WordList"
	sourceCombinations string = "Combinations"
)

func main() {
	var (
		mode       = flag.String("m", modePKPassPhrase, "'PKPassPhrase' only at this moment")
		source     = flag.String("src", sourceCombinations, "'sourceCombinations', 'sourceWordList'")
		wlist      = flag.String("wList", "wordlist.txt", "Path to wordlist file")
		filepath   = flag.String("f", "id_rsa", "Path to private key file")
		characters = flag.String("ch", "abcdefghijklmnopqrstuvwxyz", "characters used to gen suggestions")
		workers    = flag.Int("w", 10, "number of workers")
		start      = flag.Int("chMin", 3, "min length for suggestions")
		end        = flag.Int("chMax", 5, "max length for suggestions")
	)
	flag.Parse()

	logger := log.New(os.Stderr, "", log.LstdFlags)

	var (
		g    generators.Generator
		gCfg generators.GeneratorConfig
	)

	switch *source {
	case sourceCombinations:
		g = &combinations.AllPossible{Chars: *characters}
		gCfg.Length = *end
		gCfg.Count = *end - *start + 1

	case sourceWordList:
		g = &wordlist.WordList{FilePath: *wlist}

	default:
		log.Fatalf("uknown source: %v", *source)
	}

	switch *mode {
	case modePKPassPhrase:
		runner := gobruteforce.Runner{
			Generator: g,
			GCfg:      gCfg,
			Tryer:     &privatekey.PassPhrase{FilePath: *filepath},
			TCfg:      tryers.TryerConfig{Count: *workers},
			Log:       logger,
		}

		runner.Run()

	default:
		log.Fatalf("uknown mode: %v", *mode)
	}
}
