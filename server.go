package main

import (
	"crypto/tls"
	"log"
	"os"
	"strconv"

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

	app.Static("/", "./public")

	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{os.Getenv("CORS_ALLOWED_ORIGIN")},
	}))

	app.Get("/api/articles", controllers.IndexArticles)
	app.Get("/api/articles/:id", controllers.ShowArticle)
	app.Post("/api/articles", controllers.StoreArticle)
	app.Put("/api/articles/:id", controllers.UpdateArticle)
	app.Delete("/api/articles/:id", controllers.DeleteArticle)

	tlsEnabled, _ := strconv.ParseBool(os.Getenv("TLS_ENABLED"))

	if tlsEnabled == true {
		domain := os.Getenv("TLS_DOMAIN")

		cer, err := tls.LoadX509KeyPair(
			"/etc/letsencrypt/archive/"+domain+"/fullchain1.pem",
			"/etc/letsencrypt/archive/"+domain+"/privkey1.pem",
		)

		if err != nil {
			log.Fatal(err)
		}

		config := &tls.Config{Certificates: []tls.Certificate{cer}}

		app.Listen(443, config)
	} else {
		app.Listen(8080)
	}
}
