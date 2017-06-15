package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testDMessageRevisions(t *testing.T) {
	t.Parallel()

	query := DMessageRevisions(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testDMessageRevisionsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dMessageRevision := &DMessageRevision{}
	if err = randomize.Struct(seed, dMessageRevision, dMessageRevisionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DMessageRevision struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dMessageRevision.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = dMessageRevision.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := DMessageRevisions(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDMessageRevisionsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dMessageRevision := &DMessageRevision{}
	if err = randomize.Struct(seed, dMessageRevision, dMessageRevisionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DMessageRevision struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dMessageRevision.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = DMessageRevisions(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := DMessageRevisions(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDMessageRevisionsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dMessageRevision := &DMessageRevision{}
	if err = randomize.Struct(seed, dMessageRevision, dMessageRevisionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DMessageRevision struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dMessageRevision.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := DMessageRevisionSlice{dMessageRevision}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := DMessageRevisions(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testDMessageRevisionsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dMessageRevision := &DMessageRevision{}
	if err = randomize.Struct(seed, dMessageRevision, dMessageRevisionDBTypes, true, dMessageRevisionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DMessageRevision struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dMessageRevision.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := DMessageRevisionExists(tx, dMessageRevision.RevisionNum, dMessageRevision.MessageID)
	if err != nil {
		t.Errorf("Unable to check if DMessageRevision exists: %s", err)
	}
	if !e {
		t.Errorf("Expected DMessageRevisionExistsG to return true, but got false.")
	}
}
func testDMessageRevisionsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dMessageRevision := &DMessageRevision{}
	if err = randomize.Struct(seed, dMessageRevision, dMessageRevisionDBTypes, true, dMessageRevisionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DMessageRevision struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dMessageRevision.Insert(tx); err != nil {
		t.Error(err)
	}

	dMessageRevisionFound, err := FindDMessageRevision(tx, dMessageRevision.RevisionNum, dMessageRevision.MessageID)
	if err != nil {
		t.Error(err)
	}

	if dMessageRevisionFound == nil {
		t.Error("want a record, got nil")
	}
}
func testDMessageRevisionsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dMessageRevision := &DMessageRevision{}
	if err = randomize.Struct(seed, dMessageRevision, dMessageRevisionDBTypes, true, dMessageRevisionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DMessageRevision struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dMessageRevision.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = DMessageRevisions(tx).Bind(dMessageRevision); err != nil {
		t.Error(err)
	}
}

func testDMessageRevisionsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dMessageRevision := &DMessageRevision{}
	if err = randomize.Struct(seed, dMessageRevision, dMessageRevisionDBTypes, true, dMessageRevisionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DMessageRevision struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dMessageRevision.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := DMessageRevisions(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testDMessageRevisionsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dMessageRevisionOne := &DMessageRevision{}
	dMessageRevisionTwo := &DMessageRevision{}
	if err = randomize.Struct(seed, dMessageRevisionOne, dMessageRevisionDBTypes, false, dMessageRevisionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DMessageRevision struct: %s", err)
	}
	if err = randomize.Struct(seed, dMessageRevisionTwo, dMessageRevisionDBTypes, false, dMessageRevisionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DMessageRevision struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dMessageRevisionOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = dMessageRevisionTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := DMessageRevisions(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testDMessageRevisionsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	dMessageRevisionOne := &DMessageRevision{}
	dMessageRevisionTwo := &DMessageRevision{}
	if err = randomize.Struct(seed, dMessageRevisionOne, dMessageRevisionDBTypes, false, dMessageRevisionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DMessageRevision struct: %s", err)
	}
	if err = randomize.Struct(seed, dMessageRevisionTwo, dMessageRevisionDBTypes, false, dMessageRevisionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DMessageRevision struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dMessageRevisionOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = dMessageRevisionTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DMessageRevisions(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testDMessageRevisionsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dMessageRevision := &DMessageRevision{}
	if err = randomize.Struct(seed, dMessageRevision, dMessageRevisionDBTypes, true, dMessageRevisionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DMessageRevision struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dMessageRevision.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DMessageRevisions(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDMessageRevisionsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dMessageRevision := &DMessageRevision{}
	if err = randomize.Struct(seed, dMessageRevision, dMessageRevisionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DMessageRevision struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dMessageRevision.Insert(tx, dMessageRevisionColumns...); err != nil {
		t.Error(err)
	}

	count, err := DMessageRevisions(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDMessageRevisionToOneDMessageUsingMessage(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local DMessageRevision
	var foreign DMessage

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, dMessageRevisionDBTypes, true, dMessageRevisionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DMessageRevision struct: %s", err)
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

	slice := DMessageRevisionSlice{&local}
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

func testDMessageRevisionToOneSetOpDMessageUsingMessage(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a DMessageRevision
	var b, c DMessage

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, dMessageRevisionDBTypes, false, strmangle.SetComplement(dMessageRevisionPrimaryKeyColumns, dMessageRevisionColumnsWithoutDefault)...); err != nil {
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

		if x.R.MessageDMessageRevisions[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.MessageID != x.ID {
			t.Error("foreign key was wrong value", a.MessageID)
		}

		if exists, err := DMessageRevisionExists(tx, a.RevisionNum, a.MessageID); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'a' to exist")
		}

	}
}
func testDMessageRevisionsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dMessageRevision := &DMessageRevision{}
	if err = randomize.Struct(seed, dMessageRevision, dMessageRevisionDBTypes, true, dMessageRevisionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DMessageRevision struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dMessageRevision.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = dMessageRevision.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testDMessageRevisionsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dMessageRevision := &DMessageRevision{}
	if err = randomize.Struct(seed, dMessageRevision, dMessageRevisionDBTypes, true, dMessageRevisionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DMessageRevision struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dMessageRevision.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := DMessageRevisionSlice{dMessageRevision}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testDMessageRevisionsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dMessageRevision := &DMessageRevision{}
	if err = randomize.Struct(seed, dMessageRevision, dMessageRevisionDBTypes, true, dMessageRevisionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DMessageRevision struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dMessageRevision.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := DMessageRevisions(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	dMessageRevisionDBTypes = map[string]string{`Content`: `text`, `CreatedAt`: `timestamp with time zone`, `Embeds`: `ARRAYbigint`, `MentionRoles`: `ARRAYbigint`, `Mentions`: `ARRAYbigint`, `MessageID`: `bigint`, `RevisionNum`: `integer`}
	_                       = bytes.MinRead
)

func testDMessageRevisionsUpdate(t *testing.T) {
	t.Parallel()

	if len(dMessageRevisionColumns) == len(dMessageRevisionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	dMessageRevision := &DMessageRevision{}
	if err = randomize.Struct(seed, dMessageRevision, dMessageRevisionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DMessageRevision struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dMessageRevision.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DMessageRevisions(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, dMessageRevision, dMessageRevisionDBTypes, true, dMessageRevisionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DMessageRevision struct: %s", err)
	}

	if err = dMessageRevision.Update(tx); err != nil {
		t.Error(err)
	}
}

func testDMessageRevisionsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(dMessageRevisionColumns) == len(dMessageRevisionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	dMessageRevision := &DMessageRevision{}
	if err = randomize.Struct(seed, dMessageRevision, dMessageRevisionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DMessageRevision struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dMessageRevision.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DMessageRevisions(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, dMessageRevision, dMessageRevisionDBTypes, true, dMessageRevisionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DMessageRevision struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(dMessageRevisionColumns, dMessageRevisionPrimaryKeyColumns) {
		fields = dMessageRevisionColumns
	} else {
		fields = strmangle.SetComplement(
			dMessageRevisionColumns,
			dMessageRevisionPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(dMessageRevision))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := DMessageRevisionSlice{dMessageRevision}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testDMessageRevisionsUpsert(t *testing.T) {
	t.Parallel()

	if len(dMessageRevisionColumns) == len(dMessageRevisionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	dMessageRevision := DMessageRevision{}
	if err = randomize.Struct(seed, &dMessageRevision, dMessageRevisionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DMessageRevision struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dMessageRevision.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert DMessageRevision: %s", err)
	}

	count, err := DMessageRevisions(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &dMessageRevision, dMessageRevisionDBTypes, false, dMessageRevisionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DMessageRevision struct: %s", err)
	}

	if err = dMessageRevision.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert DMessageRevision: %s", err)
	}

	count, err = DMessageRevisions(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
