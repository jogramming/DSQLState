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

// DVoiceState is an object representing the database table.
type DVoiceState struct {
	UserID    int64  `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	GuildID   int64  `boil:"guild_id" json:"guild_id" toml:"guild_id" yaml:"guild_id"`
	ChannelID int64  `boil:"channel_id" json:"channel_id" toml:"channel_id" yaml:"channel_id"`
	SessionID string `boil:"session_id" json:"session_id" toml:"session_id" yaml:"session_id"`
	Surpress  bool   `boil:"surpress" json:"surpress" toml:"surpress" yaml:"surpress"`
	SelfMute  bool   `boil:"self_mute" json:"self_mute" toml:"self_mute" yaml:"self_mute"`
	SelfDeaf  bool   `boil:"self_deaf" json:"self_deaf" toml:"self_deaf" yaml:"self_deaf"`
	Mute      bool   `boil:"mute" json:"mute" toml:"mute" yaml:"mute"`
	Deaf      bool   `boil:"deaf" json:"deaf" toml:"deaf" yaml:"deaf"`

	R *dVoiceStateR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L dVoiceStateL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// dVoiceStateR is where relationships are stored.
type dVoiceStateR struct {
	Channel *DChannel
}

// dVoiceStateL is where Load methods for each relationship are stored.
type dVoiceStateL struct{}

var (
	dVoiceStateColumns               = []string{"user_id", "guild_id", "channel_id", "session_id", "surpress", "self_mute", "self_deaf", "mute", "deaf"}
	dVoiceStateColumnsWithoutDefault = []string{"user_id", "guild_id", "channel_id", "session_id", "surpress", "self_mute", "self_deaf", "mute", "deaf"}
	dVoiceStateColumnsWithDefault    = []string{}
	dVoiceStatePrimaryKeyColumns     = []string{"user_id", "guild_id"}
)

type (
	// DVoiceStateSlice is an alias for a slice of pointers to DVoiceState.
	// This should generally be used opposed to []DVoiceState.
	DVoiceStateSlice []*DVoiceState

	dVoiceStateQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	dVoiceStateType                 = reflect.TypeOf(&DVoiceState{})
	dVoiceStateMapping              = queries.MakeStructMapping(dVoiceStateType)
	dVoiceStatePrimaryKeyMapping, _ = queries.BindMapping(dVoiceStateType, dVoiceStateMapping, dVoiceStatePrimaryKeyColumns)
	dVoiceStateInsertCacheMut       sync.RWMutex
	dVoiceStateInsertCache          = make(map[string]insertCache)
	dVoiceStateUpdateCacheMut       sync.RWMutex
	dVoiceStateUpdateCache          = make(map[string]updateCache)
	dVoiceStateUpsertCacheMut       sync.RWMutex
	dVoiceStateUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)

// OneP returns a single dVoiceState record from the query, and panics on error.
func (q dVoiceStateQuery) OneP() *DVoiceState {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single dVoiceState record from the query.
func (q dVoiceStateQuery) One() (*DVoiceState, error) {
	o := &DVoiceState{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for d_voice_states")
	}

	return o, nil
}

// AllP returns all DVoiceState records from the query, and panics on error.
func (q dVoiceStateQuery) AllP() DVoiceStateSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all DVoiceState records from the query.
func (q dVoiceStateQuery) All() (DVoiceStateSlice, error) {
	var o DVoiceStateSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to DVoiceState slice")
	}

	return o, nil
}

// CountP returns the count of all DVoiceState records in the query, and panics on error.
func (q dVoiceStateQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all DVoiceState records in the query.
func (q dVoiceStateQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count d_voice_states rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q dVoiceStateQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q dVoiceStateQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if d_voice_states exists")
	}

	return count > 0, nil
}

// ChannelG pointed to by the foreign key.
func (o *DVoiceState) ChannelG(mods ...qm.QueryMod) dChannelQuery {
	return o.Channel(boil.GetDB(), mods...)
}

// Channel pointed to by the foreign key.
func (o *DVoiceState) Channel(exec boil.Executor, mods ...qm.QueryMod) dChannelQuery {
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
func (dVoiceStateL) LoadChannel(e boil.Executor, singular bool, maybeDVoiceState interface{}) error {
	var slice []*DVoiceState
	var object *DVoiceState

	count := 1
	if singular {
		object = maybeDVoiceState.(*DVoiceState)
	} else {
		slice = *maybeDVoiceState.(*DVoiceStateSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		if object.R == nil {
			object.R = &dVoiceStateR{}
		}
		args[0] = object.ChannelID
	} else {
		for i, obj := range slice {
			if obj.R == nil {
				obj.R = &dVoiceStateR{}
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

// SetChannelG of the d_voice_state to the related item.
// Sets o.R.Channel to related.
// Adds o to related.R.ChannelDVoiceStates.
// Uses the global database handle.
func (o *DVoiceState) SetChannelG(insert bool, related *DChannel) error {
	return o.SetChannel(boil.GetDB(), insert, related)
}

// SetChannelP of the d_voice_state to the related item.
// Sets o.R.Channel to related.
// Adds o to related.R.ChannelDVoiceStates.
// Panics on error.
func (o *DVoiceState) SetChannelP(exec boil.Executor, insert bool, related *DChannel) {
	if err := o.SetChannel(exec, insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetChannelGP of the d_voice_state to the related item.
// Sets o.R.Channel to related.
// Adds o to related.R.ChannelDVoiceStates.
// Uses the global database handle and panics on error.
func (o *DVoiceState) SetChannelGP(insert bool, related *DChannel) {
	if err := o.SetChannel(boil.GetDB(), insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetChannel of the d_voice_state to the related item.
// Sets o.R.Channel to related.
// Adds o to related.R.ChannelDVoiceStates.
func (o *DVoiceState) SetChannel(exec boil.Executor, insert bool, related *DChannel) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"d_voice_states\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"channel_id"}),
		strmangle.WhereClause("\"", "\"", 2, dVoiceStatePrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.UserID, o.GuildID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.ChannelID = related.ID

	if o.R == nil {
		o.R = &dVoiceStateR{
			Channel: related,
		}
	} else {
		o.R.Channel = related
	}

	if related.R == nil {
		related.R = &dChannelR{
			ChannelDVoiceStates: DVoiceStateSlice{o},
		}
	} else {
		related.R.ChannelDVoiceStates = append(related.R.ChannelDVoiceStates, o)
	}

	return nil
}

// DVoiceStatesG retrieves all records.
func DVoiceStatesG(mods ...qm.QueryMod) dVoiceStateQuery {
	return DVoiceStates(boil.GetDB(), mods...)
}

// DVoiceStates retrieves all the records using an executor.
func DVoiceStates(exec boil.Executor, mods ...qm.QueryMod) dVoiceStateQuery {
	mods = append(mods, qm.From("\"d_voice_states\""))
	return dVoiceStateQuery{NewQuery(exec, mods...)}
}

// FindDVoiceStateG retrieves a single record by ID.
func FindDVoiceStateG(userID int64, guildID int64, selectCols ...string) (*DVoiceState, error) {
	return FindDVoiceState(boil.GetDB(), userID, guildID, selectCols...)
}

// FindDVoiceStateGP retrieves a single record by ID, and panics on error.
func FindDVoiceStateGP(userID int64, guildID int64, selectCols ...string) *DVoiceState {
	retobj, err := FindDVoiceState(boil.GetDB(), userID, guildID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindDVoiceState retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDVoiceState(exec boil.Executor, userID int64, guildID int64, selectCols ...string) (*DVoiceState, error) {
	dVoiceStateObj := &DVoiceState{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"d_voice_states\" where \"user_id\"=$1 AND \"guild_id\"=$2", sel,
	)

	q := queries.Raw(exec, query, userID, guildID)

	err := q.Bind(dVoiceStateObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from d_voice_states")
	}

	return dVoiceStateObj, nil
}

// FindDVoiceStateP retrieves a single record by ID with an executor, and panics on error.
func FindDVoiceStateP(exec boil.Executor, userID int64, guildID int64, selectCols ...string) *DVoiceState {
	retobj, err := FindDVoiceState(exec, userID, guildID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *DVoiceState) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *DVoiceState) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *DVoiceState) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *DVoiceState) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no d_voice_states provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(dVoiceStateColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	dVoiceStateInsertCacheMut.RLock()
	cache, cached := dVoiceStateInsertCache[key]
	dVoiceStateInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			dVoiceStateColumns,
			dVoiceStateColumnsWithDefault,
			dVoiceStateColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(dVoiceStateType, dVoiceStateMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(dVoiceStateType, dVoiceStateMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"d_voice_states\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "models: unable to insert into d_voice_states")
	}

	if !cached {
		dVoiceStateInsertCacheMut.Lock()
		dVoiceStateInsertCache[key] = cache
		dVoiceStateInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single DVoiceState record. See Update for
// whitelist behavior description.
func (o *DVoiceState) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single DVoiceState record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *DVoiceState) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the DVoiceState, and panics on error.
// See Update for whitelist behavior description.
func (o *DVoiceState) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the DVoiceState.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *DVoiceState) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	key := makeCacheKey(whitelist, nil)
	dVoiceStateUpdateCacheMut.RLock()
	cache, cached := dVoiceStateUpdateCache[key]
	dVoiceStateUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(dVoiceStateColumns, dVoiceStatePrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update d_voice_states, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"d_voice_states\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, dVoiceStatePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(dVoiceStateType, dVoiceStateMapping, append(wl, dVoiceStatePrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update d_voice_states row")
	}

	if !cached {
		dVoiceStateUpdateCacheMut.Lock()
		dVoiceStateUpdateCache[key] = cache
		dVoiceStateUpdateCacheMut.Unlock()
	}

	return nil
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q dVoiceStateQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q dVoiceStateQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for d_voice_states")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o DVoiceStateSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o DVoiceStateSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o DVoiceStateSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o DVoiceStateSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dVoiceStatePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"d_voice_states\" SET %s WHERE (\"user_id\",\"guild_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(dVoiceStatePrimaryKeyColumns), len(colNames)+1, len(dVoiceStatePrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in dVoiceState slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *DVoiceState) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *DVoiceState) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *DVoiceState) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *DVoiceState) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no d_voice_states provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(dVoiceStateColumnsWithDefault, o)

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

	dVoiceStateUpsertCacheMut.RLock()
	cache, cached := dVoiceStateUpsertCache[key]
	dVoiceStateUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			dVoiceStateColumns,
			dVoiceStateColumnsWithDefault,
			dVoiceStateColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			dVoiceStateColumns,
			dVoiceStatePrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert d_voice_states, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(dVoiceStatePrimaryKeyColumns))
			copy(conflict, dVoiceStatePrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"d_voice_states\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(dVoiceStateType, dVoiceStateMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(dVoiceStateType, dVoiceStateMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert d_voice_states")
	}

	if !cached {
		dVoiceStateUpsertCacheMut.Lock()
		dVoiceStateUpsertCache[key] = cache
		dVoiceStateUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteP deletes a single DVoiceState record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *DVoiceState) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single DVoiceState record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *DVoiceState) DeleteG() error {
	if o == nil {
		return errors.New("models: no DVoiceState provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single DVoiceState record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *DVoiceState) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single DVoiceState record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *DVoiceState) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no DVoiceState provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), dVoiceStatePrimaryKeyMapping)
	sql := "DELETE FROM \"d_voice_states\" WHERE \"user_id\"=$1 AND \"guild_id\"=$2"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from d_voice_states")
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q dVoiceStateQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q dVoiceStateQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no dVoiceStateQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from d_voice_states")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o DVoiceStateSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o DVoiceStateSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no DVoiceState slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o DVoiceStateSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o DVoiceStateSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no DVoiceState slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dVoiceStatePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"d_voice_states\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, dVoiceStatePrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(dVoiceStatePrimaryKeyColumns), 1, len(dVoiceStatePrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from dVoiceState slice")
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *DVoiceState) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *DVoiceState) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *DVoiceState) ReloadG() error {
	if o == nil {
		return errors.New("models: no DVoiceState provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *DVoiceState) Reload(exec boil.Executor) error {
	ret, err := FindDVoiceState(exec, o.UserID, o.GuildID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *DVoiceStateSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *DVoiceStateSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DVoiceStateSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty DVoiceStateSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DVoiceStateSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	dVoiceStates := DVoiceStateSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dVoiceStatePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"d_voice_states\".* FROM \"d_voice_states\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, dVoiceStatePrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(dVoiceStatePrimaryKeyColumns), 1, len(dVoiceStatePrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&dVoiceStates)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in DVoiceStateSlice")
	}

	*o = dVoiceStates

	return nil
}

// DVoiceStateExists checks if the DVoiceState row exists.
func DVoiceStateExists(exec boil.Executor, userID int64, guildID int64) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"d_voice_states\" where \"user_id\"=$1 AND \"guild_id\"=$2 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, userID, guildID)
	}

	row := exec.QueryRow(sql, userID, guildID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if d_voice_states exists")
	}

	return exists, nil
}

// DVoiceStateExistsG checks if the DVoiceState row exists.
func DVoiceStateExistsG(userID int64, guildID int64) (bool, error) {
	return DVoiceStateExists(boil.GetDB(), userID, guildID)
}

// DVoiceStateExistsGP checks if the DVoiceState row exists. Panics on error.
func DVoiceStateExistsGP(userID int64, guildID int64) bool {
	e, err := DVoiceStateExists(boil.GetDB(), userID, guildID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// DVoiceStateExistsP checks if the DVoiceState row exists. Panics on error.
func DVoiceStateExistsP(exec boil.Executor, userID int64, guildID int64) bool {
	e, err := DVoiceStateExists(exec, userID, guildID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
