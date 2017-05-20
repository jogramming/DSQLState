package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testDiscordGuilds(t *testing.T) {
	t.Parallel()

	query := DiscordGuilds(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testDiscordGuildsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordGuild := &DiscordGuild{}
	if err = randomize.Struct(seed, discordGuild, discordGuildDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordGuild struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordGuild.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = discordGuild.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := DiscordGuilds(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDiscordGuildsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordGuild := &DiscordGuild{}
	if err = randomize.Struct(seed, discordGuild, discordGuildDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordGuild struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordGuild.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = DiscordGuilds(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := DiscordGuilds(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDiscordGuildsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordGuild := &DiscordGuild{}
	if err = randomize.Struct(seed, discordGuild, discordGuildDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordGuild struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordGuild.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := DiscordGuildSlice{discordGuild}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := DiscordGuilds(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testDiscordGuildsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordGuild := &DiscordGuild{}
	if err = randomize.Struct(seed, discordGuild, discordGuildDBTypes, true, discordGuildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordGuild struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordGuild.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := DiscordGuildExists(tx, discordGuild.ID)
	if err != nil {
		t.Errorf("Unable to check if DiscordGuild exists: %s", err)
	}
	if !e {
		t.Errorf("Expected DiscordGuildExistsG to return true, but got false.")
	}
}
func testDiscordGuildsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordGuild := &DiscordGuild{}
	if err = randomize.Struct(seed, discordGuild, discordGuildDBTypes, true, discordGuildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordGuild struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordGuild.Insert(tx); err != nil {
		t.Error(err)
	}

	discordGuildFound, err := FindDiscordGuild(tx, discordGuild.ID)
	if err != nil {
		t.Error(err)
	}

	if discordGuildFound == nil {
		t.Error("want a record, got nil")
	}
}
func testDiscordGuildsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordGuild := &DiscordGuild{}
	if err = randomize.Struct(seed, discordGuild, discordGuildDBTypes, true, discordGuildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordGuild struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordGuild.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = DiscordGuilds(tx).Bind(discordGuild); err != nil {
		t.Error(err)
	}
}

func testDiscordGuildsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordGuild := &DiscordGuild{}
	if err = randomize.Struct(seed, discordGuild, discordGuildDBTypes, true, discordGuildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordGuild struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordGuild.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := DiscordGuilds(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testDiscordGuildsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordGuildOne := &DiscordGuild{}
	discordGuildTwo := &DiscordGuild{}
	if err = randomize.Struct(seed, discordGuildOne, discordGuildDBTypes, false, discordGuildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordGuild struct: %s", err)
	}
	if err = randomize.Struct(seed, discordGuildTwo, discordGuildDBTypes, false, discordGuildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordGuild struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordGuildOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = discordGuildTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := DiscordGuilds(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testDiscordGuildsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	discordGuildOne := &DiscordGuild{}
	discordGuildTwo := &DiscordGuild{}
	if err = randomize.Struct(seed, discordGuildOne, discordGuildDBTypes, false, discordGuildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordGuild struct: %s", err)
	}
	if err = randomize.Struct(seed, discordGuildTwo, discordGuildDBTypes, false, discordGuildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordGuild struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordGuildOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = discordGuildTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DiscordGuilds(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testDiscordGuildsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordGuild := &DiscordGuild{}
	if err = randomize.Struct(seed, discordGuild, discordGuildDBTypes, true, discordGuildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordGuild struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordGuild.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DiscordGuilds(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDiscordGuildsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordGuild := &DiscordGuild{}
	if err = randomize.Struct(seed, discordGuild, discordGuildDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordGuild struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordGuild.Insert(tx, discordGuildColumns...); err != nil {
		t.Error(err)
	}

	count, err := DiscordGuilds(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDiscordGuildsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordGuild := &DiscordGuild{}
	if err = randomize.Struct(seed, discordGuild, discordGuildDBTypes, true, discordGuildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordGuild struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordGuild.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = discordGuild.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testDiscordGuildsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordGuild := &DiscordGuild{}
	if err = randomize.Struct(seed, discordGuild, discordGuildDBTypes, true, discordGuildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordGuild struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordGuild.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := DiscordGuildSlice{discordGuild}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testDiscordGuildsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordGuild := &DiscordGuild{}
	if err = randomize.Struct(seed, discordGuild, discordGuildDBTypes, true, discordGuildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordGuild struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordGuild.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := DiscordGuilds(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	discordGuildDBTypes = map[string]string{`AfkChannelID`: `bigint`, `AfkTimeout`: `integer`, `CreatedAt`: `timestamp with time zone`, `DefaultMessageNotifications`: `smallint`, `EmbedChannelID`: `bigint`, `EmbedEnabled`: `boolean`, `ID`: `bigint`, `Icon`: `text`, `Large`: `boolean`, `LeftAt`: `timestamp with time zone`, `MemberCount`: `integer`, `Name`: `text`, `OwnerID`: `bigint`, `Region`: `text`, `Splash`: `text`, `VerificationLevel`: `smallint`}
	_                   = bytes.MinRead
)

func testDiscordGuildsUpdate(t *testing.T) {
	t.Parallel()

	if len(discordGuildColumns) == len(discordGuildPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	discordGuild := &DiscordGuild{}
	if err = randomize.Struct(seed, discordGuild, discordGuildDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordGuild struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordGuild.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DiscordGuilds(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, discordGuild, discordGuildDBTypes, true, discordGuildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordGuild struct: %s", err)
	}

	if err = discordGuild.Update(tx); err != nil {
		t.Error(err)
	}
}

func testDiscordGuildsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(discordGuildColumns) == len(discordGuildPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	discordGuild := &DiscordGuild{}
	if err = randomize.Struct(seed, discordGuild, discordGuildDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordGuild struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordGuild.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DiscordGuilds(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, discordGuild, discordGuildDBTypes, true, discordGuildPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DiscordGuild struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(discordGuildColumns, discordGuildPrimaryKeyColumns) {
		fields = discordGuildColumns
	} else {
		fields = strmangle.SetComplement(
			discordGuildColumns,
			discordGuildPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(discordGuild))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := DiscordGuildSlice{discordGuild}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testDiscordGuildsUpsert(t *testing.T) {
	t.Parallel()

	if len(discordGuildColumns) == len(discordGuildPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	discordGuild := DiscordGuild{}
	if err = randomize.Struct(seed, &discordGuild, discordGuildDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordGuild struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordGuild.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert DiscordGuild: %s", err)
	}

	count, err := DiscordGuilds(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &discordGuild, discordGuildDBTypes, false, discordGuildPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DiscordGuild struct: %s", err)
	}

	if err = discordGuild.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert DiscordGuild: %s", err)
	}

	count, err = DiscordGuilds(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
