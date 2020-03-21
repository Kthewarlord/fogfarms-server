package repository

import (
	"database/sql"
)

type pgsqlUserRepository struct {
	Conn *sql.DB
}



