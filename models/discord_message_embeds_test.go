package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testDiscordMessageEmbeds(t *testing.T) {
	t.Parallel()

	query := DiscordMessageEmbeds(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testDiscordMessageEmbedsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordMessageEmbed := &DiscordMessageEmbed{}
	if err = randomize.Struct(seed, discordMessageEmbed, discordMessageEmbedDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordMessageEmbed struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordMessageEmbed.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = discordMessageEmbed.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := DiscordMessageEmbeds(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDiscordMessageEmbedsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordMessageEmbed := &DiscordMessageEmbed{}
	if err = randomize.Struct(seed, discordMessageEmbed, discordMessageEmbedDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordMessageEmbed struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordMessageEmbed.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = DiscordMessageEmbeds(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := DiscordMessageEmbeds(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDiscordMessageEmbedsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordMessageEmbed := &DiscordMessageEmbed{}
	if err = randomize.Struct(seed, discordMessageEmbed, discordMessageEmbedDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordMessageEmbed struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordMessageEmbed.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := DiscordMessageEmbedSlice{discordMessageEmbed}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := DiscordMessageEmbeds(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testDiscordMessageEmbedsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordMessageEmbed := &DiscordMessageEmbed{}
	if err = randomize.Struct(seed, discordMessageEmbed, discordMessageEmbedDBTypes, true, discordMessageEmbedColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordMessageEmbed struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordMessageEmbed.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := DiscordMessageEmbedExists(tx, discordMessageEmbed.ID)
	if err != nil {
		t.Errorf("Unable to check if DiscordMessageEmbed exists: %s", err)
	}
	if !e {
		t.Errorf("Expected DiscordMessageEmbedExistsG to return true, but got false.")
	}
}
func testDiscordMessageEmbedsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordMessageEmbed := &DiscordMessageEmbed{}
	if err = randomize.Struct(seed, discordMessageEmbed, discordMessageEmbedDBTypes, true, discordMessageEmbedColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordMessageEmbed struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordMessageEmbed.Insert(tx); err != nil {
		t.Error(err)
	}

	discordMessageEmbedFound, err := FindDiscordMessageEmbed(tx, discordMessageEmbed.ID)
	if err != nil {
		t.Error(err)
	}

	if discordMessageEmbedFound == nil {
		t.Error("want a record, got nil")
	}
}
func testDiscordMessageEmbedsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordMessageEmbed := &DiscordMessageEmbed{}
	if err = randomize.Struct(seed, discordMessageEmbed, discordMessageEmbedDBTypes, true, discordMessageEmbedColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordMessageEmbed struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordMessageEmbed.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = DiscordMessageEmbeds(tx).Bind(discordMessageEmbed); err != nil {
		t.Error(err)
	}
}

func testDiscordMessageEmbedsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordMessageEmbed := &DiscordMessageEmbed{}
	if err = randomize.Struct(seed, discordMessageEmbed, discordMessageEmbedDBTypes, true, discordMessageEmbedColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordMessageEmbed struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordMessageEmbed.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := DiscordMessageEmbeds(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testDiscordMessageEmbedsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordMessageEmbedOne := &DiscordMessageEmbed{}
	discordMessageEmbedTwo := &DiscordMessageEmbed{}
	if err = randomize.Struct(seed, discordMessageEmbedOne, discordMessageEmbedDBTypes, false, discordMessageEmbedColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordMessageEmbed struct: %s", err)
	}
	if err = randomize.Struct(seed, discordMessageEmbedTwo, discordMessageEmbedDBTypes, false, discordMessageEmbedColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordMessageEmbed struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordMessageEmbedOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = discordMessageEmbedTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := DiscordMessageEmbeds(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testDiscordMessageEmbedsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	discordMessageEmbedOne := &DiscordMessageEmbed{}
	discordMessageEmbedTwo := &DiscordMessageEmbed{}
	if err = randomize.Struct(seed, discordMessageEmbedOne, discordMessageEmbedDBTypes, false, discordMessageEmbedColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordMessageEmbed struct: %s", err)
	}
	if err = randomize.Struct(seed, discordMessageEmbedTwo, discordMessageEmbedDBTypes, false, discordMessageEmbedColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordMessageEmbed struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordMessageEmbedOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = discordMessageEmbedTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DiscordMessageEmbeds(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testDiscordMessageEmbedsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordMessageEmbed := &DiscordMessageEmbed{}
	if err = randomize.Struct(seed, discordMessageEmbed, discordMessageEmbedDBTypes, true, discordMessageEmbedColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordMessageEmbed struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordMessageEmbed.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DiscordMessageEmbeds(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDiscordMessageEmbedsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordMessageEmbed := &DiscordMessageEmbed{}
	if err = randomize.Struct(seed, discordMessageEmbed, discordMessageEmbedDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordMessageEmbed struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordMessageEmbed.Insert(tx, discordMessageEmbedColumns...); err != nil {
		t.Error(err)
	}

	count, err := DiscordMessageEmbeds(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDiscordMessageEmbedToOneDiscordMessageUsingMessage(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local DiscordMessageEmbed
	var foreign DiscordMessage

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, discordMessageEmbedDBTypes, true, discordMessageEmbedColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordMessageEmbed struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, discordMessageDBTypes, true, discordMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordMessage struct: %s", err)
	}

	local.MessageID.Valid = true

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.MessageID.Int64 = foreign.ID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Message(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := DiscordMessageEmbedSlice{&local}
	if err = local.L.LoadMessage(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Message == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Message = nil
	if err = local.L.LoadMessage(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Message == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testDiscordMessageEmbedToOneDiscordMessageRevisionUsingRevision(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local DiscordMessageEmbed
	var foreign DiscordMessageRevision

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, discordMessageEmbedDBTypes, true, discordMessageEmbedColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordMessageEmbed struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, discordMessageRevisionDBTypes, true, discordMessageRevisionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordMessageRevision struct: %s", err)
	}

	local.RevisionID.Valid = true

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.RevisionID.Int64 = foreign.ID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Revision(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := DiscordMessageEmbedSlice{&local}
	if err = local.L.LoadRevision(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Revision == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Revision = nil
	if err = local.L.LoadRevision(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Revision == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testDiscordMessageEmbedToOneSetOpDiscordMessageUsingMessage(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a DiscordMessageEmbed
	var b, c DiscordMessage

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, discordMessageEmbedDBTypes, false, strmangle.SetComplement(discordMessageEmbedPrimaryKeyColumns, discordMessageEmbedColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, discordMessageDBTypes, false, strmangle.SetComplement(discordMessagePrimaryKeyColumns, discordMessageColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, discordMessageDBTypes, false, strmangle.SetComplement(discordMessagePrimaryKeyColumns, discordMessageColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*DiscordMessage{&b, &c} {
		err = a.SetMessage(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Message != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.MessageDiscordMessageEmbeds[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.MessageID.Int64 != x.ID {
			t.Error("foreign key was wrong value", a.MessageID.Int64)
		}

		zero := reflect.Zero(reflect.TypeOf(a.MessageID.Int64))
		reflect.Indirect(reflect.ValueOf(&a.MessageID.Int64)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.MessageID.Int64 != x.ID {
			t.Error("foreign key was wrong value", a.MessageID.Int64, x.ID)
		}
	}
}

func testDiscordMessageEmbedToOneRemoveOpDiscordMessageUsingMessage(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a DiscordMessageEmbed
	var b DiscordMessage

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, discordMessageEmbedDBTypes, false, strmangle.SetComplement(discordMessageEmbedPrimaryKeyColumns, discordMessageEmbedColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, discordMessageDBTypes, false, strmangle.SetComplement(discordMessagePrimaryKeyColumns, discordMessageColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	if err = a.SetMessage(tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveMessage(tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.Message(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.Message != nil {
		t.Error("R struct entry should be nil")
	}

	if a.MessageID.Valid {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.MessageDiscordMessageEmbeds) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testDiscordMessageEmbedToOneSetOpDiscordMessageRevisionUsingRevision(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a DiscordMessageEmbed
	var b, c DiscordMessageRevision

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, discordMessageEmbedDBTypes, false, strmangle.SetComplement(discordMessageEmbedPrimaryKeyColumns, discordMessageEmbedColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, discordMessageRevisionDBTypes, false, strmangle.SetComplement(discordMessageRevisionPrimaryKeyColumns, discordMessageRevisionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, discordMessageRevisionDBTypes, false, strmangle.SetComplement(discordMessageRevisionPrimaryKeyColumns, discordMessageRevisionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*DiscordMessageRevision{&b, &c} {
		err = a.SetRevision(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Revision != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.RevisionDiscordMessageEmbeds[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.RevisionID.Int64 != x.ID {
			t.Error("foreign key was wrong value", a.RevisionID.Int64)
		}

		zero := reflect.Zero(reflect.TypeOf(a.RevisionID.Int64))
		reflect.Indirect(reflect.ValueOf(&a.RevisionID.Int64)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.RevisionID.Int64 != x.ID {
			t.Error("foreign key was wrong value", a.RevisionID.Int64, x.ID)
		}
	}
}

func testDiscordMessageEmbedToOneRemoveOpDiscordMessageRevisionUsingRevision(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a DiscordMessageEmbed
	var b DiscordMessageRevision

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, discordMessageEmbedDBTypes, false, strmangle.SetComplement(discordMessageEmbedPrimaryKeyColumns, discordMessageEmbedColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, discordMessageRevisionDBTypes, false, strmangle.SetComplement(discordMessageRevisionPrimaryKeyColumns, discordMessageRevisionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	if err = a.SetRevision(tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveRevision(tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.Revision(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.Revision != nil {
		t.Error("R struct entry should be nil")
	}

	if a.RevisionID.Valid {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.RevisionDiscordMessageEmbeds) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testDiscordMessageEmbedsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordMessageEmbed := &DiscordMessageEmbed{}
	if err = randomize.Struct(seed, discordMessageEmbed, discordMessageEmbedDBTypes, true, discordMessageEmbedColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordMessageEmbed struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordMessageEmbed.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = discordMessageEmbed.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testDiscordMessageEmbedsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordMessageEmbed := &DiscordMessageEmbed{}
	if err = randomize.Struct(seed, discordMessageEmbed, discordMessageEmbedDBTypes, true, discordMessageEmbedColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordMessageEmbed struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordMessageEmbed.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := DiscordMessageEmbedSlice{discordMessageEmbed}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testDiscordMessageEmbedsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordMessageEmbed := &DiscordMessageEmbed{}
	if err = randomize.Struct(seed, discordMessageEmbed, discordMessageEmbedDBTypes, true, discordMessageEmbedColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordMessageEmbed struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordMessageEmbed.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := DiscordMessageEmbeds(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	discordMessageEmbedDBTypes = map[string]string{`AuthorIconURL`: `text`, `AuthorName`: `text`, `AuthorProxyIconURL`: `text`, `AuthorURL`: `text`, `Color`: `integer`, `Description`: `text`, `FieldInlines`: `ARRAYboolean`, `FieldNames`: `ARRAYtext`, `FieldValues`: `ARRAYtext`, `FooterIconURL`: `text`, `FooterProxyIconURL`: `text`, `FooterText`: `text`, `ID`: `bigint`, `ImageHeight`: `integer`, `ImageProxyURL`: `text`, `ImageURL`: `text`, `ImageWidth`: `integer`, `MessageID`: `bigint`, `ProviderName`: `text`, `ProviderURL`: `text`, `RevisionID`: `bigint`, `ThumbnailHeight`: `integer`, `ThumbnailProxyURL`: `text`, `ThumbnailURL`: `text`, `ThumbnailWidth`: `integer`, `Timestamp`: `text`, `Title`: `text`, `Type`: `text`, `URL`: `text`, `VideoHeight`: `integer`, `VideoProxyURL`: `text`, `VideoURL`: `text`, `VideoWidth`: `integer`}
	_                          = bytes.MinRead
)

func testDiscordMessageEmbedsUpdate(t *testing.T) {
	t.Parallel()

	if len(discordMessageEmbedColumns) == len(discordMessageEmbedPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	discordMessageEmbed := &DiscordMessageEmbed{}
	if err = randomize.Struct(seed, discordMessageEmbed, discordMessageEmbedDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordMessageEmbed struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordMessageEmbed.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DiscordMessageEmbeds(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, discordMessageEmbed, discordMessageEmbedDBTypes, true, discordMessageEmbedColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordMessageEmbed struct: %s", err)
	}

	if err = discordMessageEmbed.Update(tx); err != nil {
		t.Error(err)
	}
}

func testDiscordMessageEmbedsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(discordMessageEmbedColumns) == len(discordMessageEmbedPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	discordMessageEmbed := &DiscordMessageEmbed{}
	if err = randomize.Struct(seed, discordMessageEmbed, discordMessageEmbedDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordMessageEmbed struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordMessageEmbed.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DiscordMessageEmbeds(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, discordMessageEmbed, discordMessageEmbedDBTypes, true, discordMessageEmbedPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DiscordMessageEmbed struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(discordMessageEmbedColumns, discordMessageEmbedPrimaryKeyColumns) {
		fields = discordMessageEmbedColumns
	} else {
		fields = strmangle.SetComplement(
			discordMessageEmbedColumns,
			discordMessageEmbedPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(discordMessageEmbed))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := DiscordMessageEmbedSlice{discordMessageEmbed}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testDiscordMessageEmbedsUpsert(t *testing.T) {
	t.Parallel()

	if len(discordMessageEmbedColumns) == len(discordMessageEmbedPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	discordMessageEmbed := DiscordMessageEmbed{}
	if err = randomize.Struct(seed, &discordMessageEmbed, discordMessageEmbedDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordMessageEmbed struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordMessageEmbed.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert DiscordMessageEmbed: %s", err)
	}

	count, err := DiscordMessageEmbeds(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &discordMessageEmbed, discordMessageEmbedDBTypes, false, discordMessageEmbedPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DiscordMessageEmbed struct: %s", err)
	}

	if err = discordMessageEmbed.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert DiscordMessageEmbed: %s", err)
	}

	count, err = DiscordMessageEmbeds(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
