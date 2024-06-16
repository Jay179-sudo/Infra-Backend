package mail

import (
	"log"
	"net/smtp"
	"os"
)

func Send(body, to string) {
	from := "pp7405124@gmail.com"
	pass := os.Getenv("PASSWORD") // configured using secrets

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Your Request for a server has been processed\n\n" +
		body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}

	log.Print("Email sent to ", to)
}
