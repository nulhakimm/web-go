package config

import (
	"context"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

func NewDatabase() (*firestore.Client, error) {

	ctx := context.Background()

	sa := option.WithCredentialsFile("./serviceAccountKey.json")

	client, err := firestore.NewClient(ctx, "latihan-2-89ee6", sa)
	if err != nil {
		return nil, err
	}

	return client, nil

}
