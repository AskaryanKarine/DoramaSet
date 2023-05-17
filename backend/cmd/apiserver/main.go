package main

import (
	"DoramaSet/internal/handler/apiserver"
	"fmt"
	"os"
)

func main() {
	app, err := apiserver.Init()
	if err != nil {
		fmt.Printf("Initialisation application error: %s", err)
		os.Exit(1)
	}

	app.Run()
}
