package graphql

import (
	"backend/src/graphql/modules/account"
	"backend/src/graphql/modules/authentication"
)

type Resolver struct {
	account.AccountResolver
	authentication.AuthenticationResolver
}

func GraphqlResolver(params GraphqlResolverParams) *Resolver {
	r := &Resolver{}
	r.AccountResolver = account.NewAccountResolver(account.NewAccountParams{
		Db:    params.Db,
		Redis: params.Redis,
	})

	r.AuthenticationResolver = authentication.NewAuthenticationResolver(authentication.NewAuthenticationResolverParams{
		Db:    params.Db,
		Redis: params.Redis,
	})

	return r
}
