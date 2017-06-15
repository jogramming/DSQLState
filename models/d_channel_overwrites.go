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

// DChannelOverwrite is an object representing the database table.
type DChannelOverwrite struct {
	ID        int64  `boil:"id" json:"id" toml:"id" yaml:"id"`
	ChannelID int64  `boil:"channel_id" json:"channel_id" toml:"channel_id" yaml:"channel_id"`
	Type      string `boil:"type" json:"type" toml:"type" yaml:"type"`
	Allow     int    `boil:"allow" json:"allow" toml:"allow" yaml:"allow"`
	Deny      int    `boil:"deny" json:"deny" toml:"deny" yaml:"deny"`

	R *dChannelOverwriteR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L dChannelOverwriteL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// dChannelOverwriteR is where relationships are stored.
type dChannelOverwriteR struct {
	Channel *DChannel
}

// dChannelOverwriteL is where Load methods for each relationship are stored.
type dChannelOverwriteL struct{}

var (
	dChannelOverwriteColumns               = []string{"id", "channel_id", "type", "allow", "deny"}
	dChannelOverwriteColumnsWithoutDefault = []string{"id", "channel_id", "type", "allow", "deny"}
	dChannelOverwriteColumnsWithDefault    = []string{}
	dChannelOverwritePrimaryKeyColumns     = []string{"id", "channel_id"}
)

type (
	// DChannelOverwriteSlice is an alias for a slice of pointers to DChannelOverwrite.
	// This should generally be used opposed to []DChannelOverwrite.
	DChannelOverwriteSlice []*DChannelOverwrite

	dChannelOverwriteQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	dChannelOverwriteType                 = reflect.TypeOf(&DChannelOverwrite{})
	dChannelOverwriteMapping              = queries.MakeStructMapping(dChannelOverwriteType)
	dChannelOverwritePrimaryKeyMapping, _ = queries.BindMapping(dChannelOverwriteType, dChannelOverwriteMapping, dChannelOverwritePrimaryKeyColumns)
	dChannelOverwriteInsertCacheMut       sync.RWMutex
	dChannelOverwriteInsertCache          = make(map[string]insertCache)
	dChannelOverwriteUpdateCacheMut       sync.RWMutex
	dChannelOverwriteUpdateCache          = make(map[string]updateCache)
	dChannelOverwriteUpsertCacheMut       sync.RWMutex
	dChannelOverwriteUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)

// OneP returns a single dChannelOverwrite record from the query, and panics on error.
func (q dChannelOverwriteQuery) OneP() *DChannelOverwrite {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single dChannelOverwrite record from the query.
func (q dChannelOverwriteQuery) One() (*DChannelOverwrite, error) {
	o := &DChannelOverwrite{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for d_channel_overwrites")
	}

	return o, nil
}

// AllP returns all DChannelOverwrite records from the query, and panics on error.
func (q dChannelOverwriteQuery) AllP() DChannelOverwriteSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all DChannelOverwrite records from the query.
func (q dChannelOverwriteQuery) All() (DChannelOverwriteSlice, error) {
	var o DChannelOverwriteSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to DChannelOverwrite slice")
	}

	return o, nil
}

// CountP returns the count of all DChannelOverwrite records in the query, and panics on error.
func (q dChannelOverwriteQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all DChannelOverwrite records in the query.
func (q dChannelOverwriteQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count d_channel_overwrites rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q dChannelOverwriteQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q dChannelOverwriteQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if d_channel_overwrites exists")
	}

	return count > 0, nil
}

// ChannelG pointed to by the foreign key.
func (o *DChannelOverwrite) ChannelG(mods ...qm.QueryMod) dChannelQuery {
	return o.Channel(boil.GetDB(), mods...)
}

// Channel pointed to by the foreign key.
func (o *DChannelOverwrite) Channel(exec boil.Executor, mods ...qm.QueryMod) dChannelQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.ChannelID),
	}

	queryMods = append(queryMods, mods...)

	query := DChannels(exec, queryMods...)
	queries.SetFrom(query.Query, "\"d_channels\"")

	return query
}

// LoadChannel allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (dChannelOverwriteL) LoadChannel(e boil.Executor, singular bool, maybeDChannelOverwrite interface{}) error {
	var slice []*DChannelOverwrite
	var object *DChannelOverwrite

	count := 1
	if singular {
		object = maybeDChannelOverwrite.(*DChannelOverwrite)
	} else {
		slice = *maybeDChannelOverwrite.(*DChannelOverwriteSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		if object.R == nil {
			object.R = &dChannelOverwriteR{}
		}
		args[0] = object.ChannelID
	} else {
		for i, obj := range slice {
			if obj.R == nil {
				obj.R = &dChannelOverwriteR{}
			}
			args[i] = obj.ChannelID
		}
	}

	query := fmt.Sprintf(
		"select * from \"d_channels\" where \"id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load DChannel")
	}
	defer results.Close()

	var resultSlice []*DChannel
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice DChannel")
	}

	if singular && len(resultSlice) != 0 {
		object.R.Channel = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ChannelID == foreign.ID {
				local.R.Channel = foreign
				break
			}
		}
	}

	return nil
}

// SetChannelG of the d_channel_overwrite to the related item.
// Sets o.R.Channel to related.
// Adds o to related.R.ChannelDChannelOverwrites.
// Uses the global database handle.
func (o *DChannelOverwrite) SetChannelG(insert bool, related *DChannel) error {
	return o.SetChannel(boil.GetDB(), insert, related)
}

// SetChannelP of the d_channel_overwrite to the related item.
// Sets o.R.Channel to related.
// Adds o to related.R.ChannelDChannelOverwrites.
// Panics on error.
func (o *DChannelOverwrite) SetChannelP(exec boil.Executor, insert bool, related *DChannel) {
	if err := o.SetChannel(exec, insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetChannelGP of the d_channel_overwrite to the related item.
// Sets o.R.Channel to related.
// Adds o to related.R.ChannelDChannelOverwrites.
// Uses the global database handle and panics on error.
func (o *DChannelOverwrite) SetChannelGP(insert bool, related *DChannel) {
	if err := o.SetChannel(boil.GetDB(), insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetChannel of the d_channel_overwrite to the related item.
// Sets o.R.Channel to related.
// Adds o to related.R.ChannelDChannelOverwrites.
func (o *DChannelOverwrite) SetChannel(exec boil.Executor, insert bool, related *DChannel) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"d_channel_overwrites\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"channel_id"}),
		strmangle.WhereClause("\"", "\"", 2, dChannelOverwritePrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID, o.ChannelID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.ChannelID = related.ID

	if o.R == nil {
		o.R = &dChannelOverwriteR{
			Channel: related,
		}
	} else {
		o.R.Channel = related
	}

	if related.R == nil {
		related.R = &dChannelR{
			ChannelDChannelOverwrites: DChannelOverwriteSlice{o},
		}
	} else {
		related.R.ChannelDChannelOverwrites = append(related.R.ChannelDChannelOverwrites, o)
	}

	return nil
}

// DChannelOverwritesG retrieves all records.
func DChannelOverwritesG(mods ...qm.QueryMod) dChannelOverwriteQuery {
	return DChannelOverwrites(boil.GetDB(), mods...)
}

// DChannelOverwrites retrieves all the records using an executor.
func DChannelOverwrites(exec boil.Executor, mods ...qm.QueryMod) dChannelOverwriteQuery {
	mods = append(mods, qm.From("\"d_channel_overwrites\""))
	return dChannelOverwriteQuery{NewQuery(exec, mods...)}
}

// FindDChannelOverwriteG retrieves a single record by ID.
func FindDChannelOverwriteG(id int64, channelID int64, selectCols ...string) (*DChannelOverwrite, error) {
	return FindDChannelOverwrite(boil.GetDB(), id, channelID, selectCols...)
}

// FindDChannelOverwriteGP retrieves a single record by ID, and panics on error.
func FindDChannelOverwriteGP(id int64, channelID int64, selectCols ...string) *DChannelOverwrite {
	retobj, err := FindDChannelOverwrite(boil.GetDB(), id, channelID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindDChannelOverwrite retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDChannelOverwrite(exec boil.Executor, id int64, channelID int64, selectCols ...string) (*DChannelOverwrite, error) {
	dChannelOverwriteObj := &DChannelOverwrite{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"d_channel_overwrites\" where \"id\"=$1 AND \"channel_id\"=$2", sel,
	)

	q := queries.Raw(exec, query, id, channelID)

	err := q.Bind(dChannelOverwriteObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from d_channel_overwrites")
	}

	return dChannelOverwriteObj, nil
}

// FindDChannelOverwriteP retrieves a single record by ID with an executor, and panics on error.
func FindDChannelOverwriteP(exec boil.Executor, id int64, channelID int64, selectCols ...string) *DChannelOverwrite {
	retobj, err := FindDChannelOverwrite(exec, id, channelID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *DChannelOverwrite) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *DChannelOverwrite) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *DChannelOverwrite) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *DChannelOverwrite) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no d_channel_overwrites provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(dChannelOverwriteColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	dChannelOverwriteInsertCacheMut.RLock()
	cache, cached := dChannelOverwriteInsertCache[key]
	dChannelOverwriteInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			dChannelOverwriteColumns,
			dChannelOverwriteColumnsWithDefault,
			dChannelOverwriteColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(dChannelOverwriteType, dChannelOverwriteMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(dChannelOverwriteType, dChannelOverwriteMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"d_channel_overwrites\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "models: unable to insert into d_channel_overwrites")
	}

	if !cached {
		dChannelOverwriteInsertCacheMut.Lock()
		dChannelOverwriteInsertCache[key] = cache
		dChannelOverwriteInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single DChannelOverwrite record. See Update for
// whitelist behavior description.
func (o *DChannelOverwrite) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single DChannelOverwrite record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *DChannelOverwrite) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the DChannelOverwrite, and panics on error.
// See Update for whitelist behavior description.
func (o *DChannelOverwrite) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the DChannelOverwrite.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *DChannelOverwrite) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	key := makeCacheKey(whitelist, nil)
	dChannelOverwriteUpdateCacheMut.RLock()
	cache, cached := dChannelOverwriteUpdateCache[key]
	dChannelOverwriteUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(dChannelOverwriteColumns, dChannelOverwritePrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update d_channel_overwrites, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"d_channel_overwrites\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, dChannelOverwritePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(dChannelOverwriteType, dChannelOverwriteMapping, append(wl, dChannelOverwritePrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update d_channel_overwrites row")
	}

	if !cached {
		dChannelOverwriteUpdateCacheMut.Lock()
		dChannelOverwriteUpdateCache[key] = cache
		dChannelOverwriteUpdateCacheMut.Unlock()
	}

	return nil
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q dChannelOverwriteQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q dChannelOverwriteQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for d_channel_overwrites")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o DChannelOverwriteSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o DChannelOverwriteSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o DChannelOverwriteSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o DChannelOverwriteSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dChannelOverwritePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"d_channel_overwrites\" SET %s WHERE (\"id\",\"channel_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(dChannelOverwritePrimaryKeyColumns), len(colNames)+1, len(dChannelOverwritePrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in dChannelOverwrite slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *DChannelOverwrite) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *DChannelOverwrite) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *DChannelOverwrite) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *DChannelOverwrite) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no d_channel_overwrites provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(dChannelOverwriteColumnsWithDefault, o)

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

	dChannelOverwriteUpsertCacheMut.RLock()
	cache, cached := dChannelOverwriteUpsertCache[key]
	dChannelOverwriteUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			dChannelOverwriteColumns,
			dChannelOverwriteColumnsWithDefault,
			dChannelOverwriteColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			dChannelOverwriteColumns,
			dChannelOverwritePrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert d_channel_overwrites, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(dChannelOverwritePrimaryKeyColumns))
			copy(conflict, dChannelOverwritePrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"d_channel_overwrites\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(dChannelOverwriteType, dChannelOverwriteMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(dChannelOverwriteType, dChannelOverwriteMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert d_channel_overwrites")
	}

	if !cached {
		dChannelOverwriteUpsertCacheMut.Lock()
		dChannelOverwriteUpsertCache[key] = cache
		dChannelOverwriteUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteP deletes a single DChannelOverwrite record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *DChannelOverwrite) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single DChannelOverwrite record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *DChannelOverwrite) DeleteG() error {
	if o == nil {
		return errors.New("models: no DChannelOverwrite provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single DChannelOverwrite record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *DChannelOverwrite) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single DChannelOverwrite record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *DChannelOverwrite) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no DChannelOverwrite provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), dChannelOverwritePrimaryKeyMapping)
	sql := "DELETE FROM \"d_channel_overwrites\" WHERE \"id\"=$1 AND \"channel_id\"=$2"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from d_channel_overwrites")
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q dChannelOverwriteQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q dChannelOverwriteQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no dChannelOverwriteQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from d_channel_overwrites")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o DChannelOverwriteSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o DChannelOverwriteSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no DChannelOverwrite slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o DChannelOverwriteSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o DChannelOverwriteSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no DChannelOverwrite slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dChannelOverwritePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"d_channel_overwrites\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, dChannelOverwritePrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(dChannelOverwritePrimaryKeyColumns), 1, len(dChannelOverwritePrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from dChannelOverwrite slice")
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *DChannelOverwrite) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *DChannelOverwrite) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *DChannelOverwrite) ReloadG() error {
	if o == nil {
		return errors.New("models: no DChannelOverwrite provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *DChannelOverwrite) Reload(exec boil.Executor) error {
	ret, err := FindDChannelOverwrite(exec, o.ID, o.ChannelID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *DChannelOverwriteSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *DChannelOverwriteSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DChannelOverwriteSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty DChannelOverwriteSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DChannelOverwriteSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	dChannelOverwrites := DChannelOverwriteSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dChannelOverwritePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"d_channel_overwrites\".* FROM \"d_channel_overwrites\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, dChannelOverwritePrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(dChannelOverwritePrimaryKeyColumns), 1, len(dChannelOverwritePrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&dChannelOverwrites)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in DChannelOverwriteSlice")
	}

	*o = dChannelOverwrites

	return nil
}

// DChannelOverwriteExists checks if the DChannelOverwrite row exists.
func DChannelOverwriteExists(exec boil.Executor, id int64, channelID int64) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"d_channel_overwrites\" where \"id\"=$1 AND \"channel_id\"=$2 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, id, channelID)
	}

	row := exec.QueryRow(sql, id, channelID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if d_channel_overwrites exists")
	}

	return exists, nil
}

// DChannelOverwriteExistsG checks if the DChannelOverwrite row exists.
func DChannelOverwriteExistsG(id int64, channelID int64) (bool, error) {
	return DChannelOverwriteExists(boil.GetDB(), id, channelID)
}

// DChannelOverwriteExistsGP checks if the DChannelOverwrite row exists. Panics on error.
func DChannelOverwriteExistsGP(id int64, channelID int64) bool {
	e, err := DChannelOverwriteExists(boil.GetDB(), id, channelID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// DChannelOverwriteExistsP checks if the DChannelOverwrite row exists. Panics on error.
func DChannelOverwriteExistsP(exec boil.Executor, id int64, channelID int64) bool {
	e, err := DChannelOverwriteExists(exec, id, channelID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
