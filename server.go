package main

import (
    

    "github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django/v3"
)

func Server() fiber.App{

	engine := django.New("./templates", ".html")
	
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/", "./static")

	return *app
}