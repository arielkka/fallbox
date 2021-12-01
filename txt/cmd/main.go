package main

import "github.com/arielkka/fallbox/txt/internal/app"

func main() {
	err := app.Run()
	if err != nil {
		panic(err)
	}
}
