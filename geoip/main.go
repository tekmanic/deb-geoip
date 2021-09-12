package main

import (
	_ "embed"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"

	"github.com/tekmanic/deb-geoip/geoip/internal/handlers"
)

func main() {
	geoDir := os.Getenv("GEOIP_DIR")
	if geoDir == "" {
		log.Println("GEOIP_DIR environment variable not set")
		geoDir = "./"
	}

	// Create new fiber instance
	app := fiber.New(
		fiber.Config{
			// Pass view engine
			Views: html.New(geoDir+"internal/views", ".html"),
			// Pass global error handler
			ErrorHandler: handlers.Errors(geoDir + "public/500.html"),
		},
	)

	// Render index template with IP input value
	app.Get("/", handlers.Render())

	// Serve static assets
	app.Static("/", geoDir+"public", fiber.Static{
		Compress: true,
	})

	// Main GEO handler that is cached for 10 minutes
	app.Get("/geo/:ip?", handlers.Cache(10*time.Minute), handlers.GEO())

	// Maxmind Geo handler
	app.Get("/geomm/:ip?", handlers.GeoIP())

	// Handle 404 errors
	app.Use(handlers.NotFound(geoDir + "public/404.html"))

	// Listen on port :3000
	log.Fatal(app.Listen(":3000"))
}
