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

// DGuildRole is an object representing the database table.
type DGuildRole struct {
	ID          int64     `boil:"id" json:"id" toml:"id" yaml:"id"`
	GuildID     int64     `boil:"guild_id" json:"guild_id" toml:"guild_id" yaml:"guild_id"`
	CreatedAt   time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	DeletedAt   null.Time `boil:"deleted_at" json:"deleted_at,omitempty" toml:"deleted_at" yaml:"deleted_at,omitempty"`
	Synced      bool      `boil:"synced" json:"synced" toml:"synced" yaml:"synced"`
	Name        string    `boil:"name" json:"name" toml:"name" yaml:"name"`
	Managed     bool      `boil:"managed" json:"managed" toml:"managed" yaml:"managed"`
	Mentionable bool      `boil:"mentionable" json:"mentionable" toml:"mentionable" yaml:"mentionable"`
	Hoist       bool      `boil:"hoist" json:"hoist" toml:"hoist" yaml:"hoist"`
	Color       int       `boil:"color" json:"color" toml:"color" yaml:"color"`
	Position    int       `boil:"position" json:"position" toml:"position" yaml:"position"`
	Permissions int       `boil:"permissions" json:"permissions" toml:"permissions" yaml:"permissions"`

	R *dGuildRoleR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L dGuildRoleL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// dGuildRoleR is where relationships are stored.
type dGuildRoleR struct {
}

// dGuildRoleL is where Load methods for each relationship are stored.
type dGuildRoleL struct{}

var (
	dGuildRoleColumns               = []string{"id", "guild_id", "created_at", "deleted_at", "synced", "name", "managed", "mentionable", "hoist", "color", "position", "permissions"}
	dGuildRoleColumnsWithoutDefault = []string{"id", "guild_id", "created_at", "deleted_at", "synced", "name", "managed", "mentionable", "hoist", "color", "position", "permissions"}
	dGuildRoleColumnsWithDefault    = []string{}
	dGuildRolePrimaryKeyColumns     = []string{"id"}
)

type (
	// DGuildRoleSlice is an alias for a slice of pointers to DGuildRole.
	// This should generally be used opposed to []DGuildRole.
	DGuildRoleSlice []*DGuildRole

	dGuildRoleQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	dGuildRoleType                 = reflect.TypeOf(&DGuildRole{})
	dGuildRoleMapping              = queries.MakeStructMapping(dGuildRoleType)
	dGuildRolePrimaryKeyMapping, _ = queries.BindMapping(dGuildRoleType, dGuildRoleMapping, dGuildRolePrimaryKeyColumns)
	dGuildRoleInsertCacheMut       sync.RWMutex
	dGuildRoleInsertCache          = make(map[string]insertCache)
	dGuildRoleUpdateCacheMut       sync.RWMutex
	dGuildRoleUpdateCache          = make(map[string]updateCache)
	dGuildRoleUpsertCacheMut       sync.RWMutex
	dGuildRoleUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)

// OneP returns a single dGuildRole record from the query, and panics on error.
func (q dGuildRoleQuery) OneP() *DGuildRole {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single dGuildRole record from the query.
func (q dGuildRoleQuery) One() (*DGuildRole, error) {
	o := &DGuildRole{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for d_guild_roles")
	}

	return o, nil
}

// AllP returns all DGuildRole records from the query, and panics on error.
func (q dGuildRoleQuery) AllP() DGuildRoleSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all DGuildRole records from the query.
func (q dGuildRoleQuery) All() (DGuildRoleSlice, error) {
	var o DGuildRoleSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to DGuildRole slice")
	}

	return o, nil
}

// CountP returns the count of all DGuildRole records in the query, and panics on error.
func (q dGuildRoleQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all DGuildRole records in the query.
func (q dGuildRoleQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count d_guild_roles rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q dGuildRoleQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q dGuildRoleQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if d_guild_roles exists")
	}

	return count > 0, nil
}

// DGuildRolesG retrieves all records.
func DGuildRolesG(mods ...qm.QueryMod) dGuildRoleQuery {
	return DGuildRoles(boil.GetDB(), mods...)
}

// DGuildRoles retrieves all the records using an executor.
func DGuildRoles(exec boil.Executor, mods ...qm.QueryMod) dGuildRoleQuery {
	mods = append(mods, qm.From("\"d_guild_roles\""))
	return dGuildRoleQuery{NewQuery(exec, mods...)}
}

// FindDGuildRoleG retrieves a single record by ID.
func FindDGuildRoleG(id int64, selectCols ...string) (*DGuildRole, error) {
	return FindDGuildRole(boil.GetDB(), id, selectCols...)
}

// FindDGuildRoleGP retrieves a single record by ID, and panics on error.
func FindDGuildRoleGP(id int64, selectCols ...string) *DGuildRole {
	retobj, err := FindDGuildRole(boil.GetDB(), id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindDGuildRole retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDGuildRole(exec boil.Executor, id int64, selectCols ...string) (*DGuildRole, error) {
	dGuildRoleObj := &DGuildRole{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"d_guild_roles\" where \"id\"=$1", sel,
	)

	q := queries.Raw(exec, query, id)

	err := q.Bind(dGuildRoleObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from d_guild_roles")
	}

	return dGuildRoleObj, nil
}

// FindDGuildRoleP retrieves a single record by ID with an executor, and panics on error.
func FindDGuildRoleP(exec boil.Executor, id int64, selectCols ...string) *DGuildRole {
	retobj, err := FindDGuildRole(exec, id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *DGuildRole) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *DGuildRole) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *DGuildRole) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *DGuildRole) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no d_guild_roles provided for insertion")
	}

	var err error
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(dGuildRoleColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	dGuildRoleInsertCacheMut.RLock()
	cache, cached := dGuildRoleInsertCache[key]
	dGuildRoleInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			dGuildRoleColumns,
			dGuildRoleColumnsWithDefault,
			dGuildRoleColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(dGuildRoleType, dGuildRoleMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(dGuildRoleType, dGuildRoleMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"d_guild_roles\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "models: unable to insert into d_guild_roles")
	}

	if !cached {
		dGuildRoleInsertCacheMut.Lock()
		dGuildRoleInsertCache[key] = cache
		dGuildRoleInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single DGuildRole record. See Update for
// whitelist behavior description.
func (o *DGuildRole) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single DGuildRole record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *DGuildRole) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the DGuildRole, and panics on error.
// See Update for whitelist behavior description.
func (o *DGuildRole) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the DGuildRole.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *DGuildRole) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	key := makeCacheKey(whitelist, nil)
	dGuildRoleUpdateCacheMut.RLock()
	cache, cached := dGuildRoleUpdateCache[key]
	dGuildRoleUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(dGuildRoleColumns, dGuildRolePrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update d_guild_roles, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"d_guild_roles\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, dGuildRolePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(dGuildRoleType, dGuildRoleMapping, append(wl, dGuildRolePrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update d_guild_roles row")
	}

	if !cached {
		dGuildRoleUpdateCacheMut.Lock()
		dGuildRoleUpdateCache[key] = cache
		dGuildRoleUpdateCacheMut.Unlock()
	}

	return nil
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q dGuildRoleQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q dGuildRoleQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for d_guild_roles")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o DGuildRoleSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o DGuildRoleSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o DGuildRoleSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o DGuildRoleSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dGuildRolePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"d_guild_roles\" SET %s WHERE (\"id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(dGuildRolePrimaryKeyColumns), len(colNames)+1, len(dGuildRolePrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in dGuildRole slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *DGuildRole) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *DGuildRole) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *DGuildRole) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *DGuildRole) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no d_guild_roles provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(dGuildRoleColumnsWithDefault, o)

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

	dGuildRoleUpsertCacheMut.RLock()
	cache, cached := dGuildRoleUpsertCache[key]
	dGuildRoleUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			dGuildRoleColumns,
			dGuildRoleColumnsWithDefault,
			dGuildRoleColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			dGuildRoleColumns,
			dGuildRolePrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert d_guild_roles, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(dGuildRolePrimaryKeyColumns))
			copy(conflict, dGuildRolePrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"d_guild_roles\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(dGuildRoleType, dGuildRoleMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(dGuildRoleType, dGuildRoleMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert d_guild_roles")
	}

	if !cached {
		dGuildRoleUpsertCacheMut.Lock()
		dGuildRoleUpsertCache[key] = cache
		dGuildRoleUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteP deletes a single DGuildRole record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *DGuildRole) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single DGuildRole record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *DGuildRole) DeleteG() error {
	if o == nil {
		return errors.New("models: no DGuildRole provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single DGuildRole record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *DGuildRole) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single DGuildRole record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *DGuildRole) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no DGuildRole provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), dGuildRolePrimaryKeyMapping)
	sql := "DELETE FROM \"d_guild_roles\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from d_guild_roles")
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q dGuildRoleQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q dGuildRoleQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no dGuildRoleQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from d_guild_roles")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o DGuildRoleSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o DGuildRoleSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no DGuildRole slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o DGuildRoleSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o DGuildRoleSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no DGuildRole slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dGuildRolePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"d_guild_roles\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, dGuildRolePrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(dGuildRolePrimaryKeyColumns), 1, len(dGuildRolePrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from dGuildRole slice")
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *DGuildRole) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *DGuildRole) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *DGuildRole) ReloadG() error {
	if o == nil {
		return errors.New("models: no DGuildRole provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *DGuildRole) Reload(exec boil.Executor) error {
	ret, err := FindDGuildRole(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *DGuildRoleSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *DGuildRoleSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DGuildRoleSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty DGuildRoleSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DGuildRoleSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	dGuildRoles := DGuildRoleSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dGuildRolePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"d_guild_roles\".* FROM \"d_guild_roles\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, dGuildRolePrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(dGuildRolePrimaryKeyColumns), 1, len(dGuildRolePrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&dGuildRoles)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in DGuildRoleSlice")
	}

	*o = dGuildRoles

	return nil
}

// DGuildRoleExists checks if the DGuildRole row exists.
func DGuildRoleExists(exec boil.Executor, id int64) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"d_guild_roles\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, id)
	}

	row := exec.QueryRow(sql, id)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if d_guild_roles exists")
	}

	return exists, nil
}

// DGuildRoleExistsG checks if the DGuildRole row exists.
func DGuildRoleExistsG(id int64) (bool, error) {
	return DGuildRoleExists(boil.GetDB(), id)
}

// DGuildRoleExistsGP checks if the DGuildRole row exists. Panics on error.
func DGuildRoleExistsGP(id int64) bool {
	e, err := DGuildRoleExists(boil.GetDB(), id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// DGuildRoleExistsP checks if the DGuildRole row exists. Panics on error.
func DGuildRoleExistsP(exec boil.Executor, id int64) bool {
	e, err := DGuildRoleExists(exec, id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
