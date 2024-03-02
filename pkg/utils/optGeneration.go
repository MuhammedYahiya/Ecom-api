package utils

import (
	"crypto/rand"
	"crypto/tls"
	"fmt"
	"os"

	gomail "github.com/go-mail/mail"
)

// Otpgeneration generates and sends an OTP to the specified email address
func Otpgeneration(email string) (string, error) {
	// Generate OTP
	onetimepassword, err := GenCaptchaCode()
	if err != nil {
		return "", err
	}

	// Create new message
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("EMAIL"))
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Your OTP")
	m.SetBody("text/plain", fmt.Sprintf("%s is your OTP to register. Thank you for registering. Do not share this code with anyone.", onetimepassword))

	// Create new dialer
	d := gomail.NewDialer("smtp.gmail.com", 587, os.Getenv("EMAIL"), os.Getenv("PASSWORD"))
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send email
	if err := d.DialAndSend(m); err != nil {
		fmt.Println("Failed to send OTP:", err)
		return "", err
	}

	fmt.Println("OTP sent successfully")
	return onetimepassword, nil
}

// GenCaptchaCode generates a 6-digit OTP
func GenCaptchaCode() (string, error) {
	codes := make([]byte, 6)
	if _, err := rand.Read(codes); err != nil {
		return "", err
	}

	for i := 0; i < 6; i++ {
		codes[i] = uint8(48 + (codes[i] % 10))
	}

	return string(codes), nil
}
