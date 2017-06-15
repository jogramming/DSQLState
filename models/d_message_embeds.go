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

// DMessageEmbed is an object representing the database table.
type DMessageEmbed struct {
	ID                 int64             `boil:"id" json:"id" toml:"id" yaml:"id"`
	MessageID          int64             `boil:"message_id" json:"message_id" toml:"message_id" yaml:"message_id"`
	RevisionNum        int               `boil:"revision_num" json:"revision_num" toml:"revision_num" yaml:"revision_num"`
	URL                string            `boil:"url" json:"url" toml:"url" yaml:"url"`
	Type               string            `boil:"type" json:"type" toml:"type" yaml:"type"`
	Title              string            `boil:"title" json:"title" toml:"title" yaml:"title"`
	Description        string            `boil:"description" json:"description" toml:"description" yaml:"description"`
	Timestamp          string            `boil:"timestamp" json:"timestamp" toml:"timestamp" yaml:"timestamp"`
	Color              int               `boil:"color" json:"color" toml:"color" yaml:"color"`
	FieldNames         types.StringArray `boil:"field_names" json:"field_names" toml:"field_names" yaml:"field_names"`
	FieldValues        types.StringArray `boil:"field_values" json:"field_values" toml:"field_values" yaml:"field_values"`
	FieldInlines       types.BoolArray   `boil:"field_inlines" json:"field_inlines" toml:"field_inlines" yaml:"field_inlines"`
	FooterText         null.String       `boil:"footer_text" json:"footer_text,omitempty" toml:"footer_text" yaml:"footer_text,omitempty"`
	FooterIconURL      null.String       `boil:"footer_icon_url" json:"footer_icon_url,omitempty" toml:"footer_icon_url" yaml:"footer_icon_url,omitempty"`
	FooterProxyIconURL null.String       `boil:"footer_proxy_icon_url" json:"footer_proxy_icon_url,omitempty" toml:"footer_proxy_icon_url" yaml:"footer_proxy_icon_url,omitempty"`
	ImageURL           null.String       `boil:"image_url" json:"image_url,omitempty" toml:"image_url" yaml:"image_url,omitempty"`
	ImageProxyURL      null.String       `boil:"image_proxy_url" json:"image_proxy_url,omitempty" toml:"image_proxy_url" yaml:"image_proxy_url,omitempty"`
	ImageWidth         null.Int          `boil:"image_width" json:"image_width,omitempty" toml:"image_width" yaml:"image_width,omitempty"`
	ImageHeight        null.Int          `boil:"image_height" json:"image_height,omitempty" toml:"image_height" yaml:"image_height,omitempty"`
	ThumbnailURL       null.String       `boil:"thumbnail_url" json:"thumbnail_url,omitempty" toml:"thumbnail_url" yaml:"thumbnail_url,omitempty"`
	ThumbnailProxyURL  null.String       `boil:"thumbnail_proxy_url" json:"thumbnail_proxy_url,omitempty" toml:"thumbnail_proxy_url" yaml:"thumbnail_proxy_url,omitempty"`
	ThumbnailWidth     null.Int          `boil:"thumbnail_width" json:"thumbnail_width,omitempty" toml:"thumbnail_width" yaml:"thumbnail_width,omitempty"`
	ThumbnailHeight    null.Int          `boil:"thumbnail_height" json:"thumbnail_height,omitempty" toml:"thumbnail_height" yaml:"thumbnail_height,omitempty"`
	VideoURL           null.String       `boil:"video_url" json:"video_url,omitempty" toml:"video_url" yaml:"video_url,omitempty"`
	VideoProxyURL      null.String       `boil:"video_proxy_url" json:"video_proxy_url,omitempty" toml:"video_proxy_url" yaml:"video_proxy_url,omitempty"`
	VideoWidth         null.Int          `boil:"video_width" json:"video_width,omitempty" toml:"video_width" yaml:"video_width,omitempty"`
	VideoHeight        null.Int          `boil:"video_height" json:"video_height,omitempty" toml:"video_height" yaml:"video_height,omitempty"`
	ProviderURL        null.String       `boil:"provider_url" json:"provider_url,omitempty" toml:"provider_url" yaml:"provider_url,omitempty"`
	ProviderName       null.String       `boil:"provider_name" json:"provider_name,omitempty" toml:"provider_name" yaml:"provider_name,omitempty"`
	AuthorURL          null.String       `boil:"author_url" json:"author_url,omitempty" toml:"author_url" yaml:"author_url,omitempty"`
	AuthorName         null.String       `boil:"author_name" json:"author_name,omitempty" toml:"author_name" yaml:"author_name,omitempty"`
	AuthorIconURL      null.String       `boil:"author_icon_url" json:"author_icon_url,omitempty" toml:"author_icon_url" yaml:"author_icon_url,omitempty"`
	AuthorProxyIconURL null.String       `boil:"author_proxy_icon_url" json:"author_proxy_icon_url,omitempty" toml:"author_proxy_icon_url" yaml:"author_proxy_icon_url,omitempty"`

	R *dMessageEmbedR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L dMessageEmbedL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// dMessageEmbedR is where relationships are stored.
type dMessageEmbedR struct {
	Message *DMessage
}

// dMessageEmbedL is where Load methods for each relationship are stored.
type dMessageEmbedL struct{}

var (
	dMessageEmbedColumns               = []string{"id", "message_id", "revision_num", "url", "type", "title", "description", "timestamp", "color", "field_names", "field_values", "field_inlines", "footer_text", "footer_icon_url", "footer_proxy_icon_url", "image_url", "image_proxy_url", "image_width", "image_height", "thumbnail_url", "thumbnail_proxy_url", "thumbnail_width", "thumbnail_height", "video_url", "video_proxy_url", "video_width", "video_height", "provider_url", "provider_name", "author_url", "author_name", "author_icon_url", "author_proxy_icon_url"}
	dMessageEmbedColumnsWithoutDefault = []string{"message_id", "revision_num", "url", "type", "title", "description", "timestamp", "color", "field_names", "field_values", "field_inlines", "footer_text", "footer_icon_url", "footer_proxy_icon_url", "image_url", "image_proxy_url", "image_width", "image_height", "thumbnail_url", "thumbnail_proxy_url", "thumbnail_width", "thumbnail_height", "video_url", "video_proxy_url", "video_width", "video_height", "provider_url", "provider_name", "author_url", "author_name", "author_icon_url", "author_proxy_icon_url"}
	dMessageEmbedColumnsWithDefault    = []string{"id"}
	dMessageEmbedPrimaryKeyColumns     = []string{"id"}
)

type (
	// DMessageEmbedSlice is an alias for a slice of pointers to DMessageEmbed.
	// This should generally be used opposed to []DMessageEmbed.
	DMessageEmbedSlice []*DMessageEmbed

	dMessageEmbedQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	dMessageEmbedType                 = reflect.TypeOf(&DMessageEmbed{})
	dMessageEmbedMapping              = queries.MakeStructMapping(dMessageEmbedType)
	dMessageEmbedPrimaryKeyMapping, _ = queries.BindMapping(dMessageEmbedType, dMessageEmbedMapping, dMessageEmbedPrimaryKeyColumns)
	dMessageEmbedInsertCacheMut       sync.RWMutex
	dMessageEmbedInsertCache          = make(map[string]insertCache)
	dMessageEmbedUpdateCacheMut       sync.RWMutex
	dMessageEmbedUpdateCache          = make(map[string]updateCache)
	dMessageEmbedUpsertCacheMut       sync.RWMutex
	dMessageEmbedUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)

// OneP returns a single dMessageEmbed record from the query, and panics on error.
func (q dMessageEmbedQuery) OneP() *DMessageEmbed {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single dMessageEmbed record from the query.
func (q dMessageEmbedQuery) One() (*DMessageEmbed, error) {
	o := &DMessageEmbed{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for d_message_embeds")
	}

	return o, nil
}

// AllP returns all DMessageEmbed records from the query, and panics on error.
func (q dMessageEmbedQuery) AllP() DMessageEmbedSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all DMessageEmbed records from the query.
func (q dMessageEmbedQuery) All() (DMessageEmbedSlice, error) {
	var o DMessageEmbedSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to DMessageEmbed slice")
	}

	return o, nil
}

// CountP returns the count of all DMessageEmbed records in the query, and panics on error.
func (q dMessageEmbedQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all DMessageEmbed records in the query.
func (q dMessageEmbedQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count d_message_embeds rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q dMessageEmbedQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q dMessageEmbedQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if d_message_embeds exists")
	}

	return count > 0, nil
}

// MessageG pointed to by the foreign key.
func (o *DMessageEmbed) MessageG(mods ...qm.QueryMod) dMessageQuery {
	return o.Message(boil.GetDB(), mods...)
}

// Message pointed to by the foreign key.
func (o *DMessageEmbed) Message(exec boil.Executor, mods ...qm.QueryMod) dMessageQuery {
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
func (dMessageEmbedL) LoadMessage(e boil.Executor, singular bool, maybeDMessageEmbed interface{}) error {
	var slice []*DMessageEmbed
	var object *DMessageEmbed

	count := 1
	if singular {
		object = maybeDMessageEmbed.(*DMessageEmbed)
	} else {
		slice = *maybeDMessageEmbed.(*DMessageEmbedSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		if object.R == nil {
			object.R = &dMessageEmbedR{}
		}
		args[0] = object.MessageID
	} else {
		for i, obj := range slice {
			if obj.R == nil {
				obj.R = &dMessageEmbedR{}
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

// SetMessageG of the d_message_embed to the related item.
// Sets o.R.Message to related.
// Adds o to related.R.MessageDMessageEmbeds.
// Uses the global database handle.
func (o *DMessageEmbed) SetMessageG(insert bool, related *DMessage) error {
	return o.SetMessage(boil.GetDB(), insert, related)
}

// SetMessageP of the d_message_embed to the related item.
// Sets o.R.Message to related.
// Adds o to related.R.MessageDMessageEmbeds.
// Panics on error.
func (o *DMessageEmbed) SetMessageP(exec boil.Executor, insert bool, related *DMessage) {
	if err := o.SetMessage(exec, insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetMessageGP of the d_message_embed to the related item.
// Sets o.R.Message to related.
// Adds o to related.R.MessageDMessageEmbeds.
// Uses the global database handle and panics on error.
func (o *DMessageEmbed) SetMessageGP(insert bool, related *DMessage) {
	if err := o.SetMessage(boil.GetDB(), insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetMessage of the d_message_embed to the related item.
// Sets o.R.Message to related.
// Adds o to related.R.MessageDMessageEmbeds.
func (o *DMessageEmbed) SetMessage(exec boil.Executor, insert bool, related *DMessage) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"d_message_embeds\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"message_id"}),
		strmangle.WhereClause("\"", "\"", 2, dMessageEmbedPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.MessageID = related.ID

	if o.R == nil {
		o.R = &dMessageEmbedR{
			Message: related,
		}
	} else {
		o.R.Message = related
	}

	if related.R == nil {
		related.R = &dMessageR{
			MessageDMessageEmbeds: DMessageEmbedSlice{o},
		}
	} else {
		related.R.MessageDMessageEmbeds = append(related.R.MessageDMessageEmbeds, o)
	}

	return nil
}

// DMessageEmbedsG retrieves all records.
func DMessageEmbedsG(mods ...qm.QueryMod) dMessageEmbedQuery {
	return DMessageEmbeds(boil.GetDB(), mods...)
}

// DMessageEmbeds retrieves all the records using an executor.
func DMessageEmbeds(exec boil.Executor, mods ...qm.QueryMod) dMessageEmbedQuery {
	mods = append(mods, qm.From("\"d_message_embeds\""))
	return dMessageEmbedQuery{NewQuery(exec, mods...)}
}

// FindDMessageEmbedG retrieves a single record by ID.
func FindDMessageEmbedG(id int64, selectCols ...string) (*DMessageEmbed, error) {
	return FindDMessageEmbed(boil.GetDB(), id, selectCols...)
}

// FindDMessageEmbedGP retrieves a single record by ID, and panics on error.
func FindDMessageEmbedGP(id int64, selectCols ...string) *DMessageEmbed {
	retobj, err := FindDMessageEmbed(boil.GetDB(), id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindDMessageEmbed retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDMessageEmbed(exec boil.Executor, id int64, selectCols ...string) (*DMessageEmbed, error) {
	dMessageEmbedObj := &DMessageEmbed{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"d_message_embeds\" where \"id\"=$1", sel,
	)

	q := queries.Raw(exec, query, id)

	err := q.Bind(dMessageEmbedObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from d_message_embeds")
	}

	return dMessageEmbedObj, nil
}

// FindDMessageEmbedP retrieves a single record by ID with an executor, and panics on error.
func FindDMessageEmbedP(exec boil.Executor, id int64, selectCols ...string) *DMessageEmbed {
	retobj, err := FindDMessageEmbed(exec, id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *DMessageEmbed) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *DMessageEmbed) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *DMessageEmbed) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *DMessageEmbed) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no d_message_embeds provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(dMessageEmbedColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	dMessageEmbedInsertCacheMut.RLock()
	cache, cached := dMessageEmbedInsertCache[key]
	dMessageEmbedInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			dMessageEmbedColumns,
			dMessageEmbedColumnsWithDefault,
			dMessageEmbedColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(dMessageEmbedType, dMessageEmbedMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(dMessageEmbedType, dMessageEmbedMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"d_message_embeds\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "models: unable to insert into d_message_embeds")
	}

	if !cached {
		dMessageEmbedInsertCacheMut.Lock()
		dMessageEmbedInsertCache[key] = cache
		dMessageEmbedInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single DMessageEmbed record. See Update for
// whitelist behavior description.
func (o *DMessageEmbed) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single DMessageEmbed record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *DMessageEmbed) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the DMessageEmbed, and panics on error.
// See Update for whitelist behavior description.
func (o *DMessageEmbed) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the DMessageEmbed.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *DMessageEmbed) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	key := makeCacheKey(whitelist, nil)
	dMessageEmbedUpdateCacheMut.RLock()
	cache, cached := dMessageEmbedUpdateCache[key]
	dMessageEmbedUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(dMessageEmbedColumns, dMessageEmbedPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update d_message_embeds, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"d_message_embeds\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, dMessageEmbedPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(dMessageEmbedType, dMessageEmbedMapping, append(wl, dMessageEmbedPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update d_message_embeds row")
	}

	if !cached {
		dMessageEmbedUpdateCacheMut.Lock()
		dMessageEmbedUpdateCache[key] = cache
		dMessageEmbedUpdateCacheMut.Unlock()
	}

	return nil
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q dMessageEmbedQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q dMessageEmbedQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for d_message_embeds")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o DMessageEmbedSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o DMessageEmbedSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o DMessageEmbedSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o DMessageEmbedSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dMessageEmbedPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"d_message_embeds\" SET %s WHERE (\"id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(dMessageEmbedPrimaryKeyColumns), len(colNames)+1, len(dMessageEmbedPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in dMessageEmbed slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *DMessageEmbed) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *DMessageEmbed) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *DMessageEmbed) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *DMessageEmbed) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no d_message_embeds provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(dMessageEmbedColumnsWithDefault, o)

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

	dMessageEmbedUpsertCacheMut.RLock()
	cache, cached := dMessageEmbedUpsertCache[key]
	dMessageEmbedUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			dMessageEmbedColumns,
			dMessageEmbedColumnsWithDefault,
			dMessageEmbedColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			dMessageEmbedColumns,
			dMessageEmbedPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert d_message_embeds, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(dMessageEmbedPrimaryKeyColumns))
			copy(conflict, dMessageEmbedPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"d_message_embeds\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(dMessageEmbedType, dMessageEmbedMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(dMessageEmbedType, dMessageEmbedMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert d_message_embeds")
	}

	if !cached {
		dMessageEmbedUpsertCacheMut.Lock()
		dMessageEmbedUpsertCache[key] = cache
		dMessageEmbedUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteP deletes a single DMessageEmbed record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *DMessageEmbed) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single DMessageEmbed record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *DMessageEmbed) DeleteG() error {
	if o == nil {
		return errors.New("models: no DMessageEmbed provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single DMessageEmbed record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *DMessageEmbed) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single DMessageEmbed record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *DMessageEmbed) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no DMessageEmbed provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), dMessageEmbedPrimaryKeyMapping)
	sql := "DELETE FROM \"d_message_embeds\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from d_message_embeds")
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q dMessageEmbedQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q dMessageEmbedQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no dMessageEmbedQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from d_message_embeds")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o DMessageEmbedSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o DMessageEmbedSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no DMessageEmbed slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o DMessageEmbedSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o DMessageEmbedSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no DMessageEmbed slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dMessageEmbedPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"d_message_embeds\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, dMessageEmbedPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(dMessageEmbedPrimaryKeyColumns), 1, len(dMessageEmbedPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from dMessageEmbed slice")
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *DMessageEmbed) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *DMessageEmbed) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *DMessageEmbed) ReloadG() error {
	if o == nil {
		return errors.New("models: no DMessageEmbed provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *DMessageEmbed) Reload(exec boil.Executor) error {
	ret, err := FindDMessageEmbed(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *DMessageEmbedSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *DMessageEmbedSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DMessageEmbedSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty DMessageEmbedSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DMessageEmbedSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	dMessageEmbeds := DMessageEmbedSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dMessageEmbedPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"d_message_embeds\".* FROM \"d_message_embeds\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, dMessageEmbedPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(dMessageEmbedPrimaryKeyColumns), 1, len(dMessageEmbedPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&dMessageEmbeds)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in DMessageEmbedSlice")
	}

	*o = dMessageEmbeds

	return nil
}

// DMessageEmbedExists checks if the DMessageEmbed row exists.
func DMessageEmbedExists(exec boil.Executor, id int64) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"d_message_embeds\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, id)
	}

	row := exec.QueryRow(sql, id)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if d_message_embeds exists")
	}

	return exists, nil
}

// DMessageEmbedExistsG checks if the DMessageEmbed row exists.
func DMessageEmbedExistsG(id int64) (bool, error) {
	return DMessageEmbedExists(boil.GetDB(), id)
}

// DMessageEmbedExistsGP checks if the DMessageEmbed row exists. Panics on error.
func DMessageEmbedExistsGP(id int64) bool {
	e, err := DMessageEmbedExists(boil.GetDB(), id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// DMessageEmbedExistsP checks if the DMessageEmbed row exists. Panics on error.
func DMessageEmbedExistsP(exec boil.Executor, id int64) bool {
	e, err := DMessageEmbedExists(exec, id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
