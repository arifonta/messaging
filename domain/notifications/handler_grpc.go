package notifications

import (
	context "context"
)

type handler struct {
	isService service
}

func NewHandler(is service) handler {
	return handler{
		isService: is,
	}
}

func (h handler) SendEmail(ctx context.Context, req *EmailRequest) (res *EmailResponse, err error) {
	dataReq := Request{
		To:      req.To,
		From:    req.From,
		Subject: req.Subject,
		Body:    req.Body,
	}
	datares, err := h.isService.SendEmail(dataReq)

	if err != nil {
		return nil, err
	}

	res = &EmailResponse{
		Success: datares.success,
		Message: datares.message,
	}

	return res, nil
}

func (h handler) SendWhatsApp(ctx context.Context, req *WhatsAppRequest) (res *WhatsAppResponse, err error) {
	dataReq := ChatRequest{
		To:      req.To,
		Message: req.Message,
	}

	datares, err := h.isService.SendWhatsApp(dataReq)

	if err != nil {
		return nil, err
	}

	res = &WhatsAppResponse{
		Success: datares.success,
		Message: datares.message,
	}

	return nil, nil
}

func (h handler) mustEmbedUnimplementedMessagingServiceServer() {}
