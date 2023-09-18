package email

import (
	"newsletter-sub/drivers/gomailpkg"
	models_email "newsletter-sub/models/email"
)

type EmailService interface {
	SendEmail(data *models_email.EmailPayload)
}

type emailService struct {
}

// Dependency injection
func NewEmailService() EmailService {
	return &emailService{}
}

func (s *emailService) SendEmail(data *models_email.EmailPayload) {
	gomailpkg.SendViaGomail(data)
}
