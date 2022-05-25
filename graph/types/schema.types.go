package types

import "github.com/graphql-go/graphql"

var BookType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Book",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"title": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"author": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"publisher": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
	},
)

var BookInputType = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "BookInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"title": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"author": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"publisher": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
	},
) 