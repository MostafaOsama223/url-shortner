package main

import (
	"context"
	"log"
	"os"

	"github.com/MostafaOsama223/shortner-service/api"
	"github.com/MostafaOsama223/shortner-service/database"
	"github.com/MostafaOsama223/shortner-service/repo"
	"github.com/MostafaOsama223/shortner-service/service"

	"github.com/joho/godotenv"
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	loadEnv()
	uri := os.Getenv("MONGO_URI")
	log.Printf("Starting URL shortener service with MONGO_URI: %s", uri)

	mongodbClient, err := database.NewMongoDB(uri)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer mongodbClient.Disconnect(context.Background())

	db := mongodbClient.Database("url_shortener")
	log.Printf("Using database: %s", db.Name())

	urlsCollection := mongodbClient.Collection(db, "urls")
	log.Printf("Using collection: %s", "urls")

	urlRepo := repo.NewUrlRepo(urlsCollection)
	urlService := service.NewUrlService(urlRepo)

	log.Printf("Starting HTTP server...")
	urlAPIHandler := api.NewUrlAPIHandler(urlService)
	urlAPIHandler.Run()
}
