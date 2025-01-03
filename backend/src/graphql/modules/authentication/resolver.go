package authentication

import (
	error_utils "backend/src/error"
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

func (r AuthenticationResolver) Login(ctx context.Context, input LoginInput) (Authentication, error) {
	authentication, err := r.authenticationService.Login(LoginData{
		Username: input.Username,
		Password: input.Password,
	})
	if err != nil {
		return Authentication{}, error_utils.GraphqlError{
			Message: err.Error(),
		}
	}

	return authentication, nil
}
