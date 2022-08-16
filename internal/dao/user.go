package dao

import (
	"context"

	"github.com/labstack/echo"
	"google.golang.org/api/iterator"

	"github.com/meau-app/server/internal/config"
)

type User struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	City      string `json:"city"`
	CityState string `json:"state"`
	Address   string `json:"address"`
	Age       int64  `json:"age"`
	Gender    string `json:"gender"`

	ProfileImage string `json:"profile_image"` /* a base 64 image*/
	Pets         []Pet  `json:"pets"`
}

func GetUser(ctx context.Context, id string) (User, error) {
	client, err := config.Database.Firestore(ctx)
	if err != nil {
		return User{}, err
	}

	doc, err := client.Collection("users").Doc(id).Get(ctx)
	if err != nil {
		return User{}, err
	}

	user := User{}
	doc.DataTo(&user)

	return user, nil
}

func GetUsers(ctx context.Context) ([]User, error) {
	users := []User{}

	client, err := config.Database.Firestore(ctx)
	if err != nil {
		return users, err
	}

	logger := ctx.Value(ContextLoggerKey).(echo.Logger)

	it := client.Collection("users").Documents(ctx)

	for {
		doc, err := it.Next()
		if err == iterator.Done {
			break
		} else if err != nil {
			return users, err
		}

		user := User{}

		if err = doc.DataTo(&user); err != nil {
			logger.Errorf("failed to convert document, reason %v", err)
		}

		users = append(users, user)
	}

	return users, nil
}

func SaveUser(ctx context.Context, user *User) error {
	client, err := config.Database.Firestore(ctx)
	if err != nil {
		return err
	}

	logger := ctx.Value(ContextLoggerKey).(echo.Logger)

	_, err = client.Collection("users").Doc(user.Email).Set(ctx, user)
	if err != nil {
		logger.Errorf("failed to save a user, reason %v", err)
		return err
	}

	return nil
}
