package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/k2hmr/panforyou-test-2/graph/api"
	"github.com/k2hmr/panforyou-test-2/graph/generated"
	"github.com/k2hmr/panforyou-test-2/graph/model"
)

func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

func (r *Resolver) Breads(ctx context.Context) ([]*model.Bread, error) {
	breads, err := api.FetchBreads(r.firestoreClient, ctx)
	if err != nil {
		return nil, err
	}
	return breads, nil
}

func (r *Resolver) Bread(ctx context.Context, id string) (*model.Bread, error) {
	breads, err := api.FetchBreads(r.firestoreClient, ctx)
	if err != nil {
		return nil, err
	}
	for _, bread := range breads {
		if bread.ID == id {
			return bread, nil
		}
	}
	return nil, fmt.Errorf("データが見つかりませんでした。")
}

type queryResolver struct{ *Resolver }
