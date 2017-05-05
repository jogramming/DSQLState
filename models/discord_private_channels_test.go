package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testDiscordPrivateChannels(t *testing.T) {
	t.Parallel()

	query := DiscordPrivateChannels(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testDiscordPrivateChannelsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordPrivateChannel := &DiscordPrivateChannel{}
	if err = randomize.Struct(seed, discordPrivateChannel, discordPrivateChannelDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordPrivateChannel struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordPrivateChannel.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = discordPrivateChannel.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := DiscordPrivateChannels(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDiscordPrivateChannelsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordPrivateChannel := &DiscordPrivateChannel{}
	if err = randomize.Struct(seed, discordPrivateChannel, discordPrivateChannelDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordPrivateChannel struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordPrivateChannel.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = DiscordPrivateChannels(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := DiscordPrivateChannels(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDiscordPrivateChannelsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordPrivateChannel := &DiscordPrivateChannel{}
	if err = randomize.Struct(seed, discordPrivateChannel, discordPrivateChannelDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordPrivateChannel struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordPrivateChannel.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := DiscordPrivateChannelSlice{discordPrivateChannel}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := DiscordPrivateChannels(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testDiscordPrivateChannelsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordPrivateChannel := &DiscordPrivateChannel{}
	if err = randomize.Struct(seed, discordPrivateChannel, discordPrivateChannelDBTypes, true, discordPrivateChannelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordPrivateChannel struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordPrivateChannel.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := DiscordPrivateChannelExists(tx, discordPrivateChannel.ID)
	if err != nil {
		t.Errorf("Unable to check if DiscordPrivateChannel exists: %s", err)
	}
	if !e {
		t.Errorf("Expected DiscordPrivateChannelExistsG to return true, but got false.")
	}
}
func testDiscordPrivateChannelsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordPrivateChannel := &DiscordPrivateChannel{}
	if err = randomize.Struct(seed, discordPrivateChannel, discordPrivateChannelDBTypes, true, discordPrivateChannelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordPrivateChannel struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordPrivateChannel.Insert(tx); err != nil {
		t.Error(err)
	}

	discordPrivateChannelFound, err := FindDiscordPrivateChannel(tx, discordPrivateChannel.ID)
	if err != nil {
		t.Error(err)
	}

	if discordPrivateChannelFound == nil {
		t.Error("want a record, got nil")
	}
}
func testDiscordPrivateChannelsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordPrivateChannel := &DiscordPrivateChannel{}
	if err = randomize.Struct(seed, discordPrivateChannel, discordPrivateChannelDBTypes, true, discordPrivateChannelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordPrivateChannel struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordPrivateChannel.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = DiscordPrivateChannels(tx).Bind(discordPrivateChannel); err != nil {
		t.Error(err)
	}
}

func testDiscordPrivateChannelsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordPrivateChannel := &DiscordPrivateChannel{}
	if err = randomize.Struct(seed, discordPrivateChannel, discordPrivateChannelDBTypes, true, discordPrivateChannelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordPrivateChannel struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordPrivateChannel.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := DiscordPrivateChannels(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testDiscordPrivateChannelsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordPrivateChannelOne := &DiscordPrivateChannel{}
	discordPrivateChannelTwo := &DiscordPrivateChannel{}
	if err = randomize.Struct(seed, discordPrivateChannelOne, discordPrivateChannelDBTypes, false, discordPrivateChannelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordPrivateChannel struct: %s", err)
	}
	if err = randomize.Struct(seed, discordPrivateChannelTwo, discordPrivateChannelDBTypes, false, discordPrivateChannelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordPrivateChannel struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordPrivateChannelOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = discordPrivateChannelTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := DiscordPrivateChannels(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testDiscordPrivateChannelsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	discordPrivateChannelOne := &DiscordPrivateChannel{}
	discordPrivateChannelTwo := &DiscordPrivateChannel{}
	if err = randomize.Struct(seed, discordPrivateChannelOne, discordPrivateChannelDBTypes, false, discordPrivateChannelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordPrivateChannel struct: %s", err)
	}
	if err = randomize.Struct(seed, discordPrivateChannelTwo, discordPrivateChannelDBTypes, false, discordPrivateChannelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordPrivateChannel struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordPrivateChannelOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = discordPrivateChannelTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DiscordPrivateChannels(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testDiscordPrivateChannelsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordPrivateChannel := &DiscordPrivateChannel{}
	if err = randomize.Struct(seed, discordPrivateChannel, discordPrivateChannelDBTypes, true, discordPrivateChannelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordPrivateChannel struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordPrivateChannel.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DiscordPrivateChannels(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDiscordPrivateChannelsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordPrivateChannel := &DiscordPrivateChannel{}
	if err = randomize.Struct(seed, discordPrivateChannel, discordPrivateChannelDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordPrivateChannel struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordPrivateChannel.Insert(tx, discordPrivateChannelColumns...); err != nil {
		t.Error(err)
	}

	count, err := DiscordPrivateChannels(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDiscordPrivateChannelToOneDiscordUserUsingRecipient(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local DiscordPrivateChannel
	var foreign DiscordUser

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, discordPrivateChannelDBTypes, true, discordPrivateChannelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordPrivateChannel struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, discordUserDBTypes, true, discordUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordUser struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.RecipientID = foreign.ID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Recipient(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := DiscordPrivateChannelSlice{&local}
	if err = local.L.LoadRecipient(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Recipient == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Recipient = nil
	if err = local.L.LoadRecipient(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Recipient == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testDiscordPrivateChannelToOneSetOpDiscordUserUsingRecipient(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a DiscordPrivateChannel
	var b, c DiscordUser

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, discordPrivateChannelDBTypes, false, strmangle.SetComplement(discordPrivateChannelPrimaryKeyColumns, discordPrivateChannelColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, discordUserDBTypes, false, strmangle.SetComplement(discordUserPrimaryKeyColumns, discordUserColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, discordUserDBTypes, false, strmangle.SetComplement(discordUserPrimaryKeyColumns, discordUserColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*DiscordUser{&b, &c} {
		err = a.SetRecipient(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Recipient != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.RecipientDiscordPrivateChannels[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.RecipientID != x.ID {
			t.Error("foreign key was wrong value", a.RecipientID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.RecipientID))
		reflect.Indirect(reflect.ValueOf(&a.RecipientID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.RecipientID != x.ID {
			t.Error("foreign key was wrong value", a.RecipientID, x.ID)
		}
	}
}
func testDiscordPrivateChannelsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordPrivateChannel := &DiscordPrivateChannel{}
	if err = randomize.Struct(seed, discordPrivateChannel, discordPrivateChannelDBTypes, true, discordPrivateChannelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordPrivateChannel struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordPrivateChannel.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = discordPrivateChannel.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testDiscordPrivateChannelsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordPrivateChannel := &DiscordPrivateChannel{}
	if err = randomize.Struct(seed, discordPrivateChannel, discordPrivateChannelDBTypes, true, discordPrivateChannelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordPrivateChannel struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordPrivateChannel.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := DiscordPrivateChannelSlice{discordPrivateChannel}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testDiscordPrivateChannelsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordPrivateChannel := &DiscordPrivateChannel{}
	if err = randomize.Struct(seed, discordPrivateChannel, discordPrivateChannelDBTypes, true, discordPrivateChannelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordPrivateChannel struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordPrivateChannel.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := DiscordPrivateChannels(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	discordPrivateChannelDBTypes = map[string]string{`CreatedAt`: `timestamp with time zone`, `ID`: `bigint`, `LastMessageID`: `bigint`, `Name`: `text`, `RecipientID`: `bigint`, `Topic`: `text`}
	_                            = bytes.MinRead
)

func testDiscordPrivateChannelsUpdate(t *testing.T) {
	t.Parallel()

	if len(discordPrivateChannelColumns) == len(discordPrivateChannelPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	discordPrivateChannel := &DiscordPrivateChannel{}
	if err = randomize.Struct(seed, discordPrivateChannel, discordPrivateChannelDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordPrivateChannel struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordPrivateChannel.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DiscordPrivateChannels(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, discordPrivateChannel, discordPrivateChannelDBTypes, true, discordPrivateChannelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordPrivateChannel struct: %s", err)
	}

	if err = discordPrivateChannel.Update(tx); err != nil {
		t.Error(err)
	}
}

func testDiscordPrivateChannelsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(discordPrivateChannelColumns) == len(discordPrivateChannelPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	discordPrivateChannel := &DiscordPrivateChannel{}
	if err = randomize.Struct(seed, discordPrivateChannel, discordPrivateChannelDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordPrivateChannel struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordPrivateChannel.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DiscordPrivateChannels(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, discordPrivateChannel, discordPrivateChannelDBTypes, true, discordPrivateChannelPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DiscordPrivateChannel struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(discordPrivateChannelColumns, discordPrivateChannelPrimaryKeyColumns) {
		fields = discordPrivateChannelColumns
	} else {
		fields = strmangle.SetComplement(
			discordPrivateChannelColumns,
			discordPrivateChannelPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(discordPrivateChannel))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := DiscordPrivateChannelSlice{discordPrivateChannel}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testDiscordPrivateChannelsUpsert(t *testing.T) {
	t.Parallel()

	if len(discordPrivateChannelColumns) == len(discordPrivateChannelPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	discordPrivateChannel := DiscordPrivateChannel{}
	if err = randomize.Struct(seed, &discordPrivateChannel, discordPrivateChannelDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordPrivateChannel struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordPrivateChannel.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert DiscordPrivateChannel: %s", err)
	}

	count, err := DiscordPrivateChannels(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &discordPrivateChannel, discordPrivateChannelDBTypes, false, discordPrivateChannelPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DiscordPrivateChannel struct: %s", err)
	}

	if err = discordPrivateChannel.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert DiscordPrivateChannel: %s", err)
	}

	count, err = DiscordPrivateChannels(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
