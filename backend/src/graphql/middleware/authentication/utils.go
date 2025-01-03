package authentication

import (
	"context"
	"strings"
)

func GetAuthToken(authorization string) string {
	return strings.Replace(authorization, "Bearer ", "", 1)
}

func GetAuthorizationContext(ctx context.Context) AuthorizationContext {
	authContext := AuthorizationContext{}

	if ctx.Value("accountId") != nil {
		authContext.AccountId = ctx.Value("accountId").(string)
	}

	if ctx.Value("projectId") != nil {
		authContext.ProjectId = ctx.Value("projectId").(string)
	}

	if ctx.Value("token") != nil {
		authContext.Token = ctx.Value("token").(string)
	}
	return authContext
}
