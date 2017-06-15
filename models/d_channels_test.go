package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testDChannels(t *testing.T) {
	t.Parallel()

	query := DChannels(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testDChannelsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dChannel := &DChannel{}
	if err = randomize.Struct(seed, dChannel, dChannelDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DChannel struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dChannel.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = dChannel.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := DChannels(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDChannelsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dChannel := &DChannel{}
	if err = randomize.Struct(seed, dChannel, dChannelDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DChannel struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dChannel.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = DChannels(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := DChannels(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDChannelsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dChannel := &DChannel{}
	if err = randomize.Struct(seed, dChannel, dChannelDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DChannel struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dChannel.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := DChannelSlice{dChannel}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := DChannels(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testDChannelsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dChannel := &DChannel{}
	if err = randomize.Struct(seed, dChannel, dChannelDBTypes, true, dChannelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DChannel struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dChannel.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := DChannelExists(tx, dChannel.ID)
	if err != nil {
		t.Errorf("Unable to check if DChannel exists: %s", err)
	}
	if !e {
		t.Errorf("Expected DChannelExistsG to return true, but got false.")
	}
}
func testDChannelsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dChannel := &DChannel{}
	if err = randomize.Struct(seed, dChannel, dChannelDBTypes, true, dChannelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DChannel struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dChannel.Insert(tx); err != nil {
		t.Error(err)
	}

	dChannelFound, err := FindDChannel(tx, dChannel.ID)
	if err != nil {
		t.Error(err)
	}

	if dChannelFound == nil {
		t.Error("want a record, got nil")
	}
}
func testDChannelsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dChannel := &DChannel{}
	if err = randomize.Struct(seed, dChannel, dChannelDBTypes, true, dChannelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DChannel struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dChannel.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = DChannels(tx).Bind(dChannel); err != nil {
		t.Error(err)
	}
}

func testDChannelsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dChannel := &DChannel{}
	if err = randomize.Struct(seed, dChannel, dChannelDBTypes, true, dChannelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DChannel struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dChannel.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := DChannels(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testDChannelsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dChannelOne := &DChannel{}
	dChannelTwo := &DChannel{}
	if err = randomize.Struct(seed, dChannelOne, dChannelDBTypes, false, dChannelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DChannel struct: %s", err)
	}
	if err = randomize.Struct(seed, dChannelTwo, dChannelDBTypes, false, dChannelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DChannel struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dChannelOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = dChannelTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := DChannels(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testDChannelsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	dChannelOne := &DChannel{}
	dChannelTwo := &DChannel{}
	if err = randomize.Struct(seed, dChannelOne, dChannelDBTypes, false, dChannelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DChannel struct: %s", err)
	}
	if err = randomize.Struct(seed, dChannelTwo, dChannelDBTypes, false, dChannelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DChannel struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dChannelOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = dChannelTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DChannels(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testDChannelsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dChannel := &DChannel{}
	if err = randomize.Struct(seed, dChannel, dChannelDBTypes, true, dChannelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DChannel struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dChannel.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DChannels(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDChannelsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dChannel := &DChannel{}
	if err = randomize.Struct(seed, dChannel, dChannelDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DChannel struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dChannel.Insert(tx, dChannelColumns...); err != nil {
		t.Error(err)
	}

	count, err := DChannels(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDChannelToManyChannelDChannelOverwrites(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a DChannel
	var b, c DChannelOverwrite

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, dChannelDBTypes, true, dChannelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DChannel struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, dChannelOverwriteDBTypes, false, dChannelOverwriteColumnsWithDefault...)
	randomize.Struct(seed, &c, dChannelOverwriteDBTypes, false, dChannelOverwriteColumnsWithDefault...)

	b.ChannelID = a.ID
	c.ChannelID = a.ID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	dChannelOverwrite, err := a.ChannelDChannelOverwrites(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range dChannelOverwrite {
		if v.ChannelID == b.ChannelID {
			bFound = true
		}
		if v.ChannelID == c.ChannelID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := DChannelSlice{&a}
	if err = a.L.LoadChannelDChannelOverwrites(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.ChannelDChannelOverwrites); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.ChannelDChannelOverwrites = nil
	if err = a.L.LoadChannelDChannelOverwrites(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.ChannelDChannelOverwrites); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", dChannelOverwrite)
	}
}

func testDChannelToManyChannelDVoiceStates(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a DChannel
	var b, c DVoiceState

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, dChannelDBTypes, true, dChannelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DChannel struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, dVoiceStateDBTypes, false, dVoiceStateColumnsWithDefault...)
	randomize.Struct(seed, &c, dVoiceStateDBTypes, false, dVoiceStateColumnsWithDefault...)

	b.ChannelID = a.ID
	c.ChannelID = a.ID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	dVoiceState, err := a.ChannelDVoiceStates(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range dVoiceState {
		if v.ChannelID == b.ChannelID {
			bFound = true
		}
		if v.ChannelID == c.ChannelID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := DChannelSlice{&a}
	if err = a.L.LoadChannelDVoiceStates(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.ChannelDVoiceStates); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.ChannelDVoiceStates = nil
	if err = a.L.LoadChannelDVoiceStates(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.ChannelDVoiceStates); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", dVoiceState)
	}
}

func testDChannelToManyAddOpChannelDChannelOverwrites(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a DChannel
	var b, c, d, e DChannelOverwrite

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, dChannelDBTypes, false, strmangle.SetComplement(dChannelPrimaryKeyColumns, dChannelColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*DChannelOverwrite{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, dChannelOverwriteDBTypes, false, strmangle.SetComplement(dChannelOverwritePrimaryKeyColumns, dChannelOverwriteColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*DChannelOverwrite{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddChannelDChannelOverwrites(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.ChannelID {
			t.Error("foreign key was wrong value", a.ID, first.ChannelID)
		}
		if a.ID != second.ChannelID {
			t.Error("foreign key was wrong value", a.ID, second.ChannelID)
		}

		if first.R.Channel != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Channel != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.ChannelDChannelOverwrites[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.ChannelDChannelOverwrites[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.ChannelDChannelOverwrites(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testDChannelToManyAddOpChannelDVoiceStates(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a DChannel
	var b, c, d, e DVoiceState

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, dChannelDBTypes, false, strmangle.SetComplement(dChannelPrimaryKeyColumns, dChannelColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*DVoiceState{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, dVoiceStateDBTypes, false, strmangle.SetComplement(dVoiceStatePrimaryKeyColumns, dVoiceStateColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*DVoiceState{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddChannelDVoiceStates(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.ChannelID {
			t.Error("foreign key was wrong value", a.ID, first.ChannelID)
		}
		if a.ID != second.ChannelID {
			t.Error("foreign key was wrong value", a.ID, second.ChannelID)
		}

		if first.R.Channel != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Channel != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.ChannelDVoiceStates[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.ChannelDVoiceStates[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.ChannelDVoiceStates(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testDChannelsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dChannel := &DChannel{}
	if err = randomize.Struct(seed, dChannel, dChannelDBTypes, true, dChannelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DChannel struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dChannel.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = dChannel.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testDChannelsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dChannel := &DChannel{}
	if err = randomize.Struct(seed, dChannel, dChannelDBTypes, true, dChannelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DChannel struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dChannel.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := DChannelSlice{dChannel}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testDChannelsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dChannel := &DChannel{}
	if err = randomize.Struct(seed, dChannel, dChannelDBTypes, true, dChannelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DChannel struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dChannel.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := DChannels(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	dChannelDBTypes = map[string]string{`Bitrate`: `integer`, `CreatedAt`: `timestamp with time zone`, `DeletedAt`: `timestamp with time zone`, `GuildID`: `bigint`, `ID`: `bigint`, `LastMessageID`: `bigint`, `Name`: `text`, `Position`: `integer`, `RecipientID`: `bigint`, `Synced`: `boolean`, `Topic`: `text`, `Type`: `text`}
	_               = bytes.MinRead
)

func testDChannelsUpdate(t *testing.T) {
	t.Parallel()

	if len(dChannelColumns) == len(dChannelPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	dChannel := &DChannel{}
	if err = randomize.Struct(seed, dChannel, dChannelDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DChannel struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dChannel.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DChannels(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, dChannel, dChannelDBTypes, true, dChannelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DChannel struct: %s", err)
	}

	if err = dChannel.Update(tx); err != nil {
		t.Error(err)
	}
}

func testDChannelsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(dChannelColumns) == len(dChannelPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	dChannel := &DChannel{}
	if err = randomize.Struct(seed, dChannel, dChannelDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DChannel struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dChannel.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DChannels(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, dChannel, dChannelDBTypes, true, dChannelPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DChannel struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(dChannelColumns, dChannelPrimaryKeyColumns) {
		fields = dChannelColumns
	} else {
		fields = strmangle.SetComplement(
			dChannelColumns,
			dChannelPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(dChannel))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := DChannelSlice{dChannel}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testDChannelsUpsert(t *testing.T) {
	t.Parallel()

	if len(dChannelColumns) == len(dChannelPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	dChannel := DChannel{}
	if err = randomize.Struct(seed, &dChannel, dChannelDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DChannel struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dChannel.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert DChannel: %s", err)
	}

	count, err := DChannels(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &dChannel, dChannelDBTypes, false, dChannelPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DChannel struct: %s", err)
	}

	if err = dChannel.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert DChannel: %s", err)
	}

	count, err = DChannels(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
