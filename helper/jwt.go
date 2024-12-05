package helper

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTManager interface {
	Generate(string, string) (string, error)
	Verification(string) (*claims, error)
}

type jwtManager struct {
	secretKey  string
	issuer     string
	expiryTime int
}

type claims struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func NewJWTManager(secretKey, issuer string, expiryTime int) JWTManager {
	return &jwtManager{
		secretKey,
		issuer,
		expiryTime,
	}
}

func (j *jwtManager) Generate(id, username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, j.buildClaim(id, username))
	t, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		return "", err
	}

	return t, nil
}

// Verification implements JWTManager.
func (j *jwtManager) Verification(token string) (*claims, error) {
	tk, err := jwt.ParseWithClaims(
		token,
		&claims{},
		func(t *jwt.Token) (interface{}, error) {
			_, ok := t.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("unexpected token signing method")
			}

			return []byte(j.secretKey), nil
		},
	)

	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	claims, ok := tk.Claims.(*claims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}

func (j *jwtManager) buildClaim(id, username string) *claims {
	now := time.Now()
	return &claims{
		ID:       id,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    j.issuer,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Minute * 10)),
		},
	}
}
