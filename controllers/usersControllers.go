package controllers

import (
	"pertemuan4/enitity"
	"pertemuan4/models"
	"pertemuan4/pkg/database"
	"pertemuan4/pkg/tools"

	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {

	//inputan query
	pelatihan := c.Query("pelatihan")
	model := c.Query("model")

	pelatihan1 := c.Params("perlatihan")

	users := []enitity.Users{
		{
			Id:       1,
			Nama:     "Bagja",
			Email:    "rendiSejagat@gmail.com",
			Password: "rendu",
			NoTelpon: "093123129",
		},
		{
			Id:       1,
			Nama:     "Bagja",
			Email:    "rendiSejagat@gmail.com",
			Password: "rendu",
			NoTelpon: "093123129",
		},
	}

	return c.JSON(fiber.Map{
		"Pesan":       "Berhasil Mendapatkan Data",
		"dataQuery":   pelatihan,
		"pelaatihan1": pelatihan1,
		"model":       model,
		"data":        users,
	})
}

func CreateUsersFormData(c *fiber.Ctx) error {

	//Body Pars
	var reques models.Users
	if err := c.BodyParser(&reques); err != nil {
		return err
	}

	newUsers := enitity.Users{
		Nama:     reques.Nama,
		Email:    reques.Email,
		Password: reques.Password,
	}

	return c.JSON(fiber.Map{
		"Pesan": "Berhasil Menginput Data",
		"data":  newUsers,
	})
}

func CreateUsersFormDataJhon(c *fiber.Ctx) error {

	//Body Pars
	var reques models.Users
	if err := c.BodyParser(&reques); err != nil {
		return err
	}
	newUsers := enitity.Users{
		Nama:     reques.Nama,
		Email:    reques.Email,
		Password: tools.GeneratePassword(reques.Password),
	}
	//Buat databas
	database.DB.Create(&newUsers)

	return c.JSON(fiber.Map{
		"Pesan": "Berhasil Menginput Data",
		"data":  newUsers,
	})
}

func GetUsersData(c *fiber.Ctx) error {

	//
	var users []enitity.Users
	var users1Data enitity.Users

	//Memanggil semua data Data
	database.DB.Find(&users)

	//Memanggil 1 Data Aja
	database.DB.First(&users1Data)

	return c.JSON(fiber.Map{
		"Pesan":       "Berhasil Get Data",
		"data":        users,
		"data singel": users1Data,
	})
}

func GetDataById(c *fiber.Ctx) error {

	id := c.Query("id")

	var user enitity.Users

	//database
	database.DB.Where("id = ?", id).First(&user)

	return c.JSON(fiber.Map{
		"Pesan": "Berhasil Get Data",
		"data":  user,
	})
}

func UpdateData(c *fiber.Ctx) error {

	id := c.Query("id")

	var user enitity.Users

	var reques models.Users
	if err := c.BodyParser(&reques); err != nil {
		return err
	}
	//database
	database.DB.Where("id = ?", id).First(&user)
	//Pengubahan data
	user.Nama = reques.Nama
	user.Email = reques.Email
	user.NoTelpon = reques.NoTelpon
	//Update nya
	database.DB.Save(&user)

	return c.JSON(fiber.Map{
		"Pesan": "Berhasil Update Data",
		"data":  user,
	})
}

func DeleteData(c *fiber.Ctx) error {
	id := c.Query("id")

	var user enitity.Users

	database.DB.Where("id = ?", id).Delete(&user)

	return c.JSON(fiber.Map{})
}
