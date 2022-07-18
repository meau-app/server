package dao

import (
	"context"
	"fmt"

	"github.com/rafaelcn/meau/internal/config"
	"google.golang.org/api/iterator"
)

type Pet struct {
	Name     string                 `json:"name,omitempty"`
	Vaccines map[string]interface{} `json:"vaccines,omitempty"`
	Age      int64
	Race     string
	Adopted  bool `json:"adopted"`
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

	iter := client.Collection("pets").Where("adopted", "==", false).Documents(ctx)

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		} else if err != nil {
			return pets, err
		}
		fmt.Println(doc.Data())
	}

	return pets, nil
}

func SavePet(ctx context.Context, pet Pet) error {
	return nil
}
