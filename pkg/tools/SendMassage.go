package tools

import (
	"fmt"
	"gopkg.in/gomail.v2"
)

const (
	SMTPHost     = "smtp.gmail.com"
	SMTPPort     = 587
	SenderName   = "Kelas Golang <isi dengan menggunakan email mu>"
	AuthEmail    = "isi Email Kamu"
	AuthPassword = "isi Password App di Google Seting" // App Password (tanpa spasi) (Mendaappatkan nya di Google Setting)
)

// sendRegistrationSuccessEmail mengirim email notifikasi registrasi berhasil
func SendRegistrationSuccessEmail(toEmail, userName string) error {
	subject := "Selamat! Anda telah berhasil register"
	body := fmt.Sprintf(`
		<html>
		<body style="font-family: Arial, sans-serif; line-height: 1.6;">
			<h2>Halo %s,</h2>
			<p>Selamat! Anda telah berhasil register di aplikasi <b>Kelas Golang</b>.</p>
			<p>Akun Anda sudah aktif dan siap digunakan.</p>
			<hr>
			<p>Jika Anda tidak merasa melakukan registrasi ini, silakan abaikan email ini.</p>
			<br>
			<p>Salam hangat,</p>
			<p><b>Tim Kelas Golang</b></p>
		</body>
		</html>
	`, userName)

	m := gomail.NewMessage()
	m.SetHeader("From", SenderName)
	m.SetHeader("To", toEmail)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewDialer(SMTPHost, SMTPPort, AuthEmail, AuthPassword)

	return d.DialAndSend(m)
}
