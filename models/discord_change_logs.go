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

// DiscordChangeLog is an object representing the database table.
type DiscordChangeLog struct {
	ID          int64       `boil:"id" json:"id" toml:"id" yaml:"id"`
	Field       int         `boil:"field" json:"field" toml:"field" yaml:"field"`
	Valueint    null.Int64  `boil:"valueint" json:"valueint,omitempty" toml:"valueint" yaml:"valueint,omitempty"`
	Valuestring null.String `boil:"valuestring" json:"valuestring,omitempty" toml:"valuestring" yaml:"valuestring,omitempty"`
	Valuebool   null.Bool   `boil:"valuebool" json:"valuebool,omitempty" toml:"valuebool" yaml:"valuebool,omitempty"`

	R *discordChangeLogR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L discordChangeLogL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// discordChangeLogR is where relationships are stored.
type discordChangeLogR struct {
}

// discordChangeLogL is where Load methods for each relationship are stored.
type discordChangeLogL struct{}

var (
	discordChangeLogColumns               = []string{"id", "field", "valueint", "valuestring", "valuebool"}
	discordChangeLogColumnsWithoutDefault = []string{"field", "valueint", "valuestring", "valuebool"}
	discordChangeLogColumnsWithDefault    = []string{"id"}
	discordChangeLogPrimaryKeyColumns     = []string{"id"}
)

type (
	// DiscordChangeLogSlice is an alias for a slice of pointers to DiscordChangeLog.
	// This should generally be used opposed to []DiscordChangeLog.
	DiscordChangeLogSlice []*DiscordChangeLog

	discordChangeLogQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	discordChangeLogType                 = reflect.TypeOf(&DiscordChangeLog{})
	discordChangeLogMapping              = queries.MakeStructMapping(discordChangeLogType)
	discordChangeLogPrimaryKeyMapping, _ = queries.BindMapping(discordChangeLogType, discordChangeLogMapping, discordChangeLogPrimaryKeyColumns)
	discordChangeLogInsertCacheMut       sync.RWMutex
	discordChangeLogInsertCache          = make(map[string]insertCache)
	discordChangeLogUpdateCacheMut       sync.RWMutex
	discordChangeLogUpdateCache          = make(map[string]updateCache)
	discordChangeLogUpsertCacheMut       sync.RWMutex
	discordChangeLogUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)

// OneP returns a single discordChangeLog record from the query, and panics on error.
func (q discordChangeLogQuery) OneP() *DiscordChangeLog {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single discordChangeLog record from the query.
func (q discordChangeLogQuery) One() (*DiscordChangeLog, error) {
	o := &DiscordChangeLog{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for discord_change_logs")
	}

	return o, nil
}

// AllP returns all DiscordChangeLog records from the query, and panics on error.
func (q discordChangeLogQuery) AllP() DiscordChangeLogSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all DiscordChangeLog records from the query.
func (q discordChangeLogQuery) All() (DiscordChangeLogSlice, error) {
	var o DiscordChangeLogSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to DiscordChangeLog slice")
	}

	return o, nil
}

// CountP returns the count of all DiscordChangeLog records in the query, and panics on error.
func (q discordChangeLogQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all DiscordChangeLog records in the query.
func (q discordChangeLogQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count discord_change_logs rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q discordChangeLogQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q discordChangeLogQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if discord_change_logs exists")
	}

	return count > 0, nil
}

// DiscordChangeLogsG retrieves all records.
func DiscordChangeLogsG(mods ...qm.QueryMod) discordChangeLogQuery {
	return DiscordChangeLogs(boil.GetDB(), mods...)
}

// DiscordChangeLogs retrieves all the records using an executor.
func DiscordChangeLogs(exec boil.Executor, mods ...qm.QueryMod) discordChangeLogQuery {
	mods = append(mods, qm.From("\"discord_change_logs\""))
	return discordChangeLogQuery{NewQuery(exec, mods...)}
}

// FindDiscordChangeLogG retrieves a single record by ID.
func FindDiscordChangeLogG(id int64, selectCols ...string) (*DiscordChangeLog, error) {
	return FindDiscordChangeLog(boil.GetDB(), id, selectCols...)
}

// FindDiscordChangeLogGP retrieves a single record by ID, and panics on error.
func FindDiscordChangeLogGP(id int64, selectCols ...string) *DiscordChangeLog {
	retobj, err := FindDiscordChangeLog(boil.GetDB(), id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindDiscordChangeLog retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDiscordChangeLog(exec boil.Executor, id int64, selectCols ...string) (*DiscordChangeLog, error) {
	discordChangeLogObj := &DiscordChangeLog{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"discord_change_logs\" where \"id\"=$1", sel,
	)

	q := queries.Raw(exec, query, id)

	err := q.Bind(discordChangeLogObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from discord_change_logs")
	}

	return discordChangeLogObj, nil
}

// FindDiscordChangeLogP retrieves a single record by ID with an executor, and panics on error.
func FindDiscordChangeLogP(exec boil.Executor, id int64, selectCols ...string) *DiscordChangeLog {
	retobj, err := FindDiscordChangeLog(exec, id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *DiscordChangeLog) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *DiscordChangeLog) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *DiscordChangeLog) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *DiscordChangeLog) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no discord_change_logs provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(discordChangeLogColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	discordChangeLogInsertCacheMut.RLock()
	cache, cached := discordChangeLogInsertCache[key]
	discordChangeLogInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			discordChangeLogColumns,
			discordChangeLogColumnsWithDefault,
			discordChangeLogColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(discordChangeLogType, discordChangeLogMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(discordChangeLogType, discordChangeLogMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"discord_change_logs\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "models: unable to insert into discord_change_logs")
	}

	if !cached {
		discordChangeLogInsertCacheMut.Lock()
		discordChangeLogInsertCache[key] = cache
		discordChangeLogInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single DiscordChangeLog record. See Update for
// whitelist behavior description.
func (o *DiscordChangeLog) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single DiscordChangeLog record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *DiscordChangeLog) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the DiscordChangeLog, and panics on error.
// See Update for whitelist behavior description.
func (o *DiscordChangeLog) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the DiscordChangeLog.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *DiscordChangeLog) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	key := makeCacheKey(whitelist, nil)
	discordChangeLogUpdateCacheMut.RLock()
	cache, cached := discordChangeLogUpdateCache[key]
	discordChangeLogUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(discordChangeLogColumns, discordChangeLogPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update discord_change_logs, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"discord_change_logs\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, discordChangeLogPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(discordChangeLogType, discordChangeLogMapping, append(wl, discordChangeLogPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update discord_change_logs row")
	}

	if !cached {
		discordChangeLogUpdateCacheMut.Lock()
		discordChangeLogUpdateCache[key] = cache
		discordChangeLogUpdateCacheMut.Unlock()
	}

	return nil
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q discordChangeLogQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q discordChangeLogQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for discord_change_logs")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o DiscordChangeLogSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o DiscordChangeLogSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o DiscordChangeLogSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o DiscordChangeLogSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), discordChangeLogPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"discord_change_logs\" SET %s WHERE (\"id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(discordChangeLogPrimaryKeyColumns), len(colNames)+1, len(discordChangeLogPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in discordChangeLog slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *DiscordChangeLog) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *DiscordChangeLog) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *DiscordChangeLog) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *DiscordChangeLog) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no discord_change_logs provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(discordChangeLogColumnsWithDefault, o)

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

	discordChangeLogUpsertCacheMut.RLock()
	cache, cached := discordChangeLogUpsertCache[key]
	discordChangeLogUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			discordChangeLogColumns,
			discordChangeLogColumnsWithDefault,
			discordChangeLogColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			discordChangeLogColumns,
			discordChangeLogPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert discord_change_logs, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(discordChangeLogPrimaryKeyColumns))
			copy(conflict, discordChangeLogPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"discord_change_logs\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(discordChangeLogType, discordChangeLogMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(discordChangeLogType, discordChangeLogMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert discord_change_logs")
	}

	if !cached {
		discordChangeLogUpsertCacheMut.Lock()
		discordChangeLogUpsertCache[key] = cache
		discordChangeLogUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteP deletes a single DiscordChangeLog record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *DiscordChangeLog) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single DiscordChangeLog record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *DiscordChangeLog) DeleteG() error {
	if o == nil {
		return errors.New("models: no DiscordChangeLog provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single DiscordChangeLog record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *DiscordChangeLog) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single DiscordChangeLog record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *DiscordChangeLog) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no DiscordChangeLog provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), discordChangeLogPrimaryKeyMapping)
	sql := "DELETE FROM \"discord_change_logs\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from discord_change_logs")
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q discordChangeLogQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q discordChangeLogQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no discordChangeLogQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from discord_change_logs")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o DiscordChangeLogSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o DiscordChangeLogSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no DiscordChangeLog slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o DiscordChangeLogSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o DiscordChangeLogSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no DiscordChangeLog slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), discordChangeLogPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"discord_change_logs\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, discordChangeLogPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(discordChangeLogPrimaryKeyColumns), 1, len(discordChangeLogPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from discordChangeLog slice")
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *DiscordChangeLog) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *DiscordChangeLog) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *DiscordChangeLog) ReloadG() error {
	if o == nil {
		return errors.New("models: no DiscordChangeLog provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *DiscordChangeLog) Reload(exec boil.Executor) error {
	ret, err := FindDiscordChangeLog(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *DiscordChangeLogSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *DiscordChangeLogSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DiscordChangeLogSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty DiscordChangeLogSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DiscordChangeLogSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	discordChangeLogs := DiscordChangeLogSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), discordChangeLogPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"discord_change_logs\".* FROM \"discord_change_logs\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, discordChangeLogPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(discordChangeLogPrimaryKeyColumns), 1, len(discordChangeLogPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&discordChangeLogs)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in DiscordChangeLogSlice")
	}

	*o = discordChangeLogs

	return nil
}

// DiscordChangeLogExists checks if the DiscordChangeLog row exists.
func DiscordChangeLogExists(exec boil.Executor, id int64) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"discord_change_logs\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, id)
	}

	row := exec.QueryRow(sql, id)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if discord_change_logs exists")
	}

	return exists, nil
}

// DiscordChangeLogExistsG checks if the DiscordChangeLog row exists.
func DiscordChangeLogExistsG(id int64) (bool, error) {
	return DiscordChangeLogExists(boil.GetDB(), id)
}

// DiscordChangeLogExistsGP checks if the DiscordChangeLog row exists. Panics on error.
func DiscordChangeLogExistsGP(id int64) bool {
	e, err := DiscordChangeLogExists(boil.GetDB(), id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// DiscordChangeLogExistsP checks if the DiscordChangeLog row exists. Panics on error.
func DiscordChangeLogExistsP(exec boil.Executor, id int64) bool {
	e, err := DiscordChangeLogExists(exec, id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
