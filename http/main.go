package main

import (
	"log"

	"github.com/sweetie-pie/line-recommendation/internal/port/http"
	"github.com/sweetie-pie/line-recommendation/internal/port/mysql"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	db, err := mysql.New()
	if err != nil {
		panic(err)
	}

	h := http.Handler{Repository: db}

	app.Get("/routes", h.Routes)
	app.Get("/searches", h.Searches)

	log.Println("server start on")

	if err := app.Listen(":8080"); err != nil {
		panic(err)
	}
}
