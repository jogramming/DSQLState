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

// DMessage is an object representing the database table.
type DMessage struct {
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

	R *dMessageR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L dMessageL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// dMessageR is where relationships are stored.
type dMessageR struct {
	MessageDMessageRevisions DMessageRevisionSlice
	MessageDMessageEmbeds    DMessageEmbedSlice
}

// dMessageL is where Load methods for each relationship are stored.
type dMessageL struct{}

var (
	dMessageColumns               = []string{"id", "channel_id", "timestamp", "edited_timestamp", "deleted_at", "mention_roles", "mentions", "mention_everyone", "author_id", "author_username", "author_discrim", "author_avatar", "author_bot", "content", "embeds"}
	dMessageColumnsWithoutDefault = []string{"id", "channel_id", "timestamp", "edited_timestamp", "deleted_at", "mention_roles", "mentions", "mention_everyone", "author_id", "author_username", "author_discrim", "author_avatar", "author_bot", "content", "embeds"}
	dMessageColumnsWithDefault    = []string{}
	dMessagePrimaryKeyColumns     = []string{"id"}
)

type (
	// DMessageSlice is an alias for a slice of pointers to DMessage.
	// This should generally be used opposed to []DMessage.
	DMessageSlice []*DMessage

	dMessageQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	dMessageType                 = reflect.TypeOf(&DMessage{})
	dMessageMapping              = queries.MakeStructMapping(dMessageType)
	dMessagePrimaryKeyMapping, _ = queries.BindMapping(dMessageType, dMessageMapping, dMessagePrimaryKeyColumns)
	dMessageInsertCacheMut       sync.RWMutex
	dMessageInsertCache          = make(map[string]insertCache)
	dMessageUpdateCacheMut       sync.RWMutex
	dMessageUpdateCache          = make(map[string]updateCache)
	dMessageUpsertCacheMut       sync.RWMutex
	dMessageUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)

// OneP returns a single dMessage record from the query, and panics on error.
func (q dMessageQuery) OneP() *DMessage {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single dMessage record from the query.
func (q dMessageQuery) One() (*DMessage, error) {
	o := &DMessage{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for d_messages")
	}

	return o, nil
}

// AllP returns all DMessage records from the query, and panics on error.
func (q dMessageQuery) AllP() DMessageSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all DMessage records from the query.
func (q dMessageQuery) All() (DMessageSlice, error) {
	var o DMessageSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to DMessage slice")
	}

	return o, nil
}

// CountP returns the count of all DMessage records in the query, and panics on error.
func (q dMessageQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all DMessage records in the query.
func (q dMessageQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count d_messages rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q dMessageQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q dMessageQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if d_messages exists")
	}

	return count > 0, nil
}

// MessageDMessageRevisionsG retrieves all the d_message_revision's d message revisions via message_id column.
func (o *DMessage) MessageDMessageRevisionsG(mods ...qm.QueryMod) dMessageRevisionQuery {
	return o.MessageDMessageRevisions(boil.GetDB(), mods...)
}

// MessageDMessageRevisions retrieves all the d_message_revision's d message revisions with an executor via message_id column.
func (o *DMessage) MessageDMessageRevisions(exec boil.Executor, mods ...qm.QueryMod) dMessageRevisionQuery {
	queryMods := []qm.QueryMod{
		qm.Select("\"a\".*"),
	}

	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"a\".\"message_id\"=?", o.ID),
	)

	query := DMessageRevisions(exec, queryMods...)
	queries.SetFrom(query.Query, "\"d_message_revisions\" as \"a\"")
	return query
}

// MessageDMessageEmbedsG retrieves all the d_message_embed's d message embeds via message_id column.
func (o *DMessage) MessageDMessageEmbedsG(mods ...qm.QueryMod) dMessageEmbedQuery {
	return o.MessageDMessageEmbeds(boil.GetDB(), mods...)
}

// MessageDMessageEmbeds retrieves all the d_message_embed's d message embeds with an executor via message_id column.
func (o *DMessage) MessageDMessageEmbeds(exec boil.Executor, mods ...qm.QueryMod) dMessageEmbedQuery {
	queryMods := []qm.QueryMod{
		qm.Select("\"a\".*"),
	}

	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"a\".\"message_id\"=?", o.ID),
	)

	query := DMessageEmbeds(exec, queryMods...)
	queries.SetFrom(query.Query, "\"d_message_embeds\" as \"a\"")
	return query
}

// LoadMessageDMessageRevisions allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (dMessageL) LoadMessageDMessageRevisions(e boil.Executor, singular bool, maybeDMessage interface{}) error {
	var slice []*DMessage
	var object *DMessage

	count := 1
	if singular {
		object = maybeDMessage.(*DMessage)
	} else {
		slice = *maybeDMessage.(*DMessageSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		if object.R == nil {
			object.R = &dMessageR{}
		}
		args[0] = object.ID
	} else {
		for i, obj := range slice {
			if obj.R == nil {
				obj.R = &dMessageR{}
			}
			args[i] = obj.ID
		}
	}

	query := fmt.Sprintf(
		"select * from \"d_message_revisions\" where \"message_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)
	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load d_message_revisions")
	}
	defer results.Close()

	var resultSlice []*DMessageRevision
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice d_message_revisions")
	}

	if singular {
		object.R.MessageDMessageRevisions = resultSlice
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.MessageID {
				local.R.MessageDMessageRevisions = append(local.R.MessageDMessageRevisions, foreign)
				break
			}
		}
	}

	return nil
}

// LoadMessageDMessageEmbeds allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (dMessageL) LoadMessageDMessageEmbeds(e boil.Executor, singular bool, maybeDMessage interface{}) error {
	var slice []*DMessage
	var object *DMessage

	count := 1
	if singular {
		object = maybeDMessage.(*DMessage)
	} else {
		slice = *maybeDMessage.(*DMessageSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		if object.R == nil {
			object.R = &dMessageR{}
		}
		args[0] = object.ID
	} else {
		for i, obj := range slice {
			if obj.R == nil {
				obj.R = &dMessageR{}
			}
			args[i] = obj.ID
		}
	}

	query := fmt.Sprintf(
		"select * from \"d_message_embeds\" where \"message_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)
	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load d_message_embeds")
	}
	defer results.Close()

	var resultSlice []*DMessageEmbed
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice d_message_embeds")
	}

	if singular {
		object.R.MessageDMessageEmbeds = resultSlice
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.MessageID {
				local.R.MessageDMessageEmbeds = append(local.R.MessageDMessageEmbeds, foreign)
				break
			}
		}
	}

	return nil
}

// AddMessageDMessageRevisionsG adds the given related objects to the existing relationships
// of the d_message, optionally inserting them as new records.
// Appends related to o.R.MessageDMessageRevisions.
// Sets related.R.Message appropriately.
// Uses the global database handle.
func (o *DMessage) AddMessageDMessageRevisionsG(insert bool, related ...*DMessageRevision) error {
	return o.AddMessageDMessageRevisions(boil.GetDB(), insert, related...)
}

// AddMessageDMessageRevisionsP adds the given related objects to the existing relationships
// of the d_message, optionally inserting them as new records.
// Appends related to o.R.MessageDMessageRevisions.
// Sets related.R.Message appropriately.
// Panics on error.
func (o *DMessage) AddMessageDMessageRevisionsP(exec boil.Executor, insert bool, related ...*DMessageRevision) {
	if err := o.AddMessageDMessageRevisions(exec, insert, related...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// AddMessageDMessageRevisionsGP adds the given related objects to the existing relationships
// of the d_message, optionally inserting them as new records.
// Appends related to o.R.MessageDMessageRevisions.
// Sets related.R.Message appropriately.
// Uses the global database handle and panics on error.
func (o *DMessage) AddMessageDMessageRevisionsGP(insert bool, related ...*DMessageRevision) {
	if err := o.AddMessageDMessageRevisions(boil.GetDB(), insert, related...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// AddMessageDMessageRevisions adds the given related objects to the existing relationships
// of the d_message, optionally inserting them as new records.
// Appends related to o.R.MessageDMessageRevisions.
// Sets related.R.Message appropriately.
func (o *DMessage) AddMessageDMessageRevisions(exec boil.Executor, insert bool, related ...*DMessageRevision) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.MessageID = o.ID
			if err = rel.Insert(exec); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"d_message_revisions\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"message_id"}),
				strmangle.WhereClause("\"", "\"", 2, dMessageRevisionPrimaryKeyColumns),
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
		o.R = &dMessageR{
			MessageDMessageRevisions: related,
		}
	} else {
		o.R.MessageDMessageRevisions = append(o.R.MessageDMessageRevisions, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &dMessageRevisionR{
				Message: o,
			}
		} else {
			rel.R.Message = o
		}
	}
	return nil
}

// AddMessageDMessageEmbedsG adds the given related objects to the existing relationships
// of the d_message, optionally inserting them as new records.
// Appends related to o.R.MessageDMessageEmbeds.
// Sets related.R.Message appropriately.
// Uses the global database handle.
func (o *DMessage) AddMessageDMessageEmbedsG(insert bool, related ...*DMessageEmbed) error {
	return o.AddMessageDMessageEmbeds(boil.GetDB(), insert, related...)
}

// AddMessageDMessageEmbedsP adds the given related objects to the existing relationships
// of the d_message, optionally inserting them as new records.
// Appends related to o.R.MessageDMessageEmbeds.
// Sets related.R.Message appropriately.
// Panics on error.
func (o *DMessage) AddMessageDMessageEmbedsP(exec boil.Executor, insert bool, related ...*DMessageEmbed) {
	if err := o.AddMessageDMessageEmbeds(exec, insert, related...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// AddMessageDMessageEmbedsGP adds the given related objects to the existing relationships
// of the d_message, optionally inserting them as new records.
// Appends related to o.R.MessageDMessageEmbeds.
// Sets related.R.Message appropriately.
// Uses the global database handle and panics on error.
func (o *DMessage) AddMessageDMessageEmbedsGP(insert bool, related ...*DMessageEmbed) {
	if err := o.AddMessageDMessageEmbeds(boil.GetDB(), insert, related...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// AddMessageDMessageEmbeds adds the given related objects to the existing relationships
// of the d_message, optionally inserting them as new records.
// Appends related to o.R.MessageDMessageEmbeds.
// Sets related.R.Message appropriately.
func (o *DMessage) AddMessageDMessageEmbeds(exec boil.Executor, insert bool, related ...*DMessageEmbed) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.MessageID = o.ID
			if err = rel.Insert(exec); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"d_message_embeds\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"message_id"}),
				strmangle.WhereClause("\"", "\"", 2, dMessageEmbedPrimaryKeyColumns),
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
		o.R = &dMessageR{
			MessageDMessageEmbeds: related,
		}
	} else {
		o.R.MessageDMessageEmbeds = append(o.R.MessageDMessageEmbeds, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &dMessageEmbedR{
				Message: o,
			}
		} else {
			rel.R.Message = o
		}
	}
	return nil
}

// DMessagesG retrieves all records.
func DMessagesG(mods ...qm.QueryMod) dMessageQuery {
	return DMessages(boil.GetDB(), mods...)
}

// DMessages retrieves all the records using an executor.
func DMessages(exec boil.Executor, mods ...qm.QueryMod) dMessageQuery {
	mods = append(mods, qm.From("\"d_messages\""))
	return dMessageQuery{NewQuery(exec, mods...)}
}

// FindDMessageG retrieves a single record by ID.
func FindDMessageG(id int64, selectCols ...string) (*DMessage, error) {
	return FindDMessage(boil.GetDB(), id, selectCols...)
}

// FindDMessageGP retrieves a single record by ID, and panics on error.
func FindDMessageGP(id int64, selectCols ...string) *DMessage {
	retobj, err := FindDMessage(boil.GetDB(), id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindDMessage retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDMessage(exec boil.Executor, id int64, selectCols ...string) (*DMessage, error) {
	dMessageObj := &DMessage{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"d_messages\" where \"id\"=$1", sel,
	)

	q := queries.Raw(exec, query, id)

	err := q.Bind(dMessageObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from d_messages")
	}

	return dMessageObj, nil
}

// FindDMessageP retrieves a single record by ID with an executor, and panics on error.
func FindDMessageP(exec boil.Executor, id int64, selectCols ...string) *DMessage {
	retobj, err := FindDMessage(exec, id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *DMessage) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *DMessage) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *DMessage) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *DMessage) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no d_messages provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(dMessageColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	dMessageInsertCacheMut.RLock()
	cache, cached := dMessageInsertCache[key]
	dMessageInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			dMessageColumns,
			dMessageColumnsWithDefault,
			dMessageColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(dMessageType, dMessageMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(dMessageType, dMessageMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"d_messages\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "models: unable to insert into d_messages")
	}

	if !cached {
		dMessageInsertCacheMut.Lock()
		dMessageInsertCache[key] = cache
		dMessageInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single DMessage record. See Update for
// whitelist behavior description.
func (o *DMessage) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single DMessage record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *DMessage) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the DMessage, and panics on error.
// See Update for whitelist behavior description.
func (o *DMessage) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the DMessage.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *DMessage) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	key := makeCacheKey(whitelist, nil)
	dMessageUpdateCacheMut.RLock()
	cache, cached := dMessageUpdateCache[key]
	dMessageUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(dMessageColumns, dMessagePrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update d_messages, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"d_messages\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, dMessagePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(dMessageType, dMessageMapping, append(wl, dMessagePrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update d_messages row")
	}

	if !cached {
		dMessageUpdateCacheMut.Lock()
		dMessageUpdateCache[key] = cache
		dMessageUpdateCacheMut.Unlock()
	}

	return nil
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q dMessageQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q dMessageQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for d_messages")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o DMessageSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o DMessageSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o DMessageSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o DMessageSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dMessagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"d_messages\" SET %s WHERE (\"id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(dMessagePrimaryKeyColumns), len(colNames)+1, len(dMessagePrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in dMessage slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *DMessage) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *DMessage) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *DMessage) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *DMessage) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no d_messages provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(dMessageColumnsWithDefault, o)

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

	dMessageUpsertCacheMut.RLock()
	cache, cached := dMessageUpsertCache[key]
	dMessageUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			dMessageColumns,
			dMessageColumnsWithDefault,
			dMessageColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			dMessageColumns,
			dMessagePrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert d_messages, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(dMessagePrimaryKeyColumns))
			copy(conflict, dMessagePrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"d_messages\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(dMessageType, dMessageMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(dMessageType, dMessageMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert d_messages")
	}

	if !cached {
		dMessageUpsertCacheMut.Lock()
		dMessageUpsertCache[key] = cache
		dMessageUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteP deletes a single DMessage record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *DMessage) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single DMessage record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *DMessage) DeleteG() error {
	if o == nil {
		return errors.New("models: no DMessage provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single DMessage record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *DMessage) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single DMessage record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *DMessage) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no DMessage provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), dMessagePrimaryKeyMapping)
	sql := "DELETE FROM \"d_messages\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from d_messages")
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q dMessageQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q dMessageQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no dMessageQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from d_messages")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o DMessageSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o DMessageSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no DMessage slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o DMessageSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o DMessageSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no DMessage slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dMessagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"d_messages\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, dMessagePrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(dMessagePrimaryKeyColumns), 1, len(dMessagePrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from dMessage slice")
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *DMessage) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *DMessage) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *DMessage) ReloadG() error {
	if o == nil {
		return errors.New("models: no DMessage provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *DMessage) Reload(exec boil.Executor) error {
	ret, err := FindDMessage(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *DMessageSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *DMessageSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DMessageSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty DMessageSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DMessageSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	dMessages := DMessageSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dMessagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"d_messages\".* FROM \"d_messages\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, dMessagePrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(dMessagePrimaryKeyColumns), 1, len(dMessagePrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&dMessages)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in DMessageSlice")
	}

	*o = dMessages

	return nil
}

// DMessageExists checks if the DMessage row exists.
func DMessageExists(exec boil.Executor, id int64) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"d_messages\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, id)
	}

	row := exec.QueryRow(sql, id)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if d_messages exists")
	}

	return exists, nil
}

// DMessageExistsG checks if the DMessage row exists.
func DMessageExistsG(id int64) (bool, error) {
	return DMessageExists(boil.GetDB(), id)
}

// DMessageExistsGP checks if the DMessage row exists. Panics on error.
func DMessageExistsGP(id int64) bool {
	e, err := DMessageExists(boil.GetDB(), id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// DMessageExistsP checks if the DMessage row exists. Panics on error.
func DMessageExistsP(exec boil.Executor, id int64) bool {
	e, err := DMessageExists(exec, id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
