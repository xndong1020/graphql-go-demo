package main

import (
	"log"
	"net/http"
	"os"

	"acy.com/graphqlgodemo/database"
	"acy.com/graphqlgodemo/graph"
	"acy.com/graphqlgodemo/graph/schemas"
	"acy.com/graphqlgodemo/middlewares"
	"acy.com/graphqlgodemo/repositories"
	"acy.com/graphqlgodemo/services"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"

	"github.com/joho/godotenv"
)

const defaultPort = "8080"

func main() {

	godotenv.Load()

	config := &database.Config{
        Host:     os.Getenv("DB_HOST"),
        Port:     os.Getenv("DB_PORT"),
        Password: os.Getenv("DB_PASS"),
        User:     os.Getenv("DB_USER"),
        SSLMode:  os.Getenv("DB_SSLMODE"),
        DBName:   os.Getenv("DB_NAME"),
		Schema:   os.Getenv("DB_SCHEMA"),
    }

	db, err := database.NewConnection(config)
	
	if err != nil {
		panic(err)
	}

	resolver := &graph.Resolver{BookService: services.NewBookService(repositories.NewBookRepository(db))}

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: schemas.GetQueryFields(resolver)}
	rootMutation := graphql.ObjectConfig{Name: "rootMutation",Fields: schemas.GetMutationFields(resolver)}

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

	http.Handle("/graphql", middlewares.CorsMiddleware(h))
	log.Printf("connect to http://localhost:%s/graphql for GraphQL playground", defaultPort)
	log.Fatal(http.ListenAndServe(":"+defaultPort, nil))
}
