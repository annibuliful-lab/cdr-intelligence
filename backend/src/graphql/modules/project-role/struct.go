package projectrole

import "github.com/graph-gophers/graphql-go"

type ProjectRole struct {
	Id        graphql.ID
	projectId graphql.ID
	title     string
}
