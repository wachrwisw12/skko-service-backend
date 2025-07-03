package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/gofiber/fiber/v2"
	"github.com/yeawyow/gateway/db"
	"github.com/yeawyow/gateway/handler"
)

var DB *sql.DB

func main() {
	// โหลด .env
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ ไม่พบไฟล์ .env (ไม่เป็นไรถ้า set env ไว้ในระบบแล้ว)")
	}

	// เชื่อมต่อ DB
	if err := db.Connect(); err != nil {
		log.Fatal("❌ ไม่สามารถเชื่อมต่อ database ได้:", err)
	}

	// ดึง cert path จาก ENV (หรือใช้ค่าตั้งต้น)
	certFile := os.Getenv("SSL_CERT_FILE")
	if certFile == "" {
		certFile = "/etc/ssl/moph/wildcard_moph_go_th.crt"
	}
	keyFile := os.Getenv("SSL_KEY_FILE")
	if keyFile == "" {
		keyFile = "/etc/ssl/moph/wildcard_moph_go_th.key"
	}

	// สร้าง Fiber app
	app := fiber.New()

	// Auth Group
	auth := app.Group("/auth")
	auth.Get("/line", handler.LineHandler)

	// API Group
	api := app.Group("/api/v1")
	api.Post("/sendmoph", func(c *fiber.Ctx) error {
		return c.SendString("sendmoph")
	})
	api.Post("/gentoken", func(c *fiber.Ctx) error {
		return c.SendString("gentoken")
	})

	// เปิด HTTPS ที่ port 3001
	log.Println("🚀 เริ่มบริการ HTTPS ที่ https://localhost:3001/")
	if err := app.ListenTLS(":3001", certFile, keyFile); err != nil {
		log.Fatal("❌ เปิด HTTPS ไม่สำเร็จ:", err)
	}
}
