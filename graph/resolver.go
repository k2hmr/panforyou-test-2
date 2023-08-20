package graph

import (
	"context"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	firestoreClient *firestore.Client
}

func NewResolver(ctx context.Context) (*Resolver, error) {
	sa := option.WithCredentialsFile("path/to/serviceAccount.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		return nil, err
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, err
	}

	return &Resolver{
		firestoreClient: client,
	}, nil
}
