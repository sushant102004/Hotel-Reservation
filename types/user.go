package types

import (
	"fmt"
	"regexp"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

var (
	fNameLen    = 3
	lNameLen    = 3
	passwordLen = 8
)

type CreateUserParams struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type UpdateUserParams struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type User struct {
	/*
		Due to bson ID will be stored as _id in MongoDB
		omitempty will not show ID in JSON Response.
	*/
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName         string             `bson:"firstName" json:"firstName"`
	LastName          string             `bson:"lastName" json:"lastName"`
	Email             string             `bson:"email" json:"email"`
	EncryptedPassword string             `bson:"EncryptedPassword" json:"-"`
}

func NewUserFromParams(params CreateUserParams) (*User, error) {
	encryptedPw, err := bcrypt.GenerateFromPassword([]byte(params.Password), 12)
	if err != nil {
		return nil, err
	}

	return &User{
		FirstName:         params.FirstName,
		LastName:          params.LastName,
		Email:             params.Email,
		EncryptedPassword: string(encryptedPw),
	}, err
}

func (params UpdateUserParams) ToBSON() bson.M {
	m := bson.M{}

	if len(params.FirstName) > 0 && len(params.LastName) > 0 {
		m["firstName"] = params.FirstName
		m["lastName"] = params.LastName
	}
	return m
}

func (params CreateUserParams) Validate() error {
	if len(params.FirstName) < fNameLen {
		return fmt.Errorf("first Name must be %d characters long", fNameLen)
	}

	if len(params.LastName) < lNameLen {
		return fmt.Errorf("last Name must be %d characters long", lNameLen)
	}

	if len(params.Password) < passwordLen {
		return fmt.Errorf("password Name must be %d characters long", passwordLen)
	}

	if !isEmailValid(params.Email) {
		return fmt.Errorf("email is not valid")
	}

	return nil
}

func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(e)
}
