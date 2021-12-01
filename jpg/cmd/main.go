package main

import "github.com/arielkka/fallbox/jpg/internal/app"

func main() {
	err := app.Run()
	if err != nil {
		panic(err)
	}
}
