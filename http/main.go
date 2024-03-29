package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/sweetie-pie/line-recommendation/internal/port/http"
	"github.com/sweetie-pie/line-recommendation/internal/port/mysql"

	"github.com/gofiber/fiber/v2"
)

func main() {
	var (
		httpPort = flag.Int("port", 8080, "http port on localhost")
	)

	flag.Parse()

	// creating a new fiber app
	app := fiber.New()

	// initializing database
	db, err := mysql.New()
	if err != nil {
		panic(err)
	}

	// creating a handler
	h := http.Handler{Repository: db}

	app.Put("/node", h.CreateNode)
	app.Put("/route", h.CreateRoute)
	app.Put("/search", h.Search)
	app.Get("/data", h.Data)

	log.Printf("server start on %d ...\n", *httpPort)

	// starting http server
	if err := app.Listen(fmt.Sprintf(":%d", *httpPort)); err != nil {
		panic(err)
	}
}
