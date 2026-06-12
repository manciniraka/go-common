package jwt

import (
	jwtgo "github.com/golang-jwt/jwt/v5"
)

func GenerateToken(
	claims jwtgo.MapClaims,
	secret string,
) (string, error) {
	token := jwtgo.NewWithClaims(
		jwtgo.SigningMethodHS256,
		claims,
	)

	return token.SignedString(
		[]byte(secret),
	)
}

func ParseToken(
	tokenString string,
	secret string,
) (jwtgo.MapClaims, error) {
	token, err := jwtgo.Parse(
		tokenString,
		func(token *jwtgo.Token) (interface{}, error) {
			return []byte(secret), nil
		},
	)
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, jwtgo.ErrTokenSignatureInvalid
	}

	claims, ok := token.Claims.(jwtgo.MapClaims)
	if !ok {
		return nil, jwtgo.ErrTokenInvalidClaims
	}

	return claims, nil
}
