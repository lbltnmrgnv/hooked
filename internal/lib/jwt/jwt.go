package jwt

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"hello/internal/models"
	"time"
)

const SigningKey = "testkey"

var (
	ErrTypeAssert            = errors.New("type assertion failed")
	ErrInvalidSigningMethod  = errors.New("invalid signing method")
	ErrInvalidToken          = errors.New("invalid token")
	ErrTokenExpired          = errors.New("token expired")
	ErrNotFoundInTokenClaims = errors.New("not found in token claims")
)

type TokenClaims struct {
	UID      string
	Username string
	Exp      int64
}

func NewToken(user *models.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", ErrTypeAssert
	}

	claims["uid"] = user.ID
	claims["username"] = user.Username
	claims["exp"] = time.Now().Add(time.Minute + 10).Unix()

	tokenString, err := token.SignedString([]byte(SigningKey))
	if err != nil {
		return "", errors.New("parsing error")
	}

	return tokenString, nil
}

func ParseToken(tokenString, signingKey string) (*TokenClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidSigningMethod
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		if errors.Is(err, ErrTokenExpired) {
			return nil, ErrTokenExpired
		}
		fmt.Print(err)
		return nil, errors.New("parsing error")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, ErrInvalidToken
	}

	var userID string
	if uid, ok := claims["uid"]; ok {
		switch v := uid.(type) {
		case string:
			userID = v
		case float64:
			userID = fmt.Sprintf("%.0f", v)
		}
	} else {
		return nil, ErrNotFoundInTokenClaims
	}

	var exp int64
	if expClaim, ok := claims["exp"]; ok {
		switch v := expClaim.(type) {
		case float64:
			exp = int64(v)
		case int64:
			exp = v
		default:
			return nil, ErrNotFoundInTokenClaims
		}
	} else {
		return nil, ErrNotFoundInTokenClaims
	}

	var username string
	if usernameClaim, ok := claims["username"]; ok {
		switch v := usernameClaim.(type) {
		case string:
			username = v
		}
	} else {
		return nil, ErrNotFoundInTokenClaims
	}

	tc := TokenClaims{
		UID:      userID,
		Username: username,
		Exp:      exp,
	}

	return &tc, nil
}
