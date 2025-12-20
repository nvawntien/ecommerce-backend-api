package utils

import (
	"bytes"
	"go-ecommerce-backend-api/global"
	"html/template"

	"gopkg.in/gomail.v2"
)

type OTPData struct {
	FullName string
	OTPCode  string
}

func SendOTPEmail(toEmail string, fullName string, otpCode string) error {
	tmplText := `
	<!DOCTYPE html>
	<html>
	<head>
		<style>
			.container { font-family: Arial, sans-serif; max-width: 500px; margin: auto; padding: 20px; border: 1px solid #ddd; border-radius: 10px; }
			.header { color: #333; text-align: center; }
			.otp { font-size: 24px; font-weight: bold; color: #2c3e50; text-align: center; letter-spacing: 5px; padding: 10px; background: #f4f4f4; margin: 20px 0; }
			.footer { font-size: 12px; color: #888; text-align: center; margin-top: 20px; }
		</style>
	</head>
	<body>
		<div class="container">
			<h2 class="header">Xác nhận đăng ký tài khoản</h2>
			<p>Xin chào <strong>{{.FullName}}</strong>,</p>
			<p>Cảm ơn bạn đã đăng ký tại E-commerce App. Mã xác thực (OTP) của bạn là:</p>
			<div class="otp">{{.OTPCode}}</div>
			<p>Mã này sẽ hết hạn sau 5 phút. Vui lòng không chia sẻ mã này với bất kỳ ai.</p>
			<div class="footer">Đây là email tự động, vui lòng không phản hồi.</div>
		</div>
	</body>
	</html>`

	tmpl, err := template.New("otp_email").Parse(tmplText)
	if err != nil {
		return err
	}

	data := OTPData{
		FullName: fullName,
		OTPCode:  otpCode,
	}

	var body bytes.Buffer
	if err := tmpl.Execute(&body, data); err != nil {
		return err
	}

	m := gomail.NewMessage()
	m.SetHeader("From", global.Config.Email.Sender)
	m.SetHeader("To", toEmail)
	m.SetHeader("Subject", "Mã xác nhận tài khoản - E-commerce")
	m.SetBody("text/html", body.String())

	d := gomail.NewDialer(global.Config.Email.SMTPHost, global.Config.Email.SMTPPort, global.Config.Email.Sender, global.Config.Email.Password)
	return d.DialAndSend(m)
}