package main

import (
	"log"
	"os"

	"github.com/Kamva/mgm/v2"
	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	"github.com/hl-service/go-api/controllers"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	godotenv.Load()
	connectionString := os.Getenv("MONGODB_CONNECTION_STRING")

	if len(connectionString) == 0 {
		connectionString = "mongodb://mongodb:27017"
	}

	err := mgm.SetDefaultConfig(nil, "local", options.Client().ApplyURI(connectionString))

	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{os.Getenv("CORS_ALLOWED_ORIGIN")},
	}))

	app.Get("/api/articles", controllers.IndexArticles)
	app.Get("/api/articles/:id", controllers.ShowArticle)
	app.Post("/api/articles", controllers.StoreArticle)
	app.Put("/api/articles/:id", controllers.UpdateArticle)
	app.Delete("/api/articles/:id", controllers.DeleteArticle)

	app.Listen(8080)
}
