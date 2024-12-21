package service

import (
	"log"
	"os"

	"github.com/resend/resend-go/v2"
)

const From = "hi@lutfifadlan.com"
const ReplyTo = "hi@lutfifadlan.com"

type EmailService struct {
	Client resend.Client
}

func NewEmailService() *EmailService {
	return &EmailService{
		Client: *resend.NewClient(os.Getenv("RESEND_API_KEY")),
	}
}

func (s *EmailService) SendEmail(to, subject, text string) error {
	params := &resend.SendEmailRequest{
		To:      []string{to},
		From:    From,
		Text:    text,
		Subject: subject,
		ReplyTo: ReplyTo,
	}

	sent, err := s.Client.Emails.Send(params)
	if err != nil {
		return err
	}

	log.Println("Email sent: ", sent)
	return nil
}
