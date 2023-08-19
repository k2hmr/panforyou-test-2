package api

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/k2hmr/panforyou-test-2/graph/model"
)

func FetchBreads(client *firestore.Client, ctx context.Context) (breads []*model.Bread, err error) {
	docs := client.Collection("breads").Documents(ctx)
	for {
		doc, err := docs.Next()
		if err != nil {
			break
		}
		bread := &model.Bread{
			ID: doc.Data()["id"].(string),
			Name: doc.Data()["name"].(string),
			CreatedAt: doc.Data()["createdAt"].(string),
		}
		breads = append(breads, bread)
	}

  defer client.Close()
	return breads, nil
}