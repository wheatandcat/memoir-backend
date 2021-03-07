package main

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/wheatandcat/memoir-backend/auth"
	"github.com/wheatandcat/memoir-backend/graph"
	"github.com/wheatandcat/memoir-backend/graph/generated"
)

func main() {
	router := chi.NewRouter()

	router.Use(auth.NotLoginMiddleware())

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	router.Handle("/", playground.Handler("Starwars", "/query"))
	router.Handle("/query", srv)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}
}
