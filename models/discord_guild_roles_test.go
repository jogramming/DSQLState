package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testDiscordGuildRoles(t *testing.T) {
	t.Parallel()

	query := DiscordGuildRoles(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testDiscordGuildRolesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordGuildRole := &DiscordGuildRole{}
	if err = randomize.Struct(seed, discordGuildRole, discordGuildRoleDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordGuildRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordGuildRole.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = discordGuildRole.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := DiscordGuildRoles(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDiscordGuildRolesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordGuildRole := &DiscordGuildRole{}
	if err = randomize.Struct(seed, discordGuildRole, discordGuildRoleDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordGuildRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordGuildRole.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = DiscordGuildRoles(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := DiscordGuildRoles(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDiscordGuildRolesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordGuildRole := &DiscordGuildRole{}
	if err = randomize.Struct(seed, discordGuildRole, discordGuildRoleDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordGuildRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordGuildRole.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := DiscordGuildRoleSlice{discordGuildRole}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := DiscordGuildRoles(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testDiscordGuildRolesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordGuildRole := &DiscordGuildRole{}
	if err = randomize.Struct(seed, discordGuildRole, discordGuildRoleDBTypes, true, discordGuildRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordGuildRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordGuildRole.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := DiscordGuildRoleExists(tx, discordGuildRole.ID)
	if err != nil {
		t.Errorf("Unable to check if DiscordGuildRole exists: %s", err)
	}
	if !e {
		t.Errorf("Expected DiscordGuildRoleExistsG to return true, but got false.")
	}
}
func testDiscordGuildRolesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordGuildRole := &DiscordGuildRole{}
	if err = randomize.Struct(seed, discordGuildRole, discordGuildRoleDBTypes, true, discordGuildRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordGuildRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordGuildRole.Insert(tx); err != nil {
		t.Error(err)
	}

	discordGuildRoleFound, err := FindDiscordGuildRole(tx, discordGuildRole.ID)
	if err != nil {
		t.Error(err)
	}

	if discordGuildRoleFound == nil {
		t.Error("want a record, got nil")
	}
}
func testDiscordGuildRolesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordGuildRole := &DiscordGuildRole{}
	if err = randomize.Struct(seed, discordGuildRole, discordGuildRoleDBTypes, true, discordGuildRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordGuildRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordGuildRole.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = DiscordGuildRoles(tx).Bind(discordGuildRole); err != nil {
		t.Error(err)
	}
}

func testDiscordGuildRolesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordGuildRole := &DiscordGuildRole{}
	if err = randomize.Struct(seed, discordGuildRole, discordGuildRoleDBTypes, true, discordGuildRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordGuildRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordGuildRole.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := DiscordGuildRoles(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testDiscordGuildRolesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordGuildRoleOne := &DiscordGuildRole{}
	discordGuildRoleTwo := &DiscordGuildRole{}
	if err = randomize.Struct(seed, discordGuildRoleOne, discordGuildRoleDBTypes, false, discordGuildRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordGuildRole struct: %s", err)
	}
	if err = randomize.Struct(seed, discordGuildRoleTwo, discordGuildRoleDBTypes, false, discordGuildRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordGuildRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordGuildRoleOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = discordGuildRoleTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := DiscordGuildRoles(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testDiscordGuildRolesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	discordGuildRoleOne := &DiscordGuildRole{}
	discordGuildRoleTwo := &DiscordGuildRole{}
	if err = randomize.Struct(seed, discordGuildRoleOne, discordGuildRoleDBTypes, false, discordGuildRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordGuildRole struct: %s", err)
	}
	if err = randomize.Struct(seed, discordGuildRoleTwo, discordGuildRoleDBTypes, false, discordGuildRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordGuildRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordGuildRoleOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = discordGuildRoleTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DiscordGuildRoles(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testDiscordGuildRolesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordGuildRole := &DiscordGuildRole{}
	if err = randomize.Struct(seed, discordGuildRole, discordGuildRoleDBTypes, true, discordGuildRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordGuildRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordGuildRole.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DiscordGuildRoles(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDiscordGuildRolesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordGuildRole := &DiscordGuildRole{}
	if err = randomize.Struct(seed, discordGuildRole, discordGuildRoleDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordGuildRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordGuildRole.Insert(tx, discordGuildRoleColumns...); err != nil {
		t.Error(err)
	}

	count, err := DiscordGuildRoles(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDiscordGuildRoleToOneDiscordGuildUsingGuild(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local DiscordGuildRole
	var foreign DiscordGuild

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, discordGuildRoleDBTypes, true, discordGuildRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordGuildRole struct: %s", err)
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

	slice := DiscordGuildRoleSlice{&local}
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

func testDiscordGuildRoleToOneSetOpDiscordGuildUsingGuild(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a DiscordGuildRole
	var b, c DiscordGuild

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, discordGuildRoleDBTypes, false, strmangle.SetComplement(discordGuildRolePrimaryKeyColumns, discordGuildRoleColumnsWithoutDefault)...); err != nil {
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

		if x.R.GuildDiscordGuildRoles[0] != &a {
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
func testDiscordGuildRolesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordGuildRole := &DiscordGuildRole{}
	if err = randomize.Struct(seed, discordGuildRole, discordGuildRoleDBTypes, true, discordGuildRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordGuildRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordGuildRole.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = discordGuildRole.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testDiscordGuildRolesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordGuildRole := &DiscordGuildRole{}
	if err = randomize.Struct(seed, discordGuildRole, discordGuildRoleDBTypes, true, discordGuildRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordGuildRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordGuildRole.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := DiscordGuildRoleSlice{discordGuildRole}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testDiscordGuildRolesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordGuildRole := &DiscordGuildRole{}
	if err = randomize.Struct(seed, discordGuildRole, discordGuildRoleDBTypes, true, discordGuildRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordGuildRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordGuildRole.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := DiscordGuildRoles(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	discordGuildRoleDBTypes = map[string]string{`Color`: `integer`, `CreatedAt`: `timestamp with time zone`, `DeletedAt`: `timestamp with time zone`, `GuildID`: `bigint`, `Hoist`: `boolean`, `ID`: `bigint`, `Managed`: `boolean`, `Mentionable`: `boolean`, `Name`: `text`, `Permissions`: `integer`, `Position`: `integer`}
	_                       = bytes.MinRead
)

func testDiscordGuildRolesUpdate(t *testing.T) {
	t.Parallel()

	if len(discordGuildRoleColumns) == len(discordGuildRolePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	discordGuildRole := &DiscordGuildRole{}
	if err = randomize.Struct(seed, discordGuildRole, discordGuildRoleDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordGuildRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordGuildRole.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DiscordGuildRoles(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, discordGuildRole, discordGuildRoleDBTypes, true, discordGuildRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordGuildRole struct: %s", err)
	}

	if err = discordGuildRole.Update(tx); err != nil {
		t.Error(err)
	}
}

func testDiscordGuildRolesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(discordGuildRoleColumns) == len(discordGuildRolePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	discordGuildRole := &DiscordGuildRole{}
	if err = randomize.Struct(seed, discordGuildRole, discordGuildRoleDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordGuildRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordGuildRole.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DiscordGuildRoles(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, discordGuildRole, discordGuildRoleDBTypes, true, discordGuildRolePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DiscordGuildRole struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(discordGuildRoleColumns, discordGuildRolePrimaryKeyColumns) {
		fields = discordGuildRoleColumns
	} else {
		fields = strmangle.SetComplement(
			discordGuildRoleColumns,
			discordGuildRolePrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(discordGuildRole))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := DiscordGuildRoleSlice{discordGuildRole}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testDiscordGuildRolesUpsert(t *testing.T) {
	t.Parallel()

	if len(discordGuildRoleColumns) == len(discordGuildRolePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	discordGuildRole := DiscordGuildRole{}
	if err = randomize.Struct(seed, &discordGuildRole, discordGuildRoleDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordGuildRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordGuildRole.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert DiscordGuildRole: %s", err)
	}

	count, err := DiscordGuildRoles(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &discordGuildRole, discordGuildRoleDBTypes, false, discordGuildRolePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DiscordGuildRole struct: %s", err)
	}

	if err = discordGuildRole.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert DiscordGuildRole: %s", err)
	}

	count, err = DiscordGuildRoles(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
