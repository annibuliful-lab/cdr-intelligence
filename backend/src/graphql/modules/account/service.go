package account

import (
	"database/sql"
	"log"

	"backend/src/.gen/cdr-intelligence/public/model"
	"backend/src/.gen/cdr-intelligence/public/table"

	pg "github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"
	"github.com/graph-gophers/graphql-go"
	"github.com/redis/go-redis/v9"
)

type AccountService struct {
	db    *sql.DB
	redis *redis.Client
}

func NewAccountService(params NewAccountParams) AccountService {

	return AccountService{
		db:    params.Db,
		redis: params.Redis,
	}
}

func (service AccountService) GetById(id uuid.UUID) (*Account, error) {
	stmt := table.Accounts.
		SELECT(table.Accounts.AllColumns).
		WHERE(table.Accounts.ID.EQ(pg.UUID(id))).
		LIMIT(1)

	account := model.Accounts{}
	err := stmt.Query(service.db, &account)
	if err != nil {
		log.Println("get-account-profile-by-id-error", err.Error())
		return nil, err
	}

	return transformToGraphql(account), nil
}

func transformToGraphql(data model.Accounts) *Account {
	return &Account{
		Id:        graphql.ID(data.ID.String()),
		Username:  data.Username,
		CreatedAt: graphql.Time{Time: data.CreatedAt},
	}
}
