package main

import (
    "fmt"
    "log"
    "github.com/sendgrid/sendgrid-go"
    "github.com/sendgrid/sendgrid-go/helpers/mail"
)

func main() {
    from := mail.NewEmail("Rakhat", "rakhat.khamitov@nu.edu.kz") // Change to your verified sender
    subject := "Sending with Twilio SendGrid is Fun"
    to := mail.NewEmail("Sou", "rakhat.khamitov@nu.edu.kz") // Change to your recipient
    plainTextContent := "and easy to do anywhere, even with Go"
    htmlContent := "<strong>and easy to do anywhere, even with Go</strong>"
    message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
    client := sendgrid.NewSendClient("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")
    response, err := client.Send(message)
    if err != nil {
        log.Println(err)
    } else {
        fmt.Println(response.StatusCode)
        fmt.Println(response.Headers)
    }
}
