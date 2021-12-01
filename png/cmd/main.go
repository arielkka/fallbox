package main

import "gihub.com/arielkka/fallbox/png/internal/app"

func main() {
	err := app.Run()
	if err != nil {
		panic(err)
	}
}
