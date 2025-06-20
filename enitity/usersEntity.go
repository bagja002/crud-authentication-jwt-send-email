package enitity

type Users struct {
	Id       uint   `gorm:"primaryKey;autoIncrement:true"`
	Nama     string `json:"nama"`
	Email    string `json:"email"`
	Password string `json:"password"`
	NoTelpon string `json:"no_telpo"`
}
