package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testDiscordChannels(t *testing.T) {
	t.Parallel()

	query := DiscordChannels(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testDiscordChannelsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordChannel := &DiscordChannel{}
	if err = randomize.Struct(seed, discordChannel, discordChannelDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordChannel struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordChannel.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = discordChannel.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := DiscordChannels(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDiscordChannelsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordChannel := &DiscordChannel{}
	if err = randomize.Struct(seed, discordChannel, discordChannelDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordChannel struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordChannel.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = DiscordChannels(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := DiscordChannels(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDiscordChannelsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordChannel := &DiscordChannel{}
	if err = randomize.Struct(seed, discordChannel, discordChannelDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordChannel struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordChannel.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := DiscordChannelSlice{discordChannel}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := DiscordChannels(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testDiscordChannelsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordChannel := &DiscordChannel{}
	if err = randomize.Struct(seed, discordChannel, discordChannelDBTypes, true, discordChannelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordChannel struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordChannel.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := DiscordChannelExists(tx, discordChannel.ID)
	if err != nil {
		t.Errorf("Unable to check if DiscordChannel exists: %s", err)
	}
	if !e {
		t.Errorf("Expected DiscordChannelExistsG to return true, but got false.")
	}
}
func testDiscordChannelsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordChannel := &DiscordChannel{}
	if err = randomize.Struct(seed, discordChannel, discordChannelDBTypes, true, discordChannelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordChannel struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordChannel.Insert(tx); err != nil {
		t.Error(err)
	}

	discordChannelFound, err := FindDiscordChannel(tx, discordChannel.ID)
	if err != nil {
		t.Error(err)
	}

	if discordChannelFound == nil {
		t.Error("want a record, got nil")
	}
}
func testDiscordChannelsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordChannel := &DiscordChannel{}
	if err = randomize.Struct(seed, discordChannel, discordChannelDBTypes, true, discordChannelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordChannel struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordChannel.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = DiscordChannels(tx).Bind(discordChannel); err != nil {
		t.Error(err)
	}
}

func testDiscordChannelsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordChannel := &DiscordChannel{}
	if err = randomize.Struct(seed, discordChannel, discordChannelDBTypes, true, discordChannelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordChannel struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordChannel.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := DiscordChannels(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testDiscordChannelsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordChannelOne := &DiscordChannel{}
	discordChannelTwo := &DiscordChannel{}
	if err = randomize.Struct(seed, discordChannelOne, discordChannelDBTypes, false, discordChannelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordChannel struct: %s", err)
	}
	if err = randomize.Struct(seed, discordChannelTwo, discordChannelDBTypes, false, discordChannelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordChannel struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordChannelOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = discordChannelTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := DiscordChannels(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testDiscordChannelsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	discordChannelOne := &DiscordChannel{}
	discordChannelTwo := &DiscordChannel{}
	if err = randomize.Struct(seed, discordChannelOne, discordChannelDBTypes, false, discordChannelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordChannel struct: %s", err)
	}
	if err = randomize.Struct(seed, discordChannelTwo, discordChannelDBTypes, false, discordChannelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordChannel struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordChannelOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = discordChannelTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DiscordChannels(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testDiscordChannelsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordChannel := &DiscordChannel{}
	if err = randomize.Struct(seed, discordChannel, discordChannelDBTypes, true, discordChannelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordChannel struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordChannel.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DiscordChannels(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDiscordChannelsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordChannel := &DiscordChannel{}
	if err = randomize.Struct(seed, discordChannel, discordChannelDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordChannel struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordChannel.Insert(tx, discordChannelColumns...); err != nil {
		t.Error(err)
	}

	count, err := DiscordChannels(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDiscordChannelToManyChannelDiscordChannelOverwrites(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a DiscordChannel
	var b, c DiscordChannelOverwrite

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, discordChannelDBTypes, true, discordChannelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordChannel struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, discordChannelOverwriteDBTypes, false, discordChannelOverwriteColumnsWithDefault...)
	randomize.Struct(seed, &c, discordChannelOverwriteDBTypes, false, discordChannelOverwriteColumnsWithDefault...)

	b.ChannelID = a.ID
	c.ChannelID = a.ID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	discordChannelOverwrite, err := a.ChannelDiscordChannelOverwrites(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range discordChannelOverwrite {
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

	slice := DiscordChannelSlice{&a}
	if err = a.L.LoadChannelDiscordChannelOverwrites(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.ChannelDiscordChannelOverwrites); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.ChannelDiscordChannelOverwrites = nil
	if err = a.L.LoadChannelDiscordChannelOverwrites(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.ChannelDiscordChannelOverwrites); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", discordChannelOverwrite)
	}
}

func testDiscordChannelToManyChannelDiscordVoiceStates(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a DiscordChannel
	var b, c DiscordVoiceState

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, discordChannelDBTypes, true, discordChannelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordChannel struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, discordVoiceStateDBTypes, false, discordVoiceStateColumnsWithDefault...)
	randomize.Struct(seed, &c, discordVoiceStateDBTypes, false, discordVoiceStateColumnsWithDefault...)

	b.ChannelID = a.ID
	c.ChannelID = a.ID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	discordVoiceState, err := a.ChannelDiscordVoiceStates(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range discordVoiceState {
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

	slice := DiscordChannelSlice{&a}
	if err = a.L.LoadChannelDiscordVoiceStates(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.ChannelDiscordVoiceStates); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.ChannelDiscordVoiceStates = nil
	if err = a.L.LoadChannelDiscordVoiceStates(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.ChannelDiscordVoiceStates); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", discordVoiceState)
	}
}

func testDiscordChannelToManyAddOpChannelDiscordChannelOverwrites(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a DiscordChannel
	var b, c, d, e DiscordChannelOverwrite

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, discordChannelDBTypes, false, strmangle.SetComplement(discordChannelPrimaryKeyColumns, discordChannelColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*DiscordChannelOverwrite{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, discordChannelOverwriteDBTypes, false, strmangle.SetComplement(discordChannelOverwritePrimaryKeyColumns, discordChannelOverwriteColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*DiscordChannelOverwrite{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddChannelDiscordChannelOverwrites(tx, i != 0, x...)
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

		if a.R.ChannelDiscordChannelOverwrites[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.ChannelDiscordChannelOverwrites[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.ChannelDiscordChannelOverwrites(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testDiscordChannelToManyAddOpChannelDiscordVoiceStates(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a DiscordChannel
	var b, c, d, e DiscordVoiceState

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, discordChannelDBTypes, false, strmangle.SetComplement(discordChannelPrimaryKeyColumns, discordChannelColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*DiscordVoiceState{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, discordVoiceStateDBTypes, false, strmangle.SetComplement(discordVoiceStatePrimaryKeyColumns, discordVoiceStateColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*DiscordVoiceState{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddChannelDiscordVoiceStates(tx, i != 0, x...)
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

		if a.R.ChannelDiscordVoiceStates[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.ChannelDiscordVoiceStates[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.ChannelDiscordVoiceStates(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testDiscordChannelsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordChannel := &DiscordChannel{}
	if err = randomize.Struct(seed, discordChannel, discordChannelDBTypes, true, discordChannelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordChannel struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordChannel.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = discordChannel.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testDiscordChannelsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordChannel := &DiscordChannel{}
	if err = randomize.Struct(seed, discordChannel, discordChannelDBTypes, true, discordChannelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordChannel struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordChannel.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := DiscordChannelSlice{discordChannel}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testDiscordChannelsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordChannel := &DiscordChannel{}
	if err = randomize.Struct(seed, discordChannel, discordChannelDBTypes, true, discordChannelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordChannel struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordChannel.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := DiscordChannels(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	discordChannelDBTypes = map[string]string{`Bitrate`: `integer`, `CreatedAt`: `timestamp with time zone`, `DeletedAt`: `timestamp with time zone`, `GuildID`: `bigint`, `ID`: `bigint`, `LastMessageID`: `bigint`, `Name`: `text`, `Position`: `integer`, `RecipientID`: `bigint`, `Topic`: `text`, `Type`: `text`}
	_                     = bytes.MinRead
)

func testDiscordChannelsUpdate(t *testing.T) {
	t.Parallel()

	if len(discordChannelColumns) == len(discordChannelPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	discordChannel := &DiscordChannel{}
	if err = randomize.Struct(seed, discordChannel, discordChannelDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordChannel struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordChannel.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DiscordChannels(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, discordChannel, discordChannelDBTypes, true, discordChannelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordChannel struct: %s", err)
	}

	if err = discordChannel.Update(tx); err != nil {
		t.Error(err)
	}
}

func testDiscordChannelsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(discordChannelColumns) == len(discordChannelPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	discordChannel := &DiscordChannel{}
	if err = randomize.Struct(seed, discordChannel, discordChannelDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordChannel struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordChannel.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DiscordChannels(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, discordChannel, discordChannelDBTypes, true, discordChannelPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DiscordChannel struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(discordChannelColumns, discordChannelPrimaryKeyColumns) {
		fields = discordChannelColumns
	} else {
		fields = strmangle.SetComplement(
			discordChannelColumns,
			discordChannelPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(discordChannel))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := DiscordChannelSlice{discordChannel}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testDiscordChannelsUpsert(t *testing.T) {
	t.Parallel()

	if len(discordChannelColumns) == len(discordChannelPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	discordChannel := DiscordChannel{}
	if err = randomize.Struct(seed, &discordChannel, discordChannelDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordChannel struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordChannel.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert DiscordChannel: %s", err)
	}

	count, err := DiscordChannels(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &discordChannel, discordChannelDBTypes, false, discordChannelPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DiscordChannel struct: %s", err)
	}

	if err = discordChannel.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert DiscordChannel: %s", err)
	}

	count, err = DiscordChannels(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
