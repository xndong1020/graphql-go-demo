package schemas

import (
	"fmt"

	"acy.com/graphqlgodemo/dtos"
	"acy.com/graphqlgodemo/graph"
	"acy.com/graphqlgodemo/graph/types"
	"github.com/graphql-go/graphql"
)

func GetMutationFields(r *graph.Resolver) graphql.Fields {
	return graphql.Fields{
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
					bookInput := &dtos.BookInput{
						Title: input["title"].(string),
						Author: input["author"].(string),
						Publisher: input["publisher"].(string),
					}
					return r.BookService.CreateBook(bookInput)
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
					id := params.Args["id"].(int)
					input := params.Args["input"].(map[string]interface{})
					bookInput := &dtos.BookInput{
						Title: input["title"].(string),
						Author: input["author"].(string),
						Publisher: input["publisher"].(string),
					}
					err := r.BookService.UpdateBook( bookInput, id)
					if err != nil {
						return nil, err
					}
					return fmt.Sprintf("Book with id:%d has been updated successfully", id), nil
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
					id := params.Args["id"].(int)
					err := r.BookService.DeleteBook(id)
					if err != nil {
						return nil, err
					}
					return fmt.Sprintf("Book with id:%d has been deleted successfully", id), nil
				},
			},
		}
}