package main

import "gihub.com/arielkka/fallbox/excel/internal/app"

func main() {
	err := app.Run()
	if err != nil {
		panic(err)
	}
}
