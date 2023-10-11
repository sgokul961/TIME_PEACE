package helper

import (
	"time"

	"github.com/golang-jwt/jwt"
	"gokul.go/pkg/utils/models"
)

func GenerateTokenAdmin(admin models.AdminDetailsResponse) (string, error) {

	claims := &authCustomClaims{
		Id:    admin.ID,
		Email: admin.Email,
		Role:  "admin",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("comebuyjersey"))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
