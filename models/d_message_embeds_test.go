package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testDMessageEmbeds(t *testing.T) {
	t.Parallel()

	query := DMessageEmbeds(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testDMessageEmbedsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dMessageEmbed := &DMessageEmbed{}
	if err = randomize.Struct(seed, dMessageEmbed, dMessageEmbedDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DMessageEmbed struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dMessageEmbed.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = dMessageEmbed.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := DMessageEmbeds(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDMessageEmbedsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dMessageEmbed := &DMessageEmbed{}
	if err = randomize.Struct(seed, dMessageEmbed, dMessageEmbedDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DMessageEmbed struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dMessageEmbed.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = DMessageEmbeds(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := DMessageEmbeds(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDMessageEmbedsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dMessageEmbed := &DMessageEmbed{}
	if err = randomize.Struct(seed, dMessageEmbed, dMessageEmbedDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DMessageEmbed struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dMessageEmbed.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := DMessageEmbedSlice{dMessageEmbed}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := DMessageEmbeds(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testDMessageEmbedsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dMessageEmbed := &DMessageEmbed{}
	if err = randomize.Struct(seed, dMessageEmbed, dMessageEmbedDBTypes, true, dMessageEmbedColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DMessageEmbed struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dMessageEmbed.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := DMessageEmbedExists(tx, dMessageEmbed.ID)
	if err != nil {
		t.Errorf("Unable to check if DMessageEmbed exists: %s", err)
	}
	if !e {
		t.Errorf("Expected DMessageEmbedExistsG to return true, but got false.")
	}
}
func testDMessageEmbedsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dMessageEmbed := &DMessageEmbed{}
	if err = randomize.Struct(seed, dMessageEmbed, dMessageEmbedDBTypes, true, dMessageEmbedColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DMessageEmbed struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dMessageEmbed.Insert(tx); err != nil {
		t.Error(err)
	}

	dMessageEmbedFound, err := FindDMessageEmbed(tx, dMessageEmbed.ID)
	if err != nil {
		t.Error(err)
	}

	if dMessageEmbedFound == nil {
		t.Error("want a record, got nil")
	}
}
func testDMessageEmbedsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dMessageEmbed := &DMessageEmbed{}
	if err = randomize.Struct(seed, dMessageEmbed, dMessageEmbedDBTypes, true, dMessageEmbedColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DMessageEmbed struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dMessageEmbed.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = DMessageEmbeds(tx).Bind(dMessageEmbed); err != nil {
		t.Error(err)
	}
}

func testDMessageEmbedsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dMessageEmbed := &DMessageEmbed{}
	if err = randomize.Struct(seed, dMessageEmbed, dMessageEmbedDBTypes, true, dMessageEmbedColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DMessageEmbed struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dMessageEmbed.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := DMessageEmbeds(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testDMessageEmbedsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dMessageEmbedOne := &DMessageEmbed{}
	dMessageEmbedTwo := &DMessageEmbed{}
	if err = randomize.Struct(seed, dMessageEmbedOne, dMessageEmbedDBTypes, false, dMessageEmbedColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DMessageEmbed struct: %s", err)
	}
	if err = randomize.Struct(seed, dMessageEmbedTwo, dMessageEmbedDBTypes, false, dMessageEmbedColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DMessageEmbed struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dMessageEmbedOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = dMessageEmbedTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := DMessageEmbeds(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testDMessageEmbedsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	dMessageEmbedOne := &DMessageEmbed{}
	dMessageEmbedTwo := &DMessageEmbed{}
	if err = randomize.Struct(seed, dMessageEmbedOne, dMessageEmbedDBTypes, false, dMessageEmbedColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DMessageEmbed struct: %s", err)
	}
	if err = randomize.Struct(seed, dMessageEmbedTwo, dMessageEmbedDBTypes, false, dMessageEmbedColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DMessageEmbed struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dMessageEmbedOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = dMessageEmbedTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DMessageEmbeds(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testDMessageEmbedsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dMessageEmbed := &DMessageEmbed{}
	if err = randomize.Struct(seed, dMessageEmbed, dMessageEmbedDBTypes, true, dMessageEmbedColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DMessageEmbed struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dMessageEmbed.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DMessageEmbeds(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDMessageEmbedsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dMessageEmbed := &DMessageEmbed{}
	if err = randomize.Struct(seed, dMessageEmbed, dMessageEmbedDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DMessageEmbed struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dMessageEmbed.Insert(tx, dMessageEmbedColumns...); err != nil {
		t.Error(err)
	}

	count, err := DMessageEmbeds(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDMessageEmbedToOneDMessageUsingMessage(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local DMessageEmbed
	var foreign DMessage

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, dMessageEmbedDBTypes, true, dMessageEmbedColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DMessageEmbed struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, dMessageDBTypes, true, dMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DMessage struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.MessageID = foreign.ID
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

	slice := DMessageEmbedSlice{&local}
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

func testDMessageEmbedToOneSetOpDMessageUsingMessage(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a DMessageEmbed
	var b, c DMessage

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, dMessageEmbedDBTypes, false, strmangle.SetComplement(dMessageEmbedPrimaryKeyColumns, dMessageEmbedColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, dMessageDBTypes, false, strmangle.SetComplement(dMessagePrimaryKeyColumns, dMessageColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, dMessageDBTypes, false, strmangle.SetComplement(dMessagePrimaryKeyColumns, dMessageColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*DMessage{&b, &c} {
		err = a.SetMessage(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Message != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.MessageDMessageEmbeds[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.MessageID != x.ID {
			t.Error("foreign key was wrong value", a.MessageID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.MessageID))
		reflect.Indirect(reflect.ValueOf(&a.MessageID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.MessageID != x.ID {
			t.Error("foreign key was wrong value", a.MessageID, x.ID)
		}
	}
}
func testDMessageEmbedsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dMessageEmbed := &DMessageEmbed{}
	if err = randomize.Struct(seed, dMessageEmbed, dMessageEmbedDBTypes, true, dMessageEmbedColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DMessageEmbed struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dMessageEmbed.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = dMessageEmbed.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testDMessageEmbedsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dMessageEmbed := &DMessageEmbed{}
	if err = randomize.Struct(seed, dMessageEmbed, dMessageEmbedDBTypes, true, dMessageEmbedColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DMessageEmbed struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dMessageEmbed.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := DMessageEmbedSlice{dMessageEmbed}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testDMessageEmbedsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dMessageEmbed := &DMessageEmbed{}
	if err = randomize.Struct(seed, dMessageEmbed, dMessageEmbedDBTypes, true, dMessageEmbedColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DMessageEmbed struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dMessageEmbed.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := DMessageEmbeds(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	dMessageEmbedDBTypes = map[string]string{`AuthorIconURL`: `text`, `AuthorName`: `text`, `AuthorProxyIconURL`: `text`, `AuthorURL`: `text`, `Color`: `integer`, `Description`: `text`, `FieldInlines`: `ARRAYboolean`, `FieldNames`: `ARRAYtext`, `FieldValues`: `ARRAYtext`, `FooterIconURL`: `text`, `FooterProxyIconURL`: `text`, `FooterText`: `text`, `ID`: `bigint`, `ImageHeight`: `integer`, `ImageProxyURL`: `text`, `ImageURL`: `text`, `ImageWidth`: `integer`, `MessageID`: `bigint`, `ProviderName`: `text`, `ProviderURL`: `text`, `RevisionNum`: `integer`, `ThumbnailHeight`: `integer`, `ThumbnailProxyURL`: `text`, `ThumbnailURL`: `text`, `ThumbnailWidth`: `integer`, `Timestamp`: `text`, `Title`: `text`, `Type`: `text`, `URL`: `text`, `VideoHeight`: `integer`, `VideoProxyURL`: `text`, `VideoURL`: `text`, `VideoWidth`: `integer`}
	_                    = bytes.MinRead
)

func testDMessageEmbedsUpdate(t *testing.T) {
	t.Parallel()

	if len(dMessageEmbedColumns) == len(dMessageEmbedPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	dMessageEmbed := &DMessageEmbed{}
	if err = randomize.Struct(seed, dMessageEmbed, dMessageEmbedDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DMessageEmbed struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dMessageEmbed.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DMessageEmbeds(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, dMessageEmbed, dMessageEmbedDBTypes, true, dMessageEmbedColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DMessageEmbed struct: %s", err)
	}

	if err = dMessageEmbed.Update(tx); err != nil {
		t.Error(err)
	}
}

func testDMessageEmbedsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(dMessageEmbedColumns) == len(dMessageEmbedPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	dMessageEmbed := &DMessageEmbed{}
	if err = randomize.Struct(seed, dMessageEmbed, dMessageEmbedDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DMessageEmbed struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dMessageEmbed.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DMessageEmbeds(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, dMessageEmbed, dMessageEmbedDBTypes, true, dMessageEmbedPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DMessageEmbed struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(dMessageEmbedColumns, dMessageEmbedPrimaryKeyColumns) {
		fields = dMessageEmbedColumns
	} else {
		fields = strmangle.SetComplement(
			dMessageEmbedColumns,
			dMessageEmbedPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(dMessageEmbed))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := DMessageEmbedSlice{dMessageEmbed}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testDMessageEmbedsUpsert(t *testing.T) {
	t.Parallel()

	if len(dMessageEmbedColumns) == len(dMessageEmbedPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	dMessageEmbed := DMessageEmbed{}
	if err = randomize.Struct(seed, &dMessageEmbed, dMessageEmbedDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DMessageEmbed struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dMessageEmbed.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert DMessageEmbed: %s", err)
	}

	count, err := DMessageEmbeds(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &dMessageEmbed, dMessageEmbedDBTypes, false, dMessageEmbedPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DMessageEmbed struct: %s", err)
	}

	if err = dMessageEmbed.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert DMessageEmbed: %s", err)
	}

	count, err = DMessageEmbeds(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
