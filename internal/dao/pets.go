package dao

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/rafaelcn/meau/internal/config"
	"google.golang.org/api/iterator"
)

type Pet struct {
	Name     string                 `json:"name"`
	Vaccines map[string]interface{} `json:"vaccines"`
	Age      int64                  `json:"age"`
	Race     string                 `json:"race"`
	Adopted  bool                   `json:"adopted"`
}

func GetPet(ctx context.Context, id string) (Pet, error) {
	client, err := config.Database.Firestore(ctx)
	if err != nil {
		return Pet{}, err
	}

	doc, err := client.Collection("pets").Doc(id).Get(ctx)
	if err != nil {
		return Pet{}, err
	}

	pet := Pet{}
	doc.DataTo(pet)

	return pet, nil
}

func GetPets(ctx context.Context) ([]Pet, error) {
	pets := []Pet{}

	client, err := config.Database.Firestore(ctx)
	if err != nil {
		return pets, err
	}

	logger := ctx.Value(ContextLoggerKey).(echo.Logger)

	it := client.Collection("pets").Where("adopted", "==", false).Documents(ctx)

	for {
		doc, err := it.Next()
		if err == iterator.Done {
			break
		} else if err != nil {
			return pets, err
		}

		pet := Pet{}

		if err = doc.DataTo(&pet); err != nil {
			logger.Errorf("failed to convert document, reason %v", err)
		}

		pets = append(pets, pet)
	}

	return pets, nil
}

func SavePet(ctx context.Context, pet Pet) error {
	return nil
}
