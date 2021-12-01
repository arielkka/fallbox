package main

import "github.com/arielkka/fallbox/handler/internal/app"

func main() {
	err := app.Run()
	if err != nil {
		panic(err)
	}
}
