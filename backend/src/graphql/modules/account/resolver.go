package account

import (
	error_utils "backend/src/error"
	"backend/src/graphql/middleware/authentication"
	"context"

	"github.com/google/uuid"
)

type AccountResolver struct {
	accountService AccountService
}

func NewAccountResolver(params NewAccountParams) AccountResolver {

	return AccountResolver{
		accountService: NewAccountService(params),
	}
}

func (r AccountResolver) Profile(ctx context.Context) (Account, error) {
	authorizationCtx := authentication.GetAuthorizationContext(ctx)
	accountProfile, err := r.accountService.GetById(uuid.MustParse(authorizationCtx.AccountId))

	if err != nil {
		return Account{}, error_utils.GraphqlError{
			Message: err.Error(),
		}
	}
	return accountProfile, nil
}

func (r AccountResolver) Register(ctx context.Context, input RegisterInput) (Account, error) {
	account, err := r.accountService.Create(CreateAccountData{
		Username: input.Username,
		Password: input.Password,
	})

	if err != nil {
		return Account{}, error_utils.GraphqlError{
			Message: err.Error(),
		}
	}

	return account, nil
}
