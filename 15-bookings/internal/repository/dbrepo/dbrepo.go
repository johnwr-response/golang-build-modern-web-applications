package dbrepo

import (
	"database/sql"
	"github.com/johnwr-response/golang-build-modern-web-applications/15-bookings/internal/config"
	"github.com/johnwr-response/golang-build-modern-web-applications/15-bookings/internal/repository"
)

type postgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

//type mariaDBRepo struct {
//	App *config.AppConfig
//	DB  *sql.DB
//}

func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &postgresDBRepo{App: a, DB: conn}
}

//func NewMariaDBRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
//	return &mariaDBRepo{App: a, DB: conn}
//}
