package middleware

import (
	"fmt"
	"os"
	"strings"
	"time"

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

	claims, err := validateToken(token)
	if err != nil {
		return err
	}

	validTime := claims["validTill"].(float64)

	expireTime := int64(validTime)

	if time.Now().Unix() > expireTime {
		return c.JSON(map[string]string{
			"error": "Token Expried",
		})
	}

	return c.Next()
}

func validateToken(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		secret := os.Getenv("JWT_SECRET")
		return []byte(secret), nil
	})

	if err != nil {
		fmt.Println("Failed to parse token: ", err)
		return nil, err
	}

	if !token.Valid {
		fmt.Println("invalid token")
		return nil, fmt.Errorf("unauthorized")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("unauthorized")

}
