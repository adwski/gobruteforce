package privatekey

import (
	"github.com/adwski/gobruteforce/tryers"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
)

// PassPhrase tries to guess passphare of ssh private key
type PassPhrase struct {
	FilePath string
}

// Try tries to guess passphare of ssh private key
func (pp PassPhrase) Try(cfg tryers.TryerConfig, suggestions <-chan string, done chan<- struct{}, success chan<- string) {
	key, err := ioutil.ReadFile(pp.FilePath)
	if err != nil {
		cfg.Log.Fatalf("Unable to read private key: %v", err)
	}

	for sgt := range suggestions {
		if _, err = ssh.ParsePrivateKeyWithPassphrase(key, []byte(sgt)); err == nil {
			success <- sgt
			break
		}
	}
	done <- struct{}{}
}
