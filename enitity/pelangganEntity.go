package enitity

type Pelanggan struct {
	Id             uint   `gorm:"primaryKey;autoIncrement:true"`
	Nama           string `json:"nama"`
	NomorPelanggan string `json:"nomor_pelanggan"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	NoTelpon       string `json:"no_telpo"`
}
