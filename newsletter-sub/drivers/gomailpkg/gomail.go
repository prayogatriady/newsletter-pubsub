package gomailpkg

import (
	"fmt"
	"log"
	models_email "newsletter-sub/models/email"
	"newsletter-sub/utils/config"

	"gopkg.in/gomail.v2"
)

var appConfig = config.AppCfg

func SendViaGomail(data *models_email.EmailPayload) {
	d := gomail.NewDialer("smtp.gmail.com", 587, appConfig.GoMail.SenderEmail, appConfig.GoMail.SenderPassword)
	s, err := d.Dial()
	if err != nil {
		panic(err)
	}

	m := gomail.NewMessage()
	m.SetHeader("From", appConfig.GoMail.SenderEmail)
	m.SetAddressHeader("To", data.Email, data.Name)
	m.SetHeader("Subject", "Newsletter #1")
	m.SetBody("text/html", fmt.Sprintf("Hello %s!", data.Name))

	if err := gomail.Send(s, m); err != nil {
		log.Printf("Could not send email to %q: %v", data.Email, err)
	}
	m.Reset()
}
