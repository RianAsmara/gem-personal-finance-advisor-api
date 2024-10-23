package common

import (
	"errors"
	"strconv"
	"time"

	"github.com/RianAsmara/personal-finance-advisor-api/configuration"
	"github.com/golang-jwt/jwt/v4"
)

func GenerateJWT(email string, userRoles []map[string]interface{}, config configuration.Config) (string, error) {
	jwtSecret := config.Get("JWT_SECRET")
	if jwtSecret == "" {
		return "", errors.New("JWT_SECRET is not set")
	}
	jwtExpired, err := strconv.Atoi(config.Get("JWT_EXPIRED"))
	if err != nil {
		return "", errors.New("JWT_EXPIRED is not set")
	}

	claims := jwt.MapClaims{
		"email":     email,
		"userRoles": userRoles,
		"milis":     time.Now().UnixMilli(),
		"exp":       time.Now().Add(time.Duration(jwtExpired) * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
