package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/halilemincaliskan/yatirio/controllers"
	"github.com/halilemincaliskan/yatirio/model"
	"log"
)

func main() {
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	model.User{}.Migrate()
	model.Wallet{}.Migrate()
	model.Rate{}.Migrate()

	app.Static("/public", "./public")

	app.Get("/mainpage", controllers.Mainpage{}.Index)
	app.Get("/login", controllers.Login{}.Index)
	app.Get("/register", controllers.Register{}.Index)
	app.Get("/logout", controllers.Login{}.Logout)
	app.Post("/mainpage/:currency/:process", controllers.Mainpage{}.Islem)
	app.Post("/login/check", controllers.Login{}.Check)
	app.Post("/register/signup", controllers.Register{}.SignUp)

	//Port listener
	log.Fatal(app.Listen(":3000"))
}