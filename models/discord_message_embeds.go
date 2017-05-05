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

// DiscordMessageEmbed is an object representing the database table.
type DiscordMessageEmbed struct {
	ID                 int64             `boil:"id" json:"id" toml:"id" yaml:"id"`
	MessageID          null.Int64        `boil:"message_id" json:"message_id,omitempty" toml:"message_id" yaml:"message_id,omitempty"`
	RevisionID         null.Int64        `boil:"revision_id" json:"revision_id,omitempty" toml:"revision_id" yaml:"revision_id,omitempty"`
	URL                string            `boil:"url" json:"url" toml:"url" yaml:"url"`
	Type               string            `boil:"type" json:"type" toml:"type" yaml:"type"`
	Title              string            `boil:"title" json:"title" toml:"title" yaml:"title"`
	Description        string            `boil:"description" json:"description" toml:"description" yaml:"description"`
	Timestamp          string            `boil:"timestamp" json:"timestamp" toml:"timestamp" yaml:"timestamp"`
	Color              int               `boil:"color" json:"color" toml:"color" yaml:"color"`
	FieldNames         types.StringArray `boil:"field_names" json:"field_names,omitempty" toml:"field_names" yaml:"field_names,omitempty"`
	FieldValues        types.StringArray `boil:"field_values" json:"field_values,omitempty" toml:"field_values" yaml:"field_values,omitempty"`
	FieldInlines       types.BoolArray   `boil:"field_inlines" json:"field_inlines,omitempty" toml:"field_inlines" yaml:"field_inlines,omitempty"`
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

	R *discordMessageEmbedR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L discordMessageEmbedL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// discordMessageEmbedR is where relationships are stored.
type discordMessageEmbedR struct {
	Message  *DiscordMessage
	Revision *DiscordMessageRevision
}

// discordMessageEmbedL is where Load methods for each relationship are stored.
type discordMessageEmbedL struct{}

var (
	discordMessageEmbedColumns               = []string{"id", "message_id", "revision_id", "url", "type", "title", "description", "timestamp", "color", "field_names", "field_values", "field_inlines", "footer_text", "footer_icon_url", "footer_proxy_icon_url", "image_url", "image_proxy_url", "image_width", "image_height", "thumbnail_url", "thumbnail_proxy_url", "thumbnail_width", "thumbnail_height", "video_url", "video_proxy_url", "video_width", "video_height", "provider_url", "provider_name", "author_url", "author_name", "author_icon_url", "author_proxy_icon_url"}
	discordMessageEmbedColumnsWithoutDefault = []string{"message_id", "revision_id", "url", "type", "title", "description", "timestamp", "color", "field_names", "field_values", "field_inlines", "footer_text", "footer_icon_url", "footer_proxy_icon_url", "image_url", "image_proxy_url", "image_width", "image_height", "thumbnail_url", "thumbnail_proxy_url", "thumbnail_width", "thumbnail_height", "video_url", "video_proxy_url", "video_width", "video_height", "provider_url", "provider_name", "author_url", "author_name", "author_icon_url", "author_proxy_icon_url"}
	discordMessageEmbedColumnsWithDefault    = []string{"id"}
	discordMessageEmbedPrimaryKeyColumns     = []string{"id"}
)

type (
	// DiscordMessageEmbedSlice is an alias for a slice of pointers to DiscordMessageEmbed.
	// This should generally be used opposed to []DiscordMessageEmbed.
	DiscordMessageEmbedSlice []*DiscordMessageEmbed

	discordMessageEmbedQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	discordMessageEmbedType                 = reflect.TypeOf(&DiscordMessageEmbed{})
	discordMessageEmbedMapping              = queries.MakeStructMapping(discordMessageEmbedType)
	discordMessageEmbedPrimaryKeyMapping, _ = queries.BindMapping(discordMessageEmbedType, discordMessageEmbedMapping, discordMessageEmbedPrimaryKeyColumns)
	discordMessageEmbedInsertCacheMut       sync.RWMutex
	discordMessageEmbedInsertCache          = make(map[string]insertCache)
	discordMessageEmbedUpdateCacheMut       sync.RWMutex
	discordMessageEmbedUpdateCache          = make(map[string]updateCache)
	discordMessageEmbedUpsertCacheMut       sync.RWMutex
	discordMessageEmbedUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)

// OneP returns a single discordMessageEmbed record from the query, and panics on error.
func (q discordMessageEmbedQuery) OneP() *DiscordMessageEmbed {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single discordMessageEmbed record from the query.
func (q discordMessageEmbedQuery) One() (*DiscordMessageEmbed, error) {
	o := &DiscordMessageEmbed{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for discord_message_embeds")
	}

	return o, nil
}

// AllP returns all DiscordMessageEmbed records from the query, and panics on error.
func (q discordMessageEmbedQuery) AllP() DiscordMessageEmbedSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all DiscordMessageEmbed records from the query.
func (q discordMessageEmbedQuery) All() (DiscordMessageEmbedSlice, error) {
	var o DiscordMessageEmbedSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to DiscordMessageEmbed slice")
	}

	return o, nil
}

// CountP returns the count of all DiscordMessageEmbed records in the query, and panics on error.
func (q discordMessageEmbedQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all DiscordMessageEmbed records in the query.
func (q discordMessageEmbedQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count discord_message_embeds rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q discordMessageEmbedQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q discordMessageEmbedQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if discord_message_embeds exists")
	}

	return count > 0, nil
}

// MessageG pointed to by the foreign key.
func (o *DiscordMessageEmbed) MessageG(mods ...qm.QueryMod) discordMessageQuery {
	return o.Message(boil.GetDB(), mods...)
}

// Message pointed to by the foreign key.
func (o *DiscordMessageEmbed) Message(exec boil.Executor, mods ...qm.QueryMod) discordMessageQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.MessageID),
	}

	queryMods = append(queryMods, mods...)

	query := DiscordMessages(exec, queryMods...)
	queries.SetFrom(query.Query, "\"discord_messages\"")

	return query
}

// RevisionG pointed to by the foreign key.
func (o *DiscordMessageEmbed) RevisionG(mods ...qm.QueryMod) discordMessageRevisionQuery {
	return o.Revision(boil.GetDB(), mods...)
}

// Revision pointed to by the foreign key.
func (o *DiscordMessageEmbed) Revision(exec boil.Executor, mods ...qm.QueryMod) discordMessageRevisionQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.RevisionID),
	}

	queryMods = append(queryMods, mods...)

	query := DiscordMessageRevisions(exec, queryMods...)
	queries.SetFrom(query.Query, "\"discord_message_revisions\"")

	return query
}

// LoadMessage allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (discordMessageEmbedL) LoadMessage(e boil.Executor, singular bool, maybeDiscordMessageEmbed interface{}) error {
	var slice []*DiscordMessageEmbed
	var object *DiscordMessageEmbed

	count := 1
	if singular {
		object = maybeDiscordMessageEmbed.(*DiscordMessageEmbed)
	} else {
		slice = *maybeDiscordMessageEmbed.(*DiscordMessageEmbedSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		if object.R == nil {
			object.R = &discordMessageEmbedR{}
		}
		args[0] = object.MessageID
	} else {
		for i, obj := range slice {
			if obj.R == nil {
				obj.R = &discordMessageEmbedR{}
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
			if local.MessageID.Int64 == foreign.ID {
				local.R.Message = foreign
				break
			}
		}
	}

	return nil
}

// LoadRevision allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (discordMessageEmbedL) LoadRevision(e boil.Executor, singular bool, maybeDiscordMessageEmbed interface{}) error {
	var slice []*DiscordMessageEmbed
	var object *DiscordMessageEmbed

	count := 1
	if singular {
		object = maybeDiscordMessageEmbed.(*DiscordMessageEmbed)
	} else {
		slice = *maybeDiscordMessageEmbed.(*DiscordMessageEmbedSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		if object.R == nil {
			object.R = &discordMessageEmbedR{}
		}
		args[0] = object.RevisionID
	} else {
		for i, obj := range slice {
			if obj.R == nil {
				obj.R = &discordMessageEmbedR{}
			}
			args[i] = obj.RevisionID
		}
	}

	query := fmt.Sprintf(
		"select * from \"discord_message_revisions\" where \"id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load DiscordMessageRevision")
	}
	defer results.Close()

	var resultSlice []*DiscordMessageRevision
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice DiscordMessageRevision")
	}

	if singular && len(resultSlice) != 0 {
		object.R.Revision = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.RevisionID.Int64 == foreign.ID {
				local.R.Revision = foreign
				break
			}
		}
	}

	return nil
}

// SetMessageG of the discord_message_embed to the related item.
// Sets o.R.Message to related.
// Adds o to related.R.MessageDiscordMessageEmbeds.
// Uses the global database handle.
func (o *DiscordMessageEmbed) SetMessageG(insert bool, related *DiscordMessage) error {
	return o.SetMessage(boil.GetDB(), insert, related)
}

// SetMessageP of the discord_message_embed to the related item.
// Sets o.R.Message to related.
// Adds o to related.R.MessageDiscordMessageEmbeds.
// Panics on error.
func (o *DiscordMessageEmbed) SetMessageP(exec boil.Executor, insert bool, related *DiscordMessage) {
	if err := o.SetMessage(exec, insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetMessageGP of the discord_message_embed to the related item.
// Sets o.R.Message to related.
// Adds o to related.R.MessageDiscordMessageEmbeds.
// Uses the global database handle and panics on error.
func (o *DiscordMessageEmbed) SetMessageGP(insert bool, related *DiscordMessage) {
	if err := o.SetMessage(boil.GetDB(), insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetMessage of the discord_message_embed to the related item.
// Sets o.R.Message to related.
// Adds o to related.R.MessageDiscordMessageEmbeds.
func (o *DiscordMessageEmbed) SetMessage(exec boil.Executor, insert bool, related *DiscordMessage) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"discord_message_embeds\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"message_id"}),
		strmangle.WhereClause("\"", "\"", 2, discordMessageEmbedPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.MessageID.Int64 = related.ID
	o.MessageID.Valid = true

	if o.R == nil {
		o.R = &discordMessageEmbedR{
			Message: related,
		}
	} else {
		o.R.Message = related
	}

	if related.R == nil {
		related.R = &discordMessageR{
			MessageDiscordMessageEmbeds: DiscordMessageEmbedSlice{o},
		}
	} else {
		related.R.MessageDiscordMessageEmbeds = append(related.R.MessageDiscordMessageEmbeds, o)
	}

	return nil
}

// RemoveMessageG relationship.
// Sets o.R.Message to nil.
// Removes o from all passed in related items' relationships struct (Optional).
// Uses the global database handle.
func (o *DiscordMessageEmbed) RemoveMessageG(related *DiscordMessage) error {
	return o.RemoveMessage(boil.GetDB(), related)
}

// RemoveMessageP relationship.
// Sets o.R.Message to nil.
// Removes o from all passed in related items' relationships struct (Optional).
// Panics on error.
func (o *DiscordMessageEmbed) RemoveMessageP(exec boil.Executor, related *DiscordMessage) {
	if err := o.RemoveMessage(exec, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// RemoveMessageGP relationship.
// Sets o.R.Message to nil.
// Removes o from all passed in related items' relationships struct (Optional).
// Uses the global database handle and panics on error.
func (o *DiscordMessageEmbed) RemoveMessageGP(related *DiscordMessage) {
	if err := o.RemoveMessage(boil.GetDB(), related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// RemoveMessage relationship.
// Sets o.R.Message to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *DiscordMessageEmbed) RemoveMessage(exec boil.Executor, related *DiscordMessage) error {
	var err error

	o.MessageID.Valid = false
	if err = o.Update(exec, "message_id"); err != nil {
		o.MessageID.Valid = true
		return errors.Wrap(err, "failed to update local table")
	}

	o.R.Message = nil
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.MessageDiscordMessageEmbeds {
		if o.MessageID.Int64 != ri.MessageID.Int64 {
			continue
		}

		ln := len(related.R.MessageDiscordMessageEmbeds)
		if ln > 1 && i < ln-1 {
			related.R.MessageDiscordMessageEmbeds[i] = related.R.MessageDiscordMessageEmbeds[ln-1]
		}
		related.R.MessageDiscordMessageEmbeds = related.R.MessageDiscordMessageEmbeds[:ln-1]
		break
	}
	return nil
}

// SetRevisionG of the discord_message_embed to the related item.
// Sets o.R.Revision to related.
// Adds o to related.R.RevisionDiscordMessageEmbeds.
// Uses the global database handle.
func (o *DiscordMessageEmbed) SetRevisionG(insert bool, related *DiscordMessageRevision) error {
	return o.SetRevision(boil.GetDB(), insert, related)
}

// SetRevisionP of the discord_message_embed to the related item.
// Sets o.R.Revision to related.
// Adds o to related.R.RevisionDiscordMessageEmbeds.
// Panics on error.
func (o *DiscordMessageEmbed) SetRevisionP(exec boil.Executor, insert bool, related *DiscordMessageRevision) {
	if err := o.SetRevision(exec, insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetRevisionGP of the discord_message_embed to the related item.
// Sets o.R.Revision to related.
// Adds o to related.R.RevisionDiscordMessageEmbeds.
// Uses the global database handle and panics on error.
func (o *DiscordMessageEmbed) SetRevisionGP(insert bool, related *DiscordMessageRevision) {
	if err := o.SetRevision(boil.GetDB(), insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetRevision of the discord_message_embed to the related item.
// Sets o.R.Revision to related.
// Adds o to related.R.RevisionDiscordMessageEmbeds.
func (o *DiscordMessageEmbed) SetRevision(exec boil.Executor, insert bool, related *DiscordMessageRevision) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"discord_message_embeds\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"revision_id"}),
		strmangle.WhereClause("\"", "\"", 2, discordMessageEmbedPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.RevisionID.Int64 = related.ID
	o.RevisionID.Valid = true

	if o.R == nil {
		o.R = &discordMessageEmbedR{
			Revision: related,
		}
	} else {
		o.R.Revision = related
	}

	if related.R == nil {
		related.R = &discordMessageRevisionR{
			RevisionDiscordMessageEmbeds: DiscordMessageEmbedSlice{o},
		}
	} else {
		related.R.RevisionDiscordMessageEmbeds = append(related.R.RevisionDiscordMessageEmbeds, o)
	}

	return nil
}

// RemoveRevisionG relationship.
// Sets o.R.Revision to nil.
// Removes o from all passed in related items' relationships struct (Optional).
// Uses the global database handle.
func (o *DiscordMessageEmbed) RemoveRevisionG(related *DiscordMessageRevision) error {
	return o.RemoveRevision(boil.GetDB(), related)
}

// RemoveRevisionP relationship.
// Sets o.R.Revision to nil.
// Removes o from all passed in related items' relationships struct (Optional).
// Panics on error.
func (o *DiscordMessageEmbed) RemoveRevisionP(exec boil.Executor, related *DiscordMessageRevision) {
	if err := o.RemoveRevision(exec, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// RemoveRevisionGP relationship.
// Sets o.R.Revision to nil.
// Removes o from all passed in related items' relationships struct (Optional).
// Uses the global database handle and panics on error.
func (o *DiscordMessageEmbed) RemoveRevisionGP(related *DiscordMessageRevision) {
	if err := o.RemoveRevision(boil.GetDB(), related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// RemoveRevision relationship.
// Sets o.R.Revision to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *DiscordMessageEmbed) RemoveRevision(exec boil.Executor, related *DiscordMessageRevision) error {
	var err error

	o.RevisionID.Valid = false
	if err = o.Update(exec, "revision_id"); err != nil {
		o.RevisionID.Valid = true
		return errors.Wrap(err, "failed to update local table")
	}

	o.R.Revision = nil
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.RevisionDiscordMessageEmbeds {
		if o.RevisionID.Int64 != ri.RevisionID.Int64 {
			continue
		}

		ln := len(related.R.RevisionDiscordMessageEmbeds)
		if ln > 1 && i < ln-1 {
			related.R.RevisionDiscordMessageEmbeds[i] = related.R.RevisionDiscordMessageEmbeds[ln-1]
		}
		related.R.RevisionDiscordMessageEmbeds = related.R.RevisionDiscordMessageEmbeds[:ln-1]
		break
	}
	return nil
}

// DiscordMessageEmbedsG retrieves all records.
func DiscordMessageEmbedsG(mods ...qm.QueryMod) discordMessageEmbedQuery {
	return DiscordMessageEmbeds(boil.GetDB(), mods...)
}

// DiscordMessageEmbeds retrieves all the records using an executor.
func DiscordMessageEmbeds(exec boil.Executor, mods ...qm.QueryMod) discordMessageEmbedQuery {
	mods = append(mods, qm.From("\"discord_message_embeds\""))
	return discordMessageEmbedQuery{NewQuery(exec, mods...)}
}

// FindDiscordMessageEmbedG retrieves a single record by ID.
func FindDiscordMessageEmbedG(id int64, selectCols ...string) (*DiscordMessageEmbed, error) {
	return FindDiscordMessageEmbed(boil.GetDB(), id, selectCols...)
}

// FindDiscordMessageEmbedGP retrieves a single record by ID, and panics on error.
func FindDiscordMessageEmbedGP(id int64, selectCols ...string) *DiscordMessageEmbed {
	retobj, err := FindDiscordMessageEmbed(boil.GetDB(), id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindDiscordMessageEmbed retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDiscordMessageEmbed(exec boil.Executor, id int64, selectCols ...string) (*DiscordMessageEmbed, error) {
	discordMessageEmbedObj := &DiscordMessageEmbed{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"discord_message_embeds\" where \"id\"=$1", sel,
	)

	q := queries.Raw(exec, query, id)

	err := q.Bind(discordMessageEmbedObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from discord_message_embeds")
	}

	return discordMessageEmbedObj, nil
}

// FindDiscordMessageEmbedP retrieves a single record by ID with an executor, and panics on error.
func FindDiscordMessageEmbedP(exec boil.Executor, id int64, selectCols ...string) *DiscordMessageEmbed {
	retobj, err := FindDiscordMessageEmbed(exec, id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *DiscordMessageEmbed) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *DiscordMessageEmbed) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *DiscordMessageEmbed) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *DiscordMessageEmbed) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no discord_message_embeds provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(discordMessageEmbedColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	discordMessageEmbedInsertCacheMut.RLock()
	cache, cached := discordMessageEmbedInsertCache[key]
	discordMessageEmbedInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			discordMessageEmbedColumns,
			discordMessageEmbedColumnsWithDefault,
			discordMessageEmbedColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(discordMessageEmbedType, discordMessageEmbedMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(discordMessageEmbedType, discordMessageEmbedMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"discord_message_embeds\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "models: unable to insert into discord_message_embeds")
	}

	if !cached {
		discordMessageEmbedInsertCacheMut.Lock()
		discordMessageEmbedInsertCache[key] = cache
		discordMessageEmbedInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single DiscordMessageEmbed record. See Update for
// whitelist behavior description.
func (o *DiscordMessageEmbed) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single DiscordMessageEmbed record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *DiscordMessageEmbed) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the DiscordMessageEmbed, and panics on error.
// See Update for whitelist behavior description.
func (o *DiscordMessageEmbed) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the DiscordMessageEmbed.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *DiscordMessageEmbed) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	key := makeCacheKey(whitelist, nil)
	discordMessageEmbedUpdateCacheMut.RLock()
	cache, cached := discordMessageEmbedUpdateCache[key]
	discordMessageEmbedUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(discordMessageEmbedColumns, discordMessageEmbedPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update discord_message_embeds, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"discord_message_embeds\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, discordMessageEmbedPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(discordMessageEmbedType, discordMessageEmbedMapping, append(wl, discordMessageEmbedPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update discord_message_embeds row")
	}

	if !cached {
		discordMessageEmbedUpdateCacheMut.Lock()
		discordMessageEmbedUpdateCache[key] = cache
		discordMessageEmbedUpdateCacheMut.Unlock()
	}

	return nil
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q discordMessageEmbedQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q discordMessageEmbedQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for discord_message_embeds")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o DiscordMessageEmbedSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o DiscordMessageEmbedSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o DiscordMessageEmbedSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o DiscordMessageEmbedSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), discordMessageEmbedPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"discord_message_embeds\" SET %s WHERE (\"id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(discordMessageEmbedPrimaryKeyColumns), len(colNames)+1, len(discordMessageEmbedPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in discordMessageEmbed slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *DiscordMessageEmbed) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *DiscordMessageEmbed) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *DiscordMessageEmbed) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *DiscordMessageEmbed) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no discord_message_embeds provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(discordMessageEmbedColumnsWithDefault, o)

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

	discordMessageEmbedUpsertCacheMut.RLock()
	cache, cached := discordMessageEmbedUpsertCache[key]
	discordMessageEmbedUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			discordMessageEmbedColumns,
			discordMessageEmbedColumnsWithDefault,
			discordMessageEmbedColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			discordMessageEmbedColumns,
			discordMessageEmbedPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert discord_message_embeds, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(discordMessageEmbedPrimaryKeyColumns))
			copy(conflict, discordMessageEmbedPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"discord_message_embeds\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(discordMessageEmbedType, discordMessageEmbedMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(discordMessageEmbedType, discordMessageEmbedMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert discord_message_embeds")
	}

	if !cached {
		discordMessageEmbedUpsertCacheMut.Lock()
		discordMessageEmbedUpsertCache[key] = cache
		discordMessageEmbedUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteP deletes a single DiscordMessageEmbed record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *DiscordMessageEmbed) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single DiscordMessageEmbed record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *DiscordMessageEmbed) DeleteG() error {
	if o == nil {
		return errors.New("models: no DiscordMessageEmbed provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single DiscordMessageEmbed record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *DiscordMessageEmbed) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single DiscordMessageEmbed record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *DiscordMessageEmbed) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no DiscordMessageEmbed provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), discordMessageEmbedPrimaryKeyMapping)
	sql := "DELETE FROM \"discord_message_embeds\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from discord_message_embeds")
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q discordMessageEmbedQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q discordMessageEmbedQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no discordMessageEmbedQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from discord_message_embeds")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o DiscordMessageEmbedSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o DiscordMessageEmbedSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no DiscordMessageEmbed slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o DiscordMessageEmbedSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o DiscordMessageEmbedSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no DiscordMessageEmbed slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), discordMessageEmbedPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"discord_message_embeds\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, discordMessageEmbedPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(discordMessageEmbedPrimaryKeyColumns), 1, len(discordMessageEmbedPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from discordMessageEmbed slice")
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *DiscordMessageEmbed) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *DiscordMessageEmbed) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *DiscordMessageEmbed) ReloadG() error {
	if o == nil {
		return errors.New("models: no DiscordMessageEmbed provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *DiscordMessageEmbed) Reload(exec boil.Executor) error {
	ret, err := FindDiscordMessageEmbed(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *DiscordMessageEmbedSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *DiscordMessageEmbedSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DiscordMessageEmbedSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty DiscordMessageEmbedSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DiscordMessageEmbedSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	discordMessageEmbeds := DiscordMessageEmbedSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), discordMessageEmbedPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"discord_message_embeds\".* FROM \"discord_message_embeds\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, discordMessageEmbedPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(discordMessageEmbedPrimaryKeyColumns), 1, len(discordMessageEmbedPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&discordMessageEmbeds)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in DiscordMessageEmbedSlice")
	}

	*o = discordMessageEmbeds

	return nil
}

// DiscordMessageEmbedExists checks if the DiscordMessageEmbed row exists.
func DiscordMessageEmbedExists(exec boil.Executor, id int64) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"discord_message_embeds\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, id)
	}

	row := exec.QueryRow(sql, id)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if discord_message_embeds exists")
	}

	return exists, nil
}

// DiscordMessageEmbedExistsG checks if the DiscordMessageEmbed row exists.
func DiscordMessageEmbedExistsG(id int64) (bool, error) {
	return DiscordMessageEmbedExists(boil.GetDB(), id)
}

// DiscordMessageEmbedExistsGP checks if the DiscordMessageEmbed row exists. Panics on error.
func DiscordMessageEmbedExistsGP(id int64) bool {
	e, err := DiscordMessageEmbedExists(boil.GetDB(), id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// DiscordMessageEmbedExistsP checks if the DiscordMessageEmbed row exists. Panics on error.
func DiscordMessageEmbedExistsP(exec boil.Executor, id int64) bool {
	e, err := DiscordMessageEmbedExists(exec, id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
