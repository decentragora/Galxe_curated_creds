package main

import (
	"dagora_galaxe_querys/clients"
	"dagora_galaxe_querys/routes"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	_ = godotenv.Load()
	r := mux.NewRouter()

	clts, err := clients.InitClients()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	routes.HandleRoutes(r, clts)

	handler := cors.New(cors.Options{
		AllowedOrigins: []string{"https://galxe.com"}, // Adjust this to your needs
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type", "Origin", "Accept", "*"},
	}).Handler(r)

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalf("PORT must be set")
		os.Exit(1)
	}
	log.Println("Server running on port " + port)
	err = http.ListenAndServe(":"+port, handler)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
