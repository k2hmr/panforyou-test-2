package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/k2hmr/panforyou-test-2/graph/generated"
	"github.com/k2hmr/panforyou-test-2/graph/model"
)

func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

func (r *Resolver) Breads(ctx context.Context) ([]*model.Bread, error) {
	return r.breads, nil
}

func (r *Resolver) Bread(ctx context.Context, id string) (*model.Bread, error) {
	for _, bread := range r.breads {
		if bread.ID == id {
			return bread, nil
		}
	}
	return nil, fmt.Errorf("データが見つかりませんでした。")
}

type queryResolver struct{ *Resolver }