package main

import (
	"github.com/gofiber/fiber" // import the fiber package

	config "gotbotpoc/config"
	router "gotbotpoc/router"
)

func main() { // entry point to our program

	// Connect to database
	// if err := database.Connect(); err != nil {
	// 	log.Fatal(err)
	// }

	app := fiber.New() // call the New() method - used to instantiate a new Fiber App

	// app.Use(middleware.Logger())

	router.SetupRoutes(app)

	app.Listen(config.Config("PORT")) // listen/Serve the new Fiber app on port 3000

}
