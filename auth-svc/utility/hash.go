package utility

import (
	"auth-svc/internal/config"
	"auth-svc/internal/consts"
	"context"
	"errors"
	"fmt"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/golang-jwt/jwt/v5"
)

var ConfigJWTHashFunction = jwt.SigningMethodHS256

func GenJWT(payload jwt.MapClaims) (token string, err error) {
	tokenJWT := jwt.NewWithClaims(ConfigJWTHashFunction, payload)
	// Convert the secret key to []byte as jwt.SignedString expects a []byte for HMAC signing
	secretKeyBytes := []byte(config.GetConfig().Auth.SecretKey)
	token, err = tokenJWT.SignedString(secretKeyBytes)
	if err != nil {
		g.Log().Error(context.Background(), "GenJWT error: ", err)
		return
	}
	return
}

// Validate JWT token
func ParseJWT(tokenString string) (jwt.MapClaims, error) {
	tokenParse, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// Convert the secret key to []byte as jwt verification expects a []byte for HMAC
		return []byte(config.GetConfig().Auth.SecretKey), nil
	})

	var claims jwt.MapClaims

	if tokenParse != nil {
		var ok bool
		claims, ok = tokenParse.Claims.(jwt.MapClaims)
		if !ok {
			return nil, gerror.NewCode(consts.CodeInvalidToken)
		}
	}

	switch {
	case err == nil && tokenParse.Valid:
		return claims, nil
	case err != nil && (errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet)):
		return claims, jwt.ErrTokenExpired
	default:
		return nil, gerror.NewCode(consts.CodeInvalidToken)
	}
}
