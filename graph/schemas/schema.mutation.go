package schemas

import (
	"acy.com/graphqlgodemo/graph/types"
	"acy.com/graphqlgodemo/models"
	"github.com/graphql-go/graphql"
)

var MutationFields = graphql.Fields{
			"CreateBook": &graphql.Field{
				Type:        types.BookType,
				Description: "Create a new Book",
				Args: graphql.FieldConfigArgument{
					"input": &graphql.ArgumentConfig{
						Type: types.BookInputType,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					input := params.Args["input"].(map[string]interface{})
					book := &models.Book{
						ID: 5,
						Title: input["title"].(string),
						Author: input["author"].(string),
						Publisher: input["publisher"].(string),
					}
					return book, nil
				},
			},
			"UpdateBook": &graphql.Field{
				Type:        graphql.String,
				Description: "Update an existing Book",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"input": &graphql.ArgumentConfig{
						Type: types.BookInputType,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					// input := params.Args["input"].(map[string]interface{})
					// newBook := &models.Book{
					// 	ID: 5,
					// 	Title: input["title"].(string),
					// 	Author: input["author"].(string),
					// 	Publisher: input["publisher"].(string),
					// }
					return "worked!", nil
				},
			},
			"DeleteBook": &graphql.Field{
				Type:        graphql.String,
				Description: "Delete an existing Book",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					// input := params.Args["input"].(map[string]interface{})
					// newBook := &models.Book{
					// 	ID: 5,
					// 	Title: input["title"].(string),
					// 	Author: input["author"].(string),
					// 	Publisher: input["publisher"].(string),
					// }
					return "worked!", nil
				},
			},
		}