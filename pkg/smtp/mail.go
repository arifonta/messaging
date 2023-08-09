package smtp

import (
	"log"
	"messaging-interface/config"
	"strconv"

	"gopkg.in/gomail.v2"
)

type Mail struct {
	From    string   `json:"from"`
	To      []string `json:"to"`
	Subject string   `json:"subject"`
	Body    string   `json:"body"`
}

func SendMailWithGomail(req Mail) (err error) {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", req.From)
	for _, t := range req.To {
		mailer.SetHeader("To", t)
	}
	mailer.SetAddressHeader("Cc", "arifianto.budi@jec.co.id", "admin")
	mailer.SetHeader("Subject", req.Subject)
	mailer.SetBody("text/html", req.Body)
	//mailer.Attach("./notes")

	port, _ := strconv.Atoi(config.AppConfig.Smtp.Port)

	dialer := gomail.NewDialer(
		config.AppConfig.Smtp.Host,
		port,
		config.AppConfig.Email.Email,
		config.AppConfig.Email.Password,
	)

	err = dialer.DialAndSend(mailer)
	if err != nil {
		log.Println("error when try to send email from", req.From, "to", req.To)
		return
	}
	return
}
