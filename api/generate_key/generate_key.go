package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
)

// This should run only once to generate a secure SECRET_KEY on the env file. Will generate a string with 64 random characters.

func main() {
	key := make([]byte, 64)

	if _, err := rand.Read(key); err != nil {
		log.Fatal(err)
	}

	stringBase64 := base64.StdEncoding.EncodeToString(key)

	// paste this on your .env file
	fmt.Println(stringBase64)
}
