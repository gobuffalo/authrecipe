package main

import (
	"log"

	"github.com/gobuffalo/authrecipe/actions"
)

func main() {
	app := actions.App()
	log.Fatal(app.Serve())
}
