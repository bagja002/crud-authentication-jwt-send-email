# 🛡️ CRUD Authentication JWT dengan Notifikasi Email

Aplikasi backend sederhana menggunakan **Go Fiber**, menyediakan fitur:

- **CRUD User** (Create, Read, Update, Delete)  
- **Authentication** dengan JWT (JSON Web Token)  
- **Notifikasi Email** saat registrasi berhasil

---

## 🚀 Fitur Utama

- Registrasi user dengan password terenkripsi menggunakan `bcrypt`
- Login dan generate token JWT untuk autentikasi
- Proteksi endpoint menggunakan middleware JWT
- Kirim email konfirmasi registrasi secara asynchronous menggunakan `gomail`
- Contoh penggunaan goroutine untuk mengirim email tanpa blocking
- Integrasi database MySQL menggunakan GORM dengan auto migration

---

## 🛠 Teknologi

- **Go 1.20+**
- [Fiber](https://gofiber.io/) – Web framework
- [GORM](https://gorm.io/) – ORM untuk MySQL
- [gomail](https://github.com/go-gomail/gomail) – Kirim email SMTP
- [JWT](https://github.com/golang-jwt/jwt) – Token untuk autentikasi

---

## 🔧 Setup dan Menjalankan Aplikasi

### 1. Clone Repository

```bash
git clone https://github.com/bagja002/crud-authentication-jwt-send-email.git
cd crud-authentication-jwt-send-email

```
### 2.Atur Konfigurasi Database
Edit file database/connect.go:

```bash
username := "user"
password := "password"
host := "127.0.0.1"
port := "3306"
databaseName := "nama_database"

dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	username, password, host, port, databaseName)

```
### 3. Atur Konfigurasi SMTP Email
Edit file tools/email.go:


```bash
const (
	SMTPHost     = "smtp.gmail.com"
	SMTPPort     = 587
	SenderName   = "Nama Kamu <emailkamu@gmail.com>"
	AuthEmail    = "emailkamu@gmail.com"
	AuthPassword = "app_password_anda"
)

```
💡 Gunakan App Password Gmail (bukan password biasa) untuk AuthPassword.

### 4. Install Dependencies
```bash
go mod tidy

```

### 5. Jalankan Aplikasi
```bash
go run main.go

```
## 📌 API Endpoint Contoh

| Method | Endpoint           | Keterangan                 | Autentikasi JWT |
|--------|--------------------|----------------------------|-----------------|
| POST   | `/createPelanggan` | Registrasi user baru       | ❌ Tidak        |
| POST   | `/loginPelanggan`  | Login & dapatkan token JWT | ❌ Tidak        |
| GET    | `/pelangganById`   | Ambil data profil user     | ✅ Ya           |

