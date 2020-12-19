package combinations

import "github.com/adwski/gobruteforce/generators"

// AllPossible is a generator of all possible combinations of a 'characters' string
type AllPossible struct {
	Chars string
}

// Gen generates all possible combinations of a 'characters' string
// id is taken as len
func (ap *AllPossible) Gen(cfg generators.GeneratorConfig, suggestions chan<- string, done chan<- struct{}) {
	np := nextPassword(cfg.Length-cfg.ID, ap.Chars)
	for {
		pwd := np()
		if len(pwd) == 0 {
			break
		}
		suggestions <- pwd
	}
	done <- struct{}{}
}

// https://stackoverflow.com/a/22741715/9758991
func nextPassword(n int, c string) func() string {
	r := []rune(c)
	p := make([]rune, n)
	x := make([]int, len(p))
	return func() string {
		p := p[:len(x)]
		for i, xi := range x {
			p[i] = r[xi]
		}
		for i := len(x) - 1; i >= 0; i-- {
			x[i]++
			if x[i] < len(r) {
				break
			}
			x[i] = 0
			if i <= 0 {
				x = x[0:0]
				break
			}
		}
		return string(p)
	}
}
