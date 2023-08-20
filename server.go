package main

import (
	"context"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/k2hmr/panforyou-test-2/graph"
	"github.com/k2hmr/panforyou-test-2/graph/generated"

	"github.com/99designs/gqlgen/graphql/playground"
)

const port = "8080"

func main() {
	ctx := context.Background()
	rsv, err := graph.NewResolver(ctx)
	if err != nil {
		log.Fatal(err)
	}

	gqlConfig := generated.Config{Resolvers: rsv}
	gqlHandler := handler.NewDefaultServer(generated.NewExecutableSchema(gqlConfig))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", gqlHandler)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
