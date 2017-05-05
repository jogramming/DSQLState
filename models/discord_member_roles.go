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
)

// DiscordMemberRole is an object representing the database table.
type DiscordMemberRole struct {
	UserID    int64     `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	GuildID   int64     `boil:"guild_id" json:"guild_id" toml:"guild_id" yaml:"guild_id"`
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	RoleID    int64     `boil:"role_id" json:"role_id" toml:"role_id" yaml:"role_id"`

	R *discordMemberRoleR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L discordMemberRoleL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// discordMemberRoleR is where relationships are stored.
type discordMemberRoleR struct {
	User  *DiscordUser
	Guild *DiscordGuild
	Role  *DiscordGuildRole
}

// discordMemberRoleL is where Load methods for each relationship are stored.
type discordMemberRoleL struct{}

var (
	discordMemberRoleColumns               = []string{"user_id", "guild_id", "created_at", "role_id"}
	discordMemberRoleColumnsWithoutDefault = []string{"user_id", "guild_id", "created_at", "role_id"}
	discordMemberRoleColumnsWithDefault    = []string{}
	discordMemberRolePrimaryKeyColumns     = []string{"user_id", "guild_id"}
)

type (
	// DiscordMemberRoleSlice is an alias for a slice of pointers to DiscordMemberRole.
	// This should generally be used opposed to []DiscordMemberRole.
	DiscordMemberRoleSlice []*DiscordMemberRole

	discordMemberRoleQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	discordMemberRoleType                 = reflect.TypeOf(&DiscordMemberRole{})
	discordMemberRoleMapping              = queries.MakeStructMapping(discordMemberRoleType)
	discordMemberRolePrimaryKeyMapping, _ = queries.BindMapping(discordMemberRoleType, discordMemberRoleMapping, discordMemberRolePrimaryKeyColumns)
	discordMemberRoleInsertCacheMut       sync.RWMutex
	discordMemberRoleInsertCache          = make(map[string]insertCache)
	discordMemberRoleUpdateCacheMut       sync.RWMutex
	discordMemberRoleUpdateCache          = make(map[string]updateCache)
	discordMemberRoleUpsertCacheMut       sync.RWMutex
	discordMemberRoleUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)

// OneP returns a single discordMemberRole record from the query, and panics on error.
func (q discordMemberRoleQuery) OneP() *DiscordMemberRole {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single discordMemberRole record from the query.
func (q discordMemberRoleQuery) One() (*DiscordMemberRole, error) {
	o := &DiscordMemberRole{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for discord_member_roles")
	}

	return o, nil
}

// AllP returns all DiscordMemberRole records from the query, and panics on error.
func (q discordMemberRoleQuery) AllP() DiscordMemberRoleSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all DiscordMemberRole records from the query.
func (q discordMemberRoleQuery) All() (DiscordMemberRoleSlice, error) {
	var o DiscordMemberRoleSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to DiscordMemberRole slice")
	}

	return o, nil
}

// CountP returns the count of all DiscordMemberRole records in the query, and panics on error.
func (q discordMemberRoleQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all DiscordMemberRole records in the query.
func (q discordMemberRoleQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count discord_member_roles rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q discordMemberRoleQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q discordMemberRoleQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if discord_member_roles exists")
	}

	return count > 0, nil
}

// UserG pointed to by the foreign key.
func (o *DiscordMemberRole) UserG(mods ...qm.QueryMod) discordUserQuery {
	return o.User(boil.GetDB(), mods...)
}

// User pointed to by the foreign key.
func (o *DiscordMemberRole) User(exec boil.Executor, mods ...qm.QueryMod) discordUserQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.UserID),
	}

	queryMods = append(queryMods, mods...)

	query := DiscordUsers(exec, queryMods...)
	queries.SetFrom(query.Query, "\"discord_users\"")

	return query
}

// GuildG pointed to by the foreign key.
func (o *DiscordMemberRole) GuildG(mods ...qm.QueryMod) discordGuildQuery {
	return o.Guild(boil.GetDB(), mods...)
}

// Guild pointed to by the foreign key.
func (o *DiscordMemberRole) Guild(exec boil.Executor, mods ...qm.QueryMod) discordGuildQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.GuildID),
	}

	queryMods = append(queryMods, mods...)

	query := DiscordGuilds(exec, queryMods...)
	queries.SetFrom(query.Query, "\"discord_guilds\"")

	return query
}

// RoleG pointed to by the foreign key.
func (o *DiscordMemberRole) RoleG(mods ...qm.QueryMod) discordGuildRoleQuery {
	return o.Role(boil.GetDB(), mods...)
}

// Role pointed to by the foreign key.
func (o *DiscordMemberRole) Role(exec boil.Executor, mods ...qm.QueryMod) discordGuildRoleQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.RoleID),
	}

	queryMods = append(queryMods, mods...)

	query := DiscordGuildRoles(exec, queryMods...)
	queries.SetFrom(query.Query, "\"discord_guild_roles\"")

	return query
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (discordMemberRoleL) LoadUser(e boil.Executor, singular bool, maybeDiscordMemberRole interface{}) error {
	var slice []*DiscordMemberRole
	var object *DiscordMemberRole

	count := 1
	if singular {
		object = maybeDiscordMemberRole.(*DiscordMemberRole)
	} else {
		slice = *maybeDiscordMemberRole.(*DiscordMemberRoleSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		if object.R == nil {
			object.R = &discordMemberRoleR{}
		}
		args[0] = object.UserID
	} else {
		for i, obj := range slice {
			if obj.R == nil {
				obj.R = &discordMemberRoleR{}
			}
			args[i] = obj.UserID
		}
	}

	query := fmt.Sprintf(
		"select * from \"discord_users\" where \"id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load DiscordUser")
	}
	defer results.Close()

	var resultSlice []*DiscordUser
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice DiscordUser")
	}

	if singular && len(resultSlice) != 0 {
		object.R.User = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.UserID == foreign.ID {
				local.R.User = foreign
				break
			}
		}
	}

	return nil
}

// LoadGuild allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (discordMemberRoleL) LoadGuild(e boil.Executor, singular bool, maybeDiscordMemberRole interface{}) error {
	var slice []*DiscordMemberRole
	var object *DiscordMemberRole

	count := 1
	if singular {
		object = maybeDiscordMemberRole.(*DiscordMemberRole)
	} else {
		slice = *maybeDiscordMemberRole.(*DiscordMemberRoleSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		if object.R == nil {
			object.R = &discordMemberRoleR{}
		}
		args[0] = object.GuildID
	} else {
		for i, obj := range slice {
			if obj.R == nil {
				obj.R = &discordMemberRoleR{}
			}
			args[i] = obj.GuildID
		}
	}

	query := fmt.Sprintf(
		"select * from \"discord_guilds\" where \"id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load DiscordGuild")
	}
	defer results.Close()

	var resultSlice []*DiscordGuild
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice DiscordGuild")
	}

	if singular && len(resultSlice) != 0 {
		object.R.Guild = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.GuildID == foreign.ID {
				local.R.Guild = foreign
				break
			}
		}
	}

	return nil
}

// LoadRole allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (discordMemberRoleL) LoadRole(e boil.Executor, singular bool, maybeDiscordMemberRole interface{}) error {
	var slice []*DiscordMemberRole
	var object *DiscordMemberRole

	count := 1
	if singular {
		object = maybeDiscordMemberRole.(*DiscordMemberRole)
	} else {
		slice = *maybeDiscordMemberRole.(*DiscordMemberRoleSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		if object.R == nil {
			object.R = &discordMemberRoleR{}
		}
		args[0] = object.RoleID
	} else {
		for i, obj := range slice {
			if obj.R == nil {
				obj.R = &discordMemberRoleR{}
			}
			args[i] = obj.RoleID
		}
	}

	query := fmt.Sprintf(
		"select * from \"discord_guild_roles\" where \"id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load DiscordGuildRole")
	}
	defer results.Close()

	var resultSlice []*DiscordGuildRole
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice DiscordGuildRole")
	}

	if singular && len(resultSlice) != 0 {
		object.R.Role = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.RoleID == foreign.ID {
				local.R.Role = foreign
				break
			}
		}
	}

	return nil
}

// SetUserG of the discord_member_role to the related item.
// Sets o.R.User to related.
// Adds o to related.R.UserDiscordMemberRoles.
// Uses the global database handle.
func (o *DiscordMemberRole) SetUserG(insert bool, related *DiscordUser) error {
	return o.SetUser(boil.GetDB(), insert, related)
}

// SetUserP of the discord_member_role to the related item.
// Sets o.R.User to related.
// Adds o to related.R.UserDiscordMemberRoles.
// Panics on error.
func (o *DiscordMemberRole) SetUserP(exec boil.Executor, insert bool, related *DiscordUser) {
	if err := o.SetUser(exec, insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetUserGP of the discord_member_role to the related item.
// Sets o.R.User to related.
// Adds o to related.R.UserDiscordMemberRoles.
// Uses the global database handle and panics on error.
func (o *DiscordMemberRole) SetUserGP(insert bool, related *DiscordUser) {
	if err := o.SetUser(boil.GetDB(), insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetUser of the discord_member_role to the related item.
// Sets o.R.User to related.
// Adds o to related.R.UserDiscordMemberRoles.
func (o *DiscordMemberRole) SetUser(exec boil.Executor, insert bool, related *DiscordUser) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"discord_member_roles\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"user_id"}),
		strmangle.WhereClause("\"", "\"", 2, discordMemberRolePrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.UserID, o.GuildID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.UserID = related.ID

	if o.R == nil {
		o.R = &discordMemberRoleR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &discordUserR{
			UserDiscordMemberRoles: DiscordMemberRoleSlice{o},
		}
	} else {
		related.R.UserDiscordMemberRoles = append(related.R.UserDiscordMemberRoles, o)
	}

	return nil
}

// SetGuildG of the discord_member_role to the related item.
// Sets o.R.Guild to related.
// Adds o to related.R.GuildDiscordMemberRoles.
// Uses the global database handle.
func (o *DiscordMemberRole) SetGuildG(insert bool, related *DiscordGuild) error {
	return o.SetGuild(boil.GetDB(), insert, related)
}

// SetGuildP of the discord_member_role to the related item.
// Sets o.R.Guild to related.
// Adds o to related.R.GuildDiscordMemberRoles.
// Panics on error.
func (o *DiscordMemberRole) SetGuildP(exec boil.Executor, insert bool, related *DiscordGuild) {
	if err := o.SetGuild(exec, insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetGuildGP of the discord_member_role to the related item.
// Sets o.R.Guild to related.
// Adds o to related.R.GuildDiscordMemberRoles.
// Uses the global database handle and panics on error.
func (o *DiscordMemberRole) SetGuildGP(insert bool, related *DiscordGuild) {
	if err := o.SetGuild(boil.GetDB(), insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetGuild of the discord_member_role to the related item.
// Sets o.R.Guild to related.
// Adds o to related.R.GuildDiscordMemberRoles.
func (o *DiscordMemberRole) SetGuild(exec boil.Executor, insert bool, related *DiscordGuild) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"discord_member_roles\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"guild_id"}),
		strmangle.WhereClause("\"", "\"", 2, discordMemberRolePrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.UserID, o.GuildID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.GuildID = related.ID

	if o.R == nil {
		o.R = &discordMemberRoleR{
			Guild: related,
		}
	} else {
		o.R.Guild = related
	}

	if related.R == nil {
		related.R = &discordGuildR{
			GuildDiscordMemberRoles: DiscordMemberRoleSlice{o},
		}
	} else {
		related.R.GuildDiscordMemberRoles = append(related.R.GuildDiscordMemberRoles, o)
	}

	return nil
}

// SetRoleG of the discord_member_role to the related item.
// Sets o.R.Role to related.
// Adds o to related.R.RoleDiscordMemberRoles.
// Uses the global database handle.
func (o *DiscordMemberRole) SetRoleG(insert bool, related *DiscordGuildRole) error {
	return o.SetRole(boil.GetDB(), insert, related)
}

// SetRoleP of the discord_member_role to the related item.
// Sets o.R.Role to related.
// Adds o to related.R.RoleDiscordMemberRoles.
// Panics on error.
func (o *DiscordMemberRole) SetRoleP(exec boil.Executor, insert bool, related *DiscordGuildRole) {
	if err := o.SetRole(exec, insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetRoleGP of the discord_member_role to the related item.
// Sets o.R.Role to related.
// Adds o to related.R.RoleDiscordMemberRoles.
// Uses the global database handle and panics on error.
func (o *DiscordMemberRole) SetRoleGP(insert bool, related *DiscordGuildRole) {
	if err := o.SetRole(boil.GetDB(), insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetRole of the discord_member_role to the related item.
// Sets o.R.Role to related.
// Adds o to related.R.RoleDiscordMemberRoles.
func (o *DiscordMemberRole) SetRole(exec boil.Executor, insert bool, related *DiscordGuildRole) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"discord_member_roles\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"role_id"}),
		strmangle.WhereClause("\"", "\"", 2, discordMemberRolePrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.UserID, o.GuildID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.RoleID = related.ID

	if o.R == nil {
		o.R = &discordMemberRoleR{
			Role: related,
		}
	} else {
		o.R.Role = related
	}

	if related.R == nil {
		related.R = &discordGuildRoleR{
			RoleDiscordMemberRoles: DiscordMemberRoleSlice{o},
		}
	} else {
		related.R.RoleDiscordMemberRoles = append(related.R.RoleDiscordMemberRoles, o)
	}

	return nil
}

// DiscordMemberRolesG retrieves all records.
func DiscordMemberRolesG(mods ...qm.QueryMod) discordMemberRoleQuery {
	return DiscordMemberRoles(boil.GetDB(), mods...)
}

// DiscordMemberRoles retrieves all the records using an executor.
func DiscordMemberRoles(exec boil.Executor, mods ...qm.QueryMod) discordMemberRoleQuery {
	mods = append(mods, qm.From("\"discord_member_roles\""))
	return discordMemberRoleQuery{NewQuery(exec, mods...)}
}

// FindDiscordMemberRoleG retrieves a single record by ID.
func FindDiscordMemberRoleG(userID int64, guildID int64, selectCols ...string) (*DiscordMemberRole, error) {
	return FindDiscordMemberRole(boil.GetDB(), userID, guildID, selectCols...)
}

// FindDiscordMemberRoleGP retrieves a single record by ID, and panics on error.
func FindDiscordMemberRoleGP(userID int64, guildID int64, selectCols ...string) *DiscordMemberRole {
	retobj, err := FindDiscordMemberRole(boil.GetDB(), userID, guildID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindDiscordMemberRole retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDiscordMemberRole(exec boil.Executor, userID int64, guildID int64, selectCols ...string) (*DiscordMemberRole, error) {
	discordMemberRoleObj := &DiscordMemberRole{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"discord_member_roles\" where \"user_id\"=$1 AND \"guild_id\"=$2", sel,
	)

	q := queries.Raw(exec, query, userID, guildID)

	err := q.Bind(discordMemberRoleObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from discord_member_roles")
	}

	return discordMemberRoleObj, nil
}

// FindDiscordMemberRoleP retrieves a single record by ID with an executor, and panics on error.
func FindDiscordMemberRoleP(exec boil.Executor, userID int64, guildID int64, selectCols ...string) *DiscordMemberRole {
	retobj, err := FindDiscordMemberRole(exec, userID, guildID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *DiscordMemberRole) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *DiscordMemberRole) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *DiscordMemberRole) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *DiscordMemberRole) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no discord_member_roles provided for insertion")
	}

	var err error
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(discordMemberRoleColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	discordMemberRoleInsertCacheMut.RLock()
	cache, cached := discordMemberRoleInsertCache[key]
	discordMemberRoleInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			discordMemberRoleColumns,
			discordMemberRoleColumnsWithDefault,
			discordMemberRoleColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(discordMemberRoleType, discordMemberRoleMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(discordMemberRoleType, discordMemberRoleMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"discord_member_roles\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "models: unable to insert into discord_member_roles")
	}

	if !cached {
		discordMemberRoleInsertCacheMut.Lock()
		discordMemberRoleInsertCache[key] = cache
		discordMemberRoleInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single DiscordMemberRole record. See Update for
// whitelist behavior description.
func (o *DiscordMemberRole) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single DiscordMemberRole record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *DiscordMemberRole) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the DiscordMemberRole, and panics on error.
// See Update for whitelist behavior description.
func (o *DiscordMemberRole) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the DiscordMemberRole.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *DiscordMemberRole) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	key := makeCacheKey(whitelist, nil)
	discordMemberRoleUpdateCacheMut.RLock()
	cache, cached := discordMemberRoleUpdateCache[key]
	discordMemberRoleUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(discordMemberRoleColumns, discordMemberRolePrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update discord_member_roles, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"discord_member_roles\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, discordMemberRolePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(discordMemberRoleType, discordMemberRoleMapping, append(wl, discordMemberRolePrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update discord_member_roles row")
	}

	if !cached {
		discordMemberRoleUpdateCacheMut.Lock()
		discordMemberRoleUpdateCache[key] = cache
		discordMemberRoleUpdateCacheMut.Unlock()
	}

	return nil
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q discordMemberRoleQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q discordMemberRoleQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for discord_member_roles")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o DiscordMemberRoleSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o DiscordMemberRoleSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o DiscordMemberRoleSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o DiscordMemberRoleSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), discordMemberRolePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"discord_member_roles\" SET %s WHERE (\"user_id\",\"guild_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(discordMemberRolePrimaryKeyColumns), len(colNames)+1, len(discordMemberRolePrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in discordMemberRole slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *DiscordMemberRole) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *DiscordMemberRole) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *DiscordMemberRole) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *DiscordMemberRole) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no discord_member_roles provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(discordMemberRoleColumnsWithDefault, o)

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

	discordMemberRoleUpsertCacheMut.RLock()
	cache, cached := discordMemberRoleUpsertCache[key]
	discordMemberRoleUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			discordMemberRoleColumns,
			discordMemberRoleColumnsWithDefault,
			discordMemberRoleColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			discordMemberRoleColumns,
			discordMemberRolePrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert discord_member_roles, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(discordMemberRolePrimaryKeyColumns))
			copy(conflict, discordMemberRolePrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"discord_member_roles\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(discordMemberRoleType, discordMemberRoleMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(discordMemberRoleType, discordMemberRoleMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert discord_member_roles")
	}

	if !cached {
		discordMemberRoleUpsertCacheMut.Lock()
		discordMemberRoleUpsertCache[key] = cache
		discordMemberRoleUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteP deletes a single DiscordMemberRole record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *DiscordMemberRole) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single DiscordMemberRole record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *DiscordMemberRole) DeleteG() error {
	if o == nil {
		return errors.New("models: no DiscordMemberRole provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single DiscordMemberRole record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *DiscordMemberRole) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single DiscordMemberRole record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *DiscordMemberRole) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no DiscordMemberRole provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), discordMemberRolePrimaryKeyMapping)
	sql := "DELETE FROM \"discord_member_roles\" WHERE \"user_id\"=$1 AND \"guild_id\"=$2"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from discord_member_roles")
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q discordMemberRoleQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q discordMemberRoleQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no discordMemberRoleQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from discord_member_roles")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o DiscordMemberRoleSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o DiscordMemberRoleSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no DiscordMemberRole slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o DiscordMemberRoleSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o DiscordMemberRoleSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no DiscordMemberRole slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), discordMemberRolePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"discord_member_roles\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, discordMemberRolePrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(discordMemberRolePrimaryKeyColumns), 1, len(discordMemberRolePrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from discordMemberRole slice")
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *DiscordMemberRole) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *DiscordMemberRole) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *DiscordMemberRole) ReloadG() error {
	if o == nil {
		return errors.New("models: no DiscordMemberRole provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *DiscordMemberRole) Reload(exec boil.Executor) error {
	ret, err := FindDiscordMemberRole(exec, o.UserID, o.GuildID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *DiscordMemberRoleSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *DiscordMemberRoleSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DiscordMemberRoleSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty DiscordMemberRoleSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DiscordMemberRoleSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	discordMemberRoles := DiscordMemberRoleSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), discordMemberRolePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"discord_member_roles\".* FROM \"discord_member_roles\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, discordMemberRolePrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(discordMemberRolePrimaryKeyColumns), 1, len(discordMemberRolePrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&discordMemberRoles)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in DiscordMemberRoleSlice")
	}

	*o = discordMemberRoles

	return nil
}

// DiscordMemberRoleExists checks if the DiscordMemberRole row exists.
func DiscordMemberRoleExists(exec boil.Executor, userID int64, guildID int64) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"discord_member_roles\" where \"user_id\"=$1 AND \"guild_id\"=$2 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, userID, guildID)
	}

	row := exec.QueryRow(sql, userID, guildID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if discord_member_roles exists")
	}

	return exists, nil
}

// DiscordMemberRoleExistsG checks if the DiscordMemberRole row exists.
func DiscordMemberRoleExistsG(userID int64, guildID int64) (bool, error) {
	return DiscordMemberRoleExists(boil.GetDB(), userID, guildID)
}

// DiscordMemberRoleExistsGP checks if the DiscordMemberRole row exists. Panics on error.
func DiscordMemberRoleExistsGP(userID int64, guildID int64) bool {
	e, err := DiscordMemberRoleExists(boil.GetDB(), userID, guildID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// DiscordMemberRoleExistsP checks if the DiscordMemberRole row exists. Panics on error.
func DiscordMemberRoleExistsP(exec boil.Executor, userID int64, guildID int64) bool {
	e, err := DiscordMemberRoleExists(exec, userID, guildID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
