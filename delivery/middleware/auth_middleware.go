package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jutionck/golang-upskilling-agt/utils/service"
)

type AuthMiddleware interface {
	RequireToken(userRole ...string) gin.HandlerFunc
	RefreshToken() gin.HandlerFunc
}

type authMiddleware struct {
	jwtService service.JwtService
}

type authHeader struct {
	AuthorizationHeader string `header:"Authorization"`
}

func (a *authMiddleware) RequireToken(userRole ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
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
		claims, err := a.jwtService.VerifyAccessToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			return
		}

		c.Set("claims", claims)

		// Pengecekan role
		validRole := false
		for _, role := range userRole {
			if role == claims["role"] {
				validRole = true
				break
			}
		}

		if !validRole {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Forbidden Resource"})
			return
		}
		c.Next()
	}
}

func (a *authMiddleware) RefreshToken() gin.HandlerFunc {
	panic("unimplemented")
}

func NewAuthMiddleware(jwtService service.JwtService) AuthMiddleware {
	return &authMiddleware{jwtService: jwtService}
}
