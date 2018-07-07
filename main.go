package main

import (
	"log"

	"./refreshToken"
)

func main() {
	err := refreshToken.RefreshToken()
	log.Fatal(err)
}
