package account

import (
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

func (r AccountResolver) Profile(ctx context.Context) Account {
	authorizationCtx := authentication.GetAuthorizationContext(ctx)
	accountProfile, _ := r.accountService.GetById(uuid.MustParse(authorizationCtx.AccountId))
	return *accountProfile
}
