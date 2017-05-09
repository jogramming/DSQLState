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

// DiscordUser is an object representing the database table.
type DiscordUser struct {
	ID            int64       `boil:"id" json:"id" toml:"id" yaml:"id"`
	CreatedAt     time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	Username      string      `boil:"username" json:"username" toml:"username" yaml:"username"`
	Discriminator string      `boil:"discriminator" json:"discriminator" toml:"discriminator" yaml:"discriminator"`
	Bot           bool        `boil:"bot" json:"bot" toml:"bot" yaml:"bot"`
	Avatar        string      `boil:"avatar" json:"avatar" toml:"avatar" yaml:"avatar"`
	Status        string      `boil:"status" json:"status" toml:"status" yaml:"status"`
	GameName      null.String `boil:"game_name" json:"game_name,omitempty" toml:"game_name" yaml:"game_name,omitempty"`
	GameType      null.Int    `boil:"game_type" json:"game_type,omitempty" toml:"game_type" yaml:"game_type,omitempty"`
	GameURL       null.String `boil:"game_url" json:"game_url,omitempty" toml:"game_url" yaml:"game_url,omitempty"`

	R *discordUserR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L discordUserL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// discordUserR is where relationships are stored.
type discordUserR struct {
	RecipientDiscordPrivateChannels DiscordPrivateChannelSlice
	UserDiscordMembers              DiscordMemberSlice
}

// discordUserL is where Load methods for each relationship are stored.
type discordUserL struct{}

var (
	discordUserColumns               = []string{"id", "created_at", "username", "discriminator", "bot", "avatar", "status", "game_name", "game_type", "game_url"}
	discordUserColumnsWithoutDefault = []string{"id", "created_at", "username", "discriminator", "bot", "avatar", "status", "game_name", "game_type", "game_url"}
	discordUserColumnsWithDefault    = []string{}
	discordUserPrimaryKeyColumns     = []string{"id"}
)

type (
	// DiscordUserSlice is an alias for a slice of pointers to DiscordUser.
	// This should generally be used opposed to []DiscordUser.
	DiscordUserSlice []*DiscordUser

	discordUserQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	discordUserType                 = reflect.TypeOf(&DiscordUser{})
	discordUserMapping              = queries.MakeStructMapping(discordUserType)
	discordUserPrimaryKeyMapping, _ = queries.BindMapping(discordUserType, discordUserMapping, discordUserPrimaryKeyColumns)
	discordUserInsertCacheMut       sync.RWMutex
	discordUserInsertCache          = make(map[string]insertCache)
	discordUserUpdateCacheMut       sync.RWMutex
	discordUserUpdateCache          = make(map[string]updateCache)
	discordUserUpsertCacheMut       sync.RWMutex
	discordUserUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)

// OneP returns a single discordUser record from the query, and panics on error.
func (q discordUserQuery) OneP() *DiscordUser {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single discordUser record from the query.
func (q discordUserQuery) One() (*DiscordUser, error) {
	o := &DiscordUser{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for discord_users")
	}

	return o, nil
}

// AllP returns all DiscordUser records from the query, and panics on error.
func (q discordUserQuery) AllP() DiscordUserSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all DiscordUser records from the query.
func (q discordUserQuery) All() (DiscordUserSlice, error) {
	var o DiscordUserSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to DiscordUser slice")
	}

	return o, nil
}

// CountP returns the count of all DiscordUser records in the query, and panics on error.
func (q discordUserQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all DiscordUser records in the query.
func (q discordUserQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count discord_users rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q discordUserQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q discordUserQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if discord_users exists")
	}

	return count > 0, nil
}

// RecipientDiscordPrivateChannelsG retrieves all the discord_private_channel's discord private channels via recipient_id column.
func (o *DiscordUser) RecipientDiscordPrivateChannelsG(mods ...qm.QueryMod) discordPrivateChannelQuery {
	return o.RecipientDiscordPrivateChannels(boil.GetDB(), mods...)
}

// RecipientDiscordPrivateChannels retrieves all the discord_private_channel's discord private channels with an executor via recipient_id column.
func (o *DiscordUser) RecipientDiscordPrivateChannels(exec boil.Executor, mods ...qm.QueryMod) discordPrivateChannelQuery {
	queryMods := []qm.QueryMod{
		qm.Select("\"a\".*"),
	}

	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"a\".\"recipient_id\"=?", o.ID),
	)

	query := DiscordPrivateChannels(exec, queryMods...)
	queries.SetFrom(query.Query, "\"discord_private_channels\" as \"a\"")
	return query
}

// UserDiscordMembersG retrieves all the discord_member's discord members via user_id column.
func (o *DiscordUser) UserDiscordMembersG(mods ...qm.QueryMod) discordMemberQuery {
	return o.UserDiscordMembers(boil.GetDB(), mods...)
}

// UserDiscordMembers retrieves all the discord_member's discord members with an executor via user_id column.
func (o *DiscordUser) UserDiscordMembers(exec boil.Executor, mods ...qm.QueryMod) discordMemberQuery {
	queryMods := []qm.QueryMod{
		qm.Select("\"a\".*"),
	}

	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"a\".\"user_id\"=?", o.ID),
	)

	query := DiscordMembers(exec, queryMods...)
	queries.SetFrom(query.Query, "\"discord_members\" as \"a\"")
	return query
}

// LoadRecipientDiscordPrivateChannels allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (discordUserL) LoadRecipientDiscordPrivateChannels(e boil.Executor, singular bool, maybeDiscordUser interface{}) error {
	var slice []*DiscordUser
	var object *DiscordUser

	count := 1
	if singular {
		object = maybeDiscordUser.(*DiscordUser)
	} else {
		slice = *maybeDiscordUser.(*DiscordUserSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		if object.R == nil {
			object.R = &discordUserR{}
		}
		args[0] = object.ID
	} else {
		for i, obj := range slice {
			if obj.R == nil {
				obj.R = &discordUserR{}
			}
			args[i] = obj.ID
		}
	}

	query := fmt.Sprintf(
		"select * from \"discord_private_channels\" where \"recipient_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)
	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load discord_private_channels")
	}
	defer results.Close()

	var resultSlice []*DiscordPrivateChannel
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice discord_private_channels")
	}

	if singular {
		object.R.RecipientDiscordPrivateChannels = resultSlice
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.RecipientID {
				local.R.RecipientDiscordPrivateChannels = append(local.R.RecipientDiscordPrivateChannels, foreign)
				break
			}
		}
	}

	return nil
}

// LoadUserDiscordMembers allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (discordUserL) LoadUserDiscordMembers(e boil.Executor, singular bool, maybeDiscordUser interface{}) error {
	var slice []*DiscordUser
	var object *DiscordUser

	count := 1
	if singular {
		object = maybeDiscordUser.(*DiscordUser)
	} else {
		slice = *maybeDiscordUser.(*DiscordUserSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		if object.R == nil {
			object.R = &discordUserR{}
		}
		args[0] = object.ID
	} else {
		for i, obj := range slice {
			if obj.R == nil {
				obj.R = &discordUserR{}
			}
			args[i] = obj.ID
		}
	}

	query := fmt.Sprintf(
		"select * from \"discord_members\" where \"user_id\" in (%s)",
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
		object.R.UserDiscordMembers = resultSlice
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.UserID {
				local.R.UserDiscordMembers = append(local.R.UserDiscordMembers, foreign)
				break
			}
		}
	}

	return nil
}

// AddRecipientDiscordPrivateChannelsG adds the given related objects to the existing relationships
// of the discord_user, optionally inserting them as new records.
// Appends related to o.R.RecipientDiscordPrivateChannels.
// Sets related.R.Recipient appropriately.
// Uses the global database handle.
func (o *DiscordUser) AddRecipientDiscordPrivateChannelsG(insert bool, related ...*DiscordPrivateChannel) error {
	return o.AddRecipientDiscordPrivateChannels(boil.GetDB(), insert, related...)
}

// AddRecipientDiscordPrivateChannelsP adds the given related objects to the existing relationships
// of the discord_user, optionally inserting them as new records.
// Appends related to o.R.RecipientDiscordPrivateChannels.
// Sets related.R.Recipient appropriately.
// Panics on error.
func (o *DiscordUser) AddRecipientDiscordPrivateChannelsP(exec boil.Executor, insert bool, related ...*DiscordPrivateChannel) {
	if err := o.AddRecipientDiscordPrivateChannels(exec, insert, related...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// AddRecipientDiscordPrivateChannelsGP adds the given related objects to the existing relationships
// of the discord_user, optionally inserting them as new records.
// Appends related to o.R.RecipientDiscordPrivateChannels.
// Sets related.R.Recipient appropriately.
// Uses the global database handle and panics on error.
func (o *DiscordUser) AddRecipientDiscordPrivateChannelsGP(insert bool, related ...*DiscordPrivateChannel) {
	if err := o.AddRecipientDiscordPrivateChannels(boil.GetDB(), insert, related...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// AddRecipientDiscordPrivateChannels adds the given related objects to the existing relationships
// of the discord_user, optionally inserting them as new records.
// Appends related to o.R.RecipientDiscordPrivateChannels.
// Sets related.R.Recipient appropriately.
func (o *DiscordUser) AddRecipientDiscordPrivateChannels(exec boil.Executor, insert bool, related ...*DiscordPrivateChannel) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.RecipientID = o.ID
			if err = rel.Insert(exec); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"discord_private_channels\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"recipient_id"}),
				strmangle.WhereClause("\"", "\"", 2, discordPrivateChannelPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}

			if _, err = exec.Exec(updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.RecipientID = o.ID
		}
	}

	if o.R == nil {
		o.R = &discordUserR{
			RecipientDiscordPrivateChannels: related,
		}
	} else {
		o.R.RecipientDiscordPrivateChannels = append(o.R.RecipientDiscordPrivateChannels, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &discordPrivateChannelR{
				Recipient: o,
			}
		} else {
			rel.R.Recipient = o
		}
	}
	return nil
}

// AddUserDiscordMembersG adds the given related objects to the existing relationships
// of the discord_user, optionally inserting them as new records.
// Appends related to o.R.UserDiscordMembers.
// Sets related.R.User appropriately.
// Uses the global database handle.
func (o *DiscordUser) AddUserDiscordMembersG(insert bool, related ...*DiscordMember) error {
	return o.AddUserDiscordMembers(boil.GetDB(), insert, related...)
}

// AddUserDiscordMembersP adds the given related objects to the existing relationships
// of the discord_user, optionally inserting them as new records.
// Appends related to o.R.UserDiscordMembers.
// Sets related.R.User appropriately.
// Panics on error.
func (o *DiscordUser) AddUserDiscordMembersP(exec boil.Executor, insert bool, related ...*DiscordMember) {
	if err := o.AddUserDiscordMembers(exec, insert, related...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// AddUserDiscordMembersGP adds the given related objects to the existing relationships
// of the discord_user, optionally inserting them as new records.
// Appends related to o.R.UserDiscordMembers.
// Sets related.R.User appropriately.
// Uses the global database handle and panics on error.
func (o *DiscordUser) AddUserDiscordMembersGP(insert bool, related ...*DiscordMember) {
	if err := o.AddUserDiscordMembers(boil.GetDB(), insert, related...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// AddUserDiscordMembers adds the given related objects to the existing relationships
// of the discord_user, optionally inserting them as new records.
// Appends related to o.R.UserDiscordMembers.
// Sets related.R.User appropriately.
func (o *DiscordUser) AddUserDiscordMembers(exec boil.Executor, insert bool, related ...*DiscordMember) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.UserID = o.ID
			if err = rel.Insert(exec); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"discord_members\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"user_id"}),
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

			rel.UserID = o.ID
		}
	}

	if o.R == nil {
		o.R = &discordUserR{
			UserDiscordMembers: related,
		}
	} else {
		o.R.UserDiscordMembers = append(o.R.UserDiscordMembers, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &discordMemberR{
				User: o,
			}
		} else {
			rel.R.User = o
		}
	}
	return nil
}

// DiscordUsersG retrieves all records.
func DiscordUsersG(mods ...qm.QueryMod) discordUserQuery {
	return DiscordUsers(boil.GetDB(), mods...)
}

// DiscordUsers retrieves all the records using an executor.
func DiscordUsers(exec boil.Executor, mods ...qm.QueryMod) discordUserQuery {
	mods = append(mods, qm.From("\"discord_users\""))
	return discordUserQuery{NewQuery(exec, mods...)}
}

// FindDiscordUserG retrieves a single record by ID.
func FindDiscordUserG(id int64, selectCols ...string) (*DiscordUser, error) {
	return FindDiscordUser(boil.GetDB(), id, selectCols...)
}

// FindDiscordUserGP retrieves a single record by ID, and panics on error.
func FindDiscordUserGP(id int64, selectCols ...string) *DiscordUser {
	retobj, err := FindDiscordUser(boil.GetDB(), id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindDiscordUser retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDiscordUser(exec boil.Executor, id int64, selectCols ...string) (*DiscordUser, error) {
	discordUserObj := &DiscordUser{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"discord_users\" where \"id\"=$1", sel,
	)

	q := queries.Raw(exec, query, id)

	err := q.Bind(discordUserObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from discord_users")
	}

	return discordUserObj, nil
}

// FindDiscordUserP retrieves a single record by ID with an executor, and panics on error.
func FindDiscordUserP(exec boil.Executor, id int64, selectCols ...string) *DiscordUser {
	retobj, err := FindDiscordUser(exec, id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *DiscordUser) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *DiscordUser) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *DiscordUser) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *DiscordUser) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no discord_users provided for insertion")
	}

	var err error
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(discordUserColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	discordUserInsertCacheMut.RLock()
	cache, cached := discordUserInsertCache[key]
	discordUserInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			discordUserColumns,
			discordUserColumnsWithDefault,
			discordUserColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(discordUserType, discordUserMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(discordUserType, discordUserMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"discord_users\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "models: unable to insert into discord_users")
	}

	if !cached {
		discordUserInsertCacheMut.Lock()
		discordUserInsertCache[key] = cache
		discordUserInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single DiscordUser record. See Update for
// whitelist behavior description.
func (o *DiscordUser) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single DiscordUser record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *DiscordUser) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the DiscordUser, and panics on error.
// See Update for whitelist behavior description.
func (o *DiscordUser) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the DiscordUser.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *DiscordUser) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	key := makeCacheKey(whitelist, nil)
	discordUserUpdateCacheMut.RLock()
	cache, cached := discordUserUpdateCache[key]
	discordUserUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(discordUserColumns, discordUserPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update discord_users, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"discord_users\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, discordUserPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(discordUserType, discordUserMapping, append(wl, discordUserPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update discord_users row")
	}

	if !cached {
		discordUserUpdateCacheMut.Lock()
		discordUserUpdateCache[key] = cache
		discordUserUpdateCacheMut.Unlock()
	}

	return nil
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q discordUserQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q discordUserQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for discord_users")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o DiscordUserSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o DiscordUserSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o DiscordUserSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o DiscordUserSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), discordUserPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"discord_users\" SET %s WHERE (\"id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(discordUserPrimaryKeyColumns), len(colNames)+1, len(discordUserPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in discordUser slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *DiscordUser) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *DiscordUser) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *DiscordUser) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *DiscordUser) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no discord_users provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(discordUserColumnsWithDefault, o)

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

	discordUserUpsertCacheMut.RLock()
	cache, cached := discordUserUpsertCache[key]
	discordUserUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			discordUserColumns,
			discordUserColumnsWithDefault,
			discordUserColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			discordUserColumns,
			discordUserPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert discord_users, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(discordUserPrimaryKeyColumns))
			copy(conflict, discordUserPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"discord_users\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(discordUserType, discordUserMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(discordUserType, discordUserMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert discord_users")
	}

	if !cached {
		discordUserUpsertCacheMut.Lock()
		discordUserUpsertCache[key] = cache
		discordUserUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteP deletes a single DiscordUser record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *DiscordUser) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single DiscordUser record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *DiscordUser) DeleteG() error {
	if o == nil {
		return errors.New("models: no DiscordUser provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single DiscordUser record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *DiscordUser) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single DiscordUser record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *DiscordUser) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no DiscordUser provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), discordUserPrimaryKeyMapping)
	sql := "DELETE FROM \"discord_users\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from discord_users")
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q discordUserQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q discordUserQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no discordUserQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from discord_users")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o DiscordUserSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o DiscordUserSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no DiscordUser slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o DiscordUserSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o DiscordUserSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no DiscordUser slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), discordUserPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"discord_users\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, discordUserPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(discordUserPrimaryKeyColumns), 1, len(discordUserPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from discordUser slice")
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *DiscordUser) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *DiscordUser) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *DiscordUser) ReloadG() error {
	if o == nil {
		return errors.New("models: no DiscordUser provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *DiscordUser) Reload(exec boil.Executor) error {
	ret, err := FindDiscordUser(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *DiscordUserSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *DiscordUserSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DiscordUserSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty DiscordUserSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DiscordUserSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	discordUsers := DiscordUserSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), discordUserPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"discord_users\".* FROM \"discord_users\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, discordUserPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(discordUserPrimaryKeyColumns), 1, len(discordUserPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&discordUsers)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in DiscordUserSlice")
	}

	*o = discordUsers

	return nil
}

// DiscordUserExists checks if the DiscordUser row exists.
func DiscordUserExists(exec boil.Executor, id int64) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"discord_users\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, id)
	}

	row := exec.QueryRow(sql, id)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if discord_users exists")
	}

	return exists, nil
}

// DiscordUserExistsG checks if the DiscordUser row exists.
func DiscordUserExistsG(id int64) (bool, error) {
	return DiscordUserExists(boil.GetDB(), id)
}

// DiscordUserExistsGP checks if the DiscordUser row exists. Panics on error.
func DiscordUserExistsGP(id int64) bool {
	e, err := DiscordUserExists(boil.GetDB(), id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// DiscordUserExistsP checks if the DiscordUser row exists. Panics on error.
func DiscordUserExistsP(exec boil.Executor, id int64) bool {
	e, err := DiscordUserExists(exec, id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
