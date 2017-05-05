package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testDiscordGuildChannels(t *testing.T) {
	t.Parallel()

	query := DiscordGuildChannels(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testDiscordGuildChannelsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordGuildChannel := &DiscordGuildChannel{}
	if err = randomize.Struct(seed, discordGuildChannel, discordGuildChannelDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordGuildChannel struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordGuildChannel.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = discordGuildChannel.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := DiscordGuildChannels(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDiscordGuildChannelsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordGuildChannel := &DiscordGuildChannel{}
	if err = randomize.Struct(seed, discordGuildChannel, discordGuildChannelDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordGuildChannel struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordGuildChannel.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = DiscordGuildChannels(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := DiscordGuildChannels(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDiscordGuildChannelsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordGuildChannel := &DiscordGuildChannel{}
	if err = randomize.Struct(seed, discordGuildChannel, discordGuildChannelDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordGuildChannel struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordGuildChannel.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := DiscordGuildChannelSlice{discordGuildChannel}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := DiscordGuildChannels(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testDiscordGuildChannelsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordGuildChannel := &DiscordGuildChannel{}
	if err = randomize.Struct(seed, discordGuildChannel, discordGuildChannelDBTypes, true, discordGuildChannelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordGuildChannel struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordGuildChannel.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := DiscordGuildChannelExists(tx, discordGuildChannel.ID)
	if err != nil {
		t.Errorf("Unable to check if DiscordGuildChannel exists: %s", err)
	}
	if !e {
		t.Errorf("Expected DiscordGuildChannelExistsG to return true, but got false.")
	}
}
func testDiscordGuildChannelsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordGuildChannel := &DiscordGuildChannel{}
	if err = randomize.Struct(seed, discordGuildChannel, discordGuildChannelDBTypes, true, discordGuildChannelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordGuildChannel struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordGuildChannel.Insert(tx); err != nil {
		t.Error(err)
	}

	discordGuildChannelFound, err := FindDiscordGuildChannel(tx, discordGuildChannel.ID)
	if err != nil {
		t.Error(err)
	}

	if discordGuildChannelFound == nil {
		t.Error("want a record, got nil")
	}
}
func testDiscordGuildChannelsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordGuildChannel := &DiscordGuildChannel{}
	if err = randomize.Struct(seed, discordGuildChannel, discordGuildChannelDBTypes, true, discordGuildChannelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordGuildChannel struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordGuildChannel.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = DiscordGuildChannels(tx).Bind(discordGuildChannel); err != nil {
		t.Error(err)
	}
}

func testDiscordGuildChannelsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordGuildChannel := &DiscordGuildChannel{}
	if err = randomize.Struct(seed, discordGuildChannel, discordGuildChannelDBTypes, true, discordGuildChannelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordGuildChannel struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordGuildChannel.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := DiscordGuildChannels(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testDiscordGuildChannelsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordGuildChannelOne := &DiscordGuildChannel{}
	discordGuildChannelTwo := &DiscordGuildChannel{}
	if err = randomize.Struct(seed, discordGuildChannelOne, discordGuildChannelDBTypes, false, discordGuildChannelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordGuildChannel struct: %s", err)
	}
	if err = randomize.Struct(seed, discordGuildChannelTwo, discordGuildChannelDBTypes, false, discordGuildChannelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordGuildChannel struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordGuildChannelOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = discordGuildChannelTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := DiscordGuildChannels(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testDiscordGuildChannelsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	discordGuildChannelOne := &DiscordGuildChannel{}
	discordGuildChannelTwo := &DiscordGuildChannel{}
	if err = randomize.Struct(seed, discordGuildChannelOne, discordGuildChannelDBTypes, false, discordGuildChannelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordGuildChannel struct: %s", err)
	}
	if err = randomize.Struct(seed, discordGuildChannelTwo, discordGuildChannelDBTypes, false, discordGuildChannelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordGuildChannel struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordGuildChannelOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = discordGuildChannelTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DiscordGuildChannels(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testDiscordGuildChannelsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordGuildChannel := &DiscordGuildChannel{}
	if err = randomize.Struct(seed, discordGuildChannel, discordGuildChannelDBTypes, true, discordGuildChannelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordGuildChannel struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordGuildChannel.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DiscordGuildChannels(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDiscordGuildChannelsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordGuildChannel := &DiscordGuildChannel{}
	if err = randomize.Struct(seed, discordGuildChannel, discordGuildChannelDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordGuildChannel struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordGuildChannel.Insert(tx, discordGuildChannelColumns...); err != nil {
		t.Error(err)
	}

	count, err := DiscordGuildChannels(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDiscordGuildChannelToOneDiscordGuildUsingGuild(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local DiscordGuildChannel
	var foreign DiscordGuild

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, discordGuildChannelDBTypes, true, discordGuildChannelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordGuildChannel struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, discordGuildDBTypes, true, discordGuildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordGuild struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.GuildID = foreign.ID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Guild(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := DiscordGuildChannelSlice{&local}
	if err = local.L.LoadGuild(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Guild == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Guild = nil
	if err = local.L.LoadGuild(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Guild == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testDiscordGuildChannelToOneSetOpDiscordGuildUsingGuild(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a DiscordGuildChannel
	var b, c DiscordGuild

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, discordGuildChannelDBTypes, false, strmangle.SetComplement(discordGuildChannelPrimaryKeyColumns, discordGuildChannelColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, discordGuildDBTypes, false, strmangle.SetComplement(discordGuildPrimaryKeyColumns, discordGuildColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, discordGuildDBTypes, false, strmangle.SetComplement(discordGuildPrimaryKeyColumns, discordGuildColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*DiscordGuild{&b, &c} {
		err = a.SetGuild(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Guild != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.GuildDiscordGuildChannels[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.GuildID != x.ID {
			t.Error("foreign key was wrong value", a.GuildID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.GuildID))
		reflect.Indirect(reflect.ValueOf(&a.GuildID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.GuildID != x.ID {
			t.Error("foreign key was wrong value", a.GuildID, x.ID)
		}
	}
}
func testDiscordGuildChannelsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordGuildChannel := &DiscordGuildChannel{}
	if err = randomize.Struct(seed, discordGuildChannel, discordGuildChannelDBTypes, true, discordGuildChannelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordGuildChannel struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordGuildChannel.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = discordGuildChannel.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testDiscordGuildChannelsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordGuildChannel := &DiscordGuildChannel{}
	if err = randomize.Struct(seed, discordGuildChannel, discordGuildChannelDBTypes, true, discordGuildChannelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordGuildChannel struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordGuildChannel.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := DiscordGuildChannelSlice{discordGuildChannel}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testDiscordGuildChannelsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordGuildChannel := &DiscordGuildChannel{}
	if err = randomize.Struct(seed, discordGuildChannel, discordGuildChannelDBTypes, true, discordGuildChannelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordGuildChannel struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordGuildChannel.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := DiscordGuildChannels(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	discordGuildChannelDBTypes = map[string]string{`Bitrate`: `integer`, `CreatedAt`: `timestamp with time zone`, `DeletedAt`: `timestamp with time zone`, `GuildID`: `bigint`, `ID`: `bigint`, `LastMessageID`: `bigint`, `Name`: `text`, `Position`: `integer`, `Topic`: `text`, `Type`: `text`}
	_                          = bytes.MinRead
)

func testDiscordGuildChannelsUpdate(t *testing.T) {
	t.Parallel()

	if len(discordGuildChannelColumns) == len(discordGuildChannelPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	discordGuildChannel := &DiscordGuildChannel{}
	if err = randomize.Struct(seed, discordGuildChannel, discordGuildChannelDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordGuildChannel struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordGuildChannel.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DiscordGuildChannels(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, discordGuildChannel, discordGuildChannelDBTypes, true, discordGuildChannelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordGuildChannel struct: %s", err)
	}

	if err = discordGuildChannel.Update(tx); err != nil {
		t.Error(err)
	}
}

func testDiscordGuildChannelsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(discordGuildChannelColumns) == len(discordGuildChannelPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	discordGuildChannel := &DiscordGuildChannel{}
	if err = randomize.Struct(seed, discordGuildChannel, discordGuildChannelDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordGuildChannel struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordGuildChannel.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DiscordGuildChannels(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, discordGuildChannel, discordGuildChannelDBTypes, true, discordGuildChannelPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DiscordGuildChannel struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(discordGuildChannelColumns, discordGuildChannelPrimaryKeyColumns) {
		fields = discordGuildChannelColumns
	} else {
		fields = strmangle.SetComplement(
			discordGuildChannelColumns,
			discordGuildChannelPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(discordGuildChannel))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := DiscordGuildChannelSlice{discordGuildChannel}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testDiscordGuildChannelsUpsert(t *testing.T) {
	t.Parallel()

	if len(discordGuildChannelColumns) == len(discordGuildChannelPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	discordGuildChannel := DiscordGuildChannel{}
	if err = randomize.Struct(seed, &discordGuildChannel, discordGuildChannelDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordGuildChannel struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordGuildChannel.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert DiscordGuildChannel: %s", err)
	}

	count, err := DiscordGuildChannels(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &discordGuildChannel, discordGuildChannelDBTypes, false, discordGuildChannelPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DiscordGuildChannel struct: %s", err)
	}

	if err = discordGuildChannel.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert DiscordGuildChannel: %s", err)
	}

	count, err = DiscordGuildChannels(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
