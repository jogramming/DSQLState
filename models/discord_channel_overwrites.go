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

// DiscordChannelOverwrite is an object representing the database table.
type DiscordChannelOverwrite struct {
	ID        int64  `boil:"id" json:"id" toml:"id" yaml:"id"`
	ChannelID int64  `boil:"channel_id" json:"channel_id" toml:"channel_id" yaml:"channel_id"`
	Type      string `boil:"type" json:"type" toml:"type" yaml:"type"`
	Allow     int    `boil:"allow" json:"allow" toml:"allow" yaml:"allow"`
	Deny      int    `boil:"deny" json:"deny" toml:"deny" yaml:"deny"`

	R *discordChannelOverwriteR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L discordChannelOverwriteL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// discordChannelOverwriteR is where relationships are stored.
type discordChannelOverwriteR struct {
}

// discordChannelOverwriteL is where Load methods for each relationship are stored.
type discordChannelOverwriteL struct{}

var (
	discordChannelOverwriteColumns               = []string{"id", "channel_id", "type", "allow", "deny"}
	discordChannelOverwriteColumnsWithoutDefault = []string{"id", "channel_id", "type", "allow", "deny"}
	discordChannelOverwriteColumnsWithDefault    = []string{}
	discordChannelOverwritePrimaryKeyColumns     = []string{"id", "channel_id"}
)

type (
	// DiscordChannelOverwriteSlice is an alias for a slice of pointers to DiscordChannelOverwrite.
	// This should generally be used opposed to []DiscordChannelOverwrite.
	DiscordChannelOverwriteSlice []*DiscordChannelOverwrite

	discordChannelOverwriteQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	discordChannelOverwriteType                 = reflect.TypeOf(&DiscordChannelOverwrite{})
	discordChannelOverwriteMapping              = queries.MakeStructMapping(discordChannelOverwriteType)
	discordChannelOverwritePrimaryKeyMapping, _ = queries.BindMapping(discordChannelOverwriteType, discordChannelOverwriteMapping, discordChannelOverwritePrimaryKeyColumns)
	discordChannelOverwriteInsertCacheMut       sync.RWMutex
	discordChannelOverwriteInsertCache          = make(map[string]insertCache)
	discordChannelOverwriteUpdateCacheMut       sync.RWMutex
	discordChannelOverwriteUpdateCache          = make(map[string]updateCache)
	discordChannelOverwriteUpsertCacheMut       sync.RWMutex
	discordChannelOverwriteUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)

// OneP returns a single discordChannelOverwrite record from the query, and panics on error.
func (q discordChannelOverwriteQuery) OneP() *DiscordChannelOverwrite {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single discordChannelOverwrite record from the query.
func (q discordChannelOverwriteQuery) One() (*DiscordChannelOverwrite, error) {
	o := &DiscordChannelOverwrite{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for discord_channel_overwrites")
	}

	return o, nil
}

// AllP returns all DiscordChannelOverwrite records from the query, and panics on error.
func (q discordChannelOverwriteQuery) AllP() DiscordChannelOverwriteSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all DiscordChannelOverwrite records from the query.
func (q discordChannelOverwriteQuery) All() (DiscordChannelOverwriteSlice, error) {
	var o DiscordChannelOverwriteSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to DiscordChannelOverwrite slice")
	}

	return o, nil
}

// CountP returns the count of all DiscordChannelOverwrite records in the query, and panics on error.
func (q discordChannelOverwriteQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all DiscordChannelOverwrite records in the query.
func (q discordChannelOverwriteQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count discord_channel_overwrites rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q discordChannelOverwriteQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q discordChannelOverwriteQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if discord_channel_overwrites exists")
	}

	return count > 0, nil
}

// DiscordChannelOverwritesG retrieves all records.
func DiscordChannelOverwritesG(mods ...qm.QueryMod) discordChannelOverwriteQuery {
	return DiscordChannelOverwrites(boil.GetDB(), mods...)
}

// DiscordChannelOverwrites retrieves all the records using an executor.
func DiscordChannelOverwrites(exec boil.Executor, mods ...qm.QueryMod) discordChannelOverwriteQuery {
	mods = append(mods, qm.From("\"discord_channel_overwrites\""))
	return discordChannelOverwriteQuery{NewQuery(exec, mods...)}
}

// FindDiscordChannelOverwriteG retrieves a single record by ID.
func FindDiscordChannelOverwriteG(id int64, channelID int64, selectCols ...string) (*DiscordChannelOverwrite, error) {
	return FindDiscordChannelOverwrite(boil.GetDB(), id, channelID, selectCols...)
}

// FindDiscordChannelOverwriteGP retrieves a single record by ID, and panics on error.
func FindDiscordChannelOverwriteGP(id int64, channelID int64, selectCols ...string) *DiscordChannelOverwrite {
	retobj, err := FindDiscordChannelOverwrite(boil.GetDB(), id, channelID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindDiscordChannelOverwrite retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDiscordChannelOverwrite(exec boil.Executor, id int64, channelID int64, selectCols ...string) (*DiscordChannelOverwrite, error) {
	discordChannelOverwriteObj := &DiscordChannelOverwrite{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"discord_channel_overwrites\" where \"id\"=$1 AND \"channel_id\"=$2", sel,
	)

	q := queries.Raw(exec, query, id, channelID)

	err := q.Bind(discordChannelOverwriteObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from discord_channel_overwrites")
	}

	return discordChannelOverwriteObj, nil
}

// FindDiscordChannelOverwriteP retrieves a single record by ID with an executor, and panics on error.
func FindDiscordChannelOverwriteP(exec boil.Executor, id int64, channelID int64, selectCols ...string) *DiscordChannelOverwrite {
	retobj, err := FindDiscordChannelOverwrite(exec, id, channelID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *DiscordChannelOverwrite) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *DiscordChannelOverwrite) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *DiscordChannelOverwrite) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *DiscordChannelOverwrite) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no discord_channel_overwrites provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(discordChannelOverwriteColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	discordChannelOverwriteInsertCacheMut.RLock()
	cache, cached := discordChannelOverwriteInsertCache[key]
	discordChannelOverwriteInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			discordChannelOverwriteColumns,
			discordChannelOverwriteColumnsWithDefault,
			discordChannelOverwriteColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(discordChannelOverwriteType, discordChannelOverwriteMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(discordChannelOverwriteType, discordChannelOverwriteMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"discord_channel_overwrites\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "models: unable to insert into discord_channel_overwrites")
	}

	if !cached {
		discordChannelOverwriteInsertCacheMut.Lock()
		discordChannelOverwriteInsertCache[key] = cache
		discordChannelOverwriteInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single DiscordChannelOverwrite record. See Update for
// whitelist behavior description.
func (o *DiscordChannelOverwrite) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single DiscordChannelOverwrite record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *DiscordChannelOverwrite) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the DiscordChannelOverwrite, and panics on error.
// See Update for whitelist behavior description.
func (o *DiscordChannelOverwrite) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the DiscordChannelOverwrite.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *DiscordChannelOverwrite) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	key := makeCacheKey(whitelist, nil)
	discordChannelOverwriteUpdateCacheMut.RLock()
	cache, cached := discordChannelOverwriteUpdateCache[key]
	discordChannelOverwriteUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(discordChannelOverwriteColumns, discordChannelOverwritePrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update discord_channel_overwrites, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"discord_channel_overwrites\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, discordChannelOverwritePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(discordChannelOverwriteType, discordChannelOverwriteMapping, append(wl, discordChannelOverwritePrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update discord_channel_overwrites row")
	}

	if !cached {
		discordChannelOverwriteUpdateCacheMut.Lock()
		discordChannelOverwriteUpdateCache[key] = cache
		discordChannelOverwriteUpdateCacheMut.Unlock()
	}

	return nil
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q discordChannelOverwriteQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q discordChannelOverwriteQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for discord_channel_overwrites")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o DiscordChannelOverwriteSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o DiscordChannelOverwriteSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o DiscordChannelOverwriteSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o DiscordChannelOverwriteSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), discordChannelOverwritePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"discord_channel_overwrites\" SET %s WHERE (\"id\",\"channel_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(discordChannelOverwritePrimaryKeyColumns), len(colNames)+1, len(discordChannelOverwritePrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in discordChannelOverwrite slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *DiscordChannelOverwrite) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *DiscordChannelOverwrite) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *DiscordChannelOverwrite) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *DiscordChannelOverwrite) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no discord_channel_overwrites provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(discordChannelOverwriteColumnsWithDefault, o)

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

	discordChannelOverwriteUpsertCacheMut.RLock()
	cache, cached := discordChannelOverwriteUpsertCache[key]
	discordChannelOverwriteUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			discordChannelOverwriteColumns,
			discordChannelOverwriteColumnsWithDefault,
			discordChannelOverwriteColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			discordChannelOverwriteColumns,
			discordChannelOverwritePrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert discord_channel_overwrites, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(discordChannelOverwritePrimaryKeyColumns))
			copy(conflict, discordChannelOverwritePrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"discord_channel_overwrites\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(discordChannelOverwriteType, discordChannelOverwriteMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(discordChannelOverwriteType, discordChannelOverwriteMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert discord_channel_overwrites")
	}

	if !cached {
		discordChannelOverwriteUpsertCacheMut.Lock()
		discordChannelOverwriteUpsertCache[key] = cache
		discordChannelOverwriteUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteP deletes a single DiscordChannelOverwrite record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *DiscordChannelOverwrite) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single DiscordChannelOverwrite record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *DiscordChannelOverwrite) DeleteG() error {
	if o == nil {
		return errors.New("models: no DiscordChannelOverwrite provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single DiscordChannelOverwrite record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *DiscordChannelOverwrite) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single DiscordChannelOverwrite record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *DiscordChannelOverwrite) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no DiscordChannelOverwrite provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), discordChannelOverwritePrimaryKeyMapping)
	sql := "DELETE FROM \"discord_channel_overwrites\" WHERE \"id\"=$1 AND \"channel_id\"=$2"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from discord_channel_overwrites")
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q discordChannelOverwriteQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q discordChannelOverwriteQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no discordChannelOverwriteQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from discord_channel_overwrites")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o DiscordChannelOverwriteSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o DiscordChannelOverwriteSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no DiscordChannelOverwrite slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o DiscordChannelOverwriteSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o DiscordChannelOverwriteSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no DiscordChannelOverwrite slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), discordChannelOverwritePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"discord_channel_overwrites\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, discordChannelOverwritePrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(discordChannelOverwritePrimaryKeyColumns), 1, len(discordChannelOverwritePrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from discordChannelOverwrite slice")
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *DiscordChannelOverwrite) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *DiscordChannelOverwrite) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *DiscordChannelOverwrite) ReloadG() error {
	if o == nil {
		return errors.New("models: no DiscordChannelOverwrite provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *DiscordChannelOverwrite) Reload(exec boil.Executor) error {
	ret, err := FindDiscordChannelOverwrite(exec, o.ID, o.ChannelID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *DiscordChannelOverwriteSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *DiscordChannelOverwriteSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DiscordChannelOverwriteSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty DiscordChannelOverwriteSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DiscordChannelOverwriteSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	discordChannelOverwrites := DiscordChannelOverwriteSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), discordChannelOverwritePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"discord_channel_overwrites\".* FROM \"discord_channel_overwrites\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, discordChannelOverwritePrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(discordChannelOverwritePrimaryKeyColumns), 1, len(discordChannelOverwritePrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&discordChannelOverwrites)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in DiscordChannelOverwriteSlice")
	}

	*o = discordChannelOverwrites

	return nil
}

// DiscordChannelOverwriteExists checks if the DiscordChannelOverwrite row exists.
func DiscordChannelOverwriteExists(exec boil.Executor, id int64, channelID int64) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"discord_channel_overwrites\" where \"id\"=$1 AND \"channel_id\"=$2 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, id, channelID)
	}

	row := exec.QueryRow(sql, id, channelID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if discord_channel_overwrites exists")
	}

	return exists, nil
}

// DiscordChannelOverwriteExistsG checks if the DiscordChannelOverwrite row exists.
func DiscordChannelOverwriteExistsG(id int64, channelID int64) (bool, error) {
	return DiscordChannelOverwriteExists(boil.GetDB(), id, channelID)
}

// DiscordChannelOverwriteExistsGP checks if the DiscordChannelOverwrite row exists. Panics on error.
func DiscordChannelOverwriteExistsGP(id int64, channelID int64) bool {
	e, err := DiscordChannelOverwriteExists(boil.GetDB(), id, channelID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// DiscordChannelOverwriteExistsP checks if the DiscordChannelOverwrite row exists. Panics on error.
func DiscordChannelOverwriteExistsP(exec boil.Executor, id int64, channelID int64) bool {
	e, err := DiscordChannelOverwriteExists(exec, id, channelID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
