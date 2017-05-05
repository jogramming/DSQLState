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

// DiscordPrivateChannel is an object representing the database table.
type DiscordPrivateChannel struct {
	ID            int64     `boil:"id" json:"id" toml:"id" yaml:"id"`
	CreatedAt     time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	RecipientID   int64     `boil:"recipient_id" json:"recipient_id" toml:"recipient_id" yaml:"recipient_id"`
	Name          string    `boil:"name" json:"name" toml:"name" yaml:"name"`
	Topic         string    `boil:"topic" json:"topic" toml:"topic" yaml:"topic"`
	LastMessageID int64     `boil:"last_message_id" json:"last_message_id" toml:"last_message_id" yaml:"last_message_id"`

	R *discordPrivateChannelR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L discordPrivateChannelL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// discordPrivateChannelR is where relationships are stored.
type discordPrivateChannelR struct {
	Recipient *DiscordUser
}

// discordPrivateChannelL is where Load methods for each relationship are stored.
type discordPrivateChannelL struct{}

var (
	discordPrivateChannelColumns               = []string{"id", "created_at", "recipient_id", "name", "topic", "last_message_id"}
	discordPrivateChannelColumnsWithoutDefault = []string{"id", "created_at", "recipient_id", "name", "topic", "last_message_id"}
	discordPrivateChannelColumnsWithDefault    = []string{}
	discordPrivateChannelPrimaryKeyColumns     = []string{"id"}
)

type (
	// DiscordPrivateChannelSlice is an alias for a slice of pointers to DiscordPrivateChannel.
	// This should generally be used opposed to []DiscordPrivateChannel.
	DiscordPrivateChannelSlice []*DiscordPrivateChannel

	discordPrivateChannelQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	discordPrivateChannelType                 = reflect.TypeOf(&DiscordPrivateChannel{})
	discordPrivateChannelMapping              = queries.MakeStructMapping(discordPrivateChannelType)
	discordPrivateChannelPrimaryKeyMapping, _ = queries.BindMapping(discordPrivateChannelType, discordPrivateChannelMapping, discordPrivateChannelPrimaryKeyColumns)
	discordPrivateChannelInsertCacheMut       sync.RWMutex
	discordPrivateChannelInsertCache          = make(map[string]insertCache)
	discordPrivateChannelUpdateCacheMut       sync.RWMutex
	discordPrivateChannelUpdateCache          = make(map[string]updateCache)
	discordPrivateChannelUpsertCacheMut       sync.RWMutex
	discordPrivateChannelUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)

// OneP returns a single discordPrivateChannel record from the query, and panics on error.
func (q discordPrivateChannelQuery) OneP() *DiscordPrivateChannel {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single discordPrivateChannel record from the query.
func (q discordPrivateChannelQuery) One() (*DiscordPrivateChannel, error) {
	o := &DiscordPrivateChannel{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for discord_private_channels")
	}

	return o, nil
}

// AllP returns all DiscordPrivateChannel records from the query, and panics on error.
func (q discordPrivateChannelQuery) AllP() DiscordPrivateChannelSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all DiscordPrivateChannel records from the query.
func (q discordPrivateChannelQuery) All() (DiscordPrivateChannelSlice, error) {
	var o DiscordPrivateChannelSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to DiscordPrivateChannel slice")
	}

	return o, nil
}

// CountP returns the count of all DiscordPrivateChannel records in the query, and panics on error.
func (q discordPrivateChannelQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all DiscordPrivateChannel records in the query.
func (q discordPrivateChannelQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count discord_private_channels rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q discordPrivateChannelQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q discordPrivateChannelQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if discord_private_channels exists")
	}

	return count > 0, nil
}

// RecipientG pointed to by the foreign key.
func (o *DiscordPrivateChannel) RecipientG(mods ...qm.QueryMod) discordUserQuery {
	return o.Recipient(boil.GetDB(), mods...)
}

// Recipient pointed to by the foreign key.
func (o *DiscordPrivateChannel) Recipient(exec boil.Executor, mods ...qm.QueryMod) discordUserQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.RecipientID),
	}

	queryMods = append(queryMods, mods...)

	query := DiscordUsers(exec, queryMods...)
	queries.SetFrom(query.Query, "\"discord_users\"")

	return query
}

// LoadRecipient allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (discordPrivateChannelL) LoadRecipient(e boil.Executor, singular bool, maybeDiscordPrivateChannel interface{}) error {
	var slice []*DiscordPrivateChannel
	var object *DiscordPrivateChannel

	count := 1
	if singular {
		object = maybeDiscordPrivateChannel.(*DiscordPrivateChannel)
	} else {
		slice = *maybeDiscordPrivateChannel.(*DiscordPrivateChannelSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		if object.R == nil {
			object.R = &discordPrivateChannelR{}
		}
		args[0] = object.RecipientID
	} else {
		for i, obj := range slice {
			if obj.R == nil {
				obj.R = &discordPrivateChannelR{}
			}
			args[i] = obj.RecipientID
		}
	}

	query := fmt.Sprintf(
		"select * from \"discord_users\" where \"id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load DiscordUser")
	}
	defer results.Close()

	var resultSlice []*DiscordUser
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice DiscordUser")
	}

	if singular && len(resultSlice) != 0 {
		object.R.Recipient = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.RecipientID == foreign.ID {
				local.R.Recipient = foreign
				break
			}
		}
	}

	return nil
}

// SetRecipientG of the discord_private_channel to the related item.
// Sets o.R.Recipient to related.
// Adds o to related.R.RecipientDiscordPrivateChannels.
// Uses the global database handle.
func (o *DiscordPrivateChannel) SetRecipientG(insert bool, related *DiscordUser) error {
	return o.SetRecipient(boil.GetDB(), insert, related)
}

// SetRecipientP of the discord_private_channel to the related item.
// Sets o.R.Recipient to related.
// Adds o to related.R.RecipientDiscordPrivateChannels.
// Panics on error.
func (o *DiscordPrivateChannel) SetRecipientP(exec boil.Executor, insert bool, related *DiscordUser) {
	if err := o.SetRecipient(exec, insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetRecipientGP of the discord_private_channel to the related item.
// Sets o.R.Recipient to related.
// Adds o to related.R.RecipientDiscordPrivateChannels.
// Uses the global database handle and panics on error.
func (o *DiscordPrivateChannel) SetRecipientGP(insert bool, related *DiscordUser) {
	if err := o.SetRecipient(boil.GetDB(), insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetRecipient of the discord_private_channel to the related item.
// Sets o.R.Recipient to related.
// Adds o to related.R.RecipientDiscordPrivateChannels.
func (o *DiscordPrivateChannel) SetRecipient(exec boil.Executor, insert bool, related *DiscordUser) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"discord_private_channels\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"recipient_id"}),
		strmangle.WhereClause("\"", "\"", 2, discordPrivateChannelPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.RecipientID = related.ID

	if o.R == nil {
		o.R = &discordPrivateChannelR{
			Recipient: related,
		}
	} else {
		o.R.Recipient = related
	}

	if related.R == nil {
		related.R = &discordUserR{
			RecipientDiscordPrivateChannels: DiscordPrivateChannelSlice{o},
		}
	} else {
		related.R.RecipientDiscordPrivateChannels = append(related.R.RecipientDiscordPrivateChannels, o)
	}

	return nil
}

// DiscordPrivateChannelsG retrieves all records.
func DiscordPrivateChannelsG(mods ...qm.QueryMod) discordPrivateChannelQuery {
	return DiscordPrivateChannels(boil.GetDB(), mods...)
}

// DiscordPrivateChannels retrieves all the records using an executor.
func DiscordPrivateChannels(exec boil.Executor, mods ...qm.QueryMod) discordPrivateChannelQuery {
	mods = append(mods, qm.From("\"discord_private_channels\""))
	return discordPrivateChannelQuery{NewQuery(exec, mods...)}
}

// FindDiscordPrivateChannelG retrieves a single record by ID.
func FindDiscordPrivateChannelG(id int64, selectCols ...string) (*DiscordPrivateChannel, error) {
	return FindDiscordPrivateChannel(boil.GetDB(), id, selectCols...)
}

// FindDiscordPrivateChannelGP retrieves a single record by ID, and panics on error.
func FindDiscordPrivateChannelGP(id int64, selectCols ...string) *DiscordPrivateChannel {
	retobj, err := FindDiscordPrivateChannel(boil.GetDB(), id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindDiscordPrivateChannel retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDiscordPrivateChannel(exec boil.Executor, id int64, selectCols ...string) (*DiscordPrivateChannel, error) {
	discordPrivateChannelObj := &DiscordPrivateChannel{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"discord_private_channels\" where \"id\"=$1", sel,
	)

	q := queries.Raw(exec, query, id)

	err := q.Bind(discordPrivateChannelObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from discord_private_channels")
	}

	return discordPrivateChannelObj, nil
}

// FindDiscordPrivateChannelP retrieves a single record by ID with an executor, and panics on error.
func FindDiscordPrivateChannelP(exec boil.Executor, id int64, selectCols ...string) *DiscordPrivateChannel {
	retobj, err := FindDiscordPrivateChannel(exec, id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *DiscordPrivateChannel) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *DiscordPrivateChannel) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *DiscordPrivateChannel) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *DiscordPrivateChannel) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no discord_private_channels provided for insertion")
	}

	var err error
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(discordPrivateChannelColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	discordPrivateChannelInsertCacheMut.RLock()
	cache, cached := discordPrivateChannelInsertCache[key]
	discordPrivateChannelInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			discordPrivateChannelColumns,
			discordPrivateChannelColumnsWithDefault,
			discordPrivateChannelColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(discordPrivateChannelType, discordPrivateChannelMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(discordPrivateChannelType, discordPrivateChannelMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"discord_private_channels\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "models: unable to insert into discord_private_channels")
	}

	if !cached {
		discordPrivateChannelInsertCacheMut.Lock()
		discordPrivateChannelInsertCache[key] = cache
		discordPrivateChannelInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single DiscordPrivateChannel record. See Update for
// whitelist behavior description.
func (o *DiscordPrivateChannel) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single DiscordPrivateChannel record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *DiscordPrivateChannel) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the DiscordPrivateChannel, and panics on error.
// See Update for whitelist behavior description.
func (o *DiscordPrivateChannel) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the DiscordPrivateChannel.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *DiscordPrivateChannel) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	key := makeCacheKey(whitelist, nil)
	discordPrivateChannelUpdateCacheMut.RLock()
	cache, cached := discordPrivateChannelUpdateCache[key]
	discordPrivateChannelUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(discordPrivateChannelColumns, discordPrivateChannelPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update discord_private_channels, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"discord_private_channels\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, discordPrivateChannelPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(discordPrivateChannelType, discordPrivateChannelMapping, append(wl, discordPrivateChannelPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update discord_private_channels row")
	}

	if !cached {
		discordPrivateChannelUpdateCacheMut.Lock()
		discordPrivateChannelUpdateCache[key] = cache
		discordPrivateChannelUpdateCacheMut.Unlock()
	}

	return nil
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q discordPrivateChannelQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q discordPrivateChannelQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for discord_private_channels")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o DiscordPrivateChannelSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o DiscordPrivateChannelSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o DiscordPrivateChannelSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o DiscordPrivateChannelSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), discordPrivateChannelPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"discord_private_channels\" SET %s WHERE (\"id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(discordPrivateChannelPrimaryKeyColumns), len(colNames)+1, len(discordPrivateChannelPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in discordPrivateChannel slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *DiscordPrivateChannel) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *DiscordPrivateChannel) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *DiscordPrivateChannel) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *DiscordPrivateChannel) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no discord_private_channels provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(discordPrivateChannelColumnsWithDefault, o)

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

	discordPrivateChannelUpsertCacheMut.RLock()
	cache, cached := discordPrivateChannelUpsertCache[key]
	discordPrivateChannelUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			discordPrivateChannelColumns,
			discordPrivateChannelColumnsWithDefault,
			discordPrivateChannelColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			discordPrivateChannelColumns,
			discordPrivateChannelPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert discord_private_channels, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(discordPrivateChannelPrimaryKeyColumns))
			copy(conflict, discordPrivateChannelPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"discord_private_channels\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(discordPrivateChannelType, discordPrivateChannelMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(discordPrivateChannelType, discordPrivateChannelMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert discord_private_channels")
	}

	if !cached {
		discordPrivateChannelUpsertCacheMut.Lock()
		discordPrivateChannelUpsertCache[key] = cache
		discordPrivateChannelUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteP deletes a single DiscordPrivateChannel record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *DiscordPrivateChannel) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single DiscordPrivateChannel record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *DiscordPrivateChannel) DeleteG() error {
	if o == nil {
		return errors.New("models: no DiscordPrivateChannel provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single DiscordPrivateChannel record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *DiscordPrivateChannel) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single DiscordPrivateChannel record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *DiscordPrivateChannel) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no DiscordPrivateChannel provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), discordPrivateChannelPrimaryKeyMapping)
	sql := "DELETE FROM \"discord_private_channels\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from discord_private_channels")
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q discordPrivateChannelQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q discordPrivateChannelQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no discordPrivateChannelQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from discord_private_channels")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o DiscordPrivateChannelSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o DiscordPrivateChannelSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no DiscordPrivateChannel slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o DiscordPrivateChannelSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o DiscordPrivateChannelSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no DiscordPrivateChannel slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), discordPrivateChannelPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"discord_private_channels\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, discordPrivateChannelPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(discordPrivateChannelPrimaryKeyColumns), 1, len(discordPrivateChannelPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from discordPrivateChannel slice")
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *DiscordPrivateChannel) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *DiscordPrivateChannel) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *DiscordPrivateChannel) ReloadG() error {
	if o == nil {
		return errors.New("models: no DiscordPrivateChannel provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *DiscordPrivateChannel) Reload(exec boil.Executor) error {
	ret, err := FindDiscordPrivateChannel(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *DiscordPrivateChannelSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *DiscordPrivateChannelSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DiscordPrivateChannelSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty DiscordPrivateChannelSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DiscordPrivateChannelSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	discordPrivateChannels := DiscordPrivateChannelSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), discordPrivateChannelPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"discord_private_channels\".* FROM \"discord_private_channels\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, discordPrivateChannelPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(discordPrivateChannelPrimaryKeyColumns), 1, len(discordPrivateChannelPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&discordPrivateChannels)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in DiscordPrivateChannelSlice")
	}

	*o = discordPrivateChannels

	return nil
}

// DiscordPrivateChannelExists checks if the DiscordPrivateChannel row exists.
func DiscordPrivateChannelExists(exec boil.Executor, id int64) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"discord_private_channels\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, id)
	}

	row := exec.QueryRow(sql, id)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if discord_private_channels exists")
	}

	return exists, nil
}

// DiscordPrivateChannelExistsG checks if the DiscordPrivateChannel row exists.
func DiscordPrivateChannelExistsG(id int64) (bool, error) {
	return DiscordPrivateChannelExists(boil.GetDB(), id)
}

// DiscordPrivateChannelExistsGP checks if the DiscordPrivateChannel row exists. Panics on error.
func DiscordPrivateChannelExistsGP(id int64) bool {
	e, err := DiscordPrivateChannelExists(boil.GetDB(), id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// DiscordPrivateChannelExistsP checks if the DiscordPrivateChannel row exists. Panics on error.
func DiscordPrivateChannelExistsP(exec boil.Executor, id int64) bool {
	e, err := DiscordPrivateChannelExists(exec, id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
