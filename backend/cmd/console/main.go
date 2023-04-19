package main

import (
	"DoramaSet/internal/handler/console"
	"fmt"
	"os"
)

func main() {
	dsn := "host=localhost user=karine password=12346 dbname=DoramaSet sslmode=disable"
	secretKey := "qwerty"
	app, err := console.NewApp(dsn, secretKey)
	if err != nil {
		fmt.Printf("Ошибка инициализации приложения: %s\n", err)
		os.Exit(1)
	}

	app.Run()
}
