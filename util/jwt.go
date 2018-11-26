package util

import (
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/lisiur/daydayup/models"
)

// Claims .
type Claims struct {
	ID uint `json:"id,omitempty"`
	jwt.StandardClaims
}

var jwtSecret = []byte("salt")

// GenerateToken .
func GenerateToken(user models.User) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)
	claims := Claims{
		user.ID,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// ParseToken .
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
