package main

import (
	"log"
	"os"

	"search"
	_ "matchers"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("Johnson")
}
