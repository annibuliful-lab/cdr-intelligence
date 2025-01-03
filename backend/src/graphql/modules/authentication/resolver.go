package authentication

import (
	"context"
)

type AuthenticationResolver struct {
	authenticationService AuthenticationService
}

func NewAuthenticationResolver(params NewAuthenticationResolverParams) AuthenticationResolver {

	return AuthenticationResolver{
		authenticationService: NewAuthenticationService(params),
	}
}

func (AuthenticationResolver) Login(ctx context.Context, input LoginInput) Authentication {
	return Authentication{}
}
