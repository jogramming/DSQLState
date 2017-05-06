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

// DiscordMessageRevision is an object representing the database table.
type DiscordMessageRevision struct {
	RevisionNum int              `boil:"revision_num" json:"revision_num" toml:"revision_num" yaml:"revision_num"`
	MessageID   int64            `boil:"message_id" json:"message_id" toml:"message_id" yaml:"message_id"`
	CreatedAt   time.Time        `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	Content     string           `boil:"content" json:"content" toml:"content" yaml:"content"`
	Embeds      types.Int64Array `boil:"embeds" json:"embeds" toml:"embeds" yaml:"embeds"`

	R *discordMessageRevisionR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L discordMessageRevisionL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// discordMessageRevisionR is where relationships are stored.
type discordMessageRevisionR struct {
	Message *DiscordMessage
}

// discordMessageRevisionL is where Load methods for each relationship are stored.
type discordMessageRevisionL struct{}

var (
	discordMessageRevisionColumns               = []string{"revision_num", "message_id", "created_at", "content", "embeds"}
	discordMessageRevisionColumnsWithoutDefault = []string{"revision_num", "message_id", "created_at", "content", "embeds"}
	discordMessageRevisionColumnsWithDefault    = []string{}
	discordMessageRevisionPrimaryKeyColumns     = []string{"revision_num", "message_id"}
)

type (
	// DiscordMessageRevisionSlice is an alias for a slice of pointers to DiscordMessageRevision.
	// This should generally be used opposed to []DiscordMessageRevision.
	DiscordMessageRevisionSlice []*DiscordMessageRevision

	discordMessageRevisionQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	discordMessageRevisionType                 = reflect.TypeOf(&DiscordMessageRevision{})
	discordMessageRevisionMapping              = queries.MakeStructMapping(discordMessageRevisionType)
	discordMessageRevisionPrimaryKeyMapping, _ = queries.BindMapping(discordMessageRevisionType, discordMessageRevisionMapping, discordMessageRevisionPrimaryKeyColumns)
	discordMessageRevisionInsertCacheMut       sync.RWMutex
	discordMessageRevisionInsertCache          = make(map[string]insertCache)
	discordMessageRevisionUpdateCacheMut       sync.RWMutex
	discordMessageRevisionUpdateCache          = make(map[string]updateCache)
	discordMessageRevisionUpsertCacheMut       sync.RWMutex
	discordMessageRevisionUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)

// OneP returns a single discordMessageRevision record from the query, and panics on error.
func (q discordMessageRevisionQuery) OneP() *DiscordMessageRevision {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single discordMessageRevision record from the query.
func (q discordMessageRevisionQuery) One() (*DiscordMessageRevision, error) {
	o := &DiscordMessageRevision{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for discord_message_revisions")
	}

	return o, nil
}

// AllP returns all DiscordMessageRevision records from the query, and panics on error.
func (q discordMessageRevisionQuery) AllP() DiscordMessageRevisionSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all DiscordMessageRevision records from the query.
func (q discordMessageRevisionQuery) All() (DiscordMessageRevisionSlice, error) {
	var o DiscordMessageRevisionSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to DiscordMessageRevision slice")
	}

	return o, nil
}

// CountP returns the count of all DiscordMessageRevision records in the query, and panics on error.
func (q discordMessageRevisionQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all DiscordMessageRevision records in the query.
func (q discordMessageRevisionQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count discord_message_revisions rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q discordMessageRevisionQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q discordMessageRevisionQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if discord_message_revisions exists")
	}

	return count > 0, nil
}

// MessageG pointed to by the foreign key.
func (o *DiscordMessageRevision) MessageG(mods ...qm.QueryMod) discordMessageQuery {
	return o.Message(boil.GetDB(), mods...)
}

// Message pointed to by the foreign key.
func (o *DiscordMessageRevision) Message(exec boil.Executor, mods ...qm.QueryMod) discordMessageQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.MessageID),
	}

	queryMods = append(queryMods, mods...)

	query := DiscordMessages(exec, queryMods...)
	queries.SetFrom(query.Query, "\"discord_messages\"")

	return query
}

// LoadMessage allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (discordMessageRevisionL) LoadMessage(e boil.Executor, singular bool, maybeDiscordMessageRevision interface{}) error {
	var slice []*DiscordMessageRevision
	var object *DiscordMessageRevision

	count := 1
	if singular {
		object = maybeDiscordMessageRevision.(*DiscordMessageRevision)
	} else {
		slice = *maybeDiscordMessageRevision.(*DiscordMessageRevisionSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		if object.R == nil {
			object.R = &discordMessageRevisionR{}
		}
		args[0] = object.MessageID
	} else {
		for i, obj := range slice {
			if obj.R == nil {
				obj.R = &discordMessageRevisionR{}
			}
			args[i] = obj.MessageID
		}
	}

	query := fmt.Sprintf(
		"select * from \"discord_messages\" where \"id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load DiscordMessage")
	}
	defer results.Close()

	var resultSlice []*DiscordMessage
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice DiscordMessage")
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

// SetMessageG of the discord_message_revision to the related item.
// Sets o.R.Message to related.
// Adds o to related.R.MessageDiscordMessageRevisions.
// Uses the global database handle.
func (o *DiscordMessageRevision) SetMessageG(insert bool, related *DiscordMessage) error {
	return o.SetMessage(boil.GetDB(), insert, related)
}

// SetMessageP of the discord_message_revision to the related item.
// Sets o.R.Message to related.
// Adds o to related.R.MessageDiscordMessageRevisions.
// Panics on error.
func (o *DiscordMessageRevision) SetMessageP(exec boil.Executor, insert bool, related *DiscordMessage) {
	if err := o.SetMessage(exec, insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetMessageGP of the discord_message_revision to the related item.
// Sets o.R.Message to related.
// Adds o to related.R.MessageDiscordMessageRevisions.
// Uses the global database handle and panics on error.
func (o *DiscordMessageRevision) SetMessageGP(insert bool, related *DiscordMessage) {
	if err := o.SetMessage(boil.GetDB(), insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetMessage of the discord_message_revision to the related item.
// Sets o.R.Message to related.
// Adds o to related.R.MessageDiscordMessageRevisions.
func (o *DiscordMessageRevision) SetMessage(exec boil.Executor, insert bool, related *DiscordMessage) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"discord_message_revisions\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"message_id"}),
		strmangle.WhereClause("\"", "\"", 2, discordMessageRevisionPrimaryKeyColumns),
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
		o.R = &discordMessageRevisionR{
			Message: related,
		}
	} else {
		o.R.Message = related
	}

	if related.R == nil {
		related.R = &discordMessageR{
			MessageDiscordMessageRevisions: DiscordMessageRevisionSlice{o},
		}
	} else {
		related.R.MessageDiscordMessageRevisions = append(related.R.MessageDiscordMessageRevisions, o)
	}

	return nil
}

// DiscordMessageRevisionsG retrieves all records.
func DiscordMessageRevisionsG(mods ...qm.QueryMod) discordMessageRevisionQuery {
	return DiscordMessageRevisions(boil.GetDB(), mods...)
}

// DiscordMessageRevisions retrieves all the records using an executor.
func DiscordMessageRevisions(exec boil.Executor, mods ...qm.QueryMod) discordMessageRevisionQuery {
	mods = append(mods, qm.From("\"discord_message_revisions\""))
	return discordMessageRevisionQuery{NewQuery(exec, mods...)}
}

// FindDiscordMessageRevisionG retrieves a single record by ID.
func FindDiscordMessageRevisionG(revisionNum int, messageID int64, selectCols ...string) (*DiscordMessageRevision, error) {
	return FindDiscordMessageRevision(boil.GetDB(), revisionNum, messageID, selectCols...)
}

// FindDiscordMessageRevisionGP retrieves a single record by ID, and panics on error.
func FindDiscordMessageRevisionGP(revisionNum int, messageID int64, selectCols ...string) *DiscordMessageRevision {
	retobj, err := FindDiscordMessageRevision(boil.GetDB(), revisionNum, messageID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindDiscordMessageRevision retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDiscordMessageRevision(exec boil.Executor, revisionNum int, messageID int64, selectCols ...string) (*DiscordMessageRevision, error) {
	discordMessageRevisionObj := &DiscordMessageRevision{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"discord_message_revisions\" where \"revision_num\"=$1 AND \"message_id\"=$2", sel,
	)

	q := queries.Raw(exec, query, revisionNum, messageID)

	err := q.Bind(discordMessageRevisionObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from discord_message_revisions")
	}

	return discordMessageRevisionObj, nil
}

// FindDiscordMessageRevisionP retrieves a single record by ID with an executor, and panics on error.
func FindDiscordMessageRevisionP(exec boil.Executor, revisionNum int, messageID int64, selectCols ...string) *DiscordMessageRevision {
	retobj, err := FindDiscordMessageRevision(exec, revisionNum, messageID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *DiscordMessageRevision) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *DiscordMessageRevision) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *DiscordMessageRevision) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *DiscordMessageRevision) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no discord_message_revisions provided for insertion")
	}

	var err error
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(discordMessageRevisionColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	discordMessageRevisionInsertCacheMut.RLock()
	cache, cached := discordMessageRevisionInsertCache[key]
	discordMessageRevisionInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			discordMessageRevisionColumns,
			discordMessageRevisionColumnsWithDefault,
			discordMessageRevisionColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(discordMessageRevisionType, discordMessageRevisionMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(discordMessageRevisionType, discordMessageRevisionMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"discord_message_revisions\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "models: unable to insert into discord_message_revisions")
	}

	if !cached {
		discordMessageRevisionInsertCacheMut.Lock()
		discordMessageRevisionInsertCache[key] = cache
		discordMessageRevisionInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single DiscordMessageRevision record. See Update for
// whitelist behavior description.
func (o *DiscordMessageRevision) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single DiscordMessageRevision record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *DiscordMessageRevision) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the DiscordMessageRevision, and panics on error.
// See Update for whitelist behavior description.
func (o *DiscordMessageRevision) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the DiscordMessageRevision.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *DiscordMessageRevision) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	key := makeCacheKey(whitelist, nil)
	discordMessageRevisionUpdateCacheMut.RLock()
	cache, cached := discordMessageRevisionUpdateCache[key]
	discordMessageRevisionUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(discordMessageRevisionColumns, discordMessageRevisionPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update discord_message_revisions, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"discord_message_revisions\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, discordMessageRevisionPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(discordMessageRevisionType, discordMessageRevisionMapping, append(wl, discordMessageRevisionPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update discord_message_revisions row")
	}

	if !cached {
		discordMessageRevisionUpdateCacheMut.Lock()
		discordMessageRevisionUpdateCache[key] = cache
		discordMessageRevisionUpdateCacheMut.Unlock()
	}

	return nil
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q discordMessageRevisionQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q discordMessageRevisionQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for discord_message_revisions")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o DiscordMessageRevisionSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o DiscordMessageRevisionSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o DiscordMessageRevisionSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o DiscordMessageRevisionSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), discordMessageRevisionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"discord_message_revisions\" SET %s WHERE (\"revision_num\",\"message_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(discordMessageRevisionPrimaryKeyColumns), len(colNames)+1, len(discordMessageRevisionPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in discordMessageRevision slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *DiscordMessageRevision) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *DiscordMessageRevision) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *DiscordMessageRevision) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *DiscordMessageRevision) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no discord_message_revisions provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(discordMessageRevisionColumnsWithDefault, o)

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

	discordMessageRevisionUpsertCacheMut.RLock()
	cache, cached := discordMessageRevisionUpsertCache[key]
	discordMessageRevisionUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			discordMessageRevisionColumns,
			discordMessageRevisionColumnsWithDefault,
			discordMessageRevisionColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			discordMessageRevisionColumns,
			discordMessageRevisionPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert discord_message_revisions, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(discordMessageRevisionPrimaryKeyColumns))
			copy(conflict, discordMessageRevisionPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"discord_message_revisions\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(discordMessageRevisionType, discordMessageRevisionMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(discordMessageRevisionType, discordMessageRevisionMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert discord_message_revisions")
	}

	if !cached {
		discordMessageRevisionUpsertCacheMut.Lock()
		discordMessageRevisionUpsertCache[key] = cache
		discordMessageRevisionUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteP deletes a single DiscordMessageRevision record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *DiscordMessageRevision) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single DiscordMessageRevision record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *DiscordMessageRevision) DeleteG() error {
	if o == nil {
		return errors.New("models: no DiscordMessageRevision provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single DiscordMessageRevision record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *DiscordMessageRevision) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single DiscordMessageRevision record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *DiscordMessageRevision) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no DiscordMessageRevision provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), discordMessageRevisionPrimaryKeyMapping)
	sql := "DELETE FROM \"discord_message_revisions\" WHERE \"revision_num\"=$1 AND \"message_id\"=$2"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from discord_message_revisions")
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q discordMessageRevisionQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q discordMessageRevisionQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no discordMessageRevisionQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from discord_message_revisions")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o DiscordMessageRevisionSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o DiscordMessageRevisionSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no DiscordMessageRevision slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o DiscordMessageRevisionSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o DiscordMessageRevisionSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no DiscordMessageRevision slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), discordMessageRevisionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"discord_message_revisions\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, discordMessageRevisionPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(discordMessageRevisionPrimaryKeyColumns), 1, len(discordMessageRevisionPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from discordMessageRevision slice")
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *DiscordMessageRevision) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *DiscordMessageRevision) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *DiscordMessageRevision) ReloadG() error {
	if o == nil {
		return errors.New("models: no DiscordMessageRevision provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *DiscordMessageRevision) Reload(exec boil.Executor) error {
	ret, err := FindDiscordMessageRevision(exec, o.RevisionNum, o.MessageID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *DiscordMessageRevisionSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *DiscordMessageRevisionSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DiscordMessageRevisionSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty DiscordMessageRevisionSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DiscordMessageRevisionSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	discordMessageRevisions := DiscordMessageRevisionSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), discordMessageRevisionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"discord_message_revisions\".* FROM \"discord_message_revisions\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, discordMessageRevisionPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(discordMessageRevisionPrimaryKeyColumns), 1, len(discordMessageRevisionPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&discordMessageRevisions)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in DiscordMessageRevisionSlice")
	}

	*o = discordMessageRevisions

	return nil
}

// DiscordMessageRevisionExists checks if the DiscordMessageRevision row exists.
func DiscordMessageRevisionExists(exec boil.Executor, revisionNum int, messageID int64) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"discord_message_revisions\" where \"revision_num\"=$1 AND \"message_id\"=$2 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, revisionNum, messageID)
	}

	row := exec.QueryRow(sql, revisionNum, messageID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if discord_message_revisions exists")
	}

	return exists, nil
}

// DiscordMessageRevisionExistsG checks if the DiscordMessageRevision row exists.
func DiscordMessageRevisionExistsG(revisionNum int, messageID int64) (bool, error) {
	return DiscordMessageRevisionExists(boil.GetDB(), revisionNum, messageID)
}

// DiscordMessageRevisionExistsGP checks if the DiscordMessageRevision row exists. Panics on error.
func DiscordMessageRevisionExistsGP(revisionNum int, messageID int64) bool {
	e, err := DiscordMessageRevisionExists(boil.GetDB(), revisionNum, messageID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// DiscordMessageRevisionExistsP checks if the DiscordMessageRevision row exists. Panics on error.
func DiscordMessageRevisionExistsP(exec boil.Executor, revisionNum int, messageID int64) bool {
	e, err := DiscordMessageRevisionExists(exec, revisionNum, messageID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
