package dbrepo

import (
	"database/sql"

	"github.com/fxfrancky/bookings/internal/config"
	"github.com/fxfrancky/bookings/internal/repository"
)

type postGresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

type testDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &postGresDBRepo{
		App: a,
		DB:  conn,
	}
}
func NewTestingRepo(a *config.AppConfig) repository.DatabaseRepo {
	return &testDBRepo{
		App: a,
	}
}
