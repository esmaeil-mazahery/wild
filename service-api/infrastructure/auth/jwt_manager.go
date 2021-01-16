package auth

import (
	"fmt"
	"time"

	"github.com/EsmaeilMazahery/wild/model"
	"github.com/dgrijalva/jwt-go"
)

// JWTManager is a JSON web token manager
type JWTManager struct {
	secretKey          string
	tokenLoginDuration time.Duration
}

// MemberClaims is a custom JWT claims that contains some user's information
type MemberClaims struct {
	jwt.StandardClaims
	Mobile    string    `json:"mobile"`
	ID        string    `json:"id"`
	Role      string    `json:"role"`
	LoginDate time.Time `json:"loginDate"`
}

// NewJWTManager returns a new JWT manager
func NewJWTManager(secretKey string, tokenLoginDuration time.Duration) *JWTManager {
	return &JWTManager{secretKey, tokenLoginDuration}
}

// GenerateLogin generates and signs a new token for a user
func (manager *JWTManager) GenerateLogin(member *model.Member) (string, error) {
	claims := MemberClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(manager.tokenLoginDuration).Unix(),
		},
		ID:        member.ID.Hex(),
		Mobile:    member.Mobile,
		Role:      "user",
		LoginDate: time.Now(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(manager.secretKey))
}

// VerifyLogin verifies the access token string and return a member claim if the token is valid
func (manager *JWTManager) VerifyLogin(accessToken string) (*MemberClaims, error) {
	token, err := jwt.ParseWithClaims(
		accessToken,
		&MemberClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("unexpected token signing method")
			}

			return []byte(manager.secretKey), nil
		},
	)

	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	claims, ok := token.Claims.(*MemberClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}
