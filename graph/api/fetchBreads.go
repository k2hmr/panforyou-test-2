package api

import (
	"context"
	"time"

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
			ID:        doc.Data()["Id"].(string),
			Name:      doc.Data()["Name"].(string),
			CreatedAt: doc.Data()["CreatedAt"].(time.Time),
		}
		breads = append(breads, bread)
	}

	defer client.Close()
	return breads, nil
}
