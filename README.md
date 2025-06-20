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
