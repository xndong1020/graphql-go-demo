package schemas

import (
	"acy.com/graphqlgodemo/database"
	"acy.com/graphqlgodemo/graph/types"
	"acy.com/graphqlgodemo/repositories"
	"acy.com/graphqlgodemo/services"
	"github.com/graphql-go/graphql"
)

// get query schemas
func GetQueryFields() graphql.Fields {
	var bookService = services.NewBookService(repositories.NewBookRepository(database.DbInstance))
	return graphql.Fields{
		"GetAllBooks": &graphql.Field{
			// Description explains the field
			Description: "Query all books",
			Type: graphql.NewList(types.BookType),
			// Resolve is the function used to look up data
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return bookService.GetAllBooks()
			},
		},
		"GetOneBook": &graphql.Field{
			// Description explains the field
			Description: "Query one book by id",
			Type: types.BookType,
			Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
			// Resolve is the function used to look up data
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return bookService.GetOneBook(params.Args["id"].(int))
			},
		},
	}
}