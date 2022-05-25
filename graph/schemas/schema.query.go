package schemas

import (
	"acy.com/graphqlgodemo/graph"
	"acy.com/graphqlgodemo/graph/types"
	"github.com/graphql-go/graphql"
)

// get query schemas
func GetQueryFields(r *graph.Resolver) graphql.Fields {
	return graphql.Fields{
		"GetAllBooks": &graphql.Field{
			// Description explains the field
			Description: "Query all books",
			Type: graphql.NewNonNull(graphql.NewList(types.BookType)),
			// Resolve is the function used to look up data
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return r.BookService.GetAllBooks()
			},
		},
		"GetOneBook": &graphql.Field{
			// Description explains the field
			Description: "Query one book by id",
			Type: graphql.NewNonNull(types.BookType),
			Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
			// Resolve is the function used to look up data
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return r.BookService.GetOneBook(params.Args["id"].(int))
			},
		},
	}
}