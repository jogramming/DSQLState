package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testDiscordMemberRoles(t *testing.T) {
	t.Parallel()

	query := DiscordMemberRoles(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testDiscordMemberRolesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordMemberRole := &DiscordMemberRole{}
	if err = randomize.Struct(seed, discordMemberRole, discordMemberRoleDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordMemberRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordMemberRole.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = discordMemberRole.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := DiscordMemberRoles(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDiscordMemberRolesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordMemberRole := &DiscordMemberRole{}
	if err = randomize.Struct(seed, discordMemberRole, discordMemberRoleDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordMemberRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordMemberRole.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = DiscordMemberRoles(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := DiscordMemberRoles(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDiscordMemberRolesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordMemberRole := &DiscordMemberRole{}
	if err = randomize.Struct(seed, discordMemberRole, discordMemberRoleDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordMemberRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordMemberRole.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := DiscordMemberRoleSlice{discordMemberRole}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := DiscordMemberRoles(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testDiscordMemberRolesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordMemberRole := &DiscordMemberRole{}
	if err = randomize.Struct(seed, discordMemberRole, discordMemberRoleDBTypes, true, discordMemberRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordMemberRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordMemberRole.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := DiscordMemberRoleExists(tx, discordMemberRole.UserID, discordMemberRole.GuildID)
	if err != nil {
		t.Errorf("Unable to check if DiscordMemberRole exists: %s", err)
	}
	if !e {
		t.Errorf("Expected DiscordMemberRoleExistsG to return true, but got false.")
	}
}
func testDiscordMemberRolesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordMemberRole := &DiscordMemberRole{}
	if err = randomize.Struct(seed, discordMemberRole, discordMemberRoleDBTypes, true, discordMemberRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordMemberRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordMemberRole.Insert(tx); err != nil {
		t.Error(err)
	}

	discordMemberRoleFound, err := FindDiscordMemberRole(tx, discordMemberRole.UserID, discordMemberRole.GuildID)
	if err != nil {
		t.Error(err)
	}

	if discordMemberRoleFound == nil {
		t.Error("want a record, got nil")
	}
}
func testDiscordMemberRolesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordMemberRole := &DiscordMemberRole{}
	if err = randomize.Struct(seed, discordMemberRole, discordMemberRoleDBTypes, true, discordMemberRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordMemberRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordMemberRole.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = DiscordMemberRoles(tx).Bind(discordMemberRole); err != nil {
		t.Error(err)
	}
}

func testDiscordMemberRolesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordMemberRole := &DiscordMemberRole{}
	if err = randomize.Struct(seed, discordMemberRole, discordMemberRoleDBTypes, true, discordMemberRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordMemberRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordMemberRole.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := DiscordMemberRoles(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testDiscordMemberRolesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordMemberRoleOne := &DiscordMemberRole{}
	discordMemberRoleTwo := &DiscordMemberRole{}
	if err = randomize.Struct(seed, discordMemberRoleOne, discordMemberRoleDBTypes, false, discordMemberRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordMemberRole struct: %s", err)
	}
	if err = randomize.Struct(seed, discordMemberRoleTwo, discordMemberRoleDBTypes, false, discordMemberRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordMemberRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordMemberRoleOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = discordMemberRoleTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := DiscordMemberRoles(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testDiscordMemberRolesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	discordMemberRoleOne := &DiscordMemberRole{}
	discordMemberRoleTwo := &DiscordMemberRole{}
	if err = randomize.Struct(seed, discordMemberRoleOne, discordMemberRoleDBTypes, false, discordMemberRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordMemberRole struct: %s", err)
	}
	if err = randomize.Struct(seed, discordMemberRoleTwo, discordMemberRoleDBTypes, false, discordMemberRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordMemberRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordMemberRoleOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = discordMemberRoleTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DiscordMemberRoles(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testDiscordMemberRolesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordMemberRole := &DiscordMemberRole{}
	if err = randomize.Struct(seed, discordMemberRole, discordMemberRoleDBTypes, true, discordMemberRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordMemberRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordMemberRole.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DiscordMemberRoles(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDiscordMemberRolesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordMemberRole := &DiscordMemberRole{}
	if err = randomize.Struct(seed, discordMemberRole, discordMemberRoleDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordMemberRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordMemberRole.Insert(tx, discordMemberRoleColumns...); err != nil {
		t.Error(err)
	}

	count, err := DiscordMemberRoles(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDiscordMemberRoleToOneDiscordUserUsingUser(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local DiscordMemberRole
	var foreign DiscordUser

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, discordMemberRoleDBTypes, true, discordMemberRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordMemberRole struct: %s", err)
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

	slice := DiscordMemberRoleSlice{&local}
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

func testDiscordMemberRoleToOneDiscordGuildUsingGuild(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local DiscordMemberRole
	var foreign DiscordGuild

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, discordMemberRoleDBTypes, true, discordMemberRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordMemberRole struct: %s", err)
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

	slice := DiscordMemberRoleSlice{&local}
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

func testDiscordMemberRoleToOneDiscordGuildRoleUsingRole(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local DiscordMemberRole
	var foreign DiscordGuildRole

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, discordMemberRoleDBTypes, true, discordMemberRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordMemberRole struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, discordGuildRoleDBTypes, true, discordGuildRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordGuildRole struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.RoleID = foreign.ID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Role(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := DiscordMemberRoleSlice{&local}
	if err = local.L.LoadRole(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Role == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Role = nil
	if err = local.L.LoadRole(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Role == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testDiscordMemberRoleToOneSetOpDiscordUserUsingUser(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a DiscordMemberRole
	var b, c DiscordUser

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, discordMemberRoleDBTypes, false, strmangle.SetComplement(discordMemberRolePrimaryKeyColumns, discordMemberRoleColumnsWithoutDefault)...); err != nil {
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

		if x.R.UserDiscordMemberRoles[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.UserID != x.ID {
			t.Error("foreign key was wrong value", a.UserID)
		}

		if exists, err := DiscordMemberRoleExists(tx, a.UserID, a.GuildID); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'a' to exist")
		}

	}
}
func testDiscordMemberRoleToOneSetOpDiscordGuildUsingGuild(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a DiscordMemberRole
	var b, c DiscordGuild

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, discordMemberRoleDBTypes, false, strmangle.SetComplement(discordMemberRolePrimaryKeyColumns, discordMemberRoleColumnsWithoutDefault)...); err != nil {
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

		if x.R.GuildDiscordMemberRoles[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.GuildID != x.ID {
			t.Error("foreign key was wrong value", a.GuildID)
		}

		if exists, err := DiscordMemberRoleExists(tx, a.UserID, a.GuildID); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'a' to exist")
		}

	}
}
func testDiscordMemberRoleToOneSetOpDiscordGuildRoleUsingRole(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a DiscordMemberRole
	var b, c DiscordGuildRole

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, discordMemberRoleDBTypes, false, strmangle.SetComplement(discordMemberRolePrimaryKeyColumns, discordMemberRoleColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, discordGuildRoleDBTypes, false, strmangle.SetComplement(discordGuildRolePrimaryKeyColumns, discordGuildRoleColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, discordGuildRoleDBTypes, false, strmangle.SetComplement(discordGuildRolePrimaryKeyColumns, discordGuildRoleColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*DiscordGuildRole{&b, &c} {
		err = a.SetRole(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Role != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.RoleDiscordMemberRoles[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.RoleID != x.ID {
			t.Error("foreign key was wrong value", a.RoleID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.RoleID))
		reflect.Indirect(reflect.ValueOf(&a.RoleID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.RoleID != x.ID {
			t.Error("foreign key was wrong value", a.RoleID, x.ID)
		}
	}
}
func testDiscordMemberRolesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordMemberRole := &DiscordMemberRole{}
	if err = randomize.Struct(seed, discordMemberRole, discordMemberRoleDBTypes, true, discordMemberRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordMemberRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordMemberRole.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = discordMemberRole.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testDiscordMemberRolesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordMemberRole := &DiscordMemberRole{}
	if err = randomize.Struct(seed, discordMemberRole, discordMemberRoleDBTypes, true, discordMemberRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordMemberRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordMemberRole.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := DiscordMemberRoleSlice{discordMemberRole}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testDiscordMemberRolesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordMemberRole := &DiscordMemberRole{}
	if err = randomize.Struct(seed, discordMemberRole, discordMemberRoleDBTypes, true, discordMemberRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordMemberRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordMemberRole.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := DiscordMemberRoles(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	discordMemberRoleDBTypes = map[string]string{`CreatedAt`: `timestamp with time zone`, `GuildID`: `bigint`, `RoleID`: `bigint`, `UserID`: `bigint`}
	_                        = bytes.MinRead
)

func testDiscordMemberRolesUpdate(t *testing.T) {
	t.Parallel()

	if len(discordMemberRoleColumns) == len(discordMemberRolePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	discordMemberRole := &DiscordMemberRole{}
	if err = randomize.Struct(seed, discordMemberRole, discordMemberRoleDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordMemberRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordMemberRole.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DiscordMemberRoles(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, discordMemberRole, discordMemberRoleDBTypes, true, discordMemberRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordMemberRole struct: %s", err)
	}

	if err = discordMemberRole.Update(tx); err != nil {
		t.Error(err)
	}
}

func testDiscordMemberRolesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(discordMemberRoleColumns) == len(discordMemberRolePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	discordMemberRole := &DiscordMemberRole{}
	if err = randomize.Struct(seed, discordMemberRole, discordMemberRoleDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordMemberRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordMemberRole.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DiscordMemberRoles(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, discordMemberRole, discordMemberRoleDBTypes, true, discordMemberRolePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DiscordMemberRole struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(discordMemberRoleColumns, discordMemberRolePrimaryKeyColumns) {
		fields = discordMemberRoleColumns
	} else {
		fields = strmangle.SetComplement(
			discordMemberRoleColumns,
			discordMemberRolePrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(discordMemberRole))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := DiscordMemberRoleSlice{discordMemberRole}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testDiscordMemberRolesUpsert(t *testing.T) {
	t.Parallel()

	if len(discordMemberRoleColumns) == len(discordMemberRolePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	discordMemberRole := DiscordMemberRole{}
	if err = randomize.Struct(seed, &discordMemberRole, discordMemberRoleDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordMemberRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordMemberRole.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert DiscordMemberRole: %s", err)
	}

	count, err := DiscordMemberRoles(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &discordMemberRole, discordMemberRoleDBTypes, false, discordMemberRolePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DiscordMemberRole struct: %s", err)
	}

	if err = discordMemberRole.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert DiscordMemberRole: %s", err)
	}

	count, err = DiscordMemberRoles(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
