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

// DiscordGuildChannel is an object representing the database table.
type DiscordGuildChannel struct {
	ID            int64     `boil:"id" json:"id" toml:"id" yaml:"id"`
	GuildID       int64     `boil:"guild_id" json:"guild_id" toml:"guild_id" yaml:"guild_id"`
	CreatedAt     time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	DeletedAt     null.Time `boil:"deleted_at" json:"deleted_at,omitempty" toml:"deleted_at" yaml:"deleted_at,omitempty"`
	Name          string    `boil:"name" json:"name" toml:"name" yaml:"name"`
	Topic         string    `boil:"topic" json:"topic" toml:"topic" yaml:"topic"`
	Type          string    `boil:"type" json:"type" toml:"type" yaml:"type"`
	LastMessageID int64     `boil:"last_message_id" json:"last_message_id" toml:"last_message_id" yaml:"last_message_id"`
	Position      int       `boil:"position" json:"position" toml:"position" yaml:"position"`
	Bitrate       int       `boil:"bitrate" json:"bitrate" toml:"bitrate" yaml:"bitrate"`

	R *discordGuildChannelR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L discordGuildChannelL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// discordGuildChannelR is where relationships are stored.
type discordGuildChannelR struct {
	Guild *DiscordGuild
}

// discordGuildChannelL is where Load methods for each relationship are stored.
type discordGuildChannelL struct{}

var (
	discordGuildChannelColumns               = []string{"id", "guild_id", "created_at", "deleted_at", "name", "topic", "type", "last_message_id", "position", "bitrate"}
	discordGuildChannelColumnsWithoutDefault = []string{"id", "guild_id", "created_at", "deleted_at", "name", "topic", "type", "last_message_id", "position", "bitrate"}
	discordGuildChannelColumnsWithDefault    = []string{}
	discordGuildChannelPrimaryKeyColumns     = []string{"id"}
)

type (
	// DiscordGuildChannelSlice is an alias for a slice of pointers to DiscordGuildChannel.
	// This should generally be used opposed to []DiscordGuildChannel.
	DiscordGuildChannelSlice []*DiscordGuildChannel

	discordGuildChannelQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	discordGuildChannelType                 = reflect.TypeOf(&DiscordGuildChannel{})
	discordGuildChannelMapping              = queries.MakeStructMapping(discordGuildChannelType)
	discordGuildChannelPrimaryKeyMapping, _ = queries.BindMapping(discordGuildChannelType, discordGuildChannelMapping, discordGuildChannelPrimaryKeyColumns)
	discordGuildChannelInsertCacheMut       sync.RWMutex
	discordGuildChannelInsertCache          = make(map[string]insertCache)
	discordGuildChannelUpdateCacheMut       sync.RWMutex
	discordGuildChannelUpdateCache          = make(map[string]updateCache)
	discordGuildChannelUpsertCacheMut       sync.RWMutex
	discordGuildChannelUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)

// OneP returns a single discordGuildChannel record from the query, and panics on error.
func (q discordGuildChannelQuery) OneP() *DiscordGuildChannel {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single discordGuildChannel record from the query.
func (q discordGuildChannelQuery) One() (*DiscordGuildChannel, error) {
	o := &DiscordGuildChannel{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for discord_guild_channels")
	}

	return o, nil
}

// AllP returns all DiscordGuildChannel records from the query, and panics on error.
func (q discordGuildChannelQuery) AllP() DiscordGuildChannelSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all DiscordGuildChannel records from the query.
func (q discordGuildChannelQuery) All() (DiscordGuildChannelSlice, error) {
	var o DiscordGuildChannelSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to DiscordGuildChannel slice")
	}

	return o, nil
}

// CountP returns the count of all DiscordGuildChannel records in the query, and panics on error.
func (q discordGuildChannelQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all DiscordGuildChannel records in the query.
func (q discordGuildChannelQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count discord_guild_channels rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q discordGuildChannelQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q discordGuildChannelQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if discord_guild_channels exists")
	}

	return count > 0, nil
}

// GuildG pointed to by the foreign key.
func (o *DiscordGuildChannel) GuildG(mods ...qm.QueryMod) discordGuildQuery {
	return o.Guild(boil.GetDB(), mods...)
}

// Guild pointed to by the foreign key.
func (o *DiscordGuildChannel) Guild(exec boil.Executor, mods ...qm.QueryMod) discordGuildQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.GuildID),
	}

	queryMods = append(queryMods, mods...)

	query := DiscordGuilds(exec, queryMods...)
	queries.SetFrom(query.Query, "\"discord_guilds\"")

	return query
}

// LoadGuild allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (discordGuildChannelL) LoadGuild(e boil.Executor, singular bool, maybeDiscordGuildChannel interface{}) error {
	var slice []*DiscordGuildChannel
	var object *DiscordGuildChannel

	count := 1
	if singular {
		object = maybeDiscordGuildChannel.(*DiscordGuildChannel)
	} else {
		slice = *maybeDiscordGuildChannel.(*DiscordGuildChannelSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		if object.R == nil {
			object.R = &discordGuildChannelR{}
		}
		args[0] = object.GuildID
	} else {
		for i, obj := range slice {
			if obj.R == nil {
				obj.R = &discordGuildChannelR{}
			}
			args[i] = obj.GuildID
		}
	}

	query := fmt.Sprintf(
		"select * from \"discord_guilds\" where \"id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load DiscordGuild")
	}
	defer results.Close()

	var resultSlice []*DiscordGuild
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice DiscordGuild")
	}

	if singular && len(resultSlice) != 0 {
		object.R.Guild = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.GuildID == foreign.ID {
				local.R.Guild = foreign
				break
			}
		}
	}

	return nil
}

// SetGuildG of the discord_guild_channel to the related item.
// Sets o.R.Guild to related.
// Adds o to related.R.GuildDiscordGuildChannels.
// Uses the global database handle.
func (o *DiscordGuildChannel) SetGuildG(insert bool, related *DiscordGuild) error {
	return o.SetGuild(boil.GetDB(), insert, related)
}

// SetGuildP of the discord_guild_channel to the related item.
// Sets o.R.Guild to related.
// Adds o to related.R.GuildDiscordGuildChannels.
// Panics on error.
func (o *DiscordGuildChannel) SetGuildP(exec boil.Executor, insert bool, related *DiscordGuild) {
	if err := o.SetGuild(exec, insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetGuildGP of the discord_guild_channel to the related item.
// Sets o.R.Guild to related.
// Adds o to related.R.GuildDiscordGuildChannels.
// Uses the global database handle and panics on error.
func (o *DiscordGuildChannel) SetGuildGP(insert bool, related *DiscordGuild) {
	if err := o.SetGuild(boil.GetDB(), insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetGuild of the discord_guild_channel to the related item.
// Sets o.R.Guild to related.
// Adds o to related.R.GuildDiscordGuildChannels.
func (o *DiscordGuildChannel) SetGuild(exec boil.Executor, insert bool, related *DiscordGuild) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"discord_guild_channels\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"guild_id"}),
		strmangle.WhereClause("\"", "\"", 2, discordGuildChannelPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.GuildID = related.ID

	if o.R == nil {
		o.R = &discordGuildChannelR{
			Guild: related,
		}
	} else {
		o.R.Guild = related
	}

	if related.R == nil {
		related.R = &discordGuildR{
			GuildDiscordGuildChannels: DiscordGuildChannelSlice{o},
		}
	} else {
		related.R.GuildDiscordGuildChannels = append(related.R.GuildDiscordGuildChannels, o)
	}

	return nil
}

// DiscordGuildChannelsG retrieves all records.
func DiscordGuildChannelsG(mods ...qm.QueryMod) discordGuildChannelQuery {
	return DiscordGuildChannels(boil.GetDB(), mods...)
}

// DiscordGuildChannels retrieves all the records using an executor.
func DiscordGuildChannels(exec boil.Executor, mods ...qm.QueryMod) discordGuildChannelQuery {
	mods = append(mods, qm.From("\"discord_guild_channels\""))
	return discordGuildChannelQuery{NewQuery(exec, mods...)}
}

// FindDiscordGuildChannelG retrieves a single record by ID.
func FindDiscordGuildChannelG(id int64, selectCols ...string) (*DiscordGuildChannel, error) {
	return FindDiscordGuildChannel(boil.GetDB(), id, selectCols...)
}

// FindDiscordGuildChannelGP retrieves a single record by ID, and panics on error.
func FindDiscordGuildChannelGP(id int64, selectCols ...string) *DiscordGuildChannel {
	retobj, err := FindDiscordGuildChannel(boil.GetDB(), id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindDiscordGuildChannel retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDiscordGuildChannel(exec boil.Executor, id int64, selectCols ...string) (*DiscordGuildChannel, error) {
	discordGuildChannelObj := &DiscordGuildChannel{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"discord_guild_channels\" where \"id\"=$1", sel,
	)

	q := queries.Raw(exec, query, id)

	err := q.Bind(discordGuildChannelObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from discord_guild_channels")
	}

	return discordGuildChannelObj, nil
}

// FindDiscordGuildChannelP retrieves a single record by ID with an executor, and panics on error.
func FindDiscordGuildChannelP(exec boil.Executor, id int64, selectCols ...string) *DiscordGuildChannel {
	retobj, err := FindDiscordGuildChannel(exec, id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *DiscordGuildChannel) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *DiscordGuildChannel) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *DiscordGuildChannel) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *DiscordGuildChannel) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no discord_guild_channels provided for insertion")
	}

	var err error
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(discordGuildChannelColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	discordGuildChannelInsertCacheMut.RLock()
	cache, cached := discordGuildChannelInsertCache[key]
	discordGuildChannelInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			discordGuildChannelColumns,
			discordGuildChannelColumnsWithDefault,
			discordGuildChannelColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(discordGuildChannelType, discordGuildChannelMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(discordGuildChannelType, discordGuildChannelMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"discord_guild_channels\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "models: unable to insert into discord_guild_channels")
	}

	if !cached {
		discordGuildChannelInsertCacheMut.Lock()
		discordGuildChannelInsertCache[key] = cache
		discordGuildChannelInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single DiscordGuildChannel record. See Update for
// whitelist behavior description.
func (o *DiscordGuildChannel) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single DiscordGuildChannel record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *DiscordGuildChannel) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the DiscordGuildChannel, and panics on error.
// See Update for whitelist behavior description.
func (o *DiscordGuildChannel) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the DiscordGuildChannel.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *DiscordGuildChannel) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	key := makeCacheKey(whitelist, nil)
	discordGuildChannelUpdateCacheMut.RLock()
	cache, cached := discordGuildChannelUpdateCache[key]
	discordGuildChannelUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(discordGuildChannelColumns, discordGuildChannelPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update discord_guild_channels, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"discord_guild_channels\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, discordGuildChannelPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(discordGuildChannelType, discordGuildChannelMapping, append(wl, discordGuildChannelPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update discord_guild_channels row")
	}

	if !cached {
		discordGuildChannelUpdateCacheMut.Lock()
		discordGuildChannelUpdateCache[key] = cache
		discordGuildChannelUpdateCacheMut.Unlock()
	}

	return nil
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q discordGuildChannelQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q discordGuildChannelQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for discord_guild_channels")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o DiscordGuildChannelSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o DiscordGuildChannelSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o DiscordGuildChannelSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o DiscordGuildChannelSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), discordGuildChannelPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"discord_guild_channels\" SET %s WHERE (\"id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(discordGuildChannelPrimaryKeyColumns), len(colNames)+1, len(discordGuildChannelPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in discordGuildChannel slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *DiscordGuildChannel) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *DiscordGuildChannel) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *DiscordGuildChannel) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *DiscordGuildChannel) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no discord_guild_channels provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(discordGuildChannelColumnsWithDefault, o)

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

	discordGuildChannelUpsertCacheMut.RLock()
	cache, cached := discordGuildChannelUpsertCache[key]
	discordGuildChannelUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			discordGuildChannelColumns,
			discordGuildChannelColumnsWithDefault,
			discordGuildChannelColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			discordGuildChannelColumns,
			discordGuildChannelPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert discord_guild_channels, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(discordGuildChannelPrimaryKeyColumns))
			copy(conflict, discordGuildChannelPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"discord_guild_channels\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(discordGuildChannelType, discordGuildChannelMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(discordGuildChannelType, discordGuildChannelMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert discord_guild_channels")
	}

	if !cached {
		discordGuildChannelUpsertCacheMut.Lock()
		discordGuildChannelUpsertCache[key] = cache
		discordGuildChannelUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteP deletes a single DiscordGuildChannel record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *DiscordGuildChannel) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single DiscordGuildChannel record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *DiscordGuildChannel) DeleteG() error {
	if o == nil {
		return errors.New("models: no DiscordGuildChannel provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single DiscordGuildChannel record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *DiscordGuildChannel) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single DiscordGuildChannel record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *DiscordGuildChannel) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no DiscordGuildChannel provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), discordGuildChannelPrimaryKeyMapping)
	sql := "DELETE FROM \"discord_guild_channels\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from discord_guild_channels")
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q discordGuildChannelQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q discordGuildChannelQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no discordGuildChannelQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from discord_guild_channels")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o DiscordGuildChannelSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o DiscordGuildChannelSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no DiscordGuildChannel slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o DiscordGuildChannelSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o DiscordGuildChannelSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no DiscordGuildChannel slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), discordGuildChannelPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"discord_guild_channels\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, discordGuildChannelPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(discordGuildChannelPrimaryKeyColumns), 1, len(discordGuildChannelPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from discordGuildChannel slice")
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *DiscordGuildChannel) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *DiscordGuildChannel) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *DiscordGuildChannel) ReloadG() error {
	if o == nil {
		return errors.New("models: no DiscordGuildChannel provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *DiscordGuildChannel) Reload(exec boil.Executor) error {
	ret, err := FindDiscordGuildChannel(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *DiscordGuildChannelSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *DiscordGuildChannelSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DiscordGuildChannelSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty DiscordGuildChannelSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DiscordGuildChannelSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	discordGuildChannels := DiscordGuildChannelSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), discordGuildChannelPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"discord_guild_channels\".* FROM \"discord_guild_channels\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, discordGuildChannelPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(discordGuildChannelPrimaryKeyColumns), 1, len(discordGuildChannelPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&discordGuildChannels)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in DiscordGuildChannelSlice")
	}

	*o = discordGuildChannels

	return nil
}

// DiscordGuildChannelExists checks if the DiscordGuildChannel row exists.
func DiscordGuildChannelExists(exec boil.Executor, id int64) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"discord_guild_channels\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, id)
	}

	row := exec.QueryRow(sql, id)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if discord_guild_channels exists")
	}

	return exists, nil
}

// DiscordGuildChannelExistsG checks if the DiscordGuildChannel row exists.
func DiscordGuildChannelExistsG(id int64) (bool, error) {
	return DiscordGuildChannelExists(boil.GetDB(), id)
}

// DiscordGuildChannelExistsGP checks if the DiscordGuildChannel row exists. Panics on error.
func DiscordGuildChannelExistsGP(id int64) bool {
	e, err := DiscordGuildChannelExists(boil.GetDB(), id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// DiscordGuildChannelExistsP checks if the DiscordGuildChannel row exists. Panics on error.
func DiscordGuildChannelExistsP(exec boil.Executor, id int64) bool {
	e, err := DiscordGuildChannelExists(exec, id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
