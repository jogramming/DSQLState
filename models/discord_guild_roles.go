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

// DiscordGuildRole is an object representing the database table.
type DiscordGuildRole struct {
	ID          int64     `boil:"id" json:"id" toml:"id" yaml:"id"`
	GuildID     int64     `boil:"guild_id" json:"guild_id" toml:"guild_id" yaml:"guild_id"`
	CreatedAt   time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	DeletedAt   null.Time `boil:"deleted_at" json:"deleted_at,omitempty" toml:"deleted_at" yaml:"deleted_at,omitempty"`
	Name        string    `boil:"name" json:"name" toml:"name" yaml:"name"`
	Managed     bool      `boil:"managed" json:"managed" toml:"managed" yaml:"managed"`
	Mentionable bool      `boil:"mentionable" json:"mentionable" toml:"mentionable" yaml:"mentionable"`
	Hoist       bool      `boil:"hoist" json:"hoist" toml:"hoist" yaml:"hoist"`
	Color       int       `boil:"color" json:"color" toml:"color" yaml:"color"`
	Position    int       `boil:"position" json:"position" toml:"position" yaml:"position"`
	Permissions int       `boil:"permissions" json:"permissions" toml:"permissions" yaml:"permissions"`

	R *discordGuildRoleR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L discordGuildRoleL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// discordGuildRoleR is where relationships are stored.
type discordGuildRoleR struct {
}

// discordGuildRoleL is where Load methods for each relationship are stored.
type discordGuildRoleL struct{}

var (
	discordGuildRoleColumns               = []string{"id", "guild_id", "created_at", "deleted_at", "name", "managed", "mentionable", "hoist", "color", "position", "permissions"}
	discordGuildRoleColumnsWithoutDefault = []string{"id", "guild_id", "created_at", "deleted_at", "name", "managed", "mentionable", "hoist", "color", "position", "permissions"}
	discordGuildRoleColumnsWithDefault    = []string{}
	discordGuildRolePrimaryKeyColumns     = []string{"id"}
)

type (
	// DiscordGuildRoleSlice is an alias for a slice of pointers to DiscordGuildRole.
	// This should generally be used opposed to []DiscordGuildRole.
	DiscordGuildRoleSlice []*DiscordGuildRole

	discordGuildRoleQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	discordGuildRoleType                 = reflect.TypeOf(&DiscordGuildRole{})
	discordGuildRoleMapping              = queries.MakeStructMapping(discordGuildRoleType)
	discordGuildRolePrimaryKeyMapping, _ = queries.BindMapping(discordGuildRoleType, discordGuildRoleMapping, discordGuildRolePrimaryKeyColumns)
	discordGuildRoleInsertCacheMut       sync.RWMutex
	discordGuildRoleInsertCache          = make(map[string]insertCache)
	discordGuildRoleUpdateCacheMut       sync.RWMutex
	discordGuildRoleUpdateCache          = make(map[string]updateCache)
	discordGuildRoleUpsertCacheMut       sync.RWMutex
	discordGuildRoleUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)

// OneP returns a single discordGuildRole record from the query, and panics on error.
func (q discordGuildRoleQuery) OneP() *DiscordGuildRole {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single discordGuildRole record from the query.
func (q discordGuildRoleQuery) One() (*DiscordGuildRole, error) {
	o := &DiscordGuildRole{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for discord_guild_roles")
	}

	return o, nil
}

// AllP returns all DiscordGuildRole records from the query, and panics on error.
func (q discordGuildRoleQuery) AllP() DiscordGuildRoleSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all DiscordGuildRole records from the query.
func (q discordGuildRoleQuery) All() (DiscordGuildRoleSlice, error) {
	var o DiscordGuildRoleSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to DiscordGuildRole slice")
	}

	return o, nil
}

// CountP returns the count of all DiscordGuildRole records in the query, and panics on error.
func (q discordGuildRoleQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all DiscordGuildRole records in the query.
func (q discordGuildRoleQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count discord_guild_roles rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q discordGuildRoleQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q discordGuildRoleQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if discord_guild_roles exists")
	}

	return count > 0, nil
}

// DiscordGuildRolesG retrieves all records.
func DiscordGuildRolesG(mods ...qm.QueryMod) discordGuildRoleQuery {
	return DiscordGuildRoles(boil.GetDB(), mods...)
}

// DiscordGuildRoles retrieves all the records using an executor.
func DiscordGuildRoles(exec boil.Executor, mods ...qm.QueryMod) discordGuildRoleQuery {
	mods = append(mods, qm.From("\"discord_guild_roles\""))
	return discordGuildRoleQuery{NewQuery(exec, mods...)}
}

// FindDiscordGuildRoleG retrieves a single record by ID.
func FindDiscordGuildRoleG(id int64, selectCols ...string) (*DiscordGuildRole, error) {
	return FindDiscordGuildRole(boil.GetDB(), id, selectCols...)
}

// FindDiscordGuildRoleGP retrieves a single record by ID, and panics on error.
func FindDiscordGuildRoleGP(id int64, selectCols ...string) *DiscordGuildRole {
	retobj, err := FindDiscordGuildRole(boil.GetDB(), id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindDiscordGuildRole retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDiscordGuildRole(exec boil.Executor, id int64, selectCols ...string) (*DiscordGuildRole, error) {
	discordGuildRoleObj := &DiscordGuildRole{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"discord_guild_roles\" where \"id\"=$1", sel,
	)

	q := queries.Raw(exec, query, id)

	err := q.Bind(discordGuildRoleObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from discord_guild_roles")
	}

	return discordGuildRoleObj, nil
}

// FindDiscordGuildRoleP retrieves a single record by ID with an executor, and panics on error.
func FindDiscordGuildRoleP(exec boil.Executor, id int64, selectCols ...string) *DiscordGuildRole {
	retobj, err := FindDiscordGuildRole(exec, id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *DiscordGuildRole) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *DiscordGuildRole) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *DiscordGuildRole) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *DiscordGuildRole) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no discord_guild_roles provided for insertion")
	}

	var err error
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(discordGuildRoleColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	discordGuildRoleInsertCacheMut.RLock()
	cache, cached := discordGuildRoleInsertCache[key]
	discordGuildRoleInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			discordGuildRoleColumns,
			discordGuildRoleColumnsWithDefault,
			discordGuildRoleColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(discordGuildRoleType, discordGuildRoleMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(discordGuildRoleType, discordGuildRoleMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"discord_guild_roles\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "models: unable to insert into discord_guild_roles")
	}

	if !cached {
		discordGuildRoleInsertCacheMut.Lock()
		discordGuildRoleInsertCache[key] = cache
		discordGuildRoleInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single DiscordGuildRole record. See Update for
// whitelist behavior description.
func (o *DiscordGuildRole) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single DiscordGuildRole record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *DiscordGuildRole) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the DiscordGuildRole, and panics on error.
// See Update for whitelist behavior description.
func (o *DiscordGuildRole) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the DiscordGuildRole.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *DiscordGuildRole) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	key := makeCacheKey(whitelist, nil)
	discordGuildRoleUpdateCacheMut.RLock()
	cache, cached := discordGuildRoleUpdateCache[key]
	discordGuildRoleUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(discordGuildRoleColumns, discordGuildRolePrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update discord_guild_roles, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"discord_guild_roles\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, discordGuildRolePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(discordGuildRoleType, discordGuildRoleMapping, append(wl, discordGuildRolePrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update discord_guild_roles row")
	}

	if !cached {
		discordGuildRoleUpdateCacheMut.Lock()
		discordGuildRoleUpdateCache[key] = cache
		discordGuildRoleUpdateCacheMut.Unlock()
	}

	return nil
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q discordGuildRoleQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q discordGuildRoleQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for discord_guild_roles")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o DiscordGuildRoleSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o DiscordGuildRoleSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o DiscordGuildRoleSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o DiscordGuildRoleSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), discordGuildRolePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"discord_guild_roles\" SET %s WHERE (\"id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(discordGuildRolePrimaryKeyColumns), len(colNames)+1, len(discordGuildRolePrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in discordGuildRole slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *DiscordGuildRole) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *DiscordGuildRole) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *DiscordGuildRole) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *DiscordGuildRole) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no discord_guild_roles provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(discordGuildRoleColumnsWithDefault, o)

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

	discordGuildRoleUpsertCacheMut.RLock()
	cache, cached := discordGuildRoleUpsertCache[key]
	discordGuildRoleUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			discordGuildRoleColumns,
			discordGuildRoleColumnsWithDefault,
			discordGuildRoleColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			discordGuildRoleColumns,
			discordGuildRolePrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert discord_guild_roles, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(discordGuildRolePrimaryKeyColumns))
			copy(conflict, discordGuildRolePrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"discord_guild_roles\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(discordGuildRoleType, discordGuildRoleMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(discordGuildRoleType, discordGuildRoleMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert discord_guild_roles")
	}

	if !cached {
		discordGuildRoleUpsertCacheMut.Lock()
		discordGuildRoleUpsertCache[key] = cache
		discordGuildRoleUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteP deletes a single DiscordGuildRole record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *DiscordGuildRole) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single DiscordGuildRole record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *DiscordGuildRole) DeleteG() error {
	if o == nil {
		return errors.New("models: no DiscordGuildRole provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single DiscordGuildRole record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *DiscordGuildRole) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single DiscordGuildRole record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *DiscordGuildRole) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no DiscordGuildRole provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), discordGuildRolePrimaryKeyMapping)
	sql := "DELETE FROM \"discord_guild_roles\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from discord_guild_roles")
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q discordGuildRoleQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q discordGuildRoleQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no discordGuildRoleQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from discord_guild_roles")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o DiscordGuildRoleSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o DiscordGuildRoleSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no DiscordGuildRole slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o DiscordGuildRoleSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o DiscordGuildRoleSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no DiscordGuildRole slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), discordGuildRolePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"discord_guild_roles\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, discordGuildRolePrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(discordGuildRolePrimaryKeyColumns), 1, len(discordGuildRolePrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from discordGuildRole slice")
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *DiscordGuildRole) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *DiscordGuildRole) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *DiscordGuildRole) ReloadG() error {
	if o == nil {
		return errors.New("models: no DiscordGuildRole provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *DiscordGuildRole) Reload(exec boil.Executor) error {
	ret, err := FindDiscordGuildRole(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *DiscordGuildRoleSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *DiscordGuildRoleSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DiscordGuildRoleSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty DiscordGuildRoleSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DiscordGuildRoleSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	discordGuildRoles := DiscordGuildRoleSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), discordGuildRolePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"discord_guild_roles\".* FROM \"discord_guild_roles\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, discordGuildRolePrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(discordGuildRolePrimaryKeyColumns), 1, len(discordGuildRolePrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&discordGuildRoles)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in DiscordGuildRoleSlice")
	}

	*o = discordGuildRoles

	return nil
}

// DiscordGuildRoleExists checks if the DiscordGuildRole row exists.
func DiscordGuildRoleExists(exec boil.Executor, id int64) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"discord_guild_roles\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, id)
	}

	row := exec.QueryRow(sql, id)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if discord_guild_roles exists")
	}

	return exists, nil
}

// DiscordGuildRoleExistsG checks if the DiscordGuildRole row exists.
func DiscordGuildRoleExistsG(id int64) (bool, error) {
	return DiscordGuildRoleExists(boil.GetDB(), id)
}

// DiscordGuildRoleExistsGP checks if the DiscordGuildRole row exists. Panics on error.
func DiscordGuildRoleExistsGP(id int64) bool {
	e, err := DiscordGuildRoleExists(boil.GetDB(), id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// DiscordGuildRoleExistsP checks if the DiscordGuildRole row exists. Panics on error.
func DiscordGuildRoleExistsP(exec boil.Executor, id int64) bool {
	e, err := DiscordGuildRoleExists(exec, id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
