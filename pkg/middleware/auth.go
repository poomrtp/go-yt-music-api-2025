package middleware

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"

	"github.com/poomrtp/go-yt-music/pkg/utils"
)

type TokenMetadata struct {
	UserID    string
	ExpiresAt time.Time
}

type Claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

var (
	JWTSecretKey []byte
)

func init() {
	JWTSecretKey = []byte(getJWTSecret())
}

func getJWTSecret() string {
	if os.Getenv("JWT_SECRET_KEY") == "" {
		if err := godotenv.Load(); err != nil {
			fmt.Printf("[ERROR] Error loading .env file: %v\n", err)
		}
	}

	secret := os.Getenv("JWT_SECRET_KEY")
	if secret == "" {
		fmt.Printf("[ERROR] JWT_SECRET_KEY environment variable is not set")
	}
	return secret
}

func Protected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenString := extractToken(c)
		if tokenString == "" {
			return utils.UnauthorizedResponse(c, "No authorization token provided", nil)
		}
		claims, err := verifyToken(tokenString)
		if err != nil {
			return utils.UnauthorizedResponse(c, "Invalid or expired token", err)
		}

		c.Locals("user", claims)
		return c.Next()
	}
}

func extractToken(c *fiber.Ctx) string {
	bearerToken := c.Get("Authorization")

	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}

	return ""
}

func verifyToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return JWTSecretKey, nil
	})

	if err != nil {
		fmt.Printf("Token verification error: %v\n", err)
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}
