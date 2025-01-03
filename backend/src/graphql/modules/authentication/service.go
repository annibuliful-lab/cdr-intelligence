package authentication

import (
	"database/sql"
	"errors"

	"backend/src/.gen/cdr-intelligence/public/model"
	"backend/src/.gen/cdr-intelligence/public/table"
	error_utils "backend/src/error"
	"backend/src/jwt"
	"backend/src/utils"

	pg "github.com/go-jet/jet/v2/postgres"
	"github.com/graph-gophers/graphql-go"
	"github.com/redis/go-redis/v9"
	"github.com/thanhpk/randstr"
)

type AuthenticationService struct {
	Db    *sql.DB
	redis *redis.Client
}

func NewAuthenticationService(params NewAuthenticationResolverParams) AuthenticationService {
	return AuthenticationService{
		Db:    params.Db,
		redis: params.redis,
	}
}

func (this AuthenticationService) Login(input LoginData) (Authentication, error) {
	getAccountStmt := table.Accounts.
		SELECT(table.Accounts.AllColumns).
		WHERE(
			table.Accounts.Username.EQ(pg.String(input.Username)),
		).
		LIMIT(1)

	account := model.Accounts{}
	err := getAccountStmt.Query(this.Db, &account)

	if err != nil && error_utils.HasNoRow(err) {
		return Authentication{}, errors.New("username or password is incorrect")
	}

	if err != nil {
		return Authentication{}, error_utils.InternalServerError
	}

	match, err := utils.ComparePasswordAndHash(input.Password, account.Password)
	if err != nil {
		return Authentication{}, errors.New("username or password is incorrect")
	}

	if !match {
		return Authentication{}, errors.New("username or password is incorrect")
	}

	token, err := jwt.SignToken(jwt.SignedTokenParams{
		AccountId: account.ID.String(),
		Nounce:    randstr.Hex(16),
	})

	if err != nil {
		return Authentication{}, error_utils.SignTokenFailed
	}

	refreshToken, err := jwt.SignRefreshToken(jwt.SignedTokenParams{
		AccountId: account.ID.String(),
		Nounce:    randstr.Hex(16),
	})

	if err != nil {
		return Authentication{}, error_utils.SignTokenFailed
	}

	return Authentication{
		Token:        token,
		RefreshToken: refreshToken,
		AccountId:    graphql.ID(account.ID.String()),
	}, nil
}
