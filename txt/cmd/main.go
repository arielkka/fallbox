package main

import "gihub.com/arielkka/fallbox/txt/internal/app"

func main() {
	err := app.Run()
	if err != nil {
		panic(err)
	}
}
