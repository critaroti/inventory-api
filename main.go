package main

import (
	"log"
	"net/http"
	"os"

	"inventory.api/config"
	"inventory.api/handler"
	"inventory.api/repository"
	"inventory.api/routes"
	"inventory.api/service"
)

func main() {
	config.ConnectDB()

	repo := repository.NewInventoryPostgres(config.DB)
	svc := service.NewInventoryService(repo)
	h := handler.NewInventoryHandler(svc)

	routes.RegisterRoutes(h)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server running on", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
