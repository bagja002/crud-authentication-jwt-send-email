package tools

import (
	"fmt"
	"pertemuan4/enitity"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtSecret = []byte("rahasia-super-aman") // Ganti dengan secret yang lebih kuat

// Struktur payload custom (opsional)
type MyCustomClaims struct {
	Username string `json:"username"`
	Id       int
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateJWT(user enitity.Pelanggan) (string, error) {
	claims := MyCustomClaims{
		Username: user.Email,
		Role:     "Pelanggan",
		Id:       int(user.Id),
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "aplikasi-saya",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   fmt.Sprint(user.Id), // bisa juga simpan ID user
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
