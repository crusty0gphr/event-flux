package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"

	eventflux "github.com/event-flux"
	"github.com/event-flux/internal"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	configs := eventflux.LoadConfigs()

	repo, err := internal.BuildRepository(configs)
	if err != nil {
		log.Fatal(err)
		return
	}
	srv := internal.NewService(repo)
	handler := internal.NewHandler(srv)

	app.Get("/", handler.NotAllowed)
	app.Get("/events/filter", handler.GetByFilters)
	app.Get("/events", handler.GetAll)
	app.Get("/events/:event_id", handler.GetByID)

	log.Fatal(
		app.Listen(
			configs.BuildAppHostUrl(),
		),
	)
}
