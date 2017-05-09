package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testDiscordMembers(t *testing.T) {
	t.Parallel()

	query := DiscordMembers(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testDiscordMembersDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordMember := &DiscordMember{}
	if err = randomize.Struct(seed, discordMember, discordMemberDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordMember struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordMember.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = discordMember.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := DiscordMembers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDiscordMembersQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordMember := &DiscordMember{}
	if err = randomize.Struct(seed, discordMember, discordMemberDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordMember struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordMember.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = DiscordMembers(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := DiscordMembers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDiscordMembersSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordMember := &DiscordMember{}
	if err = randomize.Struct(seed, discordMember, discordMemberDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordMember struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordMember.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := DiscordMemberSlice{discordMember}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := DiscordMembers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testDiscordMembersExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordMember := &DiscordMember{}
	if err = randomize.Struct(seed, discordMember, discordMemberDBTypes, true, discordMemberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordMember struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordMember.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := DiscordMemberExists(tx, discordMember.UserID, discordMember.GuildID)
	if err != nil {
		t.Errorf("Unable to check if DiscordMember exists: %s", err)
	}
	if !e {
		t.Errorf("Expected DiscordMemberExistsG to return true, but got false.")
	}
}
func testDiscordMembersFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordMember := &DiscordMember{}
	if err = randomize.Struct(seed, discordMember, discordMemberDBTypes, true, discordMemberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordMember struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordMember.Insert(tx); err != nil {
		t.Error(err)
	}

	discordMemberFound, err := FindDiscordMember(tx, discordMember.UserID, discordMember.GuildID)
	if err != nil {
		t.Error(err)
	}

	if discordMemberFound == nil {
		t.Error("want a record, got nil")
	}
}
func testDiscordMembersBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordMember := &DiscordMember{}
	if err = randomize.Struct(seed, discordMember, discordMemberDBTypes, true, discordMemberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordMember struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordMember.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = DiscordMembers(tx).Bind(discordMember); err != nil {
		t.Error(err)
	}
}

func testDiscordMembersOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordMember := &DiscordMember{}
	if err = randomize.Struct(seed, discordMember, discordMemberDBTypes, true, discordMemberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordMember struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordMember.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := DiscordMembers(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testDiscordMembersAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordMemberOne := &DiscordMember{}
	discordMemberTwo := &DiscordMember{}
	if err = randomize.Struct(seed, discordMemberOne, discordMemberDBTypes, false, discordMemberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordMember struct: %s", err)
	}
	if err = randomize.Struct(seed, discordMemberTwo, discordMemberDBTypes, false, discordMemberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordMember struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordMemberOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = discordMemberTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := DiscordMembers(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testDiscordMembersCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	discordMemberOne := &DiscordMember{}
	discordMemberTwo := &DiscordMember{}
	if err = randomize.Struct(seed, discordMemberOne, discordMemberDBTypes, false, discordMemberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordMember struct: %s", err)
	}
	if err = randomize.Struct(seed, discordMemberTwo, discordMemberDBTypes, false, discordMemberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordMember struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordMemberOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = discordMemberTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DiscordMembers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testDiscordMembersInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordMember := &DiscordMember{}
	if err = randomize.Struct(seed, discordMember, discordMemberDBTypes, true, discordMemberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordMember struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordMember.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DiscordMembers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDiscordMembersInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordMember := &DiscordMember{}
	if err = randomize.Struct(seed, discordMember, discordMemberDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordMember struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordMember.Insert(tx, discordMemberColumns...); err != nil {
		t.Error(err)
	}

	count, err := DiscordMembers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDiscordMemberToOneDiscordUserUsingUser(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local DiscordMember
	var foreign DiscordUser

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, discordMemberDBTypes, true, discordMemberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordMember struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, discordUserDBTypes, true, discordUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordUser struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.UserID = foreign.ID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.User(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := DiscordMemberSlice{&local}
	if err = local.L.LoadUser(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.User = nil
	if err = local.L.LoadUser(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testDiscordMemberToOneDiscordGuildUsingGuild(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local DiscordMember
	var foreign DiscordGuild

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, discordMemberDBTypes, true, discordMemberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordMember struct: %s", err)
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

	slice := DiscordMemberSlice{&local}
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

func testDiscordMemberToOneSetOpDiscordUserUsingUser(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a DiscordMember
	var b, c DiscordUser

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, discordMemberDBTypes, false, strmangle.SetComplement(discordMemberPrimaryKeyColumns, discordMemberColumnsWithoutDefault)...); err != nil {
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
		err = a.SetUser(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.User != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.UserDiscordMembers[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.UserID != x.ID {
			t.Error("foreign key was wrong value", a.UserID)
		}

		if exists, err := DiscordMemberExists(tx, a.UserID, a.GuildID); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'a' to exist")
		}

	}
}
func testDiscordMemberToOneSetOpDiscordGuildUsingGuild(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a DiscordMember
	var b, c DiscordGuild

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, discordMemberDBTypes, false, strmangle.SetComplement(discordMemberPrimaryKeyColumns, discordMemberColumnsWithoutDefault)...); err != nil {
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

		if x.R.GuildDiscordMembers[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.GuildID != x.ID {
			t.Error("foreign key was wrong value", a.GuildID)
		}

		if exists, err := DiscordMemberExists(tx, a.UserID, a.GuildID); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'a' to exist")
		}

	}
}
func testDiscordMembersReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordMember := &DiscordMember{}
	if err = randomize.Struct(seed, discordMember, discordMemberDBTypes, true, discordMemberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordMember struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordMember.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = discordMember.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testDiscordMembersReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordMember := &DiscordMember{}
	if err = randomize.Struct(seed, discordMember, discordMemberDBTypes, true, discordMemberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordMember struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordMember.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := DiscordMemberSlice{discordMember}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testDiscordMembersSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordMember := &DiscordMember{}
	if err = randomize.Struct(seed, discordMember, discordMemberDBTypes, true, discordMemberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordMember struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordMember.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := DiscordMembers(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	discordMemberDBTypes = map[string]string{`CreatedAt`: `timestamp with time zone`, `Deaf`: `boolean`, `GuildID`: `bigint`, `JoinedAt`: `timestamp with time zone`, `LeftAt`: `timestamp with time zone`, `Mute`: `boolean`, `Nick`: `character varying`, `Roles`: `ARRAYbigint`, `UserID`: `bigint`}
	_                    = bytes.MinRead
)

func testDiscordMembersUpdate(t *testing.T) {
	t.Parallel()

	if len(discordMemberColumns) == len(discordMemberPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	discordMember := &DiscordMember{}
	if err = randomize.Struct(seed, discordMember, discordMemberDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordMember struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordMember.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DiscordMembers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, discordMember, discordMemberDBTypes, true, discordMemberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordMember struct: %s", err)
	}

	if err = discordMember.Update(tx); err != nil {
		t.Error(err)
	}
}

func testDiscordMembersSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(discordMemberColumns) == len(discordMemberPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	discordMember := &DiscordMember{}
	if err = randomize.Struct(seed, discordMember, discordMemberDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordMember struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordMember.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DiscordMembers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, discordMember, discordMemberDBTypes, true, discordMemberPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DiscordMember struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(discordMemberColumns, discordMemberPrimaryKeyColumns) {
		fields = discordMemberColumns
	} else {
		fields = strmangle.SetComplement(
			discordMemberColumns,
			discordMemberPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(discordMember))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := DiscordMemberSlice{discordMember}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testDiscordMembersUpsert(t *testing.T) {
	t.Parallel()

	if len(discordMemberColumns) == len(discordMemberPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	discordMember := DiscordMember{}
	if err = randomize.Struct(seed, &discordMember, discordMemberDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordMember struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordMember.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert DiscordMember: %s", err)
	}

	count, err := DiscordMembers(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &discordMember, discordMemberDBTypes, false, discordMemberPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DiscordMember struct: %s", err)
	}

	if err = discordMember.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert DiscordMember: %s", err)
	}

	count, err = DiscordMembers(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
