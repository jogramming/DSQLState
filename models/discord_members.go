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
	"github.com/vattle/sqlboiler/types"
	"gopkg.in/nullbio/null.v6"
)

// DiscordMember is an object representing the database table.
type DiscordMember struct {
	UserID    int64            `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	GuildID   int64            `boil:"guild_id" json:"guild_id" toml:"guild_id" yaml:"guild_id"`
	CreatedAt time.Time        `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	LeftAt    null.Time        `boil:"left_at" json:"left_at,omitempty" toml:"left_at" yaml:"left_at,omitempty"`
	JoinedAt  time.Time        `boil:"joined_at" json:"joined_at" toml:"joined_at" yaml:"joined_at"`
	Nick      string           `boil:"nick" json:"nick" toml:"nick" yaml:"nick"`
	Deaf      bool             `boil:"deaf" json:"deaf" toml:"deaf" yaml:"deaf"`
	Mute      bool             `boil:"mute" json:"mute" toml:"mute" yaml:"mute"`
	Roles     types.Int64Array `boil:"roles" json:"roles" toml:"roles" yaml:"roles"`

	R *discordMemberR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L discordMemberL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// discordMemberR is where relationships are stored.
type discordMemberR struct {
	User  *DiscordUser
	Guild *DiscordGuild
}

// discordMemberL is where Load methods for each relationship are stored.
type discordMemberL struct{}

var (
	discordMemberColumns               = []string{"user_id", "guild_id", "created_at", "left_at", "joined_at", "nick", "deaf", "mute", "roles"}
	discordMemberColumnsWithoutDefault = []string{"user_id", "guild_id", "created_at", "left_at", "joined_at", "nick", "deaf", "mute", "roles"}
	discordMemberColumnsWithDefault    = []string{}
	discordMemberPrimaryKeyColumns     = []string{"user_id", "guild_id"}
)

type (
	// DiscordMemberSlice is an alias for a slice of pointers to DiscordMember.
	// This should generally be used opposed to []DiscordMember.
	DiscordMemberSlice []*DiscordMember

	discordMemberQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	discordMemberType                 = reflect.TypeOf(&DiscordMember{})
	discordMemberMapping              = queries.MakeStructMapping(discordMemberType)
	discordMemberPrimaryKeyMapping, _ = queries.BindMapping(discordMemberType, discordMemberMapping, discordMemberPrimaryKeyColumns)
	discordMemberInsertCacheMut       sync.RWMutex
	discordMemberInsertCache          = make(map[string]insertCache)
	discordMemberUpdateCacheMut       sync.RWMutex
	discordMemberUpdateCache          = make(map[string]updateCache)
	discordMemberUpsertCacheMut       sync.RWMutex
	discordMemberUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)

// OneP returns a single discordMember record from the query, and panics on error.
func (q discordMemberQuery) OneP() *DiscordMember {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single discordMember record from the query.
func (q discordMemberQuery) One() (*DiscordMember, error) {
	o := &DiscordMember{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for discord_members")
	}

	return o, nil
}

// AllP returns all DiscordMember records from the query, and panics on error.
func (q discordMemberQuery) AllP() DiscordMemberSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all DiscordMember records from the query.
func (q discordMemberQuery) All() (DiscordMemberSlice, error) {
	var o DiscordMemberSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to DiscordMember slice")
	}

	return o, nil
}

// CountP returns the count of all DiscordMember records in the query, and panics on error.
func (q discordMemberQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all DiscordMember records in the query.
func (q discordMemberQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count discord_members rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q discordMemberQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q discordMemberQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if discord_members exists")
	}

	return count > 0, nil
}

// UserG pointed to by the foreign key.
func (o *DiscordMember) UserG(mods ...qm.QueryMod) discordUserQuery {
	return o.User(boil.GetDB(), mods...)
}

// User pointed to by the foreign key.
func (o *DiscordMember) User(exec boil.Executor, mods ...qm.QueryMod) discordUserQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.UserID),
	}

	queryMods = append(queryMods, mods...)

	query := DiscordUsers(exec, queryMods...)
	queries.SetFrom(query.Query, "\"discord_users\"")

	return query
}

// GuildG pointed to by the foreign key.
func (o *DiscordMember) GuildG(mods ...qm.QueryMod) discordGuildQuery {
	return o.Guild(boil.GetDB(), mods...)
}

// Guild pointed to by the foreign key.
func (o *DiscordMember) Guild(exec boil.Executor, mods ...qm.QueryMod) discordGuildQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.GuildID),
	}

	queryMods = append(queryMods, mods...)

	query := DiscordGuilds(exec, queryMods...)
	queries.SetFrom(query.Query, "\"discord_guilds\"")

	return query
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (discordMemberL) LoadUser(e boil.Executor, singular bool, maybeDiscordMember interface{}) error {
	var slice []*DiscordMember
	var object *DiscordMember

	count := 1
	if singular {
		object = maybeDiscordMember.(*DiscordMember)
	} else {
		slice = *maybeDiscordMember.(*DiscordMemberSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		if object.R == nil {
			object.R = &discordMemberR{}
		}
		args[0] = object.UserID
	} else {
		for i, obj := range slice {
			if obj.R == nil {
				obj.R = &discordMemberR{}
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
func (discordMemberL) LoadGuild(e boil.Executor, singular bool, maybeDiscordMember interface{}) error {
	var slice []*DiscordMember
	var object *DiscordMember

	count := 1
	if singular {
		object = maybeDiscordMember.(*DiscordMember)
	} else {
		slice = *maybeDiscordMember.(*DiscordMemberSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		if object.R == nil {
			object.R = &discordMemberR{}
		}
		args[0] = object.GuildID
	} else {
		for i, obj := range slice {
			if obj.R == nil {
				obj.R = &discordMemberR{}
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

// SetUserG of the discord_member to the related item.
// Sets o.R.User to related.
// Adds o to related.R.UserDiscordMembers.
// Uses the global database handle.
func (o *DiscordMember) SetUserG(insert bool, related *DiscordUser) error {
	return o.SetUser(boil.GetDB(), insert, related)
}

// SetUserP of the discord_member to the related item.
// Sets o.R.User to related.
// Adds o to related.R.UserDiscordMembers.
// Panics on error.
func (o *DiscordMember) SetUserP(exec boil.Executor, insert bool, related *DiscordUser) {
	if err := o.SetUser(exec, insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetUserGP of the discord_member to the related item.
// Sets o.R.User to related.
// Adds o to related.R.UserDiscordMembers.
// Uses the global database handle and panics on error.
func (o *DiscordMember) SetUserGP(insert bool, related *DiscordUser) {
	if err := o.SetUser(boil.GetDB(), insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetUser of the discord_member to the related item.
// Sets o.R.User to related.
// Adds o to related.R.UserDiscordMembers.
func (o *DiscordMember) SetUser(exec boil.Executor, insert bool, related *DiscordUser) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"discord_members\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"user_id"}),
		strmangle.WhereClause("\"", "\"", 2, discordMemberPrimaryKeyColumns),
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
		o.R = &discordMemberR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &discordUserR{
			UserDiscordMembers: DiscordMemberSlice{o},
		}
	} else {
		related.R.UserDiscordMembers = append(related.R.UserDiscordMembers, o)
	}

	return nil
}

// SetGuildG of the discord_member to the related item.
// Sets o.R.Guild to related.
// Adds o to related.R.GuildDiscordMembers.
// Uses the global database handle.
func (o *DiscordMember) SetGuildG(insert bool, related *DiscordGuild) error {
	return o.SetGuild(boil.GetDB(), insert, related)
}

// SetGuildP of the discord_member to the related item.
// Sets o.R.Guild to related.
// Adds o to related.R.GuildDiscordMembers.
// Panics on error.
func (o *DiscordMember) SetGuildP(exec boil.Executor, insert bool, related *DiscordGuild) {
	if err := o.SetGuild(exec, insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetGuildGP of the discord_member to the related item.
// Sets o.R.Guild to related.
// Adds o to related.R.GuildDiscordMembers.
// Uses the global database handle and panics on error.
func (o *DiscordMember) SetGuildGP(insert bool, related *DiscordGuild) {
	if err := o.SetGuild(boil.GetDB(), insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetGuild of the discord_member to the related item.
// Sets o.R.Guild to related.
// Adds o to related.R.GuildDiscordMembers.
func (o *DiscordMember) SetGuild(exec boil.Executor, insert bool, related *DiscordGuild) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"discord_members\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"guild_id"}),
		strmangle.WhereClause("\"", "\"", 2, discordMemberPrimaryKeyColumns),
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
		o.R = &discordMemberR{
			Guild: related,
		}
	} else {
		o.R.Guild = related
	}

	if related.R == nil {
		related.R = &discordGuildR{
			GuildDiscordMembers: DiscordMemberSlice{o},
		}
	} else {
		related.R.GuildDiscordMembers = append(related.R.GuildDiscordMembers, o)
	}

	return nil
}

// DiscordMembersG retrieves all records.
func DiscordMembersG(mods ...qm.QueryMod) discordMemberQuery {
	return DiscordMembers(boil.GetDB(), mods...)
}

// DiscordMembers retrieves all the records using an executor.
func DiscordMembers(exec boil.Executor, mods ...qm.QueryMod) discordMemberQuery {
	mods = append(mods, qm.From("\"discord_members\""))
	return discordMemberQuery{NewQuery(exec, mods...)}
}

// FindDiscordMemberG retrieves a single record by ID.
func FindDiscordMemberG(userID int64, guildID int64, selectCols ...string) (*DiscordMember, error) {
	return FindDiscordMember(boil.GetDB(), userID, guildID, selectCols...)
}

// FindDiscordMemberGP retrieves a single record by ID, and panics on error.
func FindDiscordMemberGP(userID int64, guildID int64, selectCols ...string) *DiscordMember {
	retobj, err := FindDiscordMember(boil.GetDB(), userID, guildID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindDiscordMember retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDiscordMember(exec boil.Executor, userID int64, guildID int64, selectCols ...string) (*DiscordMember, error) {
	discordMemberObj := &DiscordMember{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"discord_members\" where \"user_id\"=$1 AND \"guild_id\"=$2", sel,
	)

	q := queries.Raw(exec, query, userID, guildID)

	err := q.Bind(discordMemberObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from discord_members")
	}

	return discordMemberObj, nil
}

// FindDiscordMemberP retrieves a single record by ID with an executor, and panics on error.
func FindDiscordMemberP(exec boil.Executor, userID int64, guildID int64, selectCols ...string) *DiscordMember {
	retobj, err := FindDiscordMember(exec, userID, guildID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *DiscordMember) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *DiscordMember) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *DiscordMember) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *DiscordMember) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no discord_members provided for insertion")
	}

	var err error
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(discordMemberColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	discordMemberInsertCacheMut.RLock()
	cache, cached := discordMemberInsertCache[key]
	discordMemberInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			discordMemberColumns,
			discordMemberColumnsWithDefault,
			discordMemberColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(discordMemberType, discordMemberMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(discordMemberType, discordMemberMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"discord_members\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "models: unable to insert into discord_members")
	}

	if !cached {
		discordMemberInsertCacheMut.Lock()
		discordMemberInsertCache[key] = cache
		discordMemberInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single DiscordMember record. See Update for
// whitelist behavior description.
func (o *DiscordMember) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single DiscordMember record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *DiscordMember) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the DiscordMember, and panics on error.
// See Update for whitelist behavior description.
func (o *DiscordMember) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the DiscordMember.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *DiscordMember) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	key := makeCacheKey(whitelist, nil)
	discordMemberUpdateCacheMut.RLock()
	cache, cached := discordMemberUpdateCache[key]
	discordMemberUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(discordMemberColumns, discordMemberPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update discord_members, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"discord_members\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, discordMemberPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(discordMemberType, discordMemberMapping, append(wl, discordMemberPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update discord_members row")
	}

	if !cached {
		discordMemberUpdateCacheMut.Lock()
		discordMemberUpdateCache[key] = cache
		discordMemberUpdateCacheMut.Unlock()
	}

	return nil
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q discordMemberQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q discordMemberQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for discord_members")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o DiscordMemberSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o DiscordMemberSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o DiscordMemberSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o DiscordMemberSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), discordMemberPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"discord_members\" SET %s WHERE (\"user_id\",\"guild_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(discordMemberPrimaryKeyColumns), len(colNames)+1, len(discordMemberPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in discordMember slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *DiscordMember) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *DiscordMember) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *DiscordMember) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *DiscordMember) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no discord_members provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(discordMemberColumnsWithDefault, o)

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

	discordMemberUpsertCacheMut.RLock()
	cache, cached := discordMemberUpsertCache[key]
	discordMemberUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			discordMemberColumns,
			discordMemberColumnsWithDefault,
			discordMemberColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			discordMemberColumns,
			discordMemberPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert discord_members, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(discordMemberPrimaryKeyColumns))
			copy(conflict, discordMemberPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"discord_members\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(discordMemberType, discordMemberMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(discordMemberType, discordMemberMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert discord_members")
	}

	if !cached {
		discordMemberUpsertCacheMut.Lock()
		discordMemberUpsertCache[key] = cache
		discordMemberUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteP deletes a single DiscordMember record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *DiscordMember) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single DiscordMember record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *DiscordMember) DeleteG() error {
	if o == nil {
		return errors.New("models: no DiscordMember provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single DiscordMember record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *DiscordMember) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single DiscordMember record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *DiscordMember) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no DiscordMember provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), discordMemberPrimaryKeyMapping)
	sql := "DELETE FROM \"discord_members\" WHERE \"user_id\"=$1 AND \"guild_id\"=$2"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from discord_members")
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q discordMemberQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q discordMemberQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no discordMemberQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from discord_members")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o DiscordMemberSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o DiscordMemberSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no DiscordMember slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o DiscordMemberSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o DiscordMemberSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no DiscordMember slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), discordMemberPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"discord_members\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, discordMemberPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(discordMemberPrimaryKeyColumns), 1, len(discordMemberPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from discordMember slice")
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *DiscordMember) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *DiscordMember) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *DiscordMember) ReloadG() error {
	if o == nil {
		return errors.New("models: no DiscordMember provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *DiscordMember) Reload(exec boil.Executor) error {
	ret, err := FindDiscordMember(exec, o.UserID, o.GuildID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *DiscordMemberSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *DiscordMemberSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DiscordMemberSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty DiscordMemberSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DiscordMemberSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	discordMembers := DiscordMemberSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), discordMemberPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"discord_members\".* FROM \"discord_members\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, discordMemberPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(discordMemberPrimaryKeyColumns), 1, len(discordMemberPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&discordMembers)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in DiscordMemberSlice")
	}

	*o = discordMembers

	return nil
}

// DiscordMemberExists checks if the DiscordMember row exists.
func DiscordMemberExists(exec boil.Executor, userID int64, guildID int64) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"discord_members\" where \"user_id\"=$1 AND \"guild_id\"=$2 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, userID, guildID)
	}

	row := exec.QueryRow(sql, userID, guildID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if discord_members exists")
	}

	return exists, nil
}

// DiscordMemberExistsG checks if the DiscordMember row exists.
func DiscordMemberExistsG(userID int64, guildID int64) (bool, error) {
	return DiscordMemberExists(boil.GetDB(), userID, guildID)
}

// DiscordMemberExistsGP checks if the DiscordMember row exists. Panics on error.
func DiscordMemberExistsGP(userID int64, guildID int64) bool {
	e, err := DiscordMemberExists(boil.GetDB(), userID, guildID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// DiscordMemberExistsP checks if the DiscordMember row exists. Panics on error.
func DiscordMemberExistsP(exec boil.Executor, userID int64, guildID int64) bool {
	e, err := DiscordMemberExists(exec, userID, guildID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
