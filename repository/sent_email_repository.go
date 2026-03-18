package repository

import (
	"fmt"

	"gopkg.in/gomail.v2"
)

type SentEmailRepository interface {
	SentEmail(email string) bool
}

type sentEmailRepository struct{}

func NewSentEmailRepository() SentEmailRepository {
	return &sentEmailRepository{}
}

func (r *sentEmailRepository) SentEmail(email string) bool {
	fmt.Print("sent email ...")
	m := gomail.NewMessage()
	m.SetHeader("From", "rppratama1771@gmail.com")
	m.SetHeader("To", "ryanjoker87@gmail.com")
	m.SetHeader("Subject", "Hello, from Share Recipes Foods")
	// m.SetBody("text/plain", "Duis sunt magna esse id culpa Lorem proident ex pariatur id aute. Ipsum sint elit cupidatat dolore dolor quis ipsum ad tempor voluptate et occaecat pariatur. Ea fugiat ad id officia ut eu fugiat elit mollit labore quis enim sint aliquip.")
	m.SetBody("text/html", `
  		<h2>Hi!</h2>
  		<p>Terima kasih telah mendaftar ke Share Recipes.</p>
	`)

	d := gomail.NewDialer("smtp-relay.brevo.com", 587, "rppratama1771@gmail.com", "hmX0IfAF4pywNPBd")

	if err := d.DialAndSend(m); err != nil {
		fmt.Printf("Failed to send email to %s: %v\n", email, err)
		return false
	}

	fmt.Print("sent success ...")
	return true
}
