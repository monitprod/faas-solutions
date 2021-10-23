package service

import (
	"bytes"
	"fmt"
	"log"
	"net/smtp"
	"os"
	"text/template"

	m "github.com/monitprod/core/pkg/models"
)

type EmailOptions struct {
	Subject string
	Body    string
}

type EmailService interface {
	SendToMany(recipients []m.User, opts EmailOptions) error
	send(recipient m.User, opts EmailOptions) error
}

type EmailServiceImp struct {
}

func NewEmailService() EmailService {
	return &EmailServiceImp{}
}

func (e *EmailServiceImp) SendToMany(recipients []m.User, opts EmailOptions) error {
	opts.Subject = "Sua lista de produtos 📭"

	for _, user := range recipients {
		e.send(user, opts)
	}

	return nil
}

func (e *EmailServiceImp) send(recipient m.User, opts EmailOptions) error {
	// Sender data.
	from := os.Getenv("SE_MAIL_FROM")
	password := os.Getenv("SE_MAIL_FROM_PASSWORD")

	// smtp server configuration.
	smtpHost := os.Getenv("SE_MAIL_SMTP_HOST")
	smtpPort := os.Getenv("SE_MAIL_SMTP_PORT")

	// Receiver email address.
	to := []string{
		recipient.Email,
	}

	// TODO: Find another authentication method most secure
	auth := smtp.PlainAuth("", from, password, smtpHost)

	t, _ := template.New("").Parse(opts.Body)

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: %s \n%s\n\n", opts.Subject, mimeHeaders)))

	//TODO: Find best shapes to replace data from template
	t.Execute(&body, nil)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())

	if err != nil {
		log.Fatalln("Error on send mail:\n", err)
		return nil
	}

	log.Println("Email Sent!")
	return nil
}
