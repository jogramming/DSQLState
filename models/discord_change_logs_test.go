package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testDiscordChangeLogs(t *testing.T) {
	t.Parallel()

	query := DiscordChangeLogs(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testDiscordChangeLogsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordChangeLog := &DiscordChangeLog{}
	if err = randomize.Struct(seed, discordChangeLog, discordChangeLogDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordChangeLog struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordChangeLog.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = discordChangeLog.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := DiscordChangeLogs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDiscordChangeLogsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordChangeLog := &DiscordChangeLog{}
	if err = randomize.Struct(seed, discordChangeLog, discordChangeLogDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordChangeLog struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordChangeLog.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = DiscordChangeLogs(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := DiscordChangeLogs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDiscordChangeLogsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordChangeLog := &DiscordChangeLog{}
	if err = randomize.Struct(seed, discordChangeLog, discordChangeLogDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordChangeLog struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordChangeLog.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := DiscordChangeLogSlice{discordChangeLog}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := DiscordChangeLogs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testDiscordChangeLogsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordChangeLog := &DiscordChangeLog{}
	if err = randomize.Struct(seed, discordChangeLog, discordChangeLogDBTypes, true, discordChangeLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordChangeLog struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordChangeLog.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := DiscordChangeLogExists(tx, discordChangeLog.ID)
	if err != nil {
		t.Errorf("Unable to check if DiscordChangeLog exists: %s", err)
	}
	if !e {
		t.Errorf("Expected DiscordChangeLogExistsG to return true, but got false.")
	}
}
func testDiscordChangeLogsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordChangeLog := &DiscordChangeLog{}
	if err = randomize.Struct(seed, discordChangeLog, discordChangeLogDBTypes, true, discordChangeLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordChangeLog struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordChangeLog.Insert(tx); err != nil {
		t.Error(err)
	}

	discordChangeLogFound, err := FindDiscordChangeLog(tx, discordChangeLog.ID)
	if err != nil {
		t.Error(err)
	}

	if discordChangeLogFound == nil {
		t.Error("want a record, got nil")
	}
}
func testDiscordChangeLogsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordChangeLog := &DiscordChangeLog{}
	if err = randomize.Struct(seed, discordChangeLog, discordChangeLogDBTypes, true, discordChangeLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordChangeLog struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordChangeLog.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = DiscordChangeLogs(tx).Bind(discordChangeLog); err != nil {
		t.Error(err)
	}
}

func testDiscordChangeLogsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordChangeLog := &DiscordChangeLog{}
	if err = randomize.Struct(seed, discordChangeLog, discordChangeLogDBTypes, true, discordChangeLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordChangeLog struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordChangeLog.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := DiscordChangeLogs(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testDiscordChangeLogsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordChangeLogOne := &DiscordChangeLog{}
	discordChangeLogTwo := &DiscordChangeLog{}
	if err = randomize.Struct(seed, discordChangeLogOne, discordChangeLogDBTypes, false, discordChangeLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordChangeLog struct: %s", err)
	}
	if err = randomize.Struct(seed, discordChangeLogTwo, discordChangeLogDBTypes, false, discordChangeLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordChangeLog struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordChangeLogOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = discordChangeLogTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := DiscordChangeLogs(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testDiscordChangeLogsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	discordChangeLogOne := &DiscordChangeLog{}
	discordChangeLogTwo := &DiscordChangeLog{}
	if err = randomize.Struct(seed, discordChangeLogOne, discordChangeLogDBTypes, false, discordChangeLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordChangeLog struct: %s", err)
	}
	if err = randomize.Struct(seed, discordChangeLogTwo, discordChangeLogDBTypes, false, discordChangeLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordChangeLog struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordChangeLogOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = discordChangeLogTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DiscordChangeLogs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testDiscordChangeLogsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordChangeLog := &DiscordChangeLog{}
	if err = randomize.Struct(seed, discordChangeLog, discordChangeLogDBTypes, true, discordChangeLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordChangeLog struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordChangeLog.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DiscordChangeLogs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDiscordChangeLogsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordChangeLog := &DiscordChangeLog{}
	if err = randomize.Struct(seed, discordChangeLog, discordChangeLogDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordChangeLog struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordChangeLog.Insert(tx, discordChangeLogColumns...); err != nil {
		t.Error(err)
	}

	count, err := DiscordChangeLogs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDiscordChangeLogsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordChangeLog := &DiscordChangeLog{}
	if err = randomize.Struct(seed, discordChangeLog, discordChangeLogDBTypes, true, discordChangeLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordChangeLog struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordChangeLog.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = discordChangeLog.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testDiscordChangeLogsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordChangeLog := &DiscordChangeLog{}
	if err = randomize.Struct(seed, discordChangeLog, discordChangeLogDBTypes, true, discordChangeLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordChangeLog struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordChangeLog.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := DiscordChangeLogSlice{discordChangeLog}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testDiscordChangeLogsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordChangeLog := &DiscordChangeLog{}
	if err = randomize.Struct(seed, discordChangeLog, discordChangeLogDBTypes, true, discordChangeLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordChangeLog struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordChangeLog.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := DiscordChangeLogs(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	discordChangeLogDBTypes = map[string]string{`Field`: `integer`, `ID`: `bigint`, `Valuebool`: `boolean`, `Valueint`: `bigint`, `Valuestring`: `text`}
	_                       = bytes.MinRead
)

func testDiscordChangeLogsUpdate(t *testing.T) {
	t.Parallel()

	if len(discordChangeLogColumns) == len(discordChangeLogPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	discordChangeLog := &DiscordChangeLog{}
	if err = randomize.Struct(seed, discordChangeLog, discordChangeLogDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordChangeLog struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordChangeLog.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DiscordChangeLogs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, discordChangeLog, discordChangeLogDBTypes, true, discordChangeLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordChangeLog struct: %s", err)
	}

	if err = discordChangeLog.Update(tx); err != nil {
		t.Error(err)
	}
}

func testDiscordChangeLogsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(discordChangeLogColumns) == len(discordChangeLogPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	discordChangeLog := &DiscordChangeLog{}
	if err = randomize.Struct(seed, discordChangeLog, discordChangeLogDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordChangeLog struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordChangeLog.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DiscordChangeLogs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, discordChangeLog, discordChangeLogDBTypes, true, discordChangeLogPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DiscordChangeLog struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(discordChangeLogColumns, discordChangeLogPrimaryKeyColumns) {
		fields = discordChangeLogColumns
	} else {
		fields = strmangle.SetComplement(
			discordChangeLogColumns,
			discordChangeLogPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(discordChangeLog))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := DiscordChangeLogSlice{discordChangeLog}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testDiscordChangeLogsUpsert(t *testing.T) {
	t.Parallel()

	if len(discordChangeLogColumns) == len(discordChangeLogPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	discordChangeLog := DiscordChangeLog{}
	if err = randomize.Struct(seed, &discordChangeLog, discordChangeLogDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordChangeLog struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordChangeLog.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert DiscordChangeLog: %s", err)
	}

	count, err := DiscordChangeLogs(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &discordChangeLog, discordChangeLogDBTypes, false, discordChangeLogPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DiscordChangeLog struct: %s", err)
	}

	if err = discordChangeLog.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert DiscordChangeLog: %s", err)
	}

	count, err = DiscordChangeLogs(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
