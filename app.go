package main

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jutionck/golang-upskilling-agt/model"
)

var ApplicationName = "USPKILLINGGO"
var JwtSigningMethod = jwt.SigningMethodHS256
var JwtSignatureKey = []byte("JozzG4nD0ZZ!!!")

type MyClaim struct {
	jwt.RegisteredClaims
	Username string `json:"username"`
	Role     string `json:"role"`
}

type authHeader struct {
	AuthorizationHeader string `header:"Authorization"`
}

// middleware
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path == "/api/v1/auth/login" {
			c.Next()
		} else {
			var aH authHeader
			if err := c.ShouldBindHeader(&aH); err != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
				return
			}

			// kita akan dapatkan tokennya dari Authorization Header
			// Replace => Bearer token menjadi token ==== menggunakan strings.Replace
			tokenString := strings.Replace(aH.AuthorizationHeader, "Bearer ", "", -1)
			if tokenString == "" {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
				return
			}

			// Verifikasi dulu apakah token nya valid atau tidak
			claims, err := VerifyAccessToken(tokenString)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
				return
			}

			c.Set("claims", claims)
			fmt.Println("claims:", claims)

			c.Next()
		}
	}
}

func CreateAccessToken(user model.User) (string, error) {
	now := time.Now().UTC()
	end := now.Add(5 * time.Minute)
	claims := MyClaim{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    ApplicationName,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(end),
		},
		Username: user.Username,
		Role:     user.Role,
	}

	jwtNewClaim := jwt.NewWithClaims(JwtSigningMethod, claims)
	token, err := jwtNewClaim.SignedString(JwtSignatureKey)
	if err != nil {
		return "", errors.New("failed to create access token")
	}

	return token, nil
}

func VerifyAccessToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if method, ok := t.Method.(*jwt.SigningMethodHMAC); !ok || method != JwtSigningMethod {
			return nil, errors.New("signing method invalid")
		}
		return JwtSignatureKey, nil
	})

	if err != nil {
		return nil, errors.New("failed to verify access token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid || claims["iss"] != ApplicationName {
		return nil, errors.New("invalid token claims")
	}

	return claims, nil
}

func main() {
	// delivery.NewServer().Run()

	r := gin.Default()
	rg := r.Group("/api/v1")
	rg.Use(AuthMiddleware())
	rg.POST("/auth/login", func(c *gin.Context) {
		var userCredential model.User
		if err := c.ShouldBindJSON(&userCredential); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		if userCredential.Username == "admin" && userCredential.Password == "password" {
			userCredential.Role = "admin"
			token, err := CreateAccessToken(userCredential)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
				return
			}
			c.JSON(http.StatusCreated, gin.H{"data": token})
		}
	})

	rg.GET("/users", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Ok"})
	})

	r.Run(":8888")
}
