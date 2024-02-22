package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/makarellav/phogo/models"
	"os"
	"strconv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println(err)

		return
	}

	host := os.Getenv("SMTP_HOST")
	port, err := strconv.Atoi(os.Getenv("SMTP_PORT"))

	if err != nil {
		fmt.Println(err)

		return
	}

	username := os.Getenv("SMTP_USERNAME")
	password := os.Getenv("SMTP_PASSWORD")

	es := models.NewEmailService(models.SMTPConfig{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
	})

	err = es.ForgotPassword("makarellads@gmail.com", "https://phogo.com/reset-pw?token=abc123")

	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("message sent")
}
