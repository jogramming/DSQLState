package models

import (
	"bytes"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/queries"
	"github.com/vattle/sqlboiler/queries/qm"
	"github.com/vattle/sqlboiler/strmangle"
	"gopkg.in/nullbio/null.v6"
)

// DiscordGuild is an object representing the database table.
type DiscordGuild struct {
	ID                          int64     `boil:"id" json:"id" toml:"id" yaml:"id"`
	CreatedAt                   time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	LeftAt                      null.Time `boil:"left_at" json:"left_at,omitempty" toml:"left_at" yaml:"left_at,omitempty"`
	Name                        string    `boil:"name" json:"name" toml:"name" yaml:"name"`
	Icon                        string    `boil:"icon" json:"icon" toml:"icon" yaml:"icon"`
	Region                      string    `boil:"region" json:"region" toml:"region" yaml:"region"`
	AfkChannelID                int64     `boil:"afk_channel_id" json:"afk_channel_id" toml:"afk_channel_id" yaml:"afk_channel_id"`
	EmbedChannelID              int64     `boil:"embed_channel_id" json:"embed_channel_id" toml:"embed_channel_id" yaml:"embed_channel_id"`
	OwnerID                     int64     `boil:"owner_id" json:"owner_id" toml:"owner_id" yaml:"owner_id"`
	Splash                      string    `boil:"splash" json:"splash" toml:"splash" yaml:"splash"`
	AfkTimeout                  int       `boil:"afk_timeout" json:"afk_timeout" toml:"afk_timeout" yaml:"afk_timeout"`
	MemberCount                 int       `boil:"member_count" json:"member_count" toml:"member_count" yaml:"member_count"`
	VerificationLevel           int16     `boil:"verification_level" json:"verification_level" toml:"verification_level" yaml:"verification_level"`
	EmbedEnabled                bool      `boil:"embed_enabled" json:"embed_enabled" toml:"embed_enabled" yaml:"embed_enabled"`
	Large                       bool      `boil:"large" json:"large" toml:"large" yaml:"large"`
	DefaultMessageNotifications int16     `boil:"default_message_notifications" json:"default_message_notifications" toml:"default_message_notifications" yaml:"default_message_notifications"`

	R *discordGuildR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L discordGuildL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// discordGuildR is where relationships are stored.
type discordGuildR struct {
	GuildDiscordGuildChannels DiscordGuildChannelSlice
	GuildDiscordMembers       DiscordMemberSlice
	GuildDiscordGuildRoles    DiscordGuildRoleSlice
}

// discordGuildL is where Load methods for each relationship are stored.
type discordGuildL struct{}

var (
	discordGuildColumns               = []string{"id", "created_at", "left_at", "name", "icon", "region", "afk_channel_id", "embed_channel_id", "owner_id", "splash", "afk_timeout", "member_count", "verification_level", "embed_enabled", "large", "default_message_notifications"}
	discordGuildColumnsWithoutDefault = []string{"id", "created_at", "left_at", "name", "icon", "region", "afk_channel_id", "embed_channel_id", "owner_id", "splash", "afk_timeout", "member_count", "verification_level", "embed_enabled", "large", "default_message_notifications"}
	discordGuildColumnsWithDefault    = []string{}
	discordGuildPrimaryKeyColumns     = []string{"id"}
)

type (
	// DiscordGuildSlice is an alias for a slice of pointers to DiscordGuild.
	// This should generally be used opposed to []DiscordGuild.
	DiscordGuildSlice []*DiscordGuild

	discordGuildQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	discordGuildType                 = reflect.TypeOf(&DiscordGuild{})
	discordGuildMapping              = queries.MakeStructMapping(discordGuildType)
	discordGuildPrimaryKeyMapping, _ = queries.BindMapping(discordGuildType, discordGuildMapping, discordGuildPrimaryKeyColumns)
	discordGuildInsertCacheMut       sync.RWMutex
	discordGuildInsertCache          = make(map[string]insertCache)
	discordGuildUpdateCacheMut       sync.RWMutex
	discordGuildUpdateCache          = make(map[string]updateCache)
	discordGuildUpsertCacheMut       sync.RWMutex
	discordGuildUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)

// OneP returns a single discordGuild record from the query, and panics on error.
func (q discordGuildQuery) OneP() *DiscordGuild {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single discordGuild record from the query.
func (q discordGuildQuery) One() (*DiscordGuild, error) {
	o := &DiscordGuild{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for discord_guilds")
	}

	return o, nil
}

// AllP returns all DiscordGuild records from the query, and panics on error.
func (q discordGuildQuery) AllP() DiscordGuildSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all DiscordGuild records from the query.
func (q discordGuildQuery) All() (DiscordGuildSlice, error) {
	var o DiscordGuildSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to DiscordGuild slice")
	}

	return o, nil
}

// CountP returns the count of all DiscordGuild records in the query, and panics on error.
func (q discordGuildQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all DiscordGuild records in the query.
func (q discordGuildQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count discord_guilds rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q discordGuildQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q discordGuildQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if discord_guilds exists")
	}

	return count > 0, nil
}

// GuildDiscordGuildChannelsG retrieves all the discord_guild_channel's discord guild channels via guild_id column.
func (o *DiscordGuild) GuildDiscordGuildChannelsG(mods ...qm.QueryMod) discordGuildChannelQuery {
	return o.GuildDiscordGuildChannels(boil.GetDB(), mods...)
}

// GuildDiscordGuildChannels retrieves all the discord_guild_channel's discord guild channels with an executor via guild_id column.
func (o *DiscordGuild) GuildDiscordGuildChannels(exec boil.Executor, mods ...qm.QueryMod) discordGuildChannelQuery {
	queryMods := []qm.QueryMod{
		qm.Select("\"a\".*"),
	}

	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"a\".\"guild_id\"=?", o.ID),
	)

	query := DiscordGuildChannels(exec, queryMods...)
	queries.SetFrom(query.Query, "\"discord_guild_channels\" as \"a\"")
	return query
}

// GuildDiscordMembersG retrieves all the discord_member's discord members via guild_id column.
func (o *DiscordGuild) GuildDiscordMembersG(mods ...qm.QueryMod) discordMemberQuery {
	return o.GuildDiscordMembers(boil.GetDB(), mods...)
}

// GuildDiscordMembers retrieves all the discord_member's discord members with an executor via guild_id column.
func (o *DiscordGuild) GuildDiscordMembers(exec boil.Executor, mods ...qm.QueryMod) discordMemberQuery {
	queryMods := []qm.QueryMod{
		qm.Select("\"a\".*"),
	}

	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"a\".\"guild_id\"=?", o.ID),
	)

	query := DiscordMembers(exec, queryMods...)
	queries.SetFrom(query.Query, "\"discord_members\" as \"a\"")
	return query
}

// GuildDiscordGuildRolesG retrieves all the discord_guild_role's discord guild roles via guild_id column.
func (o *DiscordGuild) GuildDiscordGuildRolesG(mods ...qm.QueryMod) discordGuildRoleQuery {
	return o.GuildDiscordGuildRoles(boil.GetDB(), mods...)
}

// GuildDiscordGuildRoles retrieves all the discord_guild_role's discord guild roles with an executor via guild_id column.
func (o *DiscordGuild) GuildDiscordGuildRoles(exec boil.Executor, mods ...qm.QueryMod) discordGuildRoleQuery {
	queryMods := []qm.QueryMod{
		qm.Select("\"a\".*"),
	}

	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"a\".\"guild_id\"=?", o.ID),
	)

	query := DiscordGuildRoles(exec, queryMods...)
	queries.SetFrom(query.Query, "\"discord_guild_roles\" as \"a\"")
	return query
}

// LoadGuildDiscordGuildChannels allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (discordGuildL) LoadGuildDiscordGuildChannels(e boil.Executor, singular bool, maybeDiscordGuild interface{}) error {
	var slice []*DiscordGuild
	var object *DiscordGuild

	count := 1
	if singular {
		object = maybeDiscordGuild.(*DiscordGuild)
	} else {
		slice = *maybeDiscordGuild.(*DiscordGuildSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		if object.R == nil {
			object.R = &discordGuildR{}
		}
		args[0] = object.ID
	} else {
		for i, obj := range slice {
			if obj.R == nil {
				obj.R = &discordGuildR{}
			}
			args[i] = obj.ID
		}
	}

	query := fmt.Sprintf(
		"select * from \"discord_guild_channels\" where \"guild_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)
	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load discord_guild_channels")
	}
	defer results.Close()

	var resultSlice []*DiscordGuildChannel
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice discord_guild_channels")
	}

	if singular {
		object.R.GuildDiscordGuildChannels = resultSlice
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.GuildID {
				local.R.GuildDiscordGuildChannels = append(local.R.GuildDiscordGuildChannels, foreign)
				break
			}
		}
	}

	return nil
}

// LoadGuildDiscordMembers allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (discordGuildL) LoadGuildDiscordMembers(e boil.Executor, singular bool, maybeDiscordGuild interface{}) error {
	var slice []*DiscordGuild
	var object *DiscordGuild

	count := 1
	if singular {
		object = maybeDiscordGuild.(*DiscordGuild)
	} else {
		slice = *maybeDiscordGuild.(*DiscordGuildSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		if object.R == nil {
			object.R = &discordGuildR{}
		}
		args[0] = object.ID
	} else {
		for i, obj := range slice {
			if obj.R == nil {
				obj.R = &discordGuildR{}
			}
			args[i] = obj.ID
		}
	}

	query := fmt.Sprintf(
		"select * from \"discord_members\" where \"guild_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)
	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load discord_members")
	}
	defer results.Close()

	var resultSlice []*DiscordMember
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice discord_members")
	}

	if singular {
		object.R.GuildDiscordMembers = resultSlice
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.GuildID {
				local.R.GuildDiscordMembers = append(local.R.GuildDiscordMembers, foreign)
				break
			}
		}
	}

	return nil
}

// LoadGuildDiscordGuildRoles allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (discordGuildL) LoadGuildDiscordGuildRoles(e boil.Executor, singular bool, maybeDiscordGuild interface{}) error {
	var slice []*DiscordGuild
	var object *DiscordGuild

	count := 1
	if singular {
		object = maybeDiscordGuild.(*DiscordGuild)
	} else {
		slice = *maybeDiscordGuild.(*DiscordGuildSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		if object.R == nil {
			object.R = &discordGuildR{}
		}
		args[0] = object.ID
	} else {
		for i, obj := range slice {
			if obj.R == nil {
				obj.R = &discordGuildR{}
			}
			args[i] = obj.ID
		}
	}

	query := fmt.Sprintf(
		"select * from \"discord_guild_roles\" where \"guild_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)
	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load discord_guild_roles")
	}
	defer results.Close()

	var resultSlice []*DiscordGuildRole
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice discord_guild_roles")
	}

	if singular {
		object.R.GuildDiscordGuildRoles = resultSlice
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.GuildID {
				local.R.GuildDiscordGuildRoles = append(local.R.GuildDiscordGuildRoles, foreign)
				break
			}
		}
	}

	return nil
}

// AddGuildDiscordGuildChannelsG adds the given related objects to the existing relationships
// of the discord_guild, optionally inserting them as new records.
// Appends related to o.R.GuildDiscordGuildChannels.
// Sets related.R.Guild appropriately.
// Uses the global database handle.
func (o *DiscordGuild) AddGuildDiscordGuildChannelsG(insert bool, related ...*DiscordGuildChannel) error {
	return o.AddGuildDiscordGuildChannels(boil.GetDB(), insert, related...)
}

// AddGuildDiscordGuildChannelsP adds the given related objects to the existing relationships
// of the discord_guild, optionally inserting them as new records.
// Appends related to o.R.GuildDiscordGuildChannels.
// Sets related.R.Guild appropriately.
// Panics on error.
func (o *DiscordGuild) AddGuildDiscordGuildChannelsP(exec boil.Executor, insert bool, related ...*DiscordGuildChannel) {
	if err := o.AddGuildDiscordGuildChannels(exec, insert, related...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// AddGuildDiscordGuildChannelsGP adds the given related objects to the existing relationships
// of the discord_guild, optionally inserting them as new records.
// Appends related to o.R.GuildDiscordGuildChannels.
// Sets related.R.Guild appropriately.
// Uses the global database handle and panics on error.
func (o *DiscordGuild) AddGuildDiscordGuildChannelsGP(insert bool, related ...*DiscordGuildChannel) {
	if err := o.AddGuildDiscordGuildChannels(boil.GetDB(), insert, related...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// AddGuildDiscordGuildChannels adds the given related objects to the existing relationships
// of the discord_guild, optionally inserting them as new records.
// Appends related to o.R.GuildDiscordGuildChannels.
// Sets related.R.Guild appropriately.
func (o *DiscordGuild) AddGuildDiscordGuildChannels(exec boil.Executor, insert bool, related ...*DiscordGuildChannel) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.GuildID = o.ID
			if err = rel.Insert(exec); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"discord_guild_channels\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"guild_id"}),
				strmangle.WhereClause("\"", "\"", 2, discordGuildChannelPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}

			if _, err = exec.Exec(updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.GuildID = o.ID
		}
	}

	if o.R == nil {
		o.R = &discordGuildR{
			GuildDiscordGuildChannels: related,
		}
	} else {
		o.R.GuildDiscordGuildChannels = append(o.R.GuildDiscordGuildChannels, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &discordGuildChannelR{
				Guild: o,
			}
		} else {
			rel.R.Guild = o
		}
	}
	return nil
}

// AddGuildDiscordMembersG adds the given related objects to the existing relationships
// of the discord_guild, optionally inserting them as new records.
// Appends related to o.R.GuildDiscordMembers.
// Sets related.R.Guild appropriately.
// Uses the global database handle.
func (o *DiscordGuild) AddGuildDiscordMembersG(insert bool, related ...*DiscordMember) error {
	return o.AddGuildDiscordMembers(boil.GetDB(), insert, related...)
}

// AddGuildDiscordMembersP adds the given related objects to the existing relationships
// of the discord_guild, optionally inserting them as new records.
// Appends related to o.R.GuildDiscordMembers.
// Sets related.R.Guild appropriately.
// Panics on error.
func (o *DiscordGuild) AddGuildDiscordMembersP(exec boil.Executor, insert bool, related ...*DiscordMember) {
	if err := o.AddGuildDiscordMembers(exec, insert, related...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// AddGuildDiscordMembersGP adds the given related objects to the existing relationships
// of the discord_guild, optionally inserting them as new records.
// Appends related to o.R.GuildDiscordMembers.
// Sets related.R.Guild appropriately.
// Uses the global database handle and panics on error.
func (o *DiscordGuild) AddGuildDiscordMembersGP(insert bool, related ...*DiscordMember) {
	if err := o.AddGuildDiscordMembers(boil.GetDB(), insert, related...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// AddGuildDiscordMembers adds the given related objects to the existing relationships
// of the discord_guild, optionally inserting them as new records.
// Appends related to o.R.GuildDiscordMembers.
// Sets related.R.Guild appropriately.
func (o *DiscordGuild) AddGuildDiscordMembers(exec boil.Executor, insert bool, related ...*DiscordMember) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.GuildID = o.ID
			if err = rel.Insert(exec); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"discord_members\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"guild_id"}),
				strmangle.WhereClause("\"", "\"", 2, discordMemberPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.UserID, rel.GuildID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}

			if _, err = exec.Exec(updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.GuildID = o.ID
		}
	}

	if o.R == nil {
		o.R = &discordGuildR{
			GuildDiscordMembers: related,
		}
	} else {
		o.R.GuildDiscordMembers = append(o.R.GuildDiscordMembers, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &discordMemberR{
				Guild: o,
			}
		} else {
			rel.R.Guild = o
		}
	}
	return nil
}

// AddGuildDiscordGuildRolesG adds the given related objects to the existing relationships
// of the discord_guild, optionally inserting them as new records.
// Appends related to o.R.GuildDiscordGuildRoles.
// Sets related.R.Guild appropriately.
// Uses the global database handle.
func (o *DiscordGuild) AddGuildDiscordGuildRolesG(insert bool, related ...*DiscordGuildRole) error {
	return o.AddGuildDiscordGuildRoles(boil.GetDB(), insert, related...)
}

// AddGuildDiscordGuildRolesP adds the given related objects to the existing relationships
// of the discord_guild, optionally inserting them as new records.
// Appends related to o.R.GuildDiscordGuildRoles.
// Sets related.R.Guild appropriately.
// Panics on error.
func (o *DiscordGuild) AddGuildDiscordGuildRolesP(exec boil.Executor, insert bool, related ...*DiscordGuildRole) {
	if err := o.AddGuildDiscordGuildRoles(exec, insert, related...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// AddGuildDiscordGuildRolesGP adds the given related objects to the existing relationships
// of the discord_guild, optionally inserting them as new records.
// Appends related to o.R.GuildDiscordGuildRoles.
// Sets related.R.Guild appropriately.
// Uses the global database handle and panics on error.
func (o *DiscordGuild) AddGuildDiscordGuildRolesGP(insert bool, related ...*DiscordGuildRole) {
	if err := o.AddGuildDiscordGuildRoles(boil.GetDB(), insert, related...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// AddGuildDiscordGuildRoles adds the given related objects to the existing relationships
// of the discord_guild, optionally inserting them as new records.
// Appends related to o.R.GuildDiscordGuildRoles.
// Sets related.R.Guild appropriately.
func (o *DiscordGuild) AddGuildDiscordGuildRoles(exec boil.Executor, insert bool, related ...*DiscordGuildRole) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.GuildID = o.ID
			if err = rel.Insert(exec); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"discord_guild_roles\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"guild_id"}),
				strmangle.WhereClause("\"", "\"", 2, discordGuildRolePrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}

			if _, err = exec.Exec(updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.GuildID = o.ID
		}
	}

	if o.R == nil {
		o.R = &discordGuildR{
			GuildDiscordGuildRoles: related,
		}
	} else {
		o.R.GuildDiscordGuildRoles = append(o.R.GuildDiscordGuildRoles, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &discordGuildRoleR{
				Guild: o,
			}
		} else {
			rel.R.Guild = o
		}
	}
	return nil
}

// DiscordGuildsG retrieves all records.
func DiscordGuildsG(mods ...qm.QueryMod) discordGuildQuery {
	return DiscordGuilds(boil.GetDB(), mods...)
}

// DiscordGuilds retrieves all the records using an executor.
func DiscordGuilds(exec boil.Executor, mods ...qm.QueryMod) discordGuildQuery {
	mods = append(mods, qm.From("\"discord_guilds\""))
	return discordGuildQuery{NewQuery(exec, mods...)}
}

// FindDiscordGuildG retrieves a single record by ID.
func FindDiscordGuildG(id int64, selectCols ...string) (*DiscordGuild, error) {
	return FindDiscordGuild(boil.GetDB(), id, selectCols...)
}

// FindDiscordGuildGP retrieves a single record by ID, and panics on error.
func FindDiscordGuildGP(id int64, selectCols ...string) *DiscordGuild {
	retobj, err := FindDiscordGuild(boil.GetDB(), id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindDiscordGuild retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDiscordGuild(exec boil.Executor, id int64, selectCols ...string) (*DiscordGuild, error) {
	discordGuildObj := &DiscordGuild{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"discord_guilds\" where \"id\"=$1", sel,
	)

	q := queries.Raw(exec, query, id)

	err := q.Bind(discordGuildObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from discord_guilds")
	}

	return discordGuildObj, nil
}

// FindDiscordGuildP retrieves a single record by ID with an executor, and panics on error.
func FindDiscordGuildP(exec boil.Executor, id int64, selectCols ...string) *DiscordGuild {
	retobj, err := FindDiscordGuild(exec, id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *DiscordGuild) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *DiscordGuild) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *DiscordGuild) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *DiscordGuild) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no discord_guilds provided for insertion")
	}

	var err error
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(discordGuildColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	discordGuildInsertCacheMut.RLock()
	cache, cached := discordGuildInsertCache[key]
	discordGuildInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			discordGuildColumns,
			discordGuildColumnsWithDefault,
			discordGuildColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(discordGuildType, discordGuildMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(discordGuildType, discordGuildMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"discord_guilds\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

		if len(cache.retMapping) != 0 {
			cache.query += fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into discord_guilds")
	}

	if !cached {
		discordGuildInsertCacheMut.Lock()
		discordGuildInsertCache[key] = cache
		discordGuildInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single DiscordGuild record. See Update for
// whitelist behavior description.
func (o *DiscordGuild) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single DiscordGuild record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *DiscordGuild) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the DiscordGuild, and panics on error.
// See Update for whitelist behavior description.
func (o *DiscordGuild) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the DiscordGuild.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *DiscordGuild) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	key := makeCacheKey(whitelist, nil)
	discordGuildUpdateCacheMut.RLock()
	cache, cached := discordGuildUpdateCache[key]
	discordGuildUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(discordGuildColumns, discordGuildPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update discord_guilds, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"discord_guilds\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, discordGuildPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(discordGuildType, discordGuildMapping, append(wl, discordGuildPrimaryKeyColumns...))
		if err != nil {
			return err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err = exec.Exec(cache.query, values...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update discord_guilds row")
	}

	if !cached {
		discordGuildUpdateCacheMut.Lock()
		discordGuildUpdateCache[key] = cache
		discordGuildUpdateCacheMut.Unlock()
	}

	return nil
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q discordGuildQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q discordGuildQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for discord_guilds")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o DiscordGuildSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o DiscordGuildSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o DiscordGuildSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o DiscordGuildSlice) UpdateAll(exec boil.Executor, cols M) error {
	ln := int64(len(o))
	if ln == 0 {
		return nil
	}

	if len(cols) == 0 {
		return errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), discordGuildPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"discord_guilds\" SET %s WHERE (\"id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(discordGuildPrimaryKeyColumns), len(colNames)+1, len(discordGuildPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in discordGuild slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *DiscordGuild) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *DiscordGuild) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *DiscordGuild) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *DiscordGuild) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no discord_guilds provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(discordGuildColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs postgres problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range updateColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range whitelist {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	discordGuildUpsertCacheMut.RLock()
	cache, cached := discordGuildUpsertCache[key]
	discordGuildUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			discordGuildColumns,
			discordGuildColumnsWithDefault,
			discordGuildColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			discordGuildColumns,
			discordGuildPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert discord_guilds, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(discordGuildPrimaryKeyColumns))
			copy(conflict, discordGuildPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"discord_guilds\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(discordGuildType, discordGuildMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(discordGuildType, discordGuildMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert discord_guilds")
	}

	if !cached {
		discordGuildUpsertCacheMut.Lock()
		discordGuildUpsertCache[key] = cache
		discordGuildUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteP deletes a single DiscordGuild record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *DiscordGuild) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single DiscordGuild record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *DiscordGuild) DeleteG() error {
	if o == nil {
		return errors.New("models: no DiscordGuild provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single DiscordGuild record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *DiscordGuild) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single DiscordGuild record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *DiscordGuild) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no DiscordGuild provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), discordGuildPrimaryKeyMapping)
	sql := "DELETE FROM \"discord_guilds\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from discord_guilds")
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q discordGuildQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q discordGuildQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no discordGuildQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from discord_guilds")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o DiscordGuildSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o DiscordGuildSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no DiscordGuild slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o DiscordGuildSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o DiscordGuildSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no DiscordGuild slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), discordGuildPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"discord_guilds\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, discordGuildPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(discordGuildPrimaryKeyColumns), 1, len(discordGuildPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from discordGuild slice")
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *DiscordGuild) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *DiscordGuild) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *DiscordGuild) ReloadG() error {
	if o == nil {
		return errors.New("models: no DiscordGuild provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *DiscordGuild) Reload(exec boil.Executor) error {
	ret, err := FindDiscordGuild(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *DiscordGuildSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *DiscordGuildSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DiscordGuildSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty DiscordGuildSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DiscordGuildSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	discordGuilds := DiscordGuildSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), discordGuildPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"discord_guilds\".* FROM \"discord_guilds\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, discordGuildPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(discordGuildPrimaryKeyColumns), 1, len(discordGuildPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&discordGuilds)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in DiscordGuildSlice")
	}

	*o = discordGuilds

	return nil
}

// DiscordGuildExists checks if the DiscordGuild row exists.
func DiscordGuildExists(exec boil.Executor, id int64) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"discord_guilds\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, id)
	}

	row := exec.QueryRow(sql, id)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if discord_guilds exists")
	}

	return exists, nil
}

// DiscordGuildExistsG checks if the DiscordGuild row exists.
func DiscordGuildExistsG(id int64) (bool, error) {
	return DiscordGuildExists(boil.GetDB(), id)
}

// DiscordGuildExistsGP checks if the DiscordGuild row exists. Panics on error.
func DiscordGuildExistsGP(id int64) bool {
	e, err := DiscordGuildExists(boil.GetDB(), id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// DiscordGuildExistsP checks if the DiscordGuild row exists. Panics on error.
func DiscordGuildExistsP(exec boil.Executor, id int64) bool {
	e, err := DiscordGuildExists(exec, id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
