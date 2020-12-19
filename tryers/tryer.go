package tryers

import "log"

// TryerConfig provides additional params for Generator
type TryerConfig struct {
	Count int
	ID    int
	Log   *log.Logger
}

// Tryer takes suggested strings and tries them as a password
type Tryer interface {
	Try(TryerConfig, <-chan string, chan<- struct{}, chan<- string)
}
