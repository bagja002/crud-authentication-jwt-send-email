package main

import (
	"github.com/gofiber/fiber/v2/middleware/logger"
	"pertemuan4/controllers"
	"pertemuan4/middleware"
	"pertemuan4/pkg/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	database.ConnectDatabase()

	//route + controllers

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	//Rekomendasi Penulisa Route Agar Bisa di baca
	app.Get("/getUsers/:perlatihan", controllers.GetUsers)
	app.Post("/createDataUsers", controllers.CreateUsersFormData)
	app.Post("creteDataUsersFromJson", controllers.CreateUsersFormDataJhon)
	app.Get("/getDataUsers", controllers.GetUsersData)

	app.Get("/usersById", controllers.GetDataById)
	app.Put("/updateUsers", controllers.UpdateData)
	app.Delete("/deleteData", controllers.DeleteData)

	//Pelanggan Route
	app.Post("/createPelanggan", controllers.RegisterPelanggan)
	app.Post("/loginPelanggan", controllers.LoginPelanggan)
	app.Get("/pelangganById", middleware.JwtProtect(), controllers.GetDataPelangganById)

	app.Listen(":9000")

}
