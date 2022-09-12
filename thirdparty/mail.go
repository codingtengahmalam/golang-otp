package thirdparty

import (
	"context"
	"fmt"
	"net/smtp"
	"os"
)

type SendEmailRequest struct {
	To      string `json:"to"`
	Body    string `json:"body"`
	Subject string `json:"subject"`
}

type emailSmtp struct {
	Port        string `json:"port"`
	Host        string `json:"host"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	SenderEmail string `json:"sender_email"`
	SenderName  string `json:"sender_name"`
}

type EmailSmtp interface {
	SendEmail(ctx context.Context, request SendEmailRequest) error
}

func NewEmailSmtp() EmailSmtp {

	return &emailSmtp{
		Port:        os.Getenv("SMTP_PORT"),
		Host:        os.Getenv("SMTP_HOST"),
		Username:    os.Getenv("SMTP_USERNAME"),
		Password:    os.Getenv("SMTP_PASSWORD"),
		SenderEmail: os.Getenv("SMTP_SENDER_EMAIL"),
		SenderName:  os.Getenv("SMTP_SENDER_NAME"),
	}
}

func (e *emailSmtp) SendEmail(ctx context.Context, request SendEmailRequest) error {
	auth := smtp.PlainAuth("", e.Username, e.Password, e.Host)

	message := e.buildMessage(request)

	mailServerAddress := e.Host + ":" + e.Port

	err := smtp.SendMail(mailServerAddress, auth, e.SenderEmail, []string{request.To}, []byte(message))
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (e *emailSmtp) buildMessage(request SendEmailRequest) string {
	message := fmt.Sprintf("From: %s\r\n", e.SenderEmail)
	message += fmt.Sprintf("To: %s\r\n", request.To)
	message += fmt.Sprintf("Subject: %s\r\n", request.Subject)
	message += fmt.Sprintf("\r\n%s\r\n", request.Body)

	return message
}
