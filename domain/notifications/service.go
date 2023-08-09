package notifications

import (
	"fmt"
	customvalidator "messaging-interface/pkg/customValidator"
	"messaging-interface/pkg/smtp"
	"messaging-interface/pkg/whatsapp"
)

type service struct {
	iRepo iRepo
}

type iRepo interface {
	getTemplate(string) (string, error)
}

func NewService(_iRepo iRepo) service {
	return service{iRepo: _iRepo}
}

func (s service) SendWhatsApp(req ChatRequest) (res Response, err error) {

	err = whatsapp.CallAPIChat(whatsapp.Chat{
		To:      req.To,
		Message: req.Message,
	})

	if err != nil {
		return
	}

	return res, nil
}

func (s service) SendEmail(req Request) (res Response, err error) {
	v := customvalidator.NewCustomValidator()

	err = v.Validate(&req)
	if err != nil {
		return res, customvalidator.CustomErrorMessage(err)
	}

	err = smtp.SendMailWithGomail(smtp.Mail{
		From:    req.From,
		To:      req.To,
		Subject: req.Subject,
		Body:    req.Body,
	})

	if err != nil {
		return
	}

	res = Response{
		success: true,
		message: fmt.Sprintf("Email sent successfully to %v", req.To),
	}

	return res, nil
}
