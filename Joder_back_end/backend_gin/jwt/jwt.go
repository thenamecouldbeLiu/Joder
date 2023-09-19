package jwt

import (
	"errors"
	"joder/config"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	UserId string `json:"userId"`
	jwt.RegisteredClaims
}

func getSecretKey() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Val.JWT_KEY), nil
	}
}

func GenerateClaim(userId string, expireTime int64) Claims {
	return Claims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expireTime*24) * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

}
func GenerateJWT(userId string) (string, error) {
	claims := GenerateClaim(userId, config.Val.JWT_TOKEN_LIFE)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.Val.JWT_KEY))

	return tokenString, err
}

func ParseToken(token string) (claims *Claims, err error) {
	parsedToken, err := jwt.ParseWithClaims(token, &Claims{}, getSecretKey())

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("not a token")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New("token expired")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New("token not active")
			} else {
				return nil, errors.New("can't handle the token")
			}
		}
	}
	if claims, ok := parsedToken.Claims.(*Claims); ok && parsedToken.Valid {
		return claims, nil
	}

	return nil, errors.New("can't handle the token")

}
