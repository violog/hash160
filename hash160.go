package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/ripemd160"
)

func main() {
	input := os.Args[1]
	hash160 := sha256.New()
	decoded, e := hex.DecodeString(input)
	if e != nil {
		log.Fatalln(e)
	}
	hash160.Write(decoded)
	sum := hash160.Sum(nil)
	fmt.Printf("SHA-256: %x\n", sum)
	hash160 = ripemd160.New()
	decoded, e = hex.DecodeString(fmt.Sprintf("%x", sum))
	if e != nil {
		log.Fatalln(e)
	}
	hash160.Write(decoded)
	sum = hash160.Sum(nil)
	fmt.Printf("HASH160: %x\n", sum)
}
