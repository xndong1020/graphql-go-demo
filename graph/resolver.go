package graph

import "acy.com/graphqlgodemo/services"

type Resolver struct{
	BookService services.IBookService
}