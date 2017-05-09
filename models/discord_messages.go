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
	"gopkg.in/nullbio/null.v6"
)

// DiscordMessage is an object representing the database table.
type DiscordMessage struct {
	ID              int64            `boil:"id" json:"id" toml:"id" yaml:"id"`
	ChannelID       int64            `boil:"channel_id" json:"channel_id" toml:"channel_id" yaml:"channel_id"`
	Timestamp       time.Time        `boil:"timestamp" json:"timestamp" toml:"timestamp" yaml:"timestamp"`
	EditedTimestamp time.Time        `boil:"edited_timestamp" json:"edited_timestamp" toml:"edited_timestamp" yaml:"edited_timestamp"`
	DeletedAt       null.Time        `boil:"deleted_at" json:"deleted_at,omitempty" toml:"deleted_at" yaml:"deleted_at,omitempty"`
	MentionRoles    types.Int64Array `boil:"mention_roles" json:"mention_roles" toml:"mention_roles" yaml:"mention_roles"`
	Mentions        types.Int64Array `boil:"mentions" json:"mentions" toml:"mentions" yaml:"mentions"`
	MentionEveryone bool             `boil:"mention_everyone" json:"mention_everyone" toml:"mention_everyone" yaml:"mention_everyone"`
	AuthorID        int64            `boil:"author_id" json:"author_id" toml:"author_id" yaml:"author_id"`
	AuthorUsername  string           `boil:"author_username" json:"author_username" toml:"author_username" yaml:"author_username"`
	AuthorDiscrim   int              `boil:"author_discrim" json:"author_discrim" toml:"author_discrim" yaml:"author_discrim"`
	AuthorAvatar    string           `boil:"author_avatar" json:"author_avatar" toml:"author_avatar" yaml:"author_avatar"`
	AuthorBot       bool             `boil:"author_bot" json:"author_bot" toml:"author_bot" yaml:"author_bot"`
	Content         string           `boil:"content" json:"content" toml:"content" yaml:"content"`
	Embeds          types.Int64Array `boil:"embeds" json:"embeds" toml:"embeds" yaml:"embeds"`

	R *discordMessageR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L discordMessageL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// discordMessageR is where relationships are stored.
type discordMessageR struct {
	MessageDiscordMessageEmbeds    DiscordMessageEmbedSlice
	MessageDiscordMessageRevisions DiscordMessageRevisionSlice
}

// discordMessageL is where Load methods for each relationship are stored.
type discordMessageL struct{}

var (
	discordMessageColumns               = []string{"id", "channel_id", "timestamp", "edited_timestamp", "deleted_at", "mention_roles", "mentions", "mention_everyone", "author_id", "author_username", "author_discrim", "author_avatar", "author_bot", "content", "embeds"}
	discordMessageColumnsWithoutDefault = []string{"id", "channel_id", "timestamp", "edited_timestamp", "deleted_at", "mention_roles", "mentions", "mention_everyone", "author_id", "author_username", "author_discrim", "author_avatar", "author_bot", "content", "embeds"}
	discordMessageColumnsWithDefault    = []string{}
	discordMessagePrimaryKeyColumns     = []string{"id"}
)

type (
	// DiscordMessageSlice is an alias for a slice of pointers to DiscordMessage.
	// This should generally be used opposed to []DiscordMessage.
	DiscordMessageSlice []*DiscordMessage

	discordMessageQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	discordMessageType                 = reflect.TypeOf(&DiscordMessage{})
	discordMessageMapping              = queries.MakeStructMapping(discordMessageType)
	discordMessagePrimaryKeyMapping, _ = queries.BindMapping(discordMessageType, discordMessageMapping, discordMessagePrimaryKeyColumns)
	discordMessageInsertCacheMut       sync.RWMutex
	discordMessageInsertCache          = make(map[string]insertCache)
	discordMessageUpdateCacheMut       sync.RWMutex
	discordMessageUpdateCache          = make(map[string]updateCache)
	discordMessageUpsertCacheMut       sync.RWMutex
	discordMessageUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)

// OneP returns a single discordMessage record from the query, and panics on error.
func (q discordMessageQuery) OneP() *DiscordMessage {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single discordMessage record from the query.
func (q discordMessageQuery) One() (*DiscordMessage, error) {
	o := &DiscordMessage{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for discord_messages")
	}

	return o, nil
}

// AllP returns all DiscordMessage records from the query, and panics on error.
func (q discordMessageQuery) AllP() DiscordMessageSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all DiscordMessage records from the query.
func (q discordMessageQuery) All() (DiscordMessageSlice, error) {
	var o DiscordMessageSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to DiscordMessage slice")
	}

	return o, nil
}

// CountP returns the count of all DiscordMessage records in the query, and panics on error.
func (q discordMessageQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all DiscordMessage records in the query.
func (q discordMessageQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count discord_messages rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q discordMessageQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q discordMessageQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if discord_messages exists")
	}

	return count > 0, nil
}

// MessageDiscordMessageEmbedsG retrieves all the discord_message_embed's discord message embeds via message_id column.
func (o *DiscordMessage) MessageDiscordMessageEmbedsG(mods ...qm.QueryMod) discordMessageEmbedQuery {
	return o.MessageDiscordMessageEmbeds(boil.GetDB(), mods...)
}

// MessageDiscordMessageEmbeds retrieves all the discord_message_embed's discord message embeds with an executor via message_id column.
func (o *DiscordMessage) MessageDiscordMessageEmbeds(exec boil.Executor, mods ...qm.QueryMod) discordMessageEmbedQuery {
	queryMods := []qm.QueryMod{
		qm.Select("\"a\".*"),
	}

	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"a\".\"message_id\"=?", o.ID),
	)

	query := DiscordMessageEmbeds(exec, queryMods...)
	queries.SetFrom(query.Query, "\"discord_message_embeds\" as \"a\"")
	return query
}

// MessageDiscordMessageRevisionsG retrieves all the discord_message_revision's discord message revisions via message_id column.
func (o *DiscordMessage) MessageDiscordMessageRevisionsG(mods ...qm.QueryMod) discordMessageRevisionQuery {
	return o.MessageDiscordMessageRevisions(boil.GetDB(), mods...)
}

// MessageDiscordMessageRevisions retrieves all the discord_message_revision's discord message revisions with an executor via message_id column.
func (o *DiscordMessage) MessageDiscordMessageRevisions(exec boil.Executor, mods ...qm.QueryMod) discordMessageRevisionQuery {
	queryMods := []qm.QueryMod{
		qm.Select("\"a\".*"),
	}

	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"a\".\"message_id\"=?", o.ID),
	)

	query := DiscordMessageRevisions(exec, queryMods...)
	queries.SetFrom(query.Query, "\"discord_message_revisions\" as \"a\"")
	return query
}

// LoadMessageDiscordMessageEmbeds allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (discordMessageL) LoadMessageDiscordMessageEmbeds(e boil.Executor, singular bool, maybeDiscordMessage interface{}) error {
	var slice []*DiscordMessage
	var object *DiscordMessage

	count := 1
	if singular {
		object = maybeDiscordMessage.(*DiscordMessage)
	} else {
		slice = *maybeDiscordMessage.(*DiscordMessageSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		if object.R == nil {
			object.R = &discordMessageR{}
		}
		args[0] = object.ID
	} else {
		for i, obj := range slice {
			if obj.R == nil {
				obj.R = &discordMessageR{}
			}
			args[i] = obj.ID
		}
	}

	query := fmt.Sprintf(
		"select * from \"discord_message_embeds\" where \"message_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)
	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load discord_message_embeds")
	}
	defer results.Close()

	var resultSlice []*DiscordMessageEmbed
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice discord_message_embeds")
	}

	if singular {
		object.R.MessageDiscordMessageEmbeds = resultSlice
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.MessageID {
				local.R.MessageDiscordMessageEmbeds = append(local.R.MessageDiscordMessageEmbeds, foreign)
				break
			}
		}
	}

	return nil
}

// LoadMessageDiscordMessageRevisions allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (discordMessageL) LoadMessageDiscordMessageRevisions(e boil.Executor, singular bool, maybeDiscordMessage interface{}) error {
	var slice []*DiscordMessage
	var object *DiscordMessage

	count := 1
	if singular {
		object = maybeDiscordMessage.(*DiscordMessage)
	} else {
		slice = *maybeDiscordMessage.(*DiscordMessageSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		if object.R == nil {
			object.R = &discordMessageR{}
		}
		args[0] = object.ID
	} else {
		for i, obj := range slice {
			if obj.R == nil {
				obj.R = &discordMessageR{}
			}
			args[i] = obj.ID
		}
	}

	query := fmt.Sprintf(
		"select * from \"discord_message_revisions\" where \"message_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)
	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load discord_message_revisions")
	}
	defer results.Close()

	var resultSlice []*DiscordMessageRevision
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice discord_message_revisions")
	}

	if singular {
		object.R.MessageDiscordMessageRevisions = resultSlice
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.MessageID {
				local.R.MessageDiscordMessageRevisions = append(local.R.MessageDiscordMessageRevisions, foreign)
				break
			}
		}
	}

	return nil
}

// AddMessageDiscordMessageEmbedsG adds the given related objects to the existing relationships
// of the discord_message, optionally inserting them as new records.
// Appends related to o.R.MessageDiscordMessageEmbeds.
// Sets related.R.Message appropriately.
// Uses the global database handle.
func (o *DiscordMessage) AddMessageDiscordMessageEmbedsG(insert bool, related ...*DiscordMessageEmbed) error {
	return o.AddMessageDiscordMessageEmbeds(boil.GetDB(), insert, related...)
}

// AddMessageDiscordMessageEmbedsP adds the given related objects to the existing relationships
// of the discord_message, optionally inserting them as new records.
// Appends related to o.R.MessageDiscordMessageEmbeds.
// Sets related.R.Message appropriately.
// Panics on error.
func (o *DiscordMessage) AddMessageDiscordMessageEmbedsP(exec boil.Executor, insert bool, related ...*DiscordMessageEmbed) {
	if err := o.AddMessageDiscordMessageEmbeds(exec, insert, related...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// AddMessageDiscordMessageEmbedsGP adds the given related objects to the existing relationships
// of the discord_message, optionally inserting them as new records.
// Appends related to o.R.MessageDiscordMessageEmbeds.
// Sets related.R.Message appropriately.
// Uses the global database handle and panics on error.
func (o *DiscordMessage) AddMessageDiscordMessageEmbedsGP(insert bool, related ...*DiscordMessageEmbed) {
	if err := o.AddMessageDiscordMessageEmbeds(boil.GetDB(), insert, related...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// AddMessageDiscordMessageEmbeds adds the given related objects to the existing relationships
// of the discord_message, optionally inserting them as new records.
// Appends related to o.R.MessageDiscordMessageEmbeds.
// Sets related.R.Message appropriately.
func (o *DiscordMessage) AddMessageDiscordMessageEmbeds(exec boil.Executor, insert bool, related ...*DiscordMessageEmbed) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.MessageID = o.ID
			if err = rel.Insert(exec); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"discord_message_embeds\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"message_id"}),
				strmangle.WhereClause("\"", "\"", 2, discordMessageEmbedPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}

			if _, err = exec.Exec(updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.MessageID = o.ID
		}
	}

	if o.R == nil {
		o.R = &discordMessageR{
			MessageDiscordMessageEmbeds: related,
		}
	} else {
		o.R.MessageDiscordMessageEmbeds = append(o.R.MessageDiscordMessageEmbeds, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &discordMessageEmbedR{
				Message: o,
			}
		} else {
			rel.R.Message = o
		}
	}
	return nil
}

// AddMessageDiscordMessageRevisionsG adds the given related objects to the existing relationships
// of the discord_message, optionally inserting them as new records.
// Appends related to o.R.MessageDiscordMessageRevisions.
// Sets related.R.Message appropriately.
// Uses the global database handle.
func (o *DiscordMessage) AddMessageDiscordMessageRevisionsG(insert bool, related ...*DiscordMessageRevision) error {
	return o.AddMessageDiscordMessageRevisions(boil.GetDB(), insert, related...)
}

// AddMessageDiscordMessageRevisionsP adds the given related objects to the existing relationships
// of the discord_message, optionally inserting them as new records.
// Appends related to o.R.MessageDiscordMessageRevisions.
// Sets related.R.Message appropriately.
// Panics on error.
func (o *DiscordMessage) AddMessageDiscordMessageRevisionsP(exec boil.Executor, insert bool, related ...*DiscordMessageRevision) {
	if err := o.AddMessageDiscordMessageRevisions(exec, insert, related...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// AddMessageDiscordMessageRevisionsGP adds the given related objects to the existing relationships
// of the discord_message, optionally inserting them as new records.
// Appends related to o.R.MessageDiscordMessageRevisions.
// Sets related.R.Message appropriately.
// Uses the global database handle and panics on error.
func (o *DiscordMessage) AddMessageDiscordMessageRevisionsGP(insert bool, related ...*DiscordMessageRevision) {
	if err := o.AddMessageDiscordMessageRevisions(boil.GetDB(), insert, related...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// AddMessageDiscordMessageRevisions adds the given related objects to the existing relationships
// of the discord_message, optionally inserting them as new records.
// Appends related to o.R.MessageDiscordMessageRevisions.
// Sets related.R.Message appropriately.
func (o *DiscordMessage) AddMessageDiscordMessageRevisions(exec boil.Executor, insert bool, related ...*DiscordMessageRevision) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.MessageID = o.ID
			if err = rel.Insert(exec); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"discord_message_revisions\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"message_id"}),
				strmangle.WhereClause("\"", "\"", 2, discordMessageRevisionPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.RevisionNum, rel.MessageID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}

			if _, err = exec.Exec(updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.MessageID = o.ID
		}
	}

	if o.R == nil {
		o.R = &discordMessageR{
			MessageDiscordMessageRevisions: related,
		}
	} else {
		o.R.MessageDiscordMessageRevisions = append(o.R.MessageDiscordMessageRevisions, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &discordMessageRevisionR{
				Message: o,
			}
		} else {
			rel.R.Message = o
		}
	}
	return nil
}

// DiscordMessagesG retrieves all records.
func DiscordMessagesG(mods ...qm.QueryMod) discordMessageQuery {
	return DiscordMessages(boil.GetDB(), mods...)
}

// DiscordMessages retrieves all the records using an executor.
func DiscordMessages(exec boil.Executor, mods ...qm.QueryMod) discordMessageQuery {
	mods = append(mods, qm.From("\"discord_messages\""))
	return discordMessageQuery{NewQuery(exec, mods...)}
}

// FindDiscordMessageG retrieves a single record by ID.
func FindDiscordMessageG(id int64, selectCols ...string) (*DiscordMessage, error) {
	return FindDiscordMessage(boil.GetDB(), id, selectCols...)
}

// FindDiscordMessageGP retrieves a single record by ID, and panics on error.
func FindDiscordMessageGP(id int64, selectCols ...string) *DiscordMessage {
	retobj, err := FindDiscordMessage(boil.GetDB(), id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindDiscordMessage retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDiscordMessage(exec boil.Executor, id int64, selectCols ...string) (*DiscordMessage, error) {
	discordMessageObj := &DiscordMessage{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"discord_messages\" where \"id\"=$1", sel,
	)

	q := queries.Raw(exec, query, id)

	err := q.Bind(discordMessageObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from discord_messages")
	}

	return discordMessageObj, nil
}

// FindDiscordMessageP retrieves a single record by ID with an executor, and panics on error.
func FindDiscordMessageP(exec boil.Executor, id int64, selectCols ...string) *DiscordMessage {
	retobj, err := FindDiscordMessage(exec, id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *DiscordMessage) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *DiscordMessage) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *DiscordMessage) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *DiscordMessage) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no discord_messages provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(discordMessageColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	discordMessageInsertCacheMut.RLock()
	cache, cached := discordMessageInsertCache[key]
	discordMessageInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			discordMessageColumns,
			discordMessageColumnsWithDefault,
			discordMessageColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(discordMessageType, discordMessageMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(discordMessageType, discordMessageMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"discord_messages\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "models: unable to insert into discord_messages")
	}

	if !cached {
		discordMessageInsertCacheMut.Lock()
		discordMessageInsertCache[key] = cache
		discordMessageInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single DiscordMessage record. See Update for
// whitelist behavior description.
func (o *DiscordMessage) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single DiscordMessage record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *DiscordMessage) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the DiscordMessage, and panics on error.
// See Update for whitelist behavior description.
func (o *DiscordMessage) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the DiscordMessage.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *DiscordMessage) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	key := makeCacheKey(whitelist, nil)
	discordMessageUpdateCacheMut.RLock()
	cache, cached := discordMessageUpdateCache[key]
	discordMessageUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(discordMessageColumns, discordMessagePrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update discord_messages, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"discord_messages\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, discordMessagePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(discordMessageType, discordMessageMapping, append(wl, discordMessagePrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update discord_messages row")
	}

	if !cached {
		discordMessageUpdateCacheMut.Lock()
		discordMessageUpdateCache[key] = cache
		discordMessageUpdateCacheMut.Unlock()
	}

	return nil
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q discordMessageQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q discordMessageQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for discord_messages")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o DiscordMessageSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o DiscordMessageSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o DiscordMessageSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o DiscordMessageSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), discordMessagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"discord_messages\" SET %s WHERE (\"id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(discordMessagePrimaryKeyColumns), len(colNames)+1, len(discordMessagePrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in discordMessage slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *DiscordMessage) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *DiscordMessage) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *DiscordMessage) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *DiscordMessage) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no discord_messages provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(discordMessageColumnsWithDefault, o)

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

	discordMessageUpsertCacheMut.RLock()
	cache, cached := discordMessageUpsertCache[key]
	discordMessageUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			discordMessageColumns,
			discordMessageColumnsWithDefault,
			discordMessageColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			discordMessageColumns,
			discordMessagePrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert discord_messages, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(discordMessagePrimaryKeyColumns))
			copy(conflict, discordMessagePrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"discord_messages\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(discordMessageType, discordMessageMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(discordMessageType, discordMessageMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert discord_messages")
	}

	if !cached {
		discordMessageUpsertCacheMut.Lock()
		discordMessageUpsertCache[key] = cache
		discordMessageUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteP deletes a single DiscordMessage record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *DiscordMessage) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single DiscordMessage record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *DiscordMessage) DeleteG() error {
	if o == nil {
		return errors.New("models: no DiscordMessage provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single DiscordMessage record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *DiscordMessage) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single DiscordMessage record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *DiscordMessage) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no DiscordMessage provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), discordMessagePrimaryKeyMapping)
	sql := "DELETE FROM \"discord_messages\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from discord_messages")
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q discordMessageQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q discordMessageQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no discordMessageQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from discord_messages")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o DiscordMessageSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o DiscordMessageSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no DiscordMessage slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o DiscordMessageSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o DiscordMessageSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no DiscordMessage slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), discordMessagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"discord_messages\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, discordMessagePrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(discordMessagePrimaryKeyColumns), 1, len(discordMessagePrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from discordMessage slice")
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *DiscordMessage) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *DiscordMessage) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *DiscordMessage) ReloadG() error {
	if o == nil {
		return errors.New("models: no DiscordMessage provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *DiscordMessage) Reload(exec boil.Executor) error {
	ret, err := FindDiscordMessage(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *DiscordMessageSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *DiscordMessageSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DiscordMessageSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty DiscordMessageSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DiscordMessageSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	discordMessages := DiscordMessageSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), discordMessagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"discord_messages\".* FROM \"discord_messages\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, discordMessagePrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(discordMessagePrimaryKeyColumns), 1, len(discordMessagePrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&discordMessages)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in DiscordMessageSlice")
	}

	*o = discordMessages

	return nil
}

// DiscordMessageExists checks if the DiscordMessage row exists.
func DiscordMessageExists(exec boil.Executor, id int64) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"discord_messages\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, id)
	}

	row := exec.QueryRow(sql, id)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if discord_messages exists")
	}

	return exists, nil
}

// DiscordMessageExistsG checks if the DiscordMessage row exists.
func DiscordMessageExistsG(id int64) (bool, error) {
	return DiscordMessageExists(boil.GetDB(), id)
}

// DiscordMessageExistsGP checks if the DiscordMessage row exists. Panics on error.
func DiscordMessageExistsGP(id int64) bool {
	e, err := DiscordMessageExists(boil.GetDB(), id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// DiscordMessageExistsP checks if the DiscordMessage row exists. Panics on error.
func DiscordMessageExistsP(exec boil.Executor, id int64) bool {
	e, err := DiscordMessageExists(exec, id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
