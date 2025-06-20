package tools

import "github.com/gofiber/fiber/v2"

func ValidationJwtPelanggan(c *fiber.Ctx, role string, id_admin int, names string) *fiber.Map {
	if role != "Pelanggan" {
		return &fiber.Map{
			"pesan": "Role Bukan Pelanggan",
		}
	}
	if id_admin == 0 {
		return &fiber.Map{
			"pesan": "Admin tidak terdaftar",
		}
	}
	if names == "" {
		return &fiber.Map{
			"pesan": "Tidak ada Nama di dalam Jwt",
		}
	}
	// Jika semuanya valid, kembalikan nilai null (tidak ada pesan kesalahan)
	return nil
}
