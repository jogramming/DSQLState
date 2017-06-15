package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testDVoiceStates(t *testing.T) {
	t.Parallel()

	query := DVoiceStates(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testDVoiceStatesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dVoiceState := &DVoiceState{}
	if err = randomize.Struct(seed, dVoiceState, dVoiceStateDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DVoiceState struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dVoiceState.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = dVoiceState.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := DVoiceStates(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDVoiceStatesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dVoiceState := &DVoiceState{}
	if err = randomize.Struct(seed, dVoiceState, dVoiceStateDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DVoiceState struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dVoiceState.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = DVoiceStates(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := DVoiceStates(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDVoiceStatesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dVoiceState := &DVoiceState{}
	if err = randomize.Struct(seed, dVoiceState, dVoiceStateDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DVoiceState struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dVoiceState.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := DVoiceStateSlice{dVoiceState}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := DVoiceStates(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testDVoiceStatesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dVoiceState := &DVoiceState{}
	if err = randomize.Struct(seed, dVoiceState, dVoiceStateDBTypes, true, dVoiceStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DVoiceState struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dVoiceState.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := DVoiceStateExists(tx, dVoiceState.UserID, dVoiceState.GuildID)
	if err != nil {
		t.Errorf("Unable to check if DVoiceState exists: %s", err)
	}
	if !e {
		t.Errorf("Expected DVoiceStateExistsG to return true, but got false.")
	}
}
func testDVoiceStatesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dVoiceState := &DVoiceState{}
	if err = randomize.Struct(seed, dVoiceState, dVoiceStateDBTypes, true, dVoiceStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DVoiceState struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dVoiceState.Insert(tx); err != nil {
		t.Error(err)
	}

	dVoiceStateFound, err := FindDVoiceState(tx, dVoiceState.UserID, dVoiceState.GuildID)
	if err != nil {
		t.Error(err)
	}

	if dVoiceStateFound == nil {
		t.Error("want a record, got nil")
	}
}
func testDVoiceStatesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dVoiceState := &DVoiceState{}
	if err = randomize.Struct(seed, dVoiceState, dVoiceStateDBTypes, true, dVoiceStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DVoiceState struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dVoiceState.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = DVoiceStates(tx).Bind(dVoiceState); err != nil {
		t.Error(err)
	}
}

func testDVoiceStatesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dVoiceState := &DVoiceState{}
	if err = randomize.Struct(seed, dVoiceState, dVoiceStateDBTypes, true, dVoiceStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DVoiceState struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dVoiceState.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := DVoiceStates(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testDVoiceStatesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dVoiceStateOne := &DVoiceState{}
	dVoiceStateTwo := &DVoiceState{}
	if err = randomize.Struct(seed, dVoiceStateOne, dVoiceStateDBTypes, false, dVoiceStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DVoiceState struct: %s", err)
	}
	if err = randomize.Struct(seed, dVoiceStateTwo, dVoiceStateDBTypes, false, dVoiceStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DVoiceState struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dVoiceStateOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = dVoiceStateTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := DVoiceStates(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testDVoiceStatesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	dVoiceStateOne := &DVoiceState{}
	dVoiceStateTwo := &DVoiceState{}
	if err = randomize.Struct(seed, dVoiceStateOne, dVoiceStateDBTypes, false, dVoiceStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DVoiceState struct: %s", err)
	}
	if err = randomize.Struct(seed, dVoiceStateTwo, dVoiceStateDBTypes, false, dVoiceStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DVoiceState struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dVoiceStateOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = dVoiceStateTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DVoiceStates(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testDVoiceStatesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dVoiceState := &DVoiceState{}
	if err = randomize.Struct(seed, dVoiceState, dVoiceStateDBTypes, true, dVoiceStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DVoiceState struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dVoiceState.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DVoiceStates(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDVoiceStatesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dVoiceState := &DVoiceState{}
	if err = randomize.Struct(seed, dVoiceState, dVoiceStateDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DVoiceState struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dVoiceState.Insert(tx, dVoiceStateColumns...); err != nil {
		t.Error(err)
	}

	count, err := DVoiceStates(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDVoiceStateToOneDChannelUsingChannel(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local DVoiceState
	var foreign DChannel

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, dVoiceStateDBTypes, true, dVoiceStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DVoiceState struct: %s", err)
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

	slice := DVoiceStateSlice{&local}
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

func testDVoiceStateToOneSetOpDChannelUsingChannel(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a DVoiceState
	var b, c DChannel

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, dVoiceStateDBTypes, false, strmangle.SetComplement(dVoiceStatePrimaryKeyColumns, dVoiceStateColumnsWithoutDefault)...); err != nil {
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

		if x.R.ChannelDVoiceStates[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.ChannelID != x.ID {
			t.Error("foreign key was wrong value", a.ChannelID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.ChannelID))
		reflect.Indirect(reflect.ValueOf(&a.ChannelID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.ChannelID != x.ID {
			t.Error("foreign key was wrong value", a.ChannelID, x.ID)
		}
	}
}
func testDVoiceStatesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dVoiceState := &DVoiceState{}
	if err = randomize.Struct(seed, dVoiceState, dVoiceStateDBTypes, true, dVoiceStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DVoiceState struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dVoiceState.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = dVoiceState.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testDVoiceStatesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dVoiceState := &DVoiceState{}
	if err = randomize.Struct(seed, dVoiceState, dVoiceStateDBTypes, true, dVoiceStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DVoiceState struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dVoiceState.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := DVoiceStateSlice{dVoiceState}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testDVoiceStatesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dVoiceState := &DVoiceState{}
	if err = randomize.Struct(seed, dVoiceState, dVoiceStateDBTypes, true, dVoiceStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DVoiceState struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dVoiceState.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := DVoiceStates(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	dVoiceStateDBTypes = map[string]string{`ChannelID`: `bigint`, `Deaf`: `boolean`, `GuildID`: `bigint`, `Mute`: `boolean`, `SelfDeaf`: `boolean`, `SelfMute`: `boolean`, `SessionID`: `text`, `Surpress`: `boolean`, `UserID`: `bigint`}
	_                  = bytes.MinRead
)

func testDVoiceStatesUpdate(t *testing.T) {
	t.Parallel()

	if len(dVoiceStateColumns) == len(dVoiceStatePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	dVoiceState := &DVoiceState{}
	if err = randomize.Struct(seed, dVoiceState, dVoiceStateDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DVoiceState struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dVoiceState.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DVoiceStates(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, dVoiceState, dVoiceStateDBTypes, true, dVoiceStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DVoiceState struct: %s", err)
	}

	if err = dVoiceState.Update(tx); err != nil {
		t.Error(err)
	}
}

func testDVoiceStatesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(dVoiceStateColumns) == len(dVoiceStatePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	dVoiceState := &DVoiceState{}
	if err = randomize.Struct(seed, dVoiceState, dVoiceStateDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DVoiceState struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dVoiceState.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DVoiceStates(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, dVoiceState, dVoiceStateDBTypes, true, dVoiceStatePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DVoiceState struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(dVoiceStateColumns, dVoiceStatePrimaryKeyColumns) {
		fields = dVoiceStateColumns
	} else {
		fields = strmangle.SetComplement(
			dVoiceStateColumns,
			dVoiceStatePrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(dVoiceState))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := DVoiceStateSlice{dVoiceState}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testDVoiceStatesUpsert(t *testing.T) {
	t.Parallel()

	if len(dVoiceStateColumns) == len(dVoiceStatePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	dVoiceState := DVoiceState{}
	if err = randomize.Struct(seed, &dVoiceState, dVoiceStateDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DVoiceState struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dVoiceState.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert DVoiceState: %s", err)
	}

	count, err := DVoiceStates(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &dVoiceState, dVoiceStateDBTypes, false, dVoiceStatePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DVoiceState struct: %s", err)
	}

	if err = dVoiceState.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert DVoiceState: %s", err)
	}

	count, err = DVoiceStates(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
