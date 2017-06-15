package dsqlstate

import (
	"database/sql"
	"github.com/bwmarrin/discordgo"
)

// TOOD, use dsqlstate/models to interact with the database directly for the moment.
// First priority is to get the tracker itself working nicely.
// This is gonna contain wrappers
type Client struct {
	db *sql.DB
}

// SelfUser returns the bot's selfuser from the meta table
// an error is returned if there is no bot user stored
// or there is an issue connecting to the database
func (c *Client) SelfUser() (*discordgo.User, error) {
	return nil, nil
}
