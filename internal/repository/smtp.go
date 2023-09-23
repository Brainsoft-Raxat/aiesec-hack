package repository

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"net/smtp"

	"github.com/Brainsoft-Raxat/aiesec-hack/internal/app/config"
)

type smtpRepo struct {
	fromEmail string
	smtpHost  string
	smtpPort  string
	smtpUser  string
	smtpPass  string
}

func NewSMTPRepository(cfg config.SMTP) SMTP {
	return &smtpRepo{
		fromEmail: cfg.FromEmail,
		smtpHost:  cfg.SMTPHost,
		smtpPort:  cfg.SMTPPort,
		smtpUser:  cfg.SMTPUser,
		smtpPass:  cfg.SMTPPass,
	}
}

func (r *smtpRepo) SendEmailWithAttachment(ctx context.Context, fileData []byte, fileName, toEmail string) error {
	// Create a buffer to store the message content
	var msgBuffer bytes.Buffer

	// Create a boundary for separating parts of the email
	boundary := "myboundary"

	// Compose the email headers
	headers := fmt.Sprintf("From: %s\r\n", r.fromEmail)
	headers += fmt.Sprintf("To: %s\r\n", toEmail)
	headers += fmt.Sprintf("Subject: Citix PhotoShoot\r\n")
	headers += fmt.Sprintf("MIME-Version: 1.0\r\n")
	headers += fmt.Sprintf("Content-Type: multipart/mixed; boundary=%s\r\n\r\n", boundary)

	// Write the headers to the message buffer
	msgBuffer.WriteString(headers)

	// Write the email text
	emailText := "Email content"
	textPart := fmt.Sprintf("--%s\r\n", boundary)
	textPart += fmt.Sprintf("Content-Type: text/plain; charset=\"utf-8\"\r\n\r\n")
	textPart += fmt.Sprintf("%s\r\n\r\n", emailText)
	msgBuffer.WriteString(textPart)

	// Encode the file data to base64
	encodedFileData := base64.StdEncoding.EncodeToString(fileData)

	// Write the attachment
	attachmentPart := fmt.Sprintf("--%s\r\n", boundary)
	attachmentPart += fmt.Sprintf("Content-Type: application/octet-stream\r\n")
	attachmentPart += fmt.Sprintf("Content-Disposition: attachment; filename=\"%s\"\r\n", fileName)
	attachmentPart += fmt.Sprintf("Content-Transfer-Encoding: base64\r\n\r\n")
	attachmentPart += fmt.Sprintf("%s\r\n\r\n", encodedFileData)
	msgBuffer.WriteString(attachmentPart)

	// Final boundary to mark the end of the message
	finalBoundary := fmt.Sprintf("--%s--\r\n", boundary)
	msgBuffer.WriteString(finalBoundary)

	// Convert the message buffer to a string
	messageBody := msgBuffer.String()

	// Send the email using SMTP
	auth := smtp.PlainAuth("", r.smtpUser, r.smtpPass, r.smtpHost)
	err := smtp.SendMail(
		fmt.Sprintf("%s:%s", r.smtpHost, r.smtpPort),
		auth,
		r.fromEmail,
		[]string{toEmail},
		[]byte(messageBody),
	)
	if err != nil {
		return err
	}

	return nil
}
