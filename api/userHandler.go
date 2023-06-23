package api

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/sushant102004/Hotel-Reservation-System/db"
	"github.com/sushant102004/Hotel-Reservation-System/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserHandler struct {
	userStore db.UserStore
}

func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{userStore: userStore}
}

func (h *UserHandler) HandleGetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := h.userStore.GetUserByID(c.Context(), id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.JSON(map[string]string{
				"message": "User Not Found",
			})
		}

		return err
	}
	return c.JSON(user)
}

func (h *UserHandler) HandleGetUsers(c *fiber.Ctx) error {
	users, err := h.userStore.GetUsers(c.Context())
	if err != nil {
		return err
	}
	return c.JSON(users)
}

func (h *UserHandler) HandleInsertUser(c *fiber.Ctx) error {
	var params types.CreateUserParams

	err := c.BodyParser(&params)
	if err != nil {
		return err
	}

	err = params.Validate()
	if err != nil {
		return err
	}

	user, err := types.NewUserFromParams(params)
	if err != nil {
		return err
	}

	insertedUser, err := h.userStore.InsertUser(c.Context(), user)
	if err != nil {
		return err
	}
	return c.JSON(insertedUser)
}

func (h *UserHandler) HandleDeleteUser(c *fiber.Ctx) error {
	userID := c.Params("id")
	err := h.userStore.DeleteUser(c.Context(), userID)
	if err != nil {
		return err
	}
	return c.JSON(map[string]string{
		"message": "User Deleted Successfully",
	})
}

func (h *UserHandler) HandleUpdateUser(c *fiber.Ctx) error {
	var (
		updateRequest types.UpdateUserParams
		userID        = c.Params("id")
	)
	oID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}
	if err := c.BodyParser(&updateRequest); err != nil {
		return err
	}
	if err := h.userStore.UpdateUser(c.Context(), bson.M{"_id": oID}, bson.M{"$set": updateRequest}); err != nil {
		return err
	}

	return c.JSON(map[string]string{
		"message": "User Updated Successfully",
	})
}
