package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/k2hmr/panforyou-test-2/graph"
	"github.com/k2hmr/panforyou-test-2/graph/generated"
	"github.com/k2hmr/panforyou-test-2/graph/model"

	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	breads := []*model.Bread{
		{ ID: "1111111111111111111111", Name: "パン1", CreatedAt: "2023-07-06T08:32:22.090Z" },
		{ ID: "1111111111111111111112", Name: "パン2", CreatedAt: "2023-07-06T08:32:22.090Z" },
		{ ID: "1111111111111111111113", Name: "パン3", CreatedAt: "2023-07-06T08:32:22.090Z" },
	}

	rsv := graph.NewResolver(breads)

	gqlConfig := generated.Config{Resolvers: rsv}
	gqlHandler := handler.NewDefaultServer(generated.NewExecutableSchema(gqlConfig))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", gqlHandler)
	

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
