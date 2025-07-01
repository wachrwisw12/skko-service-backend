package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/yeawyow/gateway/handler"
)

func main() {
	app:=fiber.New()
	//auth Group
	auth:=app.Group("/auth")
	auth.Get("/line",handler.LineHandler)
   fmt.Print("test")
	log.Fatal(app.Listen(":8080"))
}
