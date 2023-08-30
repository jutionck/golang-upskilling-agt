package service

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jutionck/golang-upskilling-agt/config"
	"github.com/jutionck/golang-upskilling-agt/model"
	modelutil "github.com/jutionck/golang-upskilling-agt/utils/model_util"
)

type JwtService interface {
	CreateAccessToken(credential model.User) (string, error)
	VerifyAccessToken(tokenString string) (jwt.MapClaims, error)
}

type jwtService struct {
	cfg config.JwtConfig
}

func (j *jwtService) CreateAccessToken(credential model.User) (string, error) {
	now := time.Now().UTC()
	end := now.Add(j.cfg.JwtLifeTime)
	claims := modelutil.MyClaim{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    j.cfg.ApplicationName,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(end),
		},
		Username: credential.Username,
		Role:     credential.Role,
	}

	jwtNewClaim := jwt.NewWithClaims(j.cfg.JwtSigningMethod, claims)
	token, err := jwtNewClaim.SignedString(j.cfg.JwtSignatureKey)
	if err != nil {
		return "", errors.New("failed to create access token")
	}

	return token, nil
}

func (j *jwtService) VerifyAccessToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if method, ok := t.Method.(*jwt.SigningMethodHMAC); !ok || method != j.cfg.JwtSigningMethod {
			return nil, errors.New("signing method invalid")
		}
		return j.cfg.JwtSignatureKey, nil
	})

	if err != nil {
		return nil, errors.New("failed to verify access token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid || claims["iss"] != j.cfg.ApplicationName {
		return nil, errors.New("invalid token claims")
	}

	return claims, nil
}

func NewJwtService(cfg config.JwtConfig) JwtService {
	return &jwtService{cfg: cfg}
}
