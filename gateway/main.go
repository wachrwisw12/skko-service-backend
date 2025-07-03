package main

import (
	"database/sql"
	"log"

	"github.com/joho/godotenv"

	"github.com/gofiber/fiber/v2"
	"github.com/yeawyow/gateway/db"
	"github.com/yeawyow/gateway/handler"
)

var DB *sql.DB
func main() {

	_ = godotenv.Load()

	if err :=db.Connect(); err !=nil{
		log.Fatal("ไม่สามารถเชื่อมต่อ db ได้",err)
	}
	app:=fiber.New()
	//auth Group
	auth:=app.Group("/auth")
	auth.Get("/line",handler.LineHandler)
    
   //skktomorprom 
   mophnot:=app.Group("/api/v1")
  
   mophnot.Post("/sendmoph",func (c *fiber.Ctx) error  {
	return c.SendString("sendmoph")
   })
   mophnot.Post("/gentoken",func (c *fiber.Ctx) error  {
	return c.SendString("gentoken")
   })
	log.Fatal(app.Listen(":3001"))
}
