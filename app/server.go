package main

import (
	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	println("Server.go")
}

func loadEnv() {
	err := godotenv.Load("../.env")

	if err != nil {
		fmt.Printf("Load Failed ENV file %v", err)
	}
}
