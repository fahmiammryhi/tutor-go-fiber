package routers

import (
	"github.com/fahmiammryhi/tutor-go-fiber/controllers"
	"github.com/gofiber/fiber/v2"
)

func RouterApp(c *fiber.App){
	c.Get("/", controllers.UserControllerShow) // cara membuat route
	
}