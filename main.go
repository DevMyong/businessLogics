package main

import (
	"businessLogics/route"
	"businessLogics/services/price"
)

func main() {
	price.Begin()

	e := route.Router()
	e.Logger.Fatal(e.Start(":8080"))
}
