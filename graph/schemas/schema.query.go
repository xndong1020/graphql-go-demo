package schemas

import (
	"acy.com/graphqlgodemo/graph/types"
	"acy.com/graphqlgodemo/models"
	"github.com/graphql-go/graphql"
)

// query schemas
var QueryFields = graphql.Fields{
		"GetAllBooks": &graphql.Field{
			// Description explains the field
			Description: "Query all books",
			Type: graphql.NewList(types.BookType),
			// Resolve is the function used to look up data
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				allBooks := []*models.Book{}
				book := &models.Book{
					ID: 5,
                    Title: "test 01",
					Author: "test 01",
					Publisher: "test 01",
                }
				allBooks = append(allBooks, book)
				return allBooks, nil
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
				book := &models.Book{
					ID: params.Args["id"].(int),
                    Title: "test 01",
					Author: "test 01",
					Publisher: "test 01",
                }
				return book, nil
			},
		},
	}