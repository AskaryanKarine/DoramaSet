package main

import (
	"DoramaSet/internal/handler/console"
	"fmt"
	"os"
)

func main() {
	app, err := console.NewApp()
	if err != nil {
		fmt.Printf("Ошибка инициализации приложения: %s\n", err)
		os.Exit(1)
	}

	app.Run()
}
