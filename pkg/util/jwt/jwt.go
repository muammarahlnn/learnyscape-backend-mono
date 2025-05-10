package jwtutil

import (
	"errors"
	"learnyscape-backend-mono/pkg/config"
	"learnyscape-backend-mono/pkg/httperror"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWTUtil interface {
	SignAccess(payload *JWTPayload) (string, error)
	SignRefresh(payload *JWTPayload) (string, error)
	ParseAccess(token string) (*JWTClaims, error)
	ParseRefresh(token string) (*JWTClaims, error)
}

type JWTPayload struct {
	UserID int64
	Role   string
}

type JWTClaims struct {
	jwt.RegisteredClaims
	UserID int64  `json:"user_id"`
	Role   string `json:"role"`
}

type jwtUtil struct {
	config *config.JWTConfig
}

func NewJWTUtil() JWTUtil {
	return &jwtUtil{
		config: config.JwtConfig,
	}
}

func (j *jwtUtil) SignAccess(payload *JWTPayload) (string, error) {
	duration := time.Now().Add(time.Duration(j.config.AccessTokenDuration) * time.Minute)
	signed, err := j.sign(payload, duration, j.config.AccessSecretKey)
	if err != nil {
		return "", err
	}

	return signed, nil
}

func (j *jwtUtil) SignRefresh(payload *JWTPayload) (string, error) {
	duration := time.Now().Add(time.Duration(j.config.RefreshTokenDuration) * time.Minute)
	signed, err := j.sign(payload, duration, j.config.RefreshSecretKey)
	if err != nil {
		return "", err
	}

	return signed, nil
}

func (j *jwtUtil) sign(payload *JWTPayload, duration time.Time, key string) (string, error) {
	currentTime := time.Now()

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		JWTClaims{
			RegisteredClaims: jwt.RegisteredClaims{
				ID:        uuid.NewString(),
				Issuer:    j.config.Issuer,
				IssuedAt:  jwt.NewNumericDate(currentTime),
				ExpiresAt: jwt.NewNumericDate(duration),
			},
			UserID: payload.UserID,
			Role:   payload.Role,
		},
	)

	signedStr, err := token.SignedString([]byte(key))
	if err != nil {
		return "", err
	}

	return signedStr, nil
}

func (j *jwtUtil) ParseAccess(token string) (*JWTClaims, error) {
	claims, err := j.parse(token, j.config.AccessSecretKey)
	if err != nil {
		return nil, err
	}

	return claims, nil
}

func (j *jwtUtil) ParseRefresh(token string) (*JWTClaims, error) {
	claims, err := j.parse(token, j.config.RefreshSecretKey)
	if err != nil {
		return nil, err
	}

	return claims, nil
}

func (j *jwtUtil) parse(token, key string) (*JWTClaims, error) {
	parser := jwt.NewParser(
		jwt.WithValidMethods(j.config.AllowedAlgs),
		jwt.WithIssuer(j.config.Issuer),
		jwt.WithIssuedAt(),
		jwt.WithExpirationRequired(),
	)

	claims, err := j.parseClaims(parser, token, key)
	if err != nil {
		return nil, err
	}

	return claims, nil
}

func (j *jwtUtil) parseClaims(parser *jwt.Parser, token, key string) (*JWTClaims, error) {
	parsedToken, err := parser.ParseWithClaims(
		token,
		&JWTClaims{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(key), nil
		},
	)
	if err != nil || !parsedToken.Valid {
		return nil, err
	}

	claims, ok := parsedToken.Claims.(*JWTClaims)
	if !ok {
		return nil, httperror.NewUnauthorizedError()
	}

	currentTime := time.Now()
	if claims.ExpiresAt.Time.Before(currentTime) {
		return nil, errors.New("token expired")
	}
	if claims.Issuer != j.config.Issuer {
		return nil, errors.New("invalid issuer")
	}
	if claims.IssuedAt.Time.After(currentTime) {
		return nil, errors.New("token not yet valid")
	}
	if claims.UserID == 0 {
		return nil, errors.New("invalid user id")
	}
	if claims.Role == "" {
		return nil, errors.New("invalid role")
	}

	return claims, nil
}
