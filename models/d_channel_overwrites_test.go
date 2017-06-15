package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testDChannelOverwrites(t *testing.T) {
	t.Parallel()

	query := DChannelOverwrites(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testDChannelOverwritesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dChannelOverwrite := &DChannelOverwrite{}
	if err = randomize.Struct(seed, dChannelOverwrite, dChannelOverwriteDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DChannelOverwrite struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dChannelOverwrite.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = dChannelOverwrite.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := DChannelOverwrites(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDChannelOverwritesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dChannelOverwrite := &DChannelOverwrite{}
	if err = randomize.Struct(seed, dChannelOverwrite, dChannelOverwriteDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DChannelOverwrite struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dChannelOverwrite.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = DChannelOverwrites(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := DChannelOverwrites(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDChannelOverwritesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dChannelOverwrite := &DChannelOverwrite{}
	if err = randomize.Struct(seed, dChannelOverwrite, dChannelOverwriteDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DChannelOverwrite struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dChannelOverwrite.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := DChannelOverwriteSlice{dChannelOverwrite}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := DChannelOverwrites(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testDChannelOverwritesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dChannelOverwrite := &DChannelOverwrite{}
	if err = randomize.Struct(seed, dChannelOverwrite, dChannelOverwriteDBTypes, true, dChannelOverwriteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DChannelOverwrite struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dChannelOverwrite.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := DChannelOverwriteExists(tx, dChannelOverwrite.ID, dChannelOverwrite.ChannelID)
	if err != nil {
		t.Errorf("Unable to check if DChannelOverwrite exists: %s", err)
	}
	if !e {
		t.Errorf("Expected DChannelOverwriteExistsG to return true, but got false.")
	}
}
func testDChannelOverwritesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dChannelOverwrite := &DChannelOverwrite{}
	if err = randomize.Struct(seed, dChannelOverwrite, dChannelOverwriteDBTypes, true, dChannelOverwriteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DChannelOverwrite struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dChannelOverwrite.Insert(tx); err != nil {
		t.Error(err)
	}

	dChannelOverwriteFound, err := FindDChannelOverwrite(tx, dChannelOverwrite.ID, dChannelOverwrite.ChannelID)
	if err != nil {
		t.Error(err)
	}

	if dChannelOverwriteFound == nil {
		t.Error("want a record, got nil")
	}
}
func testDChannelOverwritesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dChannelOverwrite := &DChannelOverwrite{}
	if err = randomize.Struct(seed, dChannelOverwrite, dChannelOverwriteDBTypes, true, dChannelOverwriteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DChannelOverwrite struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dChannelOverwrite.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = DChannelOverwrites(tx).Bind(dChannelOverwrite); err != nil {
		t.Error(err)
	}
}

func testDChannelOverwritesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dChannelOverwrite := &DChannelOverwrite{}
	if err = randomize.Struct(seed, dChannelOverwrite, dChannelOverwriteDBTypes, true, dChannelOverwriteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DChannelOverwrite struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dChannelOverwrite.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := DChannelOverwrites(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testDChannelOverwritesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dChannelOverwriteOne := &DChannelOverwrite{}
	dChannelOverwriteTwo := &DChannelOverwrite{}
	if err = randomize.Struct(seed, dChannelOverwriteOne, dChannelOverwriteDBTypes, false, dChannelOverwriteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DChannelOverwrite struct: %s", err)
	}
	if err = randomize.Struct(seed, dChannelOverwriteTwo, dChannelOverwriteDBTypes, false, dChannelOverwriteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DChannelOverwrite struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dChannelOverwriteOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = dChannelOverwriteTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := DChannelOverwrites(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testDChannelOverwritesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	dChannelOverwriteOne := &DChannelOverwrite{}
	dChannelOverwriteTwo := &DChannelOverwrite{}
	if err = randomize.Struct(seed, dChannelOverwriteOne, dChannelOverwriteDBTypes, false, dChannelOverwriteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DChannelOverwrite struct: %s", err)
	}
	if err = randomize.Struct(seed, dChannelOverwriteTwo, dChannelOverwriteDBTypes, false, dChannelOverwriteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DChannelOverwrite struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dChannelOverwriteOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = dChannelOverwriteTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DChannelOverwrites(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testDChannelOverwritesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dChannelOverwrite := &DChannelOverwrite{}
	if err = randomize.Struct(seed, dChannelOverwrite, dChannelOverwriteDBTypes, true, dChannelOverwriteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DChannelOverwrite struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dChannelOverwrite.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DChannelOverwrites(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDChannelOverwritesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dChannelOverwrite := &DChannelOverwrite{}
	if err = randomize.Struct(seed, dChannelOverwrite, dChannelOverwriteDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DChannelOverwrite struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dChannelOverwrite.Insert(tx, dChannelOverwriteColumns...); err != nil {
		t.Error(err)
	}

	count, err := DChannelOverwrites(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDChannelOverwriteToOneDChannelUsingChannel(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local DChannelOverwrite
	var foreign DChannel

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, dChannelOverwriteDBTypes, true, dChannelOverwriteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DChannelOverwrite struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, dChannelDBTypes, true, dChannelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DChannel struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.ChannelID = foreign.ID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Channel(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := DChannelOverwriteSlice{&local}
	if err = local.L.LoadChannel(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Channel == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Channel = nil
	if err = local.L.LoadChannel(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Channel == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testDChannelOverwriteToOneSetOpDChannelUsingChannel(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a DChannelOverwrite
	var b, c DChannel

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, dChannelOverwriteDBTypes, false, strmangle.SetComplement(dChannelOverwritePrimaryKeyColumns, dChannelOverwriteColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, dChannelDBTypes, false, strmangle.SetComplement(dChannelPrimaryKeyColumns, dChannelColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, dChannelDBTypes, false, strmangle.SetComplement(dChannelPrimaryKeyColumns, dChannelColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*DChannel{&b, &c} {
		err = a.SetChannel(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Channel != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ChannelDChannelOverwrites[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.ChannelID != x.ID {
			t.Error("foreign key was wrong value", a.ChannelID)
		}

		if exists, err := DChannelOverwriteExists(tx, a.ID, a.ChannelID); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'a' to exist")
		}

	}
}
func testDChannelOverwritesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dChannelOverwrite := &DChannelOverwrite{}
	if err = randomize.Struct(seed, dChannelOverwrite, dChannelOverwriteDBTypes, true, dChannelOverwriteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DChannelOverwrite struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dChannelOverwrite.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = dChannelOverwrite.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testDChannelOverwritesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dChannelOverwrite := &DChannelOverwrite{}
	if err = randomize.Struct(seed, dChannelOverwrite, dChannelOverwriteDBTypes, true, dChannelOverwriteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DChannelOverwrite struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dChannelOverwrite.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := DChannelOverwriteSlice{dChannelOverwrite}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testDChannelOverwritesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dChannelOverwrite := &DChannelOverwrite{}
	if err = randomize.Struct(seed, dChannelOverwrite, dChannelOverwriteDBTypes, true, dChannelOverwriteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DChannelOverwrite struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dChannelOverwrite.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := DChannelOverwrites(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	dChannelOverwriteDBTypes = map[string]string{`Allow`: `integer`, `ChannelID`: `bigint`, `Deny`: `integer`, `ID`: `bigint`, `Type`: `character varying`}
	_                        = bytes.MinRead
)

func testDChannelOverwritesUpdate(t *testing.T) {
	t.Parallel()

	if len(dChannelOverwriteColumns) == len(dChannelOverwritePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	dChannelOverwrite := &DChannelOverwrite{}
	if err = randomize.Struct(seed, dChannelOverwrite, dChannelOverwriteDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DChannelOverwrite struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dChannelOverwrite.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DChannelOverwrites(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, dChannelOverwrite, dChannelOverwriteDBTypes, true, dChannelOverwriteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DChannelOverwrite struct: %s", err)
	}

	if err = dChannelOverwrite.Update(tx); err != nil {
		t.Error(err)
	}
}

func testDChannelOverwritesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(dChannelOverwriteColumns) == len(dChannelOverwritePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	dChannelOverwrite := &DChannelOverwrite{}
	if err = randomize.Struct(seed, dChannelOverwrite, dChannelOverwriteDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DChannelOverwrite struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dChannelOverwrite.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DChannelOverwrites(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, dChannelOverwrite, dChannelOverwriteDBTypes, true, dChannelOverwritePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DChannelOverwrite struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(dChannelOverwriteColumns, dChannelOverwritePrimaryKeyColumns) {
		fields = dChannelOverwriteColumns
	} else {
		fields = strmangle.SetComplement(
			dChannelOverwriteColumns,
			dChannelOverwritePrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(dChannelOverwrite))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := DChannelOverwriteSlice{dChannelOverwrite}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testDChannelOverwritesUpsert(t *testing.T) {
	t.Parallel()

	if len(dChannelOverwriteColumns) == len(dChannelOverwritePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	dChannelOverwrite := DChannelOverwrite{}
	if err = randomize.Struct(seed, &dChannelOverwrite, dChannelOverwriteDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DChannelOverwrite struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dChannelOverwrite.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert DChannelOverwrite: %s", err)
	}

	count, err := DChannelOverwrites(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &dChannelOverwrite, dChannelOverwriteDBTypes, false, dChannelOverwritePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DChannelOverwrite struct: %s", err)
	}

	if err = dChannelOverwrite.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert DChannelOverwrite: %s", err)
	}

	count, err = DChannelOverwrites(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
