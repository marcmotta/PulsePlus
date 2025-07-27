// cmd/pulseplus/main.go
package main

import (
	"flag"
	"log"
	"os"

	"pulseplus/internal/pulseplus"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()

	app := pulseplus.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
