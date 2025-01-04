package graphql

import (
	graphql_enum "backend/src/graphql/enum"
	"backend/src/graphql/modules/account"
	"backend/src/graphql/modules/authentication"
)

type Resolver struct {
	PermissionAbility graphql_enum.PermissionAbility
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
