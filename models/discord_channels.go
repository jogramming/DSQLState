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

// DiscordChannel is an object representing the database table.
type DiscordChannel struct {
	ID            int64      `boil:"id" json:"id" toml:"id" yaml:"id"`
	GuildID       null.Int64 `boil:"guild_id" json:"guild_id,omitempty" toml:"guild_id" yaml:"guild_id,omitempty"`
	RecipientID   null.Int64 `boil:"recipient_id" json:"recipient_id,omitempty" toml:"recipient_id" yaml:"recipient_id,omitempty"`
	CreatedAt     time.Time  `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	DeletedAt     null.Time  `boil:"deleted_at" json:"deleted_at,omitempty" toml:"deleted_at" yaml:"deleted_at,omitempty"`
	Name          string     `boil:"name" json:"name" toml:"name" yaml:"name"`
	Topic         string     `boil:"topic" json:"topic" toml:"topic" yaml:"topic"`
	Type          string     `boil:"type" json:"type" toml:"type" yaml:"type"`
	LastMessageID int64      `boil:"last_message_id" json:"last_message_id" toml:"last_message_id" yaml:"last_message_id"`
	Position      int        `boil:"position" json:"position" toml:"position" yaml:"position"`
	Bitrate       int        `boil:"bitrate" json:"bitrate" toml:"bitrate" yaml:"bitrate"`

	R *discordChannelR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L discordChannelL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// discordChannelR is where relationships are stored.
type discordChannelR struct {
	ChannelDiscordChannelOverwrites DiscordChannelOverwriteSlice
	ChannelDiscordVoiceStates       DiscordVoiceStateSlice
}

// discordChannelL is where Load methods for each relationship are stored.
type discordChannelL struct{}

var (
	discordChannelColumns               = []string{"id", "guild_id", "recipient_id", "created_at", "deleted_at", "name", "topic", "type", "last_message_id", "position", "bitrate"}
	discordChannelColumnsWithoutDefault = []string{"id", "guild_id", "recipient_id", "created_at", "deleted_at", "name", "topic", "type", "last_message_id", "position", "bitrate"}
	discordChannelColumnsWithDefault    = []string{}
	discordChannelPrimaryKeyColumns     = []string{"id"}
)

type (
	// DiscordChannelSlice is an alias for a slice of pointers to DiscordChannel.
	// This should generally be used opposed to []DiscordChannel.
	DiscordChannelSlice []*DiscordChannel

	discordChannelQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	discordChannelType                 = reflect.TypeOf(&DiscordChannel{})
	discordChannelMapping              = queries.MakeStructMapping(discordChannelType)
	discordChannelPrimaryKeyMapping, _ = queries.BindMapping(discordChannelType, discordChannelMapping, discordChannelPrimaryKeyColumns)
	discordChannelInsertCacheMut       sync.RWMutex
	discordChannelInsertCache          = make(map[string]insertCache)
	discordChannelUpdateCacheMut       sync.RWMutex
	discordChannelUpdateCache          = make(map[string]updateCache)
	discordChannelUpsertCacheMut       sync.RWMutex
	discordChannelUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)

// OneP returns a single discordChannel record from the query, and panics on error.
func (q discordChannelQuery) OneP() *DiscordChannel {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single discordChannel record from the query.
func (q discordChannelQuery) One() (*DiscordChannel, error) {
	o := &DiscordChannel{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for discord_channels")
	}

	return o, nil
}

// AllP returns all DiscordChannel records from the query, and panics on error.
func (q discordChannelQuery) AllP() DiscordChannelSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all DiscordChannel records from the query.
func (q discordChannelQuery) All() (DiscordChannelSlice, error) {
	var o DiscordChannelSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to DiscordChannel slice")
	}

	return o, nil
}

// CountP returns the count of all DiscordChannel records in the query, and panics on error.
func (q discordChannelQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all DiscordChannel records in the query.
func (q discordChannelQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count discord_channels rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q discordChannelQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q discordChannelQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if discord_channels exists")
	}

	return count > 0, nil
}

// ChannelDiscordChannelOverwritesG retrieves all the discord_channel_overwrite's discord channel overwrites via channel_id column.
func (o *DiscordChannel) ChannelDiscordChannelOverwritesG(mods ...qm.QueryMod) discordChannelOverwriteQuery {
	return o.ChannelDiscordChannelOverwrites(boil.GetDB(), mods...)
}

// ChannelDiscordChannelOverwrites retrieves all the discord_channel_overwrite's discord channel overwrites with an executor via channel_id column.
func (o *DiscordChannel) ChannelDiscordChannelOverwrites(exec boil.Executor, mods ...qm.QueryMod) discordChannelOverwriteQuery {
	queryMods := []qm.QueryMod{
		qm.Select("\"a\".*"),
	}

	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"a\".\"channel_id\"=?", o.ID),
	)

	query := DiscordChannelOverwrites(exec, queryMods...)
	queries.SetFrom(query.Query, "\"discord_channel_overwrites\" as \"a\"")
	return query
}

// ChannelDiscordVoiceStatesG retrieves all the discord_voice_state's discord voice states via channel_id column.
func (o *DiscordChannel) ChannelDiscordVoiceStatesG(mods ...qm.QueryMod) discordVoiceStateQuery {
	return o.ChannelDiscordVoiceStates(boil.GetDB(), mods...)
}

// ChannelDiscordVoiceStates retrieves all the discord_voice_state's discord voice states with an executor via channel_id column.
func (o *DiscordChannel) ChannelDiscordVoiceStates(exec boil.Executor, mods ...qm.QueryMod) discordVoiceStateQuery {
	queryMods := []qm.QueryMod{
		qm.Select("\"a\".*"),
	}

	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"a\".\"channel_id\"=?", o.ID),
	)

	query := DiscordVoiceStates(exec, queryMods...)
	queries.SetFrom(query.Query, "\"discord_voice_states\" as \"a\"")
	return query
}

// LoadChannelDiscordChannelOverwrites allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (discordChannelL) LoadChannelDiscordChannelOverwrites(e boil.Executor, singular bool, maybeDiscordChannel interface{}) error {
	var slice []*DiscordChannel
	var object *DiscordChannel

	count := 1
	if singular {
		object = maybeDiscordChannel.(*DiscordChannel)
	} else {
		slice = *maybeDiscordChannel.(*DiscordChannelSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		if object.R == nil {
			object.R = &discordChannelR{}
		}
		args[0] = object.ID
	} else {
		for i, obj := range slice {
			if obj.R == nil {
				obj.R = &discordChannelR{}
			}
			args[i] = obj.ID
		}
	}

	query := fmt.Sprintf(
		"select * from \"discord_channel_overwrites\" where \"channel_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)
	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load discord_channel_overwrites")
	}
	defer results.Close()

	var resultSlice []*DiscordChannelOverwrite
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice discord_channel_overwrites")
	}

	if singular {
		object.R.ChannelDiscordChannelOverwrites = resultSlice
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.ChannelID {
				local.R.ChannelDiscordChannelOverwrites = append(local.R.ChannelDiscordChannelOverwrites, foreign)
				break
			}
		}
	}

	return nil
}

// LoadChannelDiscordVoiceStates allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (discordChannelL) LoadChannelDiscordVoiceStates(e boil.Executor, singular bool, maybeDiscordChannel interface{}) error {
	var slice []*DiscordChannel
	var object *DiscordChannel

	count := 1
	if singular {
		object = maybeDiscordChannel.(*DiscordChannel)
	} else {
		slice = *maybeDiscordChannel.(*DiscordChannelSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		if object.R == nil {
			object.R = &discordChannelR{}
		}
		args[0] = object.ID
	} else {
		for i, obj := range slice {
			if obj.R == nil {
				obj.R = &discordChannelR{}
			}
			args[i] = obj.ID
		}
	}

	query := fmt.Sprintf(
		"select * from \"discord_voice_states\" where \"channel_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)
	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load discord_voice_states")
	}
	defer results.Close()

	var resultSlice []*DiscordVoiceState
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice discord_voice_states")
	}

	if singular {
		object.R.ChannelDiscordVoiceStates = resultSlice
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.ChannelID {
				local.R.ChannelDiscordVoiceStates = append(local.R.ChannelDiscordVoiceStates, foreign)
				break
			}
		}
	}

	return nil
}

// AddChannelDiscordChannelOverwritesG adds the given related objects to the existing relationships
// of the discord_channel, optionally inserting them as new records.
// Appends related to o.R.ChannelDiscordChannelOverwrites.
// Sets related.R.Channel appropriately.
// Uses the global database handle.
func (o *DiscordChannel) AddChannelDiscordChannelOverwritesG(insert bool, related ...*DiscordChannelOverwrite) error {
	return o.AddChannelDiscordChannelOverwrites(boil.GetDB(), insert, related...)
}

// AddChannelDiscordChannelOverwritesP adds the given related objects to the existing relationships
// of the discord_channel, optionally inserting them as new records.
// Appends related to o.R.ChannelDiscordChannelOverwrites.
// Sets related.R.Channel appropriately.
// Panics on error.
func (o *DiscordChannel) AddChannelDiscordChannelOverwritesP(exec boil.Executor, insert bool, related ...*DiscordChannelOverwrite) {
	if err := o.AddChannelDiscordChannelOverwrites(exec, insert, related...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// AddChannelDiscordChannelOverwritesGP adds the given related objects to the existing relationships
// of the discord_channel, optionally inserting them as new records.
// Appends related to o.R.ChannelDiscordChannelOverwrites.
// Sets related.R.Channel appropriately.
// Uses the global database handle and panics on error.
func (o *DiscordChannel) AddChannelDiscordChannelOverwritesGP(insert bool, related ...*DiscordChannelOverwrite) {
	if err := o.AddChannelDiscordChannelOverwrites(boil.GetDB(), insert, related...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// AddChannelDiscordChannelOverwrites adds the given related objects to the existing relationships
// of the discord_channel, optionally inserting them as new records.
// Appends related to o.R.ChannelDiscordChannelOverwrites.
// Sets related.R.Channel appropriately.
func (o *DiscordChannel) AddChannelDiscordChannelOverwrites(exec boil.Executor, insert bool, related ...*DiscordChannelOverwrite) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.ChannelID = o.ID
			if err = rel.Insert(exec); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"discord_channel_overwrites\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"channel_id"}),
				strmangle.WhereClause("\"", "\"", 2, discordChannelOverwritePrimaryKeyColumns),
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
		o.R = &discordChannelR{
			ChannelDiscordChannelOverwrites: related,
		}
	} else {
		o.R.ChannelDiscordChannelOverwrites = append(o.R.ChannelDiscordChannelOverwrites, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &discordChannelOverwriteR{
				Channel: o,
			}
		} else {
			rel.R.Channel = o
		}
	}
	return nil
}

// AddChannelDiscordVoiceStatesG adds the given related objects to the existing relationships
// of the discord_channel, optionally inserting them as new records.
// Appends related to o.R.ChannelDiscordVoiceStates.
// Sets related.R.Channel appropriately.
// Uses the global database handle.
func (o *DiscordChannel) AddChannelDiscordVoiceStatesG(insert bool, related ...*DiscordVoiceState) error {
	return o.AddChannelDiscordVoiceStates(boil.GetDB(), insert, related...)
}

// AddChannelDiscordVoiceStatesP adds the given related objects to the existing relationships
// of the discord_channel, optionally inserting them as new records.
// Appends related to o.R.ChannelDiscordVoiceStates.
// Sets related.R.Channel appropriately.
// Panics on error.
func (o *DiscordChannel) AddChannelDiscordVoiceStatesP(exec boil.Executor, insert bool, related ...*DiscordVoiceState) {
	if err := o.AddChannelDiscordVoiceStates(exec, insert, related...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// AddChannelDiscordVoiceStatesGP adds the given related objects to the existing relationships
// of the discord_channel, optionally inserting them as new records.
// Appends related to o.R.ChannelDiscordVoiceStates.
// Sets related.R.Channel appropriately.
// Uses the global database handle and panics on error.
func (o *DiscordChannel) AddChannelDiscordVoiceStatesGP(insert bool, related ...*DiscordVoiceState) {
	if err := o.AddChannelDiscordVoiceStates(boil.GetDB(), insert, related...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// AddChannelDiscordVoiceStates adds the given related objects to the existing relationships
// of the discord_channel, optionally inserting them as new records.
// Appends related to o.R.ChannelDiscordVoiceStates.
// Sets related.R.Channel appropriately.
func (o *DiscordChannel) AddChannelDiscordVoiceStates(exec boil.Executor, insert bool, related ...*DiscordVoiceState) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.ChannelID = o.ID
			if err = rel.Insert(exec); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"discord_voice_states\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"channel_id"}),
				strmangle.WhereClause("\"", "\"", 2, discordVoiceStatePrimaryKeyColumns),
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
		o.R = &discordChannelR{
			ChannelDiscordVoiceStates: related,
		}
	} else {
		o.R.ChannelDiscordVoiceStates = append(o.R.ChannelDiscordVoiceStates, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &discordVoiceStateR{
				Channel: o,
			}
		} else {
			rel.R.Channel = o
		}
	}
	return nil
}

// DiscordChannelsG retrieves all records.
func DiscordChannelsG(mods ...qm.QueryMod) discordChannelQuery {
	return DiscordChannels(boil.GetDB(), mods...)
}

// DiscordChannels retrieves all the records using an executor.
func DiscordChannels(exec boil.Executor, mods ...qm.QueryMod) discordChannelQuery {
	mods = append(mods, qm.From("\"discord_channels\""))
	return discordChannelQuery{NewQuery(exec, mods...)}
}

// FindDiscordChannelG retrieves a single record by ID.
func FindDiscordChannelG(id int64, selectCols ...string) (*DiscordChannel, error) {
	return FindDiscordChannel(boil.GetDB(), id, selectCols...)
}

// FindDiscordChannelGP retrieves a single record by ID, and panics on error.
func FindDiscordChannelGP(id int64, selectCols ...string) *DiscordChannel {
	retobj, err := FindDiscordChannel(boil.GetDB(), id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindDiscordChannel retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDiscordChannel(exec boil.Executor, id int64, selectCols ...string) (*DiscordChannel, error) {
	discordChannelObj := &DiscordChannel{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"discord_channels\" where \"id\"=$1", sel,
	)

	q := queries.Raw(exec, query, id)

	err := q.Bind(discordChannelObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from discord_channels")
	}

	return discordChannelObj, nil
}

// FindDiscordChannelP retrieves a single record by ID with an executor, and panics on error.
func FindDiscordChannelP(exec boil.Executor, id int64, selectCols ...string) *DiscordChannel {
	retobj, err := FindDiscordChannel(exec, id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *DiscordChannel) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *DiscordChannel) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *DiscordChannel) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *DiscordChannel) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no discord_channels provided for insertion")
	}

	var err error
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(discordChannelColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	discordChannelInsertCacheMut.RLock()
	cache, cached := discordChannelInsertCache[key]
	discordChannelInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			discordChannelColumns,
			discordChannelColumnsWithDefault,
			discordChannelColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(discordChannelType, discordChannelMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(discordChannelType, discordChannelMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"discord_channels\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "models: unable to insert into discord_channels")
	}

	if !cached {
		discordChannelInsertCacheMut.Lock()
		discordChannelInsertCache[key] = cache
		discordChannelInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single DiscordChannel record. See Update for
// whitelist behavior description.
func (o *DiscordChannel) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single DiscordChannel record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *DiscordChannel) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the DiscordChannel, and panics on error.
// See Update for whitelist behavior description.
func (o *DiscordChannel) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the DiscordChannel.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *DiscordChannel) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	key := makeCacheKey(whitelist, nil)
	discordChannelUpdateCacheMut.RLock()
	cache, cached := discordChannelUpdateCache[key]
	discordChannelUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(discordChannelColumns, discordChannelPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update discord_channels, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"discord_channels\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, discordChannelPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(discordChannelType, discordChannelMapping, append(wl, discordChannelPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update discord_channels row")
	}

	if !cached {
		discordChannelUpdateCacheMut.Lock()
		discordChannelUpdateCache[key] = cache
		discordChannelUpdateCacheMut.Unlock()
	}

	return nil
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q discordChannelQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q discordChannelQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for discord_channels")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o DiscordChannelSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o DiscordChannelSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o DiscordChannelSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o DiscordChannelSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), discordChannelPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"discord_channels\" SET %s WHERE (\"id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(discordChannelPrimaryKeyColumns), len(colNames)+1, len(discordChannelPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in discordChannel slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *DiscordChannel) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *DiscordChannel) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *DiscordChannel) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *DiscordChannel) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no discord_channels provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(discordChannelColumnsWithDefault, o)

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

	discordChannelUpsertCacheMut.RLock()
	cache, cached := discordChannelUpsertCache[key]
	discordChannelUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			discordChannelColumns,
			discordChannelColumnsWithDefault,
			discordChannelColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			discordChannelColumns,
			discordChannelPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert discord_channels, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(discordChannelPrimaryKeyColumns))
			copy(conflict, discordChannelPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"discord_channels\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(discordChannelType, discordChannelMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(discordChannelType, discordChannelMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert discord_channels")
	}

	if !cached {
		discordChannelUpsertCacheMut.Lock()
		discordChannelUpsertCache[key] = cache
		discordChannelUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteP deletes a single DiscordChannel record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *DiscordChannel) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single DiscordChannel record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *DiscordChannel) DeleteG() error {
	if o == nil {
		return errors.New("models: no DiscordChannel provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single DiscordChannel record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *DiscordChannel) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single DiscordChannel record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *DiscordChannel) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no DiscordChannel provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), discordChannelPrimaryKeyMapping)
	sql := "DELETE FROM \"discord_channels\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from discord_channels")
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q discordChannelQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q discordChannelQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no discordChannelQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from discord_channels")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o DiscordChannelSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o DiscordChannelSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no DiscordChannel slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o DiscordChannelSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o DiscordChannelSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no DiscordChannel slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), discordChannelPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"discord_channels\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, discordChannelPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(discordChannelPrimaryKeyColumns), 1, len(discordChannelPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from discordChannel slice")
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *DiscordChannel) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *DiscordChannel) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *DiscordChannel) ReloadG() error {
	if o == nil {
		return errors.New("models: no DiscordChannel provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *DiscordChannel) Reload(exec boil.Executor) error {
	ret, err := FindDiscordChannel(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *DiscordChannelSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *DiscordChannelSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DiscordChannelSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty DiscordChannelSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DiscordChannelSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	discordChannels := DiscordChannelSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), discordChannelPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"discord_channels\".* FROM \"discord_channels\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, discordChannelPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(discordChannelPrimaryKeyColumns), 1, len(discordChannelPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&discordChannels)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in DiscordChannelSlice")
	}

	*o = discordChannels

	return nil
}

// DiscordChannelExists checks if the DiscordChannel row exists.
func DiscordChannelExists(exec boil.Executor, id int64) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"discord_channels\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, id)
	}

	row := exec.QueryRow(sql, id)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if discord_channels exists")
	}

	return exists, nil
}

// DiscordChannelExistsG checks if the DiscordChannel row exists.
func DiscordChannelExistsG(id int64) (bool, error) {
	return DiscordChannelExists(boil.GetDB(), id)
}

// DiscordChannelExistsGP checks if the DiscordChannel row exists. Panics on error.
func DiscordChannelExistsGP(id int64) bool {
	e, err := DiscordChannelExists(boil.GetDB(), id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// DiscordChannelExistsP checks if the DiscordChannel row exists. Panics on error.
func DiscordChannelExistsP(exec boil.Executor, id int64) bool {
	e, err := DiscordChannelExists(exec, id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
