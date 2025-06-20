package controllers

import (
	"fmt"
	"log"
	"pertemuan4/enitity"
	"pertemuan4/models"
	"pertemuan4/pkg/database"
	"pertemuan4/pkg/tools"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func RegisterPelanggan(c *fiber.Ctx) error {
	var request models.Pelanggan
	if err := c.BodyParser(&request); err != nil {
		return err
	}

	newUsers := enitity.Pelanggan{
		Nama:           request.Nama,
		NomorPelanggan: "Random String", // Ganti nanti dengan generator yang benar
		Email:          request.Email,
		Password:       tools.GeneratePassword(request.Password),
		NoTelpon:       request.NoTelpon,
	}

	// Cek apakah email sudah ada
	var existingUsers enitity.Pelanggan
	database.DB.Where("email = ?", request.Email).First(&existingUsers)
	if existingUsers.Id != 0 {
		return c.Status(400).JSON(fiber.Map{
			"pesan": "Email tersebut sudah terdaftar.",
		})
	}

	// Simpan data ke database
	database.DB.Create(&newUsers)

	// Kirim Email ke Users di background (tidak blok proses)
	go func(email, nama string) {
		err := tools.SendRegistrationSuccessEmail(email, nama)
		if err != nil {
			log.Println("Gagal mengirim email:", err)
		}
	}(newUsers.Email, newUsers.Nama)

	// Response langsung tanpa menunggu email selesai dikirim
	return c.JSON(fiber.Map{
		"pesan": "Berhasil membuat akun pelanggan.",
		"data":  newUsers,
	})
}

func LoginPelanggan(c *fiber.Ctx) error {

	var request models.Pelanggan
	if err := c.BodyParser(&request); err != nil {
		return err
	}

	//tidak boleh koskng email dan password
	if request.Email == "" {
		return c.Status(400).JSON(fiber.Map{
			"Pesan": " Email Tersebut di Tidak Boleh Kosong",
		})

	}

	if request.Password == "" {
		return c.Status(400).JSON(fiber.Map{
			"Pesan": " Password Tersebut di Tidak Boleh Kosong",
		})

	}

	//cek data
	var users enitity.Pelanggan
	database.DB.Where("email = ?", request.Email).First(&users)

	if users.Id == 0 {
		return c.Status(400).JSON(fiber.Map{
			"Pesan": "Data Tidak Di Temukan",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(users.Password), []byte(request.Password)); err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"pesan": "Incorrect password!",
		})
	} else {
		t, _ := tools.GenerateJWT(users)
		return c.JSON(fiber.Map{
			"t": t,
		})
	}
}

func GetDataPelangganById(c *fiber.Ctx) error {

	id, _ := c.Locals("Id").(int)
	role, _ := c.Locals("role").(string)
	username, _ := c.Locals("username").(string)

	tools.ValidationJwtPelanggan(c, role, id, username)

	fmt.Print("Id nya itu", id)

	//Get data by Id
	var user enitity.Pelanggan
	database.DB.Where("id = ?", id).Find(&user)

	return c.JSON(fiber.Map{
		"Pesan": "Berhail Get Data Pelanggan",
		"data":  user,
	})
}
