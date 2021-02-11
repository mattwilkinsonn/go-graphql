package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Zireael13/go-graphql/graph"
	"github.com/Zireael13/go-graphql/graph/generated"
	"github.com/Zireael13/go-graphql/internal/db"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	sequel, _ := sql.Open("pgx", "postgresql://postgres:postgres@localhost:5432/gograph")

	base := db.New(sequel)

	user := &db.CreateUserParams{
		Username: "matt",
		Email:    "mattwilki17@gmail.com",
		Password: "hunter2",
	}

	base.CreateUser(context.TODO(), *user)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
