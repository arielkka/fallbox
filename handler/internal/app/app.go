package app

import (
	"fmt"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load("../../../.env"); err != nil {
		fmt.Println("No .env file found")
	}
}

func Run() error {

	return nil
}
