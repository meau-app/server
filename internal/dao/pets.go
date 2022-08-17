package dao

import (
	"context"
	"crypto/sha256"
	"encoding/hex"

	"github.com/labstack/echo/v4"
	"google.golang.org/api/iterator"

	"github.com/meau-app/server/internal/config"
)

type Pet struct {
	Name     string                 `json:"name"`
	Vaccines map[string]interface{} `json:"vaccines"`
	Age      int64                  `json:"age"`
	Species  string                 `json:"species"`
	Adopted  bool                   `json:"adopted"`
	Pictures []string               `json:"pictures"` /* base 64 images */
	Diseases []string               `json:"diseases"`
	Temper   string                 `json:"temper"`
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
	doc.DataTo(&pet)

	return pet, nil
}

func GetPets(ctx context.Context) ([]Pet, error) {
	pets := []Pet{}

	client, err := config.Database.Firestore(ctx)
	if err != nil {
		return pets, err
	}

	logger := ctx.Value(ContextLoggerKey).(echo.Logger)

	//it := client.Collection("pets").Where("adopted", "==", false).Documents(ctx)
	it := client.Collection("pets").Documents(ctx)

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

func SavePet(ctx context.Context, pet *Pet) error {
	client, err := config.Database.Firestore(ctx)
	if err != nil {
		return err
	}

	logger := ctx.Value(ContextLoggerKey).(echo.Logger)

	hasher := sha256.New().Sum([]byte(pet.Name))
	uuid := hex.EncodeToString(hasher[:10])

	_, err = client.Collection("pets").Doc(uuid).Set(ctx, pet)
	if err != nil {
		logger.Errorf("failed to save a pet, reason %v", err)
		return err
	}

	return nil
}

func DeletePet(ctx context.Context, id string) error {
	client, err := config.Database.Firestore(ctx)
	if err != nil {
		return err
	}

	logger := ctx.Value(ContextLoggerKey).(echo.Logger)

	_, err = client.Collection("pets").Doc(id).Delete(ctx)
	if err != nil {
		logger.Errorf("failed to delete pet with id %s, reason %v", id, err)
		return err
	}

	return nil
}
