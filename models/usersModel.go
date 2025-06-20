package models

type Users struct {
	Id       uint
	Nama     string `json:"nama"`
	Email    string `json:"email"`
	Password string `json:"password"`
	NoTelpon string `json:"no_telpo"`
}
