package generators

import "log"

// GeneratorConfig provides additional params for Generator
type GeneratorConfig struct {
	Count  int
	Length int
	ID     int
	Log    *log.Logger
}

// Generator produces strings to try as passwords
type Generator interface {
	Gen(GeneratorConfig, chan<- string, chan<- struct{})
}
