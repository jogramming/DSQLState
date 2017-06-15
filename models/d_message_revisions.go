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
)

// DMessageRevision is an object representing the database table.
type DMessageRevision struct {
	RevisionNum  int              `boil:"revision_num" json:"revision_num" toml:"revision_num" yaml:"revision_num"`
	MessageID    int64            `boil:"message_id" json:"message_id" toml:"message_id" yaml:"message_id"`
	CreatedAt    time.Time        `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	Content      string           `boil:"content" json:"content" toml:"content" yaml:"content"`
	Embeds       types.Int64Array `boil:"embeds" json:"embeds" toml:"embeds" yaml:"embeds"`
	Mentions     types.Int64Array `boil:"mentions" json:"mentions" toml:"mentions" yaml:"mentions"`
	MentionRoles types.Int64Array `boil:"mention_roles" json:"mention_roles" toml:"mention_roles" yaml:"mention_roles"`

	R *dMessageRevisionR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L dMessageRevisionL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// dMessageRevisionR is where relationships are stored.
type dMessageRevisionR struct {
	Message *DMessage
}

// dMessageRevisionL is where Load methods for each relationship are stored.
type dMessageRevisionL struct{}

var (
	dMessageRevisionColumns               = []string{"revision_num", "message_id", "created_at", "content", "embeds", "mentions", "mention_roles"}
	dMessageRevisionColumnsWithoutDefault = []string{"revision_num", "message_id", "created_at", "content", "embeds", "mentions", "mention_roles"}
	dMessageRevisionColumnsWithDefault    = []string{}
	dMessageRevisionPrimaryKeyColumns     = []string{"revision_num", "message_id"}
)

type (
	// DMessageRevisionSlice is an alias for a slice of pointers to DMessageRevision.
	// This should generally be used opposed to []DMessageRevision.
	DMessageRevisionSlice []*DMessageRevision

	dMessageRevisionQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	dMessageRevisionType                 = reflect.TypeOf(&DMessageRevision{})
	dMessageRevisionMapping              = queries.MakeStructMapping(dMessageRevisionType)
	dMessageRevisionPrimaryKeyMapping, _ = queries.BindMapping(dMessageRevisionType, dMessageRevisionMapping, dMessageRevisionPrimaryKeyColumns)
	dMessageRevisionInsertCacheMut       sync.RWMutex
	dMessageRevisionInsertCache          = make(map[string]insertCache)
	dMessageRevisionUpdateCacheMut       sync.RWMutex
	dMessageRevisionUpdateCache          = make(map[string]updateCache)
	dMessageRevisionUpsertCacheMut       sync.RWMutex
	dMessageRevisionUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)

// OneP returns a single dMessageRevision record from the query, and panics on error.
func (q dMessageRevisionQuery) OneP() *DMessageRevision {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single dMessageRevision record from the query.
func (q dMessageRevisionQuery) One() (*DMessageRevision, error) {
	o := &DMessageRevision{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for d_message_revisions")
	}

	return o, nil
}

// AllP returns all DMessageRevision records from the query, and panics on error.
func (q dMessageRevisionQuery) AllP() DMessageRevisionSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all DMessageRevision records from the query.
func (q dMessageRevisionQuery) All() (DMessageRevisionSlice, error) {
	var o DMessageRevisionSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to DMessageRevision slice")
	}

	return o, nil
}

// CountP returns the count of all DMessageRevision records in the query, and panics on error.
func (q dMessageRevisionQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all DMessageRevision records in the query.
func (q dMessageRevisionQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count d_message_revisions rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q dMessageRevisionQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q dMessageRevisionQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if d_message_revisions exists")
	}

	return count > 0, nil
}

// MessageG pointed to by the foreign key.
func (o *DMessageRevision) MessageG(mods ...qm.QueryMod) dMessageQuery {
	return o.Message(boil.GetDB(), mods...)
}

// Message pointed to by the foreign key.
func (o *DMessageRevision) Message(exec boil.Executor, mods ...qm.QueryMod) dMessageQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.MessageID),
	}

	queryMods = append(queryMods, mods...)

	query := DMessages(exec, queryMods...)
	queries.SetFrom(query.Query, "\"d_messages\"")

	return query
}

// LoadMessage allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (dMessageRevisionL) LoadMessage(e boil.Executor, singular bool, maybeDMessageRevision interface{}) error {
	var slice []*DMessageRevision
	var object *DMessageRevision

	count := 1
	if singular {
		object = maybeDMessageRevision.(*DMessageRevision)
	} else {
		slice = *maybeDMessageRevision.(*DMessageRevisionSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		if object.R == nil {
			object.R = &dMessageRevisionR{}
		}
		args[0] = object.MessageID
	} else {
		for i, obj := range slice {
			if obj.R == nil {
				obj.R = &dMessageRevisionR{}
			}
			args[i] = obj.MessageID
		}
	}

	query := fmt.Sprintf(
		"select * from \"d_messages\" where \"id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load DMessage")
	}
	defer results.Close()

	var resultSlice []*DMessage
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice DMessage")
	}

	if singular && len(resultSlice) != 0 {
		object.R.Message = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.MessageID == foreign.ID {
				local.R.Message = foreign
				break
			}
		}
	}

	return nil
}

// SetMessageG of the d_message_revision to the related item.
// Sets o.R.Message to related.
// Adds o to related.R.MessageDMessageRevisions.
// Uses the global database handle.
func (o *DMessageRevision) SetMessageG(insert bool, related *DMessage) error {
	return o.SetMessage(boil.GetDB(), insert, related)
}

// SetMessageP of the d_message_revision to the related item.
// Sets o.R.Message to related.
// Adds o to related.R.MessageDMessageRevisions.
// Panics on error.
func (o *DMessageRevision) SetMessageP(exec boil.Executor, insert bool, related *DMessage) {
	if err := o.SetMessage(exec, insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetMessageGP of the d_message_revision to the related item.
// Sets o.R.Message to related.
// Adds o to related.R.MessageDMessageRevisions.
// Uses the global database handle and panics on error.
func (o *DMessageRevision) SetMessageGP(insert bool, related *DMessage) {
	if err := o.SetMessage(boil.GetDB(), insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetMessage of the d_message_revision to the related item.
// Sets o.R.Message to related.
// Adds o to related.R.MessageDMessageRevisions.
func (o *DMessageRevision) SetMessage(exec boil.Executor, insert bool, related *DMessage) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"d_message_revisions\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"message_id"}),
		strmangle.WhereClause("\"", "\"", 2, dMessageRevisionPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.RevisionNum, o.MessageID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.MessageID = related.ID

	if o.R == nil {
		o.R = &dMessageRevisionR{
			Message: related,
		}
	} else {
		o.R.Message = related
	}

	if related.R == nil {
		related.R = &dMessageR{
			MessageDMessageRevisions: DMessageRevisionSlice{o},
		}
	} else {
		related.R.MessageDMessageRevisions = append(related.R.MessageDMessageRevisions, o)
	}

	return nil
}

// DMessageRevisionsG retrieves all records.
func DMessageRevisionsG(mods ...qm.QueryMod) dMessageRevisionQuery {
	return DMessageRevisions(boil.GetDB(), mods...)
}

// DMessageRevisions retrieves all the records using an executor.
func DMessageRevisions(exec boil.Executor, mods ...qm.QueryMod) dMessageRevisionQuery {
	mods = append(mods, qm.From("\"d_message_revisions\""))
	return dMessageRevisionQuery{NewQuery(exec, mods...)}
}

// FindDMessageRevisionG retrieves a single record by ID.
func FindDMessageRevisionG(revisionNum int, messageID int64, selectCols ...string) (*DMessageRevision, error) {
	return FindDMessageRevision(boil.GetDB(), revisionNum, messageID, selectCols...)
}

// FindDMessageRevisionGP retrieves a single record by ID, and panics on error.
func FindDMessageRevisionGP(revisionNum int, messageID int64, selectCols ...string) *DMessageRevision {
	retobj, err := FindDMessageRevision(boil.GetDB(), revisionNum, messageID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindDMessageRevision retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDMessageRevision(exec boil.Executor, revisionNum int, messageID int64, selectCols ...string) (*DMessageRevision, error) {
	dMessageRevisionObj := &DMessageRevision{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"d_message_revisions\" where \"revision_num\"=$1 AND \"message_id\"=$2", sel,
	)

	q := queries.Raw(exec, query, revisionNum, messageID)

	err := q.Bind(dMessageRevisionObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from d_message_revisions")
	}

	return dMessageRevisionObj, nil
}

// FindDMessageRevisionP retrieves a single record by ID with an executor, and panics on error.
func FindDMessageRevisionP(exec boil.Executor, revisionNum int, messageID int64, selectCols ...string) *DMessageRevision {
	retobj, err := FindDMessageRevision(exec, revisionNum, messageID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *DMessageRevision) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *DMessageRevision) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *DMessageRevision) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *DMessageRevision) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no d_message_revisions provided for insertion")
	}

	var err error
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(dMessageRevisionColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	dMessageRevisionInsertCacheMut.RLock()
	cache, cached := dMessageRevisionInsertCache[key]
	dMessageRevisionInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			dMessageRevisionColumns,
			dMessageRevisionColumnsWithDefault,
			dMessageRevisionColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(dMessageRevisionType, dMessageRevisionMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(dMessageRevisionType, dMessageRevisionMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"d_message_revisions\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "models: unable to insert into d_message_revisions")
	}

	if !cached {
		dMessageRevisionInsertCacheMut.Lock()
		dMessageRevisionInsertCache[key] = cache
		dMessageRevisionInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single DMessageRevision record. See Update for
// whitelist behavior description.
func (o *DMessageRevision) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single DMessageRevision record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *DMessageRevision) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the DMessageRevision, and panics on error.
// See Update for whitelist behavior description.
func (o *DMessageRevision) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the DMessageRevision.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *DMessageRevision) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	key := makeCacheKey(whitelist, nil)
	dMessageRevisionUpdateCacheMut.RLock()
	cache, cached := dMessageRevisionUpdateCache[key]
	dMessageRevisionUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(dMessageRevisionColumns, dMessageRevisionPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update d_message_revisions, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"d_message_revisions\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, dMessageRevisionPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(dMessageRevisionType, dMessageRevisionMapping, append(wl, dMessageRevisionPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update d_message_revisions row")
	}

	if !cached {
		dMessageRevisionUpdateCacheMut.Lock()
		dMessageRevisionUpdateCache[key] = cache
		dMessageRevisionUpdateCacheMut.Unlock()
	}

	return nil
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q dMessageRevisionQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q dMessageRevisionQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for d_message_revisions")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o DMessageRevisionSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o DMessageRevisionSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o DMessageRevisionSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o DMessageRevisionSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dMessageRevisionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"d_message_revisions\" SET %s WHERE (\"revision_num\",\"message_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(dMessageRevisionPrimaryKeyColumns), len(colNames)+1, len(dMessageRevisionPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in dMessageRevision slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *DMessageRevision) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *DMessageRevision) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *DMessageRevision) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *DMessageRevision) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no d_message_revisions provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(dMessageRevisionColumnsWithDefault, o)

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

	dMessageRevisionUpsertCacheMut.RLock()
	cache, cached := dMessageRevisionUpsertCache[key]
	dMessageRevisionUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			dMessageRevisionColumns,
			dMessageRevisionColumnsWithDefault,
			dMessageRevisionColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			dMessageRevisionColumns,
			dMessageRevisionPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert d_message_revisions, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(dMessageRevisionPrimaryKeyColumns))
			copy(conflict, dMessageRevisionPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"d_message_revisions\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(dMessageRevisionType, dMessageRevisionMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(dMessageRevisionType, dMessageRevisionMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert d_message_revisions")
	}

	if !cached {
		dMessageRevisionUpsertCacheMut.Lock()
		dMessageRevisionUpsertCache[key] = cache
		dMessageRevisionUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteP deletes a single DMessageRevision record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *DMessageRevision) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single DMessageRevision record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *DMessageRevision) DeleteG() error {
	if o == nil {
		return errors.New("models: no DMessageRevision provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single DMessageRevision record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *DMessageRevision) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single DMessageRevision record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *DMessageRevision) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no DMessageRevision provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), dMessageRevisionPrimaryKeyMapping)
	sql := "DELETE FROM \"d_message_revisions\" WHERE \"revision_num\"=$1 AND \"message_id\"=$2"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from d_message_revisions")
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q dMessageRevisionQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q dMessageRevisionQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no dMessageRevisionQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from d_message_revisions")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o DMessageRevisionSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o DMessageRevisionSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no DMessageRevision slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o DMessageRevisionSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o DMessageRevisionSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no DMessageRevision slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dMessageRevisionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"d_message_revisions\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, dMessageRevisionPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(dMessageRevisionPrimaryKeyColumns), 1, len(dMessageRevisionPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from dMessageRevision slice")
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *DMessageRevision) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *DMessageRevision) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *DMessageRevision) ReloadG() error {
	if o == nil {
		return errors.New("models: no DMessageRevision provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *DMessageRevision) Reload(exec boil.Executor) error {
	ret, err := FindDMessageRevision(exec, o.RevisionNum, o.MessageID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *DMessageRevisionSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *DMessageRevisionSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DMessageRevisionSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty DMessageRevisionSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DMessageRevisionSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	dMessageRevisions := DMessageRevisionSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dMessageRevisionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"d_message_revisions\".* FROM \"d_message_revisions\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, dMessageRevisionPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(dMessageRevisionPrimaryKeyColumns), 1, len(dMessageRevisionPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&dMessageRevisions)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in DMessageRevisionSlice")
	}

	*o = dMessageRevisions

	return nil
}

// DMessageRevisionExists checks if the DMessageRevision row exists.
func DMessageRevisionExists(exec boil.Executor, revisionNum int, messageID int64) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"d_message_revisions\" where \"revision_num\"=$1 AND \"message_id\"=$2 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, revisionNum, messageID)
	}

	row := exec.QueryRow(sql, revisionNum, messageID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if d_message_revisions exists")
	}

	return exists, nil
}

// DMessageRevisionExistsG checks if the DMessageRevision row exists.
func DMessageRevisionExistsG(revisionNum int, messageID int64) (bool, error) {
	return DMessageRevisionExists(boil.GetDB(), revisionNum, messageID)
}

// DMessageRevisionExistsGP checks if the DMessageRevision row exists. Panics on error.
func DMessageRevisionExistsGP(revisionNum int, messageID int64) bool {
	e, err := DMessageRevisionExists(boil.GetDB(), revisionNum, messageID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// DMessageRevisionExistsP checks if the DMessageRevision row exists. Panics on error.
func DMessageRevisionExistsP(exec boil.Executor, revisionNum int, messageID int64) bool {
	e, err := DMessageRevisionExists(exec, revisionNum, messageID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
