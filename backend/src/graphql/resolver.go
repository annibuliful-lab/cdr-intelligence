package graphql

import "backend/src/graphql/modules/account"

type Resolver struct {
	account.AccountResolver
}

func (r Resolver) Hello() string {
	return "hello"
}

func GraphqlResolver() *Resolver {
	r := &Resolver{}

	return r
}
