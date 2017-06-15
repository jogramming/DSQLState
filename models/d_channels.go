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

// DChannel is an object representing the database table.
type DChannel struct {
	ID            int64      `boil:"id" json:"id" toml:"id" yaml:"id"`
	GuildID       null.Int64 `boil:"guild_id" json:"guild_id,omitempty" toml:"guild_id" yaml:"guild_id,omitempty"`
	RecipientID   null.Int64 `boil:"recipient_id" json:"recipient_id,omitempty" toml:"recipient_id" yaml:"recipient_id,omitempty"`
	CreatedAt     time.Time  `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	DeletedAt     null.Time  `boil:"deleted_at" json:"deleted_at,omitempty" toml:"deleted_at" yaml:"deleted_at,omitempty"`
	Synced        bool       `boil:"synced" json:"synced" toml:"synced" yaml:"synced"`
	Name          string     `boil:"name" json:"name" toml:"name" yaml:"name"`
	Topic         string     `boil:"topic" json:"topic" toml:"topic" yaml:"topic"`
	Type          string     `boil:"type" json:"type" toml:"type" yaml:"type"`
	LastMessageID int64      `boil:"last_message_id" json:"last_message_id" toml:"last_message_id" yaml:"last_message_id"`
	Position      int        `boil:"position" json:"position" toml:"position" yaml:"position"`
	Bitrate       int        `boil:"bitrate" json:"bitrate" toml:"bitrate" yaml:"bitrate"`

	R *dChannelR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L dChannelL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// dChannelR is where relationships are stored.
type dChannelR struct {
	ChannelDChannelOverwrites DChannelOverwriteSlice
	ChannelDVoiceStates       DVoiceStateSlice
}

// dChannelL is where Load methods for each relationship are stored.
type dChannelL struct{}

var (
	dChannelColumns               = []string{"id", "guild_id", "recipient_id", "created_at", "deleted_at", "synced", "name", "topic", "type", "last_message_id", "position", "bitrate"}
	dChannelColumnsWithoutDefault = []string{"id", "guild_id", "recipient_id", "created_at", "deleted_at", "synced", "name", "topic", "type", "last_message_id", "position", "bitrate"}
	dChannelColumnsWithDefault    = []string{}
	dChannelPrimaryKeyColumns     = []string{"id"}
)

type (
	// DChannelSlice is an alias for a slice of pointers to DChannel.
	// This should generally be used opposed to []DChannel.
	DChannelSlice []*DChannel

	dChannelQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	dChannelType                 = reflect.TypeOf(&DChannel{})
	dChannelMapping              = queries.MakeStructMapping(dChannelType)
	dChannelPrimaryKeyMapping, _ = queries.BindMapping(dChannelType, dChannelMapping, dChannelPrimaryKeyColumns)
	dChannelInsertCacheMut       sync.RWMutex
	dChannelInsertCache          = make(map[string]insertCache)
	dChannelUpdateCacheMut       sync.RWMutex
	dChannelUpdateCache          = make(map[string]updateCache)
	dChannelUpsertCacheMut       sync.RWMutex
	dChannelUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)

// OneP returns a single dChannel record from the query, and panics on error.
func (q dChannelQuery) OneP() *DChannel {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single dChannel record from the query.
func (q dChannelQuery) One() (*DChannel, error) {
	o := &DChannel{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for d_channels")
	}

	return o, nil
}

// AllP returns all DChannel records from the query, and panics on error.
func (q dChannelQuery) AllP() DChannelSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all DChannel records from the query.
func (q dChannelQuery) All() (DChannelSlice, error) {
	var o DChannelSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to DChannel slice")
	}

	return o, nil
}

// CountP returns the count of all DChannel records in the query, and panics on error.
func (q dChannelQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all DChannel records in the query.
func (q dChannelQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count d_channels rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q dChannelQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q dChannelQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if d_channels exists")
	}

	return count > 0, nil
}

// ChannelDChannelOverwritesG retrieves all the d_channel_overwrite's d channel overwrites via channel_id column.
func (o *DChannel) ChannelDChannelOverwritesG(mods ...qm.QueryMod) dChannelOverwriteQuery {
	return o.ChannelDChannelOverwrites(boil.GetDB(), mods...)
}

// ChannelDChannelOverwrites retrieves all the d_channel_overwrite's d channel overwrites with an executor via channel_id column.
func (o *DChannel) ChannelDChannelOverwrites(exec boil.Executor, mods ...qm.QueryMod) dChannelOverwriteQuery {
	queryMods := []qm.QueryMod{
		qm.Select("\"a\".*"),
	}

	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"a\".\"channel_id\"=?", o.ID),
	)

	query := DChannelOverwrites(exec, queryMods...)
	queries.SetFrom(query.Query, "\"d_channel_overwrites\" as \"a\"")
	return query
}

// ChannelDVoiceStatesG retrieves all the d_voice_state's d voice states via channel_id column.
func (o *DChannel) ChannelDVoiceStatesG(mods ...qm.QueryMod) dVoiceStateQuery {
	return o.ChannelDVoiceStates(boil.GetDB(), mods...)
}

// ChannelDVoiceStates retrieves all the d_voice_state's d voice states with an executor via channel_id column.
func (o *DChannel) ChannelDVoiceStates(exec boil.Executor, mods ...qm.QueryMod) dVoiceStateQuery {
	queryMods := []qm.QueryMod{
		qm.Select("\"a\".*"),
	}

	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"a\".\"channel_id\"=?", o.ID),
	)

	query := DVoiceStates(exec, queryMods...)
	queries.SetFrom(query.Query, "\"d_voice_states\" as \"a\"")
	return query
}

// LoadChannelDChannelOverwrites allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (dChannelL) LoadChannelDChannelOverwrites(e boil.Executor, singular bool, maybeDChannel interface{}) error {
	var slice []*DChannel
	var object *DChannel

	count := 1
	if singular {
		object = maybeDChannel.(*DChannel)
	} else {
		slice = *maybeDChannel.(*DChannelSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		if object.R == nil {
			object.R = &dChannelR{}
		}
		args[0] = object.ID
	} else {
		for i, obj := range slice {
			if obj.R == nil {
				obj.R = &dChannelR{}
			}
			args[i] = obj.ID
		}
	}

	query := fmt.Sprintf(
		"select * from \"d_channel_overwrites\" where \"channel_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)
	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load d_channel_overwrites")
	}
	defer results.Close()

	var resultSlice []*DChannelOverwrite
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice d_channel_overwrites")
	}

	if singular {
		object.R.ChannelDChannelOverwrites = resultSlice
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.ChannelID {
				local.R.ChannelDChannelOverwrites = append(local.R.ChannelDChannelOverwrites, foreign)
				break
			}
		}
	}

	return nil
}

// LoadChannelDVoiceStates allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (dChannelL) LoadChannelDVoiceStates(e boil.Executor, singular bool, maybeDChannel interface{}) error {
	var slice []*DChannel
	var object *DChannel

	count := 1
	if singular {
		object = maybeDChannel.(*DChannel)
	} else {
		slice = *maybeDChannel.(*DChannelSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		if object.R == nil {
			object.R = &dChannelR{}
		}
		args[0] = object.ID
	} else {
		for i, obj := range slice {
			if obj.R == nil {
				obj.R = &dChannelR{}
			}
			args[i] = obj.ID
		}
	}

	query := fmt.Sprintf(
		"select * from \"d_voice_states\" where \"channel_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)
	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load d_voice_states")
	}
	defer results.Close()

	var resultSlice []*DVoiceState
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice d_voice_states")
	}

	if singular {
		object.R.ChannelDVoiceStates = resultSlice
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.ChannelID {
				local.R.ChannelDVoiceStates = append(local.R.ChannelDVoiceStates, foreign)
				break
			}
		}
	}

	return nil
}

// AddChannelDChannelOverwritesG adds the given related objects to the existing relationships
// of the d_channel, optionally inserting them as new records.
// Appends related to o.R.ChannelDChannelOverwrites.
// Sets related.R.Channel appropriately.
// Uses the global database handle.
func (o *DChannel) AddChannelDChannelOverwritesG(insert bool, related ...*DChannelOverwrite) error {
	return o.AddChannelDChannelOverwrites(boil.GetDB(), insert, related...)
}

// AddChannelDChannelOverwritesP adds the given related objects to the existing relationships
// of the d_channel, optionally inserting them as new records.
// Appends related to o.R.ChannelDChannelOverwrites.
// Sets related.R.Channel appropriately.
// Panics on error.
func (o *DChannel) AddChannelDChannelOverwritesP(exec boil.Executor, insert bool, related ...*DChannelOverwrite) {
	if err := o.AddChannelDChannelOverwrites(exec, insert, related...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// AddChannelDChannelOverwritesGP adds the given related objects to the existing relationships
// of the d_channel, optionally inserting them as new records.
// Appends related to o.R.ChannelDChannelOverwrites.
// Sets related.R.Channel appropriately.
// Uses the global database handle and panics on error.
func (o *DChannel) AddChannelDChannelOverwritesGP(insert bool, related ...*DChannelOverwrite) {
	if err := o.AddChannelDChannelOverwrites(boil.GetDB(), insert, related...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// AddChannelDChannelOverwrites adds the given related objects to the existing relationships
// of the d_channel, optionally inserting them as new records.
// Appends related to o.R.ChannelDChannelOverwrites.
// Sets related.R.Channel appropriately.
func (o *DChannel) AddChannelDChannelOverwrites(exec boil.Executor, insert bool, related ...*DChannelOverwrite) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.ChannelID = o.ID
			if err = rel.Insert(exec); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"d_channel_overwrites\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"channel_id"}),
				strmangle.WhereClause("\"", "\"", 2, dChannelOverwritePrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID, rel.ChannelID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}

			if _, err = exec.Exec(updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.ChannelID = o.ID
		}
	}

	if o.R == nil {
		o.R = &dChannelR{
			ChannelDChannelOverwrites: related,
		}
	} else {
		o.R.ChannelDChannelOverwrites = append(o.R.ChannelDChannelOverwrites, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &dChannelOverwriteR{
				Channel: o,
			}
		} else {
			rel.R.Channel = o
		}
	}
	return nil
}

// AddChannelDVoiceStatesG adds the given related objects to the existing relationships
// of the d_channel, optionally inserting them as new records.
// Appends related to o.R.ChannelDVoiceStates.
// Sets related.R.Channel appropriately.
// Uses the global database handle.
func (o *DChannel) AddChannelDVoiceStatesG(insert bool, related ...*DVoiceState) error {
	return o.AddChannelDVoiceStates(boil.GetDB(), insert, related...)
}

// AddChannelDVoiceStatesP adds the given related objects to the existing relationships
// of the d_channel, optionally inserting them as new records.
// Appends related to o.R.ChannelDVoiceStates.
// Sets related.R.Channel appropriately.
// Panics on error.
func (o *DChannel) AddChannelDVoiceStatesP(exec boil.Executor, insert bool, related ...*DVoiceState) {
	if err := o.AddChannelDVoiceStates(exec, insert, related...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// AddChannelDVoiceStatesGP adds the given related objects to the existing relationships
// of the d_channel, optionally inserting them as new records.
// Appends related to o.R.ChannelDVoiceStates.
// Sets related.R.Channel appropriately.
// Uses the global database handle and panics on error.
func (o *DChannel) AddChannelDVoiceStatesGP(insert bool, related ...*DVoiceState) {
	if err := o.AddChannelDVoiceStates(boil.GetDB(), insert, related...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// AddChannelDVoiceStates adds the given related objects to the existing relationships
// of the d_channel, optionally inserting them as new records.
// Appends related to o.R.ChannelDVoiceStates.
// Sets related.R.Channel appropriately.
func (o *DChannel) AddChannelDVoiceStates(exec boil.Executor, insert bool, related ...*DVoiceState) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.ChannelID = o.ID
			if err = rel.Insert(exec); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"d_voice_states\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"channel_id"}),
				strmangle.WhereClause("\"", "\"", 2, dVoiceStatePrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.UserID, rel.GuildID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}

			if _, err = exec.Exec(updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.ChannelID = o.ID
		}
	}

	if o.R == nil {
		o.R = &dChannelR{
			ChannelDVoiceStates: related,
		}
	} else {
		o.R.ChannelDVoiceStates = append(o.R.ChannelDVoiceStates, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &dVoiceStateR{
				Channel: o,
			}
		} else {
			rel.R.Channel = o
		}
	}
	return nil
}

// DChannelsG retrieves all records.
func DChannelsG(mods ...qm.QueryMod) dChannelQuery {
	return DChannels(boil.GetDB(), mods...)
}

// DChannels retrieves all the records using an executor.
func DChannels(exec boil.Executor, mods ...qm.QueryMod) dChannelQuery {
	mods = append(mods, qm.From("\"d_channels\""))
	return dChannelQuery{NewQuery(exec, mods...)}
}

// FindDChannelG retrieves a single record by ID.
func FindDChannelG(id int64, selectCols ...string) (*DChannel, error) {
	return FindDChannel(boil.GetDB(), id, selectCols...)
}

// FindDChannelGP retrieves a single record by ID, and panics on error.
func FindDChannelGP(id int64, selectCols ...string) *DChannel {
	retobj, err := FindDChannel(boil.GetDB(), id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindDChannel retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDChannel(exec boil.Executor, id int64, selectCols ...string) (*DChannel, error) {
	dChannelObj := &DChannel{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"d_channels\" where \"id\"=$1", sel,
	)

	q := queries.Raw(exec, query, id)

	err := q.Bind(dChannelObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from d_channels")
	}

	return dChannelObj, nil
}

// FindDChannelP retrieves a single record by ID with an executor, and panics on error.
func FindDChannelP(exec boil.Executor, id int64, selectCols ...string) *DChannel {
	retobj, err := FindDChannel(exec, id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *DChannel) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *DChannel) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *DChannel) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *DChannel) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no d_channels provided for insertion")
	}

	var err error
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(dChannelColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	dChannelInsertCacheMut.RLock()
	cache, cached := dChannelInsertCache[key]
	dChannelInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			dChannelColumns,
			dChannelColumnsWithDefault,
			dChannelColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(dChannelType, dChannelMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(dChannelType, dChannelMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"d_channels\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "models: unable to insert into d_channels")
	}

	if !cached {
		dChannelInsertCacheMut.Lock()
		dChannelInsertCache[key] = cache
		dChannelInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single DChannel record. See Update for
// whitelist behavior description.
func (o *DChannel) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single DChannel record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *DChannel) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the DChannel, and panics on error.
// See Update for whitelist behavior description.
func (o *DChannel) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the DChannel.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *DChannel) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	key := makeCacheKey(whitelist, nil)
	dChannelUpdateCacheMut.RLock()
	cache, cached := dChannelUpdateCache[key]
	dChannelUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(dChannelColumns, dChannelPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update d_channels, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"d_channels\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, dChannelPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(dChannelType, dChannelMapping, append(wl, dChannelPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update d_channels row")
	}

	if !cached {
		dChannelUpdateCacheMut.Lock()
		dChannelUpdateCache[key] = cache
		dChannelUpdateCacheMut.Unlock()
	}

	return nil
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q dChannelQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q dChannelQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for d_channels")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o DChannelSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o DChannelSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o DChannelSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o DChannelSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dChannelPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"d_channels\" SET %s WHERE (\"id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(dChannelPrimaryKeyColumns), len(colNames)+1, len(dChannelPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in dChannel slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *DChannel) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *DChannel) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *DChannel) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *DChannel) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no d_channels provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(dChannelColumnsWithDefault, o)

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

	dChannelUpsertCacheMut.RLock()
	cache, cached := dChannelUpsertCache[key]
	dChannelUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			dChannelColumns,
			dChannelColumnsWithDefault,
			dChannelColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			dChannelColumns,
			dChannelPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert d_channels, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(dChannelPrimaryKeyColumns))
			copy(conflict, dChannelPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"d_channels\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(dChannelType, dChannelMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(dChannelType, dChannelMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert d_channels")
	}

	if !cached {
		dChannelUpsertCacheMut.Lock()
		dChannelUpsertCache[key] = cache
		dChannelUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteP deletes a single DChannel record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *DChannel) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single DChannel record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *DChannel) DeleteG() error {
	if o == nil {
		return errors.New("models: no DChannel provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single DChannel record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *DChannel) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single DChannel record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *DChannel) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no DChannel provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), dChannelPrimaryKeyMapping)
	sql := "DELETE FROM \"d_channels\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from d_channels")
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q dChannelQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q dChannelQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no dChannelQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from d_channels")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o DChannelSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o DChannelSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no DChannel slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o DChannelSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o DChannelSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no DChannel slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dChannelPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"d_channels\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, dChannelPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(dChannelPrimaryKeyColumns), 1, len(dChannelPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from dChannel slice")
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *DChannel) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *DChannel) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *DChannel) ReloadG() error {
	if o == nil {
		return errors.New("models: no DChannel provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *DChannel) Reload(exec boil.Executor) error {
	ret, err := FindDChannel(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *DChannelSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *DChannelSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DChannelSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty DChannelSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DChannelSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	dChannels := DChannelSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dChannelPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"d_channels\".* FROM \"d_channels\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, dChannelPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(dChannelPrimaryKeyColumns), 1, len(dChannelPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&dChannels)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in DChannelSlice")
	}

	*o = dChannels

	return nil
}

// DChannelExists checks if the DChannel row exists.
func DChannelExists(exec boil.Executor, id int64) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"d_channels\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, id)
	}

	row := exec.QueryRow(sql, id)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if d_channels exists")
	}

	return exists, nil
}

// DChannelExistsG checks if the DChannel row exists.
func DChannelExistsG(id int64) (bool, error) {
	return DChannelExists(boil.GetDB(), id)
}

// DChannelExistsGP checks if the DChannel row exists. Panics on error.
func DChannelExistsGP(id int64) bool {
	e, err := DChannelExists(boil.GetDB(), id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// DChannelExistsP checks if the DChannel row exists. Panics on error.
func DChannelExistsP(exec boil.Executor, id int64) bool {
	e, err := DChannelExists(exec, id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
