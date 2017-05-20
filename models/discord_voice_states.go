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

// DiscordVoiceState is an object representing the database table.
type DiscordVoiceState struct {
	UserID    int64  `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	GuildID   int64  `boil:"guild_id" json:"guild_id" toml:"guild_id" yaml:"guild_id"`
	ChannelID int64  `boil:"channel_id" json:"channel_id" toml:"channel_id" yaml:"channel_id"`
	SessionID string `boil:"session_id" json:"session_id" toml:"session_id" yaml:"session_id"`
	Surpress  bool   `boil:"surpress" json:"surpress" toml:"surpress" yaml:"surpress"`
	SelfMute  bool   `boil:"self_mute" json:"self_mute" toml:"self_mute" yaml:"self_mute"`
	SelfDeaf  bool   `boil:"self_deaf" json:"self_deaf" toml:"self_deaf" yaml:"self_deaf"`
	Mute      bool   `boil:"mute" json:"mute" toml:"mute" yaml:"mute"`
	Deaf      bool   `boil:"deaf" json:"deaf" toml:"deaf" yaml:"deaf"`

	R *discordVoiceStateR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L discordVoiceStateL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// discordVoiceStateR is where relationships are stored.
type discordVoiceStateR struct {
	Channel *DiscordChannel
}

// discordVoiceStateL is where Load methods for each relationship are stored.
type discordVoiceStateL struct{}

var (
	discordVoiceStateColumns               = []string{"user_id", "guild_id", "channel_id", "session_id", "surpress", "self_mute", "self_deaf", "mute", "deaf"}
	discordVoiceStateColumnsWithoutDefault = []string{"user_id", "guild_id", "channel_id", "session_id", "surpress", "self_mute", "self_deaf", "mute", "deaf"}
	discordVoiceStateColumnsWithDefault    = []string{}
	discordVoiceStatePrimaryKeyColumns     = []string{"user_id", "guild_id"}
)

type (
	// DiscordVoiceStateSlice is an alias for a slice of pointers to DiscordVoiceState.
	// This should generally be used opposed to []DiscordVoiceState.
	DiscordVoiceStateSlice []*DiscordVoiceState

	discordVoiceStateQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	discordVoiceStateType                 = reflect.TypeOf(&DiscordVoiceState{})
	discordVoiceStateMapping              = queries.MakeStructMapping(discordVoiceStateType)
	discordVoiceStatePrimaryKeyMapping, _ = queries.BindMapping(discordVoiceStateType, discordVoiceStateMapping, discordVoiceStatePrimaryKeyColumns)
	discordVoiceStateInsertCacheMut       sync.RWMutex
	discordVoiceStateInsertCache          = make(map[string]insertCache)
	discordVoiceStateUpdateCacheMut       sync.RWMutex
	discordVoiceStateUpdateCache          = make(map[string]updateCache)
	discordVoiceStateUpsertCacheMut       sync.RWMutex
	discordVoiceStateUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)

// OneP returns a single discordVoiceState record from the query, and panics on error.
func (q discordVoiceStateQuery) OneP() *DiscordVoiceState {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single discordVoiceState record from the query.
func (q discordVoiceStateQuery) One() (*DiscordVoiceState, error) {
	o := &DiscordVoiceState{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for discord_voice_states")
	}

	return o, nil
}

// AllP returns all DiscordVoiceState records from the query, and panics on error.
func (q discordVoiceStateQuery) AllP() DiscordVoiceStateSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all DiscordVoiceState records from the query.
func (q discordVoiceStateQuery) All() (DiscordVoiceStateSlice, error) {
	var o DiscordVoiceStateSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to DiscordVoiceState slice")
	}

	return o, nil
}

// CountP returns the count of all DiscordVoiceState records in the query, and panics on error.
func (q discordVoiceStateQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all DiscordVoiceState records in the query.
func (q discordVoiceStateQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count discord_voice_states rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q discordVoiceStateQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q discordVoiceStateQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if discord_voice_states exists")
	}

	return count > 0, nil
}

// ChannelG pointed to by the foreign key.
func (o *DiscordVoiceState) ChannelG(mods ...qm.QueryMod) discordChannelQuery {
	return o.Channel(boil.GetDB(), mods...)
}

// Channel pointed to by the foreign key.
func (o *DiscordVoiceState) Channel(exec boil.Executor, mods ...qm.QueryMod) discordChannelQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.ChannelID),
	}

	queryMods = append(queryMods, mods...)

	query := DiscordChannels(exec, queryMods...)
	queries.SetFrom(query.Query, "\"discord_channels\"")

	return query
}

// LoadChannel allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (discordVoiceStateL) LoadChannel(e boil.Executor, singular bool, maybeDiscordVoiceState interface{}) error {
	var slice []*DiscordVoiceState
	var object *DiscordVoiceState

	count := 1
	if singular {
		object = maybeDiscordVoiceState.(*DiscordVoiceState)
	} else {
		slice = *maybeDiscordVoiceState.(*DiscordVoiceStateSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		if object.R == nil {
			object.R = &discordVoiceStateR{}
		}
		args[0] = object.ChannelID
	} else {
		for i, obj := range slice {
			if obj.R == nil {
				obj.R = &discordVoiceStateR{}
			}
			args[i] = obj.ChannelID
		}
	}

	query := fmt.Sprintf(
		"select * from \"discord_channels\" where \"id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load DiscordChannel")
	}
	defer results.Close()

	var resultSlice []*DiscordChannel
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice DiscordChannel")
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

// SetChannelG of the discord_voice_state to the related item.
// Sets o.R.Channel to related.
// Adds o to related.R.ChannelDiscordVoiceStates.
// Uses the global database handle.
func (o *DiscordVoiceState) SetChannelG(insert bool, related *DiscordChannel) error {
	return o.SetChannel(boil.GetDB(), insert, related)
}

// SetChannelP of the discord_voice_state to the related item.
// Sets o.R.Channel to related.
// Adds o to related.R.ChannelDiscordVoiceStates.
// Panics on error.
func (o *DiscordVoiceState) SetChannelP(exec boil.Executor, insert bool, related *DiscordChannel) {
	if err := o.SetChannel(exec, insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetChannelGP of the discord_voice_state to the related item.
// Sets o.R.Channel to related.
// Adds o to related.R.ChannelDiscordVoiceStates.
// Uses the global database handle and panics on error.
func (o *DiscordVoiceState) SetChannelGP(insert bool, related *DiscordChannel) {
	if err := o.SetChannel(boil.GetDB(), insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetChannel of the discord_voice_state to the related item.
// Sets o.R.Channel to related.
// Adds o to related.R.ChannelDiscordVoiceStates.
func (o *DiscordVoiceState) SetChannel(exec boil.Executor, insert bool, related *DiscordChannel) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"discord_voice_states\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"channel_id"}),
		strmangle.WhereClause("\"", "\"", 2, discordVoiceStatePrimaryKeyColumns),
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
		o.R = &discordVoiceStateR{
			Channel: related,
		}
	} else {
		o.R.Channel = related
	}

	if related.R == nil {
		related.R = &discordChannelR{
			ChannelDiscordVoiceStates: DiscordVoiceStateSlice{o},
		}
	} else {
		related.R.ChannelDiscordVoiceStates = append(related.R.ChannelDiscordVoiceStates, o)
	}

	return nil
}

// DiscordVoiceStatesG retrieves all records.
func DiscordVoiceStatesG(mods ...qm.QueryMod) discordVoiceStateQuery {
	return DiscordVoiceStates(boil.GetDB(), mods...)
}

// DiscordVoiceStates retrieves all the records using an executor.
func DiscordVoiceStates(exec boil.Executor, mods ...qm.QueryMod) discordVoiceStateQuery {
	mods = append(mods, qm.From("\"discord_voice_states\""))
	return discordVoiceStateQuery{NewQuery(exec, mods...)}
}

// FindDiscordVoiceStateG retrieves a single record by ID.
func FindDiscordVoiceStateG(userID int64, guildID int64, selectCols ...string) (*DiscordVoiceState, error) {
	return FindDiscordVoiceState(boil.GetDB(), userID, guildID, selectCols...)
}

// FindDiscordVoiceStateGP retrieves a single record by ID, and panics on error.
func FindDiscordVoiceStateGP(userID int64, guildID int64, selectCols ...string) *DiscordVoiceState {
	retobj, err := FindDiscordVoiceState(boil.GetDB(), userID, guildID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindDiscordVoiceState retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDiscordVoiceState(exec boil.Executor, userID int64, guildID int64, selectCols ...string) (*DiscordVoiceState, error) {
	discordVoiceStateObj := &DiscordVoiceState{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"discord_voice_states\" where \"user_id\"=$1 AND \"guild_id\"=$2", sel,
	)

	q := queries.Raw(exec, query, userID, guildID)

	err := q.Bind(discordVoiceStateObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from discord_voice_states")
	}

	return discordVoiceStateObj, nil
}

// FindDiscordVoiceStateP retrieves a single record by ID with an executor, and panics on error.
func FindDiscordVoiceStateP(exec boil.Executor, userID int64, guildID int64, selectCols ...string) *DiscordVoiceState {
	retobj, err := FindDiscordVoiceState(exec, userID, guildID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *DiscordVoiceState) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *DiscordVoiceState) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *DiscordVoiceState) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *DiscordVoiceState) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no discord_voice_states provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(discordVoiceStateColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	discordVoiceStateInsertCacheMut.RLock()
	cache, cached := discordVoiceStateInsertCache[key]
	discordVoiceStateInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			discordVoiceStateColumns,
			discordVoiceStateColumnsWithDefault,
			discordVoiceStateColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(discordVoiceStateType, discordVoiceStateMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(discordVoiceStateType, discordVoiceStateMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"discord_voice_states\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "models: unable to insert into discord_voice_states")
	}

	if !cached {
		discordVoiceStateInsertCacheMut.Lock()
		discordVoiceStateInsertCache[key] = cache
		discordVoiceStateInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single DiscordVoiceState record. See Update for
// whitelist behavior description.
func (o *DiscordVoiceState) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single DiscordVoiceState record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *DiscordVoiceState) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the DiscordVoiceState, and panics on error.
// See Update for whitelist behavior description.
func (o *DiscordVoiceState) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the DiscordVoiceState.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *DiscordVoiceState) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	key := makeCacheKey(whitelist, nil)
	discordVoiceStateUpdateCacheMut.RLock()
	cache, cached := discordVoiceStateUpdateCache[key]
	discordVoiceStateUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(discordVoiceStateColumns, discordVoiceStatePrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update discord_voice_states, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"discord_voice_states\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, discordVoiceStatePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(discordVoiceStateType, discordVoiceStateMapping, append(wl, discordVoiceStatePrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update discord_voice_states row")
	}

	if !cached {
		discordVoiceStateUpdateCacheMut.Lock()
		discordVoiceStateUpdateCache[key] = cache
		discordVoiceStateUpdateCacheMut.Unlock()
	}

	return nil
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q discordVoiceStateQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q discordVoiceStateQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for discord_voice_states")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o DiscordVoiceStateSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o DiscordVoiceStateSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o DiscordVoiceStateSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o DiscordVoiceStateSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), discordVoiceStatePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"discord_voice_states\" SET %s WHERE (\"user_id\",\"guild_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(discordVoiceStatePrimaryKeyColumns), len(colNames)+1, len(discordVoiceStatePrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in discordVoiceState slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *DiscordVoiceState) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *DiscordVoiceState) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *DiscordVoiceState) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *DiscordVoiceState) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no discord_voice_states provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(discordVoiceStateColumnsWithDefault, o)

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

	discordVoiceStateUpsertCacheMut.RLock()
	cache, cached := discordVoiceStateUpsertCache[key]
	discordVoiceStateUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			discordVoiceStateColumns,
			discordVoiceStateColumnsWithDefault,
			discordVoiceStateColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			discordVoiceStateColumns,
			discordVoiceStatePrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert discord_voice_states, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(discordVoiceStatePrimaryKeyColumns))
			copy(conflict, discordVoiceStatePrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"discord_voice_states\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(discordVoiceStateType, discordVoiceStateMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(discordVoiceStateType, discordVoiceStateMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert discord_voice_states")
	}

	if !cached {
		discordVoiceStateUpsertCacheMut.Lock()
		discordVoiceStateUpsertCache[key] = cache
		discordVoiceStateUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteP deletes a single DiscordVoiceState record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *DiscordVoiceState) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single DiscordVoiceState record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *DiscordVoiceState) DeleteG() error {
	if o == nil {
		return errors.New("models: no DiscordVoiceState provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single DiscordVoiceState record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *DiscordVoiceState) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single DiscordVoiceState record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *DiscordVoiceState) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no DiscordVoiceState provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), discordVoiceStatePrimaryKeyMapping)
	sql := "DELETE FROM \"discord_voice_states\" WHERE \"user_id\"=$1 AND \"guild_id\"=$2"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from discord_voice_states")
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q discordVoiceStateQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q discordVoiceStateQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no discordVoiceStateQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from discord_voice_states")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o DiscordVoiceStateSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o DiscordVoiceStateSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no DiscordVoiceState slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o DiscordVoiceStateSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o DiscordVoiceStateSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no DiscordVoiceState slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), discordVoiceStatePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"discord_voice_states\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, discordVoiceStatePrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(discordVoiceStatePrimaryKeyColumns), 1, len(discordVoiceStatePrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from discordVoiceState slice")
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *DiscordVoiceState) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *DiscordVoiceState) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *DiscordVoiceState) ReloadG() error {
	if o == nil {
		return errors.New("models: no DiscordVoiceState provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *DiscordVoiceState) Reload(exec boil.Executor) error {
	ret, err := FindDiscordVoiceState(exec, o.UserID, o.GuildID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *DiscordVoiceStateSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *DiscordVoiceStateSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DiscordVoiceStateSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty DiscordVoiceStateSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DiscordVoiceStateSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	discordVoiceStates := DiscordVoiceStateSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), discordVoiceStatePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"discord_voice_states\".* FROM \"discord_voice_states\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, discordVoiceStatePrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(discordVoiceStatePrimaryKeyColumns), 1, len(discordVoiceStatePrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&discordVoiceStates)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in DiscordVoiceStateSlice")
	}

	*o = discordVoiceStates

	return nil
}

// DiscordVoiceStateExists checks if the DiscordVoiceState row exists.
func DiscordVoiceStateExists(exec boil.Executor, userID int64, guildID int64) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"discord_voice_states\" where \"user_id\"=$1 AND \"guild_id\"=$2 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, userID, guildID)
	}

	row := exec.QueryRow(sql, userID, guildID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if discord_voice_states exists")
	}

	return exists, nil
}

// DiscordVoiceStateExistsG checks if the DiscordVoiceState row exists.
func DiscordVoiceStateExistsG(userID int64, guildID int64) (bool, error) {
	return DiscordVoiceStateExists(boil.GetDB(), userID, guildID)
}

// DiscordVoiceStateExistsGP checks if the DiscordVoiceState row exists. Panics on error.
func DiscordVoiceStateExistsGP(userID int64, guildID int64) bool {
	e, err := DiscordVoiceStateExists(boil.GetDB(), userID, guildID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// DiscordVoiceStateExistsP checks if the DiscordVoiceState row exists. Panics on error.
func DiscordVoiceStateExistsP(exec boil.Executor, userID int64, guildID int64) bool {
	e, err := DiscordVoiceStateExists(exec, userID, guildID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
