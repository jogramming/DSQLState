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

// DMetum is an object representing the database table.
type DMetum struct {
	Key   string `boil:"key" json:"key" toml:"key" yaml:"key"`
	Value []byte `boil:"value" json:"value" toml:"value" yaml:"value"`

	R *dMetumR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L dMetumL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// dMetumR is where relationships are stored.
type dMetumR struct {
}

// dMetumL is where Load methods for each relationship are stored.
type dMetumL struct{}

var (
	dMetumColumns               = []string{"key", "value"}
	dMetumColumnsWithoutDefault = []string{"key", "value"}
	dMetumColumnsWithDefault    = []string{}
	dMetumPrimaryKeyColumns     = []string{"key"}
)

type (
	// DMetumSlice is an alias for a slice of pointers to DMetum.
	// This should generally be used opposed to []DMetum.
	DMetumSlice []*DMetum

	dMetumQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	dMetumType                 = reflect.TypeOf(&DMetum{})
	dMetumMapping              = queries.MakeStructMapping(dMetumType)
	dMetumPrimaryKeyMapping, _ = queries.BindMapping(dMetumType, dMetumMapping, dMetumPrimaryKeyColumns)
	dMetumInsertCacheMut       sync.RWMutex
	dMetumInsertCache          = make(map[string]insertCache)
	dMetumUpdateCacheMut       sync.RWMutex
	dMetumUpdateCache          = make(map[string]updateCache)
	dMetumUpsertCacheMut       sync.RWMutex
	dMetumUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)

// OneP returns a single dMetum record from the query, and panics on error.
func (q dMetumQuery) OneP() *DMetum {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single dMetum record from the query.
func (q dMetumQuery) One() (*DMetum, error) {
	o := &DMetum{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for d_meta")
	}

	return o, nil
}

// AllP returns all DMetum records from the query, and panics on error.
func (q dMetumQuery) AllP() DMetumSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all DMetum records from the query.
func (q dMetumQuery) All() (DMetumSlice, error) {
	var o DMetumSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to DMetum slice")
	}

	return o, nil
}

// CountP returns the count of all DMetum records in the query, and panics on error.
func (q dMetumQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all DMetum records in the query.
func (q dMetumQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count d_meta rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q dMetumQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q dMetumQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if d_meta exists")
	}

	return count > 0, nil
}

// DMetaG retrieves all records.
func DMetaG(mods ...qm.QueryMod) dMetumQuery {
	return DMeta(boil.GetDB(), mods...)
}

// DMeta retrieves all the records using an executor.
func DMeta(exec boil.Executor, mods ...qm.QueryMod) dMetumQuery {
	mods = append(mods, qm.From("\"d_meta\""))
	return dMetumQuery{NewQuery(exec, mods...)}
}

// FindDMetumG retrieves a single record by ID.
func FindDMetumG(key string, selectCols ...string) (*DMetum, error) {
	return FindDMetum(boil.GetDB(), key, selectCols...)
}

// FindDMetumGP retrieves a single record by ID, and panics on error.
func FindDMetumGP(key string, selectCols ...string) *DMetum {
	retobj, err := FindDMetum(boil.GetDB(), key, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindDMetum retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDMetum(exec boil.Executor, key string, selectCols ...string) (*DMetum, error) {
	dMetumObj := &DMetum{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"d_meta\" where \"key\"=$1", sel,
	)

	q := queries.Raw(exec, query, key)

	err := q.Bind(dMetumObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from d_meta")
	}

	return dMetumObj, nil
}

// FindDMetumP retrieves a single record by ID with an executor, and panics on error.
func FindDMetumP(exec boil.Executor, key string, selectCols ...string) *DMetum {
	retobj, err := FindDMetum(exec, key, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *DMetum) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *DMetum) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *DMetum) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *DMetum) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no d_meta provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(dMetumColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	dMetumInsertCacheMut.RLock()
	cache, cached := dMetumInsertCache[key]
	dMetumInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			dMetumColumns,
			dMetumColumnsWithDefault,
			dMetumColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(dMetumType, dMetumMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(dMetumType, dMetumMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"d_meta\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "models: unable to insert into d_meta")
	}

	if !cached {
		dMetumInsertCacheMut.Lock()
		dMetumInsertCache[key] = cache
		dMetumInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single DMetum record. See Update for
// whitelist behavior description.
func (o *DMetum) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single DMetum record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *DMetum) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the DMetum, and panics on error.
// See Update for whitelist behavior description.
func (o *DMetum) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the DMetum.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *DMetum) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	key := makeCacheKey(whitelist, nil)
	dMetumUpdateCacheMut.RLock()
	cache, cached := dMetumUpdateCache[key]
	dMetumUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(dMetumColumns, dMetumPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update d_meta, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"d_meta\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, dMetumPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(dMetumType, dMetumMapping, append(wl, dMetumPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update d_meta row")
	}

	if !cached {
		dMetumUpdateCacheMut.Lock()
		dMetumUpdateCache[key] = cache
		dMetumUpdateCacheMut.Unlock()
	}

	return nil
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q dMetumQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q dMetumQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for d_meta")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o DMetumSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o DMetumSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o DMetumSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o DMetumSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dMetumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"d_meta\" SET %s WHERE (\"key\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(dMetumPrimaryKeyColumns), len(colNames)+1, len(dMetumPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in dMetum slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *DMetum) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *DMetum) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *DMetum) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *DMetum) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no d_meta provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(dMetumColumnsWithDefault, o)

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

	dMetumUpsertCacheMut.RLock()
	cache, cached := dMetumUpsertCache[key]
	dMetumUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			dMetumColumns,
			dMetumColumnsWithDefault,
			dMetumColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			dMetumColumns,
			dMetumPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert d_meta, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(dMetumPrimaryKeyColumns))
			copy(conflict, dMetumPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"d_meta\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(dMetumType, dMetumMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(dMetumType, dMetumMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert d_meta")
	}

	if !cached {
		dMetumUpsertCacheMut.Lock()
		dMetumUpsertCache[key] = cache
		dMetumUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteP deletes a single DMetum record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *DMetum) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single DMetum record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *DMetum) DeleteG() error {
	if o == nil {
		return errors.New("models: no DMetum provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single DMetum record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *DMetum) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single DMetum record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *DMetum) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no DMetum provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), dMetumPrimaryKeyMapping)
	sql := "DELETE FROM \"d_meta\" WHERE \"key\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from d_meta")
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q dMetumQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q dMetumQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no dMetumQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from d_meta")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o DMetumSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o DMetumSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no DMetum slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o DMetumSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o DMetumSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no DMetum slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dMetumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"d_meta\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, dMetumPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(dMetumPrimaryKeyColumns), 1, len(dMetumPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from dMetum slice")
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *DMetum) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *DMetum) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *DMetum) ReloadG() error {
	if o == nil {
		return errors.New("models: no DMetum provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *DMetum) Reload(exec boil.Executor) error {
	ret, err := FindDMetum(exec, o.Key)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *DMetumSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *DMetumSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DMetumSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty DMetumSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DMetumSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	dMeta := DMetumSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dMetumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"d_meta\".* FROM \"d_meta\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, dMetumPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(dMetumPrimaryKeyColumns), 1, len(dMetumPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&dMeta)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in DMetumSlice")
	}

	*o = dMeta

	return nil
}

// DMetumExists checks if the DMetum row exists.
func DMetumExists(exec boil.Executor, key string) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"d_meta\" where \"key\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, key)
	}

	row := exec.QueryRow(sql, key)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if d_meta exists")
	}

	return exists, nil
}

// DMetumExistsG checks if the DMetum row exists.
func DMetumExistsG(key string) (bool, error) {
	return DMetumExists(boil.GetDB(), key)
}

// DMetumExistsGP checks if the DMetum row exists. Panics on error.
func DMetumExistsGP(key string) bool {
	e, err := DMetumExists(boil.GetDB(), key)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// DMetumExistsP checks if the DMetum row exists. Panics on error.
func DMetumExistsP(exec boil.Executor, key string) bool {
	e, err := DMetumExists(exec, key)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
