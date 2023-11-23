package main

import (
	"os"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"github.com/ginniss2022/union/initializers"
	"github.com/ginniss2022/union/shutdown"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Repository struct {
	DB *gorm.DB
}

func main() {
	//In case of shutdown
	var exitCode int
	defer func(){
		os.Exit(exitCode)
	}()

	//Load config
	env, err := initializers.LoadConfig()
	if err != nil {
		fmt.Printf("error: %v", err)
		exitCode = 1
		return
	}
	// run the server
	cleanup, err := run(env)

	// // run the cleanup after the server is terminated
	defer cleanup()
	if err != nil {
		fmt.Printf("error: %v", err)
		exitCode = 1
		return
	}

	// ensure the server is shutdown gracefully & app runs
	shutdown.Gracefully()
}

func run(env config.EnvVars) (func(), error) {
	app, cleanup, err := buildServer(env)
	if err != nil {
		return nil, err
	}

	// start the server
	go func() {
		app.Listen("0.0.0.0:" + env.PORT)
	}()

	// return a function to close the server and database
	return func() {
		cleanup()
		app.Shutdown()
	}, nil
}

func buildServer(env config.EnvVars) (*fiber.App, func(), error) {
	// init the storage
	// db, err := storage.BootstrapMongo(env.MONGODB_URI, env.MONGODB_NAME, 10*time.Second)
	// if err != nil {
	// 	return nil, nil, err
	// }

	// create the fiber app
	app := fiber.New()

	// add middleware
	app.Use(cors.New())
	app.Use(logger.New())

	// add health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("Healthy!")
	})

	// add docs
	// app.Get("/swagger/*", swagger.HandlerDefault)

	// create the user domain
	// todoStore := todo.NewTodoStorage(db)
	// todoController := todo.NewTodoController(todoStore)
	// todo.AddTodoRoutes(app, todoController)

	return app, func() {
		// storage.CloseMongo(db)
	}, nil
}