package notifications

type Request struct {
	To      []string `validate:"required,dive,email"`
	From    string   `validate:"required"`
	Subject string   `validate:"required"`
	Body    string   `validate:"required"`
}

type ChatRequest struct {
	To      string `validate:"required,dive,number"`
	Message string `validate:"required"`
}

type Response struct {
	success bool
	message string
}
