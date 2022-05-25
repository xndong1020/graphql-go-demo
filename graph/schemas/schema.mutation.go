package schemas

import (
	"fmt"

	"acy.com/graphqlgodemo/database"
	"acy.com/graphqlgodemo/dtos"
	"acy.com/graphqlgodemo/graph/types"
	"acy.com/graphqlgodemo/repositories"
	"acy.com/graphqlgodemo/services"
	"github.com/graphql-go/graphql"
)

func GetMutationFields() graphql.Fields {
	bookService := services.NewBookService(repositories.NewBookRepository(database.DbInstance))
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
					return bookService.CreateBook(bookInput)
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
					err := bookService.UpdateBook( bookInput, id)
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
					err := bookService.DeleteBook(id)
					if err != nil {
						return nil, err
					}
					return fmt.Sprintf("Book with id:%d has been deleted successfully", id), nil
				},
			},
		}
}