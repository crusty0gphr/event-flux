package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"

	eventflux "github.com/event-flux"
	"github.com/event-flux/db"
	"github.com/event-flux/internal"
	"github.com/event-flux/migrate"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	configs := eventflux.LoadConfigs()

	cassandraSession, err := db.NewCassandraSession(configs.CassandraHost)
	if err != nil {
		log.Fatal(err)
		return
	}

	if err = migrate.CassandraMigrateUP(cassandraSession); err != nil {
		log.Fatal(err)
		return
	}

	srv := internal.NewService(nil)
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
