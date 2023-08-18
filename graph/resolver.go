package graph

import "github.com/k2hmr/panforyou-test-2/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	breads []*model.Bread
}

func NewResolver(breads []*model.Bread) *Resolver {
	return &Resolver{
		breads: breads,
	}
}
