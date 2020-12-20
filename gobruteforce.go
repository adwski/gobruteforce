package gobruteforce

import (
	"github.com/adwski/gobruteforce/generators"
	"github.com/adwski/gobruteforce/tryers"
	"log"
	"os"
	"time"
)

// Runner runs generator and a number of triers
type Runner struct {
	Generator generators.Generator
	Tryer     tryers.Tryer
	GCfg      generators.GeneratorConfig
	TCfg      tryers.TryerConfig
	Log       *log.Logger
}

func (r *Runner) Run() {

	if r.Log == nil {
		r.Log = log.New(os.Stderr, "", log.LstdFlags)
	}

	tStart := time.Now()
	log.Print("Runner started")

	suggestions := make(chan string)
	doneG := make(chan struct{})
	success := make(chan string)
	doneT := make(chan struct{})

	// Start all generators
	for i := 0; i <= r.GCfg.Count; i++ {
		gCfg := r.GCfg
		gCfg.Log = r.Log
		gCfg.ID = i
		go r.Generator.Gen(gCfg, suggestions, doneG)
	}

	// Start all tryers
	for i := 0; i <= r.TCfg.Count; i++ {
		tCfg := r.TCfg
		tCfg.Log = r.Log
		tCfg.ID = i
		go r.Tryer.Try(tCfg, suggestions, doneT, success)
	}

	wFinished := 0
	sFinished := 0

	// Wait
loop:
	for {
		select {
		case <-doneG:
			// some tryer finished
			sFinished++
			r.Log.Printf("generators finished: %d", sFinished)
			if sFinished >= r.GCfg.Count {
				log.Print("all generators finished")
				// Sleep for the case when last suggestion is successful
				time.Sleep(time.Second)
				break loop
			}
		case sccss := <-success:
			// got successful try
			r.Log.Printf("=====================================\n"+
				"success result: %s\n"+
				"=====================================", sccss)
			break loop
		case <-doneT:
			// some tryer finished
			wFinished++
			r.Log.Printf("tryers finished: %d", wFinished)
			if wFinished >= r.TCfg.Count {
				// should never happen actually
				log.Print("all tryers finished")
				break loop
			}
		}
	}

	tEnd := time.Now()
	log.Printf("Runner finished in %v", tEnd.Sub(tStart))
	return
}
