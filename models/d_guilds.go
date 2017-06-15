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

// DGuild is an object representing the database table.
type DGuild struct {
	ID                          int64     `boil:"id" json:"id" toml:"id" yaml:"id"`
	CreatedAt                   time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	LeftAt                      null.Time `boil:"left_at" json:"left_at,omitempty" toml:"left_at" yaml:"left_at,omitempty"`
	Synced                      bool      `boil:"synced" json:"synced" toml:"synced" yaml:"synced"`
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

	R *dGuildR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L dGuildL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// dGuildR is where relationships are stored.
type dGuildR struct {
}

// dGuildL is where Load methods for each relationship are stored.
type dGuildL struct{}

var (
	dGuildColumns               = []string{"id", "created_at", "left_at", "synced", "name", "icon", "region", "afk_channel_id", "embed_channel_id", "owner_id", "splash", "afk_timeout", "member_count", "verification_level", "embed_enabled", "large", "default_message_notifications"}
	dGuildColumnsWithoutDefault = []string{"id", "created_at", "left_at", "synced", "name", "icon", "region", "afk_channel_id", "embed_channel_id", "owner_id", "splash", "afk_timeout", "member_count", "verification_level", "embed_enabled", "large", "default_message_notifications"}
	dGuildColumnsWithDefault    = []string{}
	dGuildPrimaryKeyColumns     = []string{"id"}
)

type (
	// DGuildSlice is an alias for a slice of pointers to DGuild.
	// This should generally be used opposed to []DGuild.
	DGuildSlice []*DGuild

	dGuildQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	dGuildType                 = reflect.TypeOf(&DGuild{})
	dGuildMapping              = queries.MakeStructMapping(dGuildType)
	dGuildPrimaryKeyMapping, _ = queries.BindMapping(dGuildType, dGuildMapping, dGuildPrimaryKeyColumns)
	dGuildInsertCacheMut       sync.RWMutex
	dGuildInsertCache          = make(map[string]insertCache)
	dGuildUpdateCacheMut       sync.RWMutex
	dGuildUpdateCache          = make(map[string]updateCache)
	dGuildUpsertCacheMut       sync.RWMutex
	dGuildUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)

// OneP returns a single dGuild record from the query, and panics on error.
func (q dGuildQuery) OneP() *DGuild {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single dGuild record from the query.
func (q dGuildQuery) One() (*DGuild, error) {
	o := &DGuild{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for d_guilds")
	}

	return o, nil
}

// AllP returns all DGuild records from the query, and panics on error.
func (q dGuildQuery) AllP() DGuildSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all DGuild records from the query.
func (q dGuildQuery) All() (DGuildSlice, error) {
	var o DGuildSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to DGuild slice")
	}

	return o, nil
}

// CountP returns the count of all DGuild records in the query, and panics on error.
func (q dGuildQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all DGuild records in the query.
func (q dGuildQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count d_guilds rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q dGuildQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q dGuildQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if d_guilds exists")
	}

	return count > 0, nil
}

// DGuildsG retrieves all records.
func DGuildsG(mods ...qm.QueryMod) dGuildQuery {
	return DGuilds(boil.GetDB(), mods...)
}

// DGuilds retrieves all the records using an executor.
func DGuilds(exec boil.Executor, mods ...qm.QueryMod) dGuildQuery {
	mods = append(mods, qm.From("\"d_guilds\""))
	return dGuildQuery{NewQuery(exec, mods...)}
}

// FindDGuildG retrieves a single record by ID.
func FindDGuildG(id int64, selectCols ...string) (*DGuild, error) {
	return FindDGuild(boil.GetDB(), id, selectCols...)
}

// FindDGuildGP retrieves a single record by ID, and panics on error.
func FindDGuildGP(id int64, selectCols ...string) *DGuild {
	retobj, err := FindDGuild(boil.GetDB(), id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindDGuild retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDGuild(exec boil.Executor, id int64, selectCols ...string) (*DGuild, error) {
	dGuildObj := &DGuild{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"d_guilds\" where \"id\"=$1", sel,
	)

	q := queries.Raw(exec, query, id)

	err := q.Bind(dGuildObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from d_guilds")
	}

	return dGuildObj, nil
}

// FindDGuildP retrieves a single record by ID with an executor, and panics on error.
func FindDGuildP(exec boil.Executor, id int64, selectCols ...string) *DGuild {
	retobj, err := FindDGuild(exec, id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *DGuild) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *DGuild) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *DGuild) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *DGuild) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no d_guilds provided for insertion")
	}

	var err error
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(dGuildColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	dGuildInsertCacheMut.RLock()
	cache, cached := dGuildInsertCache[key]
	dGuildInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			dGuildColumns,
			dGuildColumnsWithDefault,
			dGuildColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(dGuildType, dGuildMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(dGuildType, dGuildMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"d_guilds\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "models: unable to insert into d_guilds")
	}

	if !cached {
		dGuildInsertCacheMut.Lock()
		dGuildInsertCache[key] = cache
		dGuildInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single DGuild record. See Update for
// whitelist behavior description.
func (o *DGuild) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single DGuild record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *DGuild) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the DGuild, and panics on error.
// See Update for whitelist behavior description.
func (o *DGuild) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the DGuild.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *DGuild) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	key := makeCacheKey(whitelist, nil)
	dGuildUpdateCacheMut.RLock()
	cache, cached := dGuildUpdateCache[key]
	dGuildUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(dGuildColumns, dGuildPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update d_guilds, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"d_guilds\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, dGuildPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(dGuildType, dGuildMapping, append(wl, dGuildPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update d_guilds row")
	}

	if !cached {
		dGuildUpdateCacheMut.Lock()
		dGuildUpdateCache[key] = cache
		dGuildUpdateCacheMut.Unlock()
	}

	return nil
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q dGuildQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q dGuildQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for d_guilds")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o DGuildSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o DGuildSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o DGuildSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o DGuildSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dGuildPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"d_guilds\" SET %s WHERE (\"id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(dGuildPrimaryKeyColumns), len(colNames)+1, len(dGuildPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in dGuild slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *DGuild) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *DGuild) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *DGuild) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *DGuild) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no d_guilds provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(dGuildColumnsWithDefault, o)

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

	dGuildUpsertCacheMut.RLock()
	cache, cached := dGuildUpsertCache[key]
	dGuildUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			dGuildColumns,
			dGuildColumnsWithDefault,
			dGuildColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			dGuildColumns,
			dGuildPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert d_guilds, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(dGuildPrimaryKeyColumns))
			copy(conflict, dGuildPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"d_guilds\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(dGuildType, dGuildMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(dGuildType, dGuildMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert d_guilds")
	}

	if !cached {
		dGuildUpsertCacheMut.Lock()
		dGuildUpsertCache[key] = cache
		dGuildUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteP deletes a single DGuild record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *DGuild) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single DGuild record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *DGuild) DeleteG() error {
	if o == nil {
		return errors.New("models: no DGuild provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single DGuild record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *DGuild) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single DGuild record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *DGuild) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no DGuild provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), dGuildPrimaryKeyMapping)
	sql := "DELETE FROM \"d_guilds\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from d_guilds")
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q dGuildQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q dGuildQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no dGuildQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from d_guilds")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o DGuildSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o DGuildSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no DGuild slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o DGuildSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o DGuildSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no DGuild slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dGuildPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"d_guilds\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, dGuildPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(dGuildPrimaryKeyColumns), 1, len(dGuildPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from dGuild slice")
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *DGuild) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *DGuild) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *DGuild) ReloadG() error {
	if o == nil {
		return errors.New("models: no DGuild provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *DGuild) Reload(exec boil.Executor) error {
	ret, err := FindDGuild(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *DGuildSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *DGuildSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DGuildSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty DGuildSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DGuildSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	dGuilds := DGuildSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dGuildPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"d_guilds\".* FROM \"d_guilds\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, dGuildPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(dGuildPrimaryKeyColumns), 1, len(dGuildPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&dGuilds)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in DGuildSlice")
	}

	*o = dGuilds

	return nil
}

// DGuildExists checks if the DGuild row exists.
func DGuildExists(exec boil.Executor, id int64) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"d_guilds\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, id)
	}

	row := exec.QueryRow(sql, id)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if d_guilds exists")
	}

	return exists, nil
}

// DGuildExistsG checks if the DGuild row exists.
func DGuildExistsG(id int64) (bool, error) {
	return DGuildExists(boil.GetDB(), id)
}

// DGuildExistsGP checks if the DGuild row exists. Panics on error.
func DGuildExistsGP(id int64) bool {
	e, err := DGuildExists(boil.GetDB(), id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// DGuildExistsP checks if the DGuild row exists. Panics on error.
func DGuildExistsP(exec boil.Executor, id int64) bool {
	e, err := DGuildExists(exec, id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
