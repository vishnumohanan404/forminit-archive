package main

import (
	"log"

	"github.com/vishnumohanan404/forminit/internal/app"
	"github.com/vishnumohanan404/forminit/internal/database"
)

func main() {
	err := db.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	server := app.NewAPIServer(":8080")
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
