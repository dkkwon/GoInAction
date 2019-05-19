package main

import (
	"log"
	"os"

	_ "github.com/Jpub/GoInAction/chapter2/sample/matchers"
	"github.com/Jpub/GoInAction/chapter2/sample/search"
)

// init is called prior to main.
func init() {
	// Change the device for logging to stdout.
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
}

// main is the entry point for the program.
func main() {
	// Perform the search for the specified term.
	search.Run("president")
}
