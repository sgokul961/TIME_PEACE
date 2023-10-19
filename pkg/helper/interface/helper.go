package interfaces

import (
	"mime/multipart"

	"gokul.go/pkg/utils/models"
)

type Helper interface {
	GenerateTokenAdmin(admin models.AdminDetailsResponse) (string, error)
	AddImageToS3(file *multipart.FileHeader) (string, error)
	TwilioSetup(username string, password string)
	TwilioSendOTP(phone string, serviceID string) (string, error)
	TwilioVerifyOTP(serviceID string, code string, phone string) error
	GenerateTokenClients(user models.UserDeatilsResponse) (string, error)
}
