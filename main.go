package main

import (
	"log"
	"net/http"

	"acy.com/graphqlgodemo/graph/schemas"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

const defaultPort = "8080"

func main() {
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: schemas.QueryFields}
	rootMutation := graphql.ObjectConfig{Name: "rootMutation",Fields: schemas.MutationFields}

	// Now combine into a Schema Configuration
	schemaConfig := graphql.SchemaConfig{
		Query: graphql.NewObject(rootQuery),
		Mutation: graphql.NewObject(rootMutation),
	}

	// Create a new GraphQL Schema
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	h := handler.New(&handler.Config{
		Schema: &schema,
		// Pretty print JSON response
		Pretty: true,
		// Host a GraphiQL Playground to use for testing Queries
		GraphiQL:   false,
		Playground: true,
	})

	http.Handle("/graphql", h)
	log.Printf("connect to http://localhost:%s/graphql for GraphQL playground", defaultPort)
	log.Fatal(http.ListenAndServe(":"+defaultPort, nil))
}
