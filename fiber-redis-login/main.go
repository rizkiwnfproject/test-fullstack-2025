package main

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type User struct {
	RealName string `json:"realname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})

	app := fiber.New()

	app.Post("/login", func(c *fiber.Ctx) error {
		username := c.FormValue("username")
		password := c.FormValue("password")

		val, err := rdb.Get(ctx, "login_"+username).Result()
		if err != nil {
			return c.Status(401).JSON(fiber.Map{"error": "User tidak ditemukan"})
		}

		var user User
		if err := json.Unmarshal([]byte(val), &user); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Data user corrupt"})
		}

		h := sha1.New()
		h.Write([]byte(password))
		hashedPassword := hex.EncodeToString(h.Sum(nil))

		if hashedPassword != user.Password {
			return c.Status(401).JSON(fiber.Map{"error": "Password salah"})
		}

		return c.JSON(fiber.Map{
			"message":  "Login berhasil",
			"realname": user.RealName,
			"email":    user.Email,
		})
	})

	log.Fatal(app.Listen(":3000"))
}
