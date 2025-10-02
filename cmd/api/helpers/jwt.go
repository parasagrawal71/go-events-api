package helpers

import (
	"go-events-api/cmd/api/config"
	"go-events-api/cmd/api/models"
	"go-events-api/internal/env"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Custom claims (you can also use jwt.RegisteredClaims)
type Claims struct {
	User User `json:"user"`
	jwt.RegisteredClaims
}

func GenerateJWT(user models.User) (string, error) {
	// Set claims
	claims := &Claims{
		User: User{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)), // expires in 1 hr
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    config.APP_NAME,
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token with secret key
	jwtKey := []byte(env.GetEnvString("JWT_SECRET", ""))
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyJWT(tokenStr string) (*Claims, error) {
	jwtKey := []byte(env.GetEnvString("JWT_SECRET", ""))
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}

	// // check token expiry
	// if time.Now().After(claims.ExpiresAt.Time) {
	// 	return nil, err
	// }

	return claims, nil
}
