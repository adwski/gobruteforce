### Description
This is attempt to remember passphase of my private key. 
### Examples
```bash
# gobruteforce ssh private key passkey with all possible combinations
go run cmd/bruteforcer.go -m PKPassPhrase -f id_rsa -src Combinations -ch abcdefghijklmnopqrstuvwxyz -chMin 3 -chMax 5 -w 10

# gobruteforce ssh private key passkey with word list
go run cmd/bruteforcer.go -m PKPassPhrase -f id_rsa -src WordList -wList wordlist.txt -w 10
```
