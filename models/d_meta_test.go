package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testDMeta(t *testing.T) {
	t.Parallel()

	query := DMeta(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testDMetaDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dMetum := &DMetum{}
	if err = randomize.Struct(seed, dMetum, dMetumDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DMetum struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dMetum.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = dMetum.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := DMeta(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDMetaQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dMetum := &DMetum{}
	if err = randomize.Struct(seed, dMetum, dMetumDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DMetum struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dMetum.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = DMeta(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := DMeta(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDMetaSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dMetum := &DMetum{}
	if err = randomize.Struct(seed, dMetum, dMetumDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DMetum struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dMetum.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := DMetumSlice{dMetum}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := DMeta(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testDMetaExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dMetum := &DMetum{}
	if err = randomize.Struct(seed, dMetum, dMetumDBTypes, true, dMetumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DMetum struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dMetum.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := DMetumExists(tx, dMetum.Key)
	if err != nil {
		t.Errorf("Unable to check if DMetum exists: %s", err)
	}
	if !e {
		t.Errorf("Expected DMetumExistsG to return true, but got false.")
	}
}
func testDMetaFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dMetum := &DMetum{}
	if err = randomize.Struct(seed, dMetum, dMetumDBTypes, true, dMetumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DMetum struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dMetum.Insert(tx); err != nil {
		t.Error(err)
	}

	dMetumFound, err := FindDMetum(tx, dMetum.Key)
	if err != nil {
		t.Error(err)
	}

	if dMetumFound == nil {
		t.Error("want a record, got nil")
	}
}
func testDMetaBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dMetum := &DMetum{}
	if err = randomize.Struct(seed, dMetum, dMetumDBTypes, true, dMetumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DMetum struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dMetum.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = DMeta(tx).Bind(dMetum); err != nil {
		t.Error(err)
	}
}

func testDMetaOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dMetum := &DMetum{}
	if err = randomize.Struct(seed, dMetum, dMetumDBTypes, true, dMetumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DMetum struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dMetum.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := DMeta(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testDMetaAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dMetumOne := &DMetum{}
	dMetumTwo := &DMetum{}
	if err = randomize.Struct(seed, dMetumOne, dMetumDBTypes, false, dMetumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DMetum struct: %s", err)
	}
	if err = randomize.Struct(seed, dMetumTwo, dMetumDBTypes, false, dMetumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DMetum struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dMetumOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = dMetumTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := DMeta(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testDMetaCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	dMetumOne := &DMetum{}
	dMetumTwo := &DMetum{}
	if err = randomize.Struct(seed, dMetumOne, dMetumDBTypes, false, dMetumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DMetum struct: %s", err)
	}
	if err = randomize.Struct(seed, dMetumTwo, dMetumDBTypes, false, dMetumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DMetum struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dMetumOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = dMetumTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DMeta(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testDMetaInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dMetum := &DMetum{}
	if err = randomize.Struct(seed, dMetum, dMetumDBTypes, true, dMetumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DMetum struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dMetum.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DMeta(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDMetaInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dMetum := &DMetum{}
	if err = randomize.Struct(seed, dMetum, dMetumDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DMetum struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dMetum.Insert(tx, dMetumColumns...); err != nil {
		t.Error(err)
	}

	count, err := DMeta(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDMetaReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dMetum := &DMetum{}
	if err = randomize.Struct(seed, dMetum, dMetumDBTypes, true, dMetumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DMetum struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dMetum.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = dMetum.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testDMetaReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dMetum := &DMetum{}
	if err = randomize.Struct(seed, dMetum, dMetumDBTypes, true, dMetumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DMetum struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dMetum.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := DMetumSlice{dMetum}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testDMetaSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dMetum := &DMetum{}
	if err = randomize.Struct(seed, dMetum, dMetumDBTypes, true, dMetumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DMetum struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dMetum.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := DMeta(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	dMetumDBTypes = map[string]string{`Key`: `text`, `Value`: `bytea`}
	_             = bytes.MinRead
)

func testDMetaUpdate(t *testing.T) {
	t.Parallel()

	if len(dMetumColumns) == len(dMetumPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	dMetum := &DMetum{}
	if err = randomize.Struct(seed, dMetum, dMetumDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DMetum struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dMetum.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DMeta(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, dMetum, dMetumDBTypes, true, dMetumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DMetum struct: %s", err)
	}

	if err = dMetum.Update(tx); err != nil {
		t.Error(err)
	}
}

func testDMetaSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(dMetumColumns) == len(dMetumPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	dMetum := &DMetum{}
	if err = randomize.Struct(seed, dMetum, dMetumDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DMetum struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dMetum.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DMeta(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, dMetum, dMetumDBTypes, true, dMetumPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DMetum struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(dMetumColumns, dMetumPrimaryKeyColumns) {
		fields = dMetumColumns
	} else {
		fields = strmangle.SetComplement(
			dMetumColumns,
			dMetumPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(dMetum))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := DMetumSlice{dMetum}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testDMetaUpsert(t *testing.T) {
	t.Parallel()

	if len(dMetumColumns) == len(dMetumPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	dMetum := DMetum{}
	if err = randomize.Struct(seed, &dMetum, dMetumDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DMetum struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dMetum.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert DMetum: %s", err)
	}

	count, err := DMeta(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &dMetum, dMetumDBTypes, false, dMetumPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DMetum struct: %s", err)
	}

	if err = dMetum.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert DMetum: %s", err)
	}

	count, err = DMeta(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
