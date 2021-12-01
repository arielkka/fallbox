package app

import (
	"fmt"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load("./jgo.env"); err != nil {
		fmt.Println("No .env file found", err)
	}
}

func Run() error {
	return nil
}
