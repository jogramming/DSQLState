package dsqlstate

import (
	"database/sql"
)

// TOOD, use dsqlstate/models to interact with the database directly for the moment.
// First priority is to get the tracker itself working nicely.
// This is gonna contain wrappers
type Client struct {
	db *sql.DB
}
