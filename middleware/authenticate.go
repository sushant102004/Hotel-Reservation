package middleware

import (
	"fmt"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Authenticate(c *fiber.Ctx) error {
	authHeader, ok := c.GetReqHeaders()["Authorization"]

	if !ok {
		return c.JSON(map[string]string{
			"message": "unoauthorized",
		})
	}

	token := strings.Split(authHeader, " ")[1]

	if err := parseJWTToken(token); err != nil {
		fmt.Println("Failed to parse token. Error:", err)
		return err
	}

	return nil
}

func parseJWTToken(tokenStr string) error {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		secret := os.Getenv("JWT_SECRET")
		return []byte(secret), nil
	})

	if err != nil {
		return err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["foo"], claims["nbf"])
	}
	return fmt.Errorf("unauthorized")
}
