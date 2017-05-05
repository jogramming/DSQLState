package dsqlstate

import (
	"database/sql"
)

type Client struct {
	db *sql.DB
}
