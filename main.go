package main

import (
	"fmt"
	"log"

	"github.com/alexedwards/argon2id"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	password := "helloworld"
	hash, err := argon2id.CreateHash(password, argon2id.DefaultParams)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Hashed password: ", hash)

	passwordToCheck := []string{"helloworld", "helloworld2", "helloworld3"}
	for _, password := range passwordToCheck {
		match, err := argon2id.ComparePasswordAndHash(password, hash)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("\"%s\":\t %v\n", password, match)
	}
}
