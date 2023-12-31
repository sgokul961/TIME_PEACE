package usecase

import (
	"errors"
	"fmt"

	"github.com/jinzhu/copier"
	"gokul.go/pkg/config"
	helper_interface "gokul.go/pkg/helper/interface"
	interfaces "gokul.go/pkg/repository/interface"
	"gokul.go/pkg/usecase/usecaseInterfaces"
	"gokul.go/pkg/utils/models"
)

type otpUseCase struct {
	cfg           config.Config
	otpRepository interfaces.OtpRepository
	helper        helper_interface.Helper
}

// // VarifyOtp implements usecaseInterfaces.OtpUseCase.
//
//	func (*otpUseCase) VarifyOtp(code models.VarifyData) (models.TokenUsers, error) {
//		panic("unimplemented")
//	}
func NewOtpUseCase(cfg config.Config, repo interfaces.OtpRepository, help helper_interface.Helper) usecaseInterfaces.OtpUseCase {
	return &otpUseCase{
		cfg:           cfg,
		otpRepository: repo,
		helper:        help,
	}
}

func (ot *otpUseCase) SendOTP(phone string) error {

	ok := ot.otpRepository.FindUserByMobileNumber(phone)
	if !ok {
		return errors.New("user does not exist")
	}
	ot.helper.TwilioSetup(ot.cfg.ACCOUNTSID, ot.cfg.AUTHTOKEN)

	fmt.Println("accsid:", ot.cfg.ACCOUNTSID)

	fmt.Println("auth", ot.cfg.AUTHTOKEN)

	_, err := ot.helper.TwilioSendOTP(phone, ot.cfg.SERVICESID)
	if err != nil {
		return errors.New("error occured while genarating otp")
	}
	return nil

}
func (ot *otpUseCase) VerifyOTP(code models.VarifyData) (models.TokenUsers, error) {

	ot.helper.TwilioSetup(ot.cfg.ACCOUNTSID, ot.cfg.AUTHTOKEN)
	err := ot.helper.TwilioVerifyOTP(ot.cfg.SERVICESID, code.Code, code.PhoneNumber)
	if err != nil {
		//this guard clause catches the error code runs only until here
		return models.TokenUsers{}, errors.New("error while verifying")
	}

	// if user is authenticated using OTP send back user details

	userDetails, err := ot.otpRepository.UserDeatilsUsingPhone(code.PhoneNumber)

	if err != nil {
		return models.TokenUsers{}, err
	}

	tokenString, err := ot.helper.GenerateTokenClients(userDetails)
	if err != nil {
		return models.TokenUsers{}, nil
	}
	var user models.UserDeatilsResponse

	err = copier.Copy(&user, &userDetails)
	if err != nil {
		return models.TokenUsers{}, err
	}
	return models.TokenUsers{
		Users: user,
		Token: tokenString,
	}, nil
}
