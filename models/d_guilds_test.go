package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testDGuilds(t *testing.T) {
	t.Parallel()

	query := DGuilds(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testDGuildsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dGuild := &DGuild{}
	if err = randomize.Struct(seed, dGuild, dGuildDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DGuild struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dGuild.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = dGuild.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := DGuilds(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDGuildsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dGuild := &DGuild{}
	if err = randomize.Struct(seed, dGuild, dGuildDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DGuild struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dGuild.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = DGuilds(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := DGuilds(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDGuildsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dGuild := &DGuild{}
	if err = randomize.Struct(seed, dGuild, dGuildDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DGuild struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dGuild.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := DGuildSlice{dGuild}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := DGuilds(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testDGuildsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dGuild := &DGuild{}
	if err = randomize.Struct(seed, dGuild, dGuildDBTypes, true, dGuildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DGuild struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dGuild.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := DGuildExists(tx, dGuild.ID)
	if err != nil {
		t.Errorf("Unable to check if DGuild exists: %s", err)
	}
	if !e {
		t.Errorf("Expected DGuildExistsG to return true, but got false.")
	}
}
func testDGuildsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dGuild := &DGuild{}
	if err = randomize.Struct(seed, dGuild, dGuildDBTypes, true, dGuildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DGuild struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dGuild.Insert(tx); err != nil {
		t.Error(err)
	}

	dGuildFound, err := FindDGuild(tx, dGuild.ID)
	if err != nil {
		t.Error(err)
	}

	if dGuildFound == nil {
		t.Error("want a record, got nil")
	}
}
func testDGuildsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dGuild := &DGuild{}
	if err = randomize.Struct(seed, dGuild, dGuildDBTypes, true, dGuildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DGuild struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dGuild.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = DGuilds(tx).Bind(dGuild); err != nil {
		t.Error(err)
	}
}

func testDGuildsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dGuild := &DGuild{}
	if err = randomize.Struct(seed, dGuild, dGuildDBTypes, true, dGuildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DGuild struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dGuild.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := DGuilds(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testDGuildsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dGuildOne := &DGuild{}
	dGuildTwo := &DGuild{}
	if err = randomize.Struct(seed, dGuildOne, dGuildDBTypes, false, dGuildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DGuild struct: %s", err)
	}
	if err = randomize.Struct(seed, dGuildTwo, dGuildDBTypes, false, dGuildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DGuild struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dGuildOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = dGuildTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := DGuilds(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testDGuildsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	dGuildOne := &DGuild{}
	dGuildTwo := &DGuild{}
	if err = randomize.Struct(seed, dGuildOne, dGuildDBTypes, false, dGuildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DGuild struct: %s", err)
	}
	if err = randomize.Struct(seed, dGuildTwo, dGuildDBTypes, false, dGuildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DGuild struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dGuildOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = dGuildTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DGuilds(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testDGuildsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dGuild := &DGuild{}
	if err = randomize.Struct(seed, dGuild, dGuildDBTypes, true, dGuildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DGuild struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dGuild.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DGuilds(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDGuildsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dGuild := &DGuild{}
	if err = randomize.Struct(seed, dGuild, dGuildDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DGuild struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dGuild.Insert(tx, dGuildColumns...); err != nil {
		t.Error(err)
	}

	count, err := DGuilds(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDGuildsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dGuild := &DGuild{}
	if err = randomize.Struct(seed, dGuild, dGuildDBTypes, true, dGuildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DGuild struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dGuild.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = dGuild.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testDGuildsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dGuild := &DGuild{}
	if err = randomize.Struct(seed, dGuild, dGuildDBTypes, true, dGuildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DGuild struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dGuild.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := DGuildSlice{dGuild}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testDGuildsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dGuild := &DGuild{}
	if err = randomize.Struct(seed, dGuild, dGuildDBTypes, true, dGuildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DGuild struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dGuild.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := DGuilds(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	dGuildDBTypes = map[string]string{`AfkChannelID`: `bigint`, `AfkTimeout`: `integer`, `CreatedAt`: `timestamp with time zone`, `DefaultMessageNotifications`: `smallint`, `EmbedChannelID`: `bigint`, `EmbedEnabled`: `boolean`, `ID`: `bigint`, `Icon`: `text`, `Large`: `boolean`, `LeftAt`: `timestamp with time zone`, `MemberCount`: `integer`, `Name`: `text`, `OwnerID`: `bigint`, `Region`: `text`, `Splash`: `text`, `Synced`: `boolean`, `VerificationLevel`: `smallint`}
	_             = bytes.MinRead
)

func testDGuildsUpdate(t *testing.T) {
	t.Parallel()

	if len(dGuildColumns) == len(dGuildPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	dGuild := &DGuild{}
	if err = randomize.Struct(seed, dGuild, dGuildDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DGuild struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dGuild.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DGuilds(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, dGuild, dGuildDBTypes, true, dGuildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DGuild struct: %s", err)
	}

	if err = dGuild.Update(tx); err != nil {
		t.Error(err)
	}
}

func testDGuildsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(dGuildColumns) == len(dGuildPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	dGuild := &DGuild{}
	if err = randomize.Struct(seed, dGuild, dGuildDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DGuild struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dGuild.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DGuilds(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, dGuild, dGuildDBTypes, true, dGuildPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DGuild struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(dGuildColumns, dGuildPrimaryKeyColumns) {
		fields = dGuildColumns
	} else {
		fields = strmangle.SetComplement(
			dGuildColumns,
			dGuildPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(dGuild))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := DGuildSlice{dGuild}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testDGuildsUpsert(t *testing.T) {
	t.Parallel()

	if len(dGuildColumns) == len(dGuildPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	dGuild := DGuild{}
	if err = randomize.Struct(seed, &dGuild, dGuildDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DGuild struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dGuild.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert DGuild: %s", err)
	}

	count, err := DGuilds(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &dGuild, dGuildDBTypes, false, dGuildPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DGuild struct: %s", err)
	}

	if err = dGuild.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert DGuild: %s", err)
	}

	count, err = DGuilds(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
