package auth

import (
	"errors"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	AccountId int `json:"accountId"`
	Role      int `gorm:"column:role"`
	jwt.StandardClaims
}

var (
	secretKey     []byte
	tokenLifeTime time.Duration
)

func Init(secret []byte, lifeTime time.Duration) {
	secretKey = secret
	tokenLifeTime = lifeTime
}

func Sign(accountId, role int) (string, error) {
	claims := Claims{
		AccountId: accountId,
		Role:      role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenLifeTime).Unix(),
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenClaims.SignedString(secretKey)
}

func Verify(authToken string) (claims *Claims, err error) {
	authToken = strings.Replace(authToken, "Bearer ", "", -1)
	// parse and verify the token string
	tokenClaims, err := jwt.ParseWithClaims(authToken, &Claims{}, func(token *jwt.Token) (i interface{}, err error) {
		return secretKey, nil
	})
	// detail for jwt token err message
	if err != nil {
		var message string
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				message = "token is malformed"
			} else if ve.Errors&jwt.ValidationErrorUnverifiable != 0 {
				message = "token could not be verified because of signing problems"
			} else if ve.Errors&jwt.ValidationErrorSignatureInvalid != 0 {
				message = "signature validation failed"
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				message = "token is expired"
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				message = "token is not yet valid before sometime"
			} else {
				message = "can not handle this token"
			}
		}
		err = errors.New(message)
		return nil, err
	}

	if claims, ok := tokenClaims.Claims.(*Claims); ok {
		return claims, nil
	} else {
		return nil, errors.New("token is improper")
	}

}
