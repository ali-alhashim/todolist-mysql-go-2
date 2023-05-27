package main

import "github.com/gofiber/fiber/v2"

func MyURLS(server *fiber.App) {

	app := server

	app.Get("/", Home)
	
	app.Get("/addTask/", AddTask)

	app.Post("/postTask/", PostTask)
}
