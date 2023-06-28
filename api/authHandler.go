package api

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sushant102004/Hotel-Reservation-System/db"
	"github.com/sushant102004/Hotel-Reservation-System/types"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	userStore db.UserStore
}

func NewAuthHandler(userStore db.UserStore) *AuthHandler {
	return &AuthHandler{
		userStore: userStore,
	}
}

type AuthParams struct {
	Email    string
	Password string
}

type AuthResponse struct {
	User  *types.User `json:"user"`
	Token string      `json:"token"`
}

func (h *AuthHandler) HandleAuthenticate(c *fiber.Ctx) error {
	var params AuthParams
	if err := c.BodyParser(&params); err != nil {
		return err
	}

	user, err := h.userStore.GetUserByEmail(c.Context(), params.Email)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.JSON(map[string]string{
				"error": "User not found.",
			})
		}
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.EncryptedPassword), []byte(params.Password))
	if err != nil {
		return c.JSON(map[string]string{
			"error": "Invalid password",
		})
	}

	token := createJWTToken(user)

	resp := AuthResponse{
		User:  user,
		Token: token,
	}

	return c.JSON(resp)
}

func createJWTToken(user *types.User) string {
	now := time.Now()
	validTill := now.Add(time.Hour * 6)

	claims := jwt.MapClaims{
		"id":        user.ID,
		"email":     user.Email,
		"validTill": validTill,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := []byte(os.Getenv("JWT_SECRET"))

	tokenStr, err := token.SignedString(secret)
	if err != nil {
		fmt.Println("Failed to sign token: ", err)
	}
	return tokenStr
}
