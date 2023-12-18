package main

import (
	"fmt"
	"os"

	initializer "github.com/ginniss2022/union/config"
	"github.com/ginniss2022/union/controllers"
	"github.com/ginniss2022/union/shutdown"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func main() {
	var exitCode int
	defer func() {
		os.Exit(exitCode)
	}()

	env, err := initializer.LoadConfig()
	if err != nil {
		fmt.Printf("error: %v", err)
		exitCode = 1
		return
	}

	cleanup, err := run(env)

	defer cleanup()
	if err != nil {
		fmt.Printf("error: %v", err)
		exitCode = 1
		return
	}

	shutdown.Gracefully()
}

func run(env initializer.EnvVars) (func(), error) {
	app, cleanup, err := buildServer(env)
	if err != nil {
		return nil, err
	}

	go func() {
		app.Listen("0.0.0.0:" + env.PORT)
	}()

	return func() {
		cleanup()
		app.Shutdown()
	}, nil
}

func buildServer(env initializer.EnvVars) (*fiber.App, func(), error) {
	initializer.ConnectToDatabase(
		env.PG_HOST,
		env.PG_PORT,
		env.PG_USER,
		env.PG_PASS,
		env.PG_DB,
	)

	app := fiber.New()

	app.Use(cors.New())
	app.Use(logger.New())

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("Healthy!")
	})

	app.Post("/signup", controllers.CreateNewUser)
	// add docs
	// app.Get("/swagger/*", swagger.HandlerDefault)

	return app, func() {
		// storage.CloseMongo(db)
	}, nil
}
