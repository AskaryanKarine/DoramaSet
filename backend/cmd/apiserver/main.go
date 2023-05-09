package main

import (
	"DoramaSet/internal/handler/apiserver"
	"fmt"
	"os"
)

func main() {
	app, err := apiserver.Init()
	if err != nil {
		fmt.Printf("Initialisation application error: %w", err)
		os.Exit(1)
	}

	err = app.Run()
	if err != nil {
		fmt.Printf("Application running error: %s", err)
		os.Exit(1)
	}
}
