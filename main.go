package main

import (
	"github.com/fahmiammryhi/tutor-go-fiber/database"
	"github.com/fahmiammryhi/tutor-go-fiber/entity/migration"
	"github.com/fahmiammryhi/tutor-go-fiber/routers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	database.ConnectDB()
	migration.RunMigrate() //manggil migrate
	// app.Get("/", func(c *fiber.Ctx) error { // app.Get("/", <-- ROUTERS, func -->controllersCONTROLLERS
	// 	return c.SendString("Hello, World!")
	// })

	routers.RouterApp(app) //fiber dimasukan kedalam router agar lgsg terpanggil ke server
	//test client menggunakan thunder client

	app.Listen(":3306")
}
