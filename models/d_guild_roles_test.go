package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testDGuildRoles(t *testing.T) {
	t.Parallel()

	query := DGuildRoles(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testDGuildRolesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dGuildRole := &DGuildRole{}
	if err = randomize.Struct(seed, dGuildRole, dGuildRoleDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DGuildRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dGuildRole.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = dGuildRole.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := DGuildRoles(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDGuildRolesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dGuildRole := &DGuildRole{}
	if err = randomize.Struct(seed, dGuildRole, dGuildRoleDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DGuildRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dGuildRole.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = DGuildRoles(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := DGuildRoles(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDGuildRolesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dGuildRole := &DGuildRole{}
	if err = randomize.Struct(seed, dGuildRole, dGuildRoleDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DGuildRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dGuildRole.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := DGuildRoleSlice{dGuildRole}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := DGuildRoles(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testDGuildRolesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dGuildRole := &DGuildRole{}
	if err = randomize.Struct(seed, dGuildRole, dGuildRoleDBTypes, true, dGuildRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DGuildRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dGuildRole.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := DGuildRoleExists(tx, dGuildRole.ID)
	if err != nil {
		t.Errorf("Unable to check if DGuildRole exists: %s", err)
	}
	if !e {
		t.Errorf("Expected DGuildRoleExistsG to return true, but got false.")
	}
}
func testDGuildRolesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dGuildRole := &DGuildRole{}
	if err = randomize.Struct(seed, dGuildRole, dGuildRoleDBTypes, true, dGuildRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DGuildRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dGuildRole.Insert(tx); err != nil {
		t.Error(err)
	}

	dGuildRoleFound, err := FindDGuildRole(tx, dGuildRole.ID)
	if err != nil {
		t.Error(err)
	}

	if dGuildRoleFound == nil {
		t.Error("want a record, got nil")
	}
}
func testDGuildRolesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dGuildRole := &DGuildRole{}
	if err = randomize.Struct(seed, dGuildRole, dGuildRoleDBTypes, true, dGuildRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DGuildRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dGuildRole.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = DGuildRoles(tx).Bind(dGuildRole); err != nil {
		t.Error(err)
	}
}

func testDGuildRolesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dGuildRole := &DGuildRole{}
	if err = randomize.Struct(seed, dGuildRole, dGuildRoleDBTypes, true, dGuildRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DGuildRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dGuildRole.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := DGuildRoles(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testDGuildRolesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dGuildRoleOne := &DGuildRole{}
	dGuildRoleTwo := &DGuildRole{}
	if err = randomize.Struct(seed, dGuildRoleOne, dGuildRoleDBTypes, false, dGuildRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DGuildRole struct: %s", err)
	}
	if err = randomize.Struct(seed, dGuildRoleTwo, dGuildRoleDBTypes, false, dGuildRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DGuildRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dGuildRoleOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = dGuildRoleTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := DGuildRoles(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testDGuildRolesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	dGuildRoleOne := &DGuildRole{}
	dGuildRoleTwo := &DGuildRole{}
	if err = randomize.Struct(seed, dGuildRoleOne, dGuildRoleDBTypes, false, dGuildRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DGuildRole struct: %s", err)
	}
	if err = randomize.Struct(seed, dGuildRoleTwo, dGuildRoleDBTypes, false, dGuildRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DGuildRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dGuildRoleOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = dGuildRoleTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DGuildRoles(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testDGuildRolesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dGuildRole := &DGuildRole{}
	if err = randomize.Struct(seed, dGuildRole, dGuildRoleDBTypes, true, dGuildRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DGuildRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dGuildRole.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DGuildRoles(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDGuildRolesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dGuildRole := &DGuildRole{}
	if err = randomize.Struct(seed, dGuildRole, dGuildRoleDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DGuildRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dGuildRole.Insert(tx, dGuildRoleColumns...); err != nil {
		t.Error(err)
	}

	count, err := DGuildRoles(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDGuildRolesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dGuildRole := &DGuildRole{}
	if err = randomize.Struct(seed, dGuildRole, dGuildRoleDBTypes, true, dGuildRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DGuildRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dGuildRole.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = dGuildRole.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testDGuildRolesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dGuildRole := &DGuildRole{}
	if err = randomize.Struct(seed, dGuildRole, dGuildRoleDBTypes, true, dGuildRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DGuildRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dGuildRole.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := DGuildRoleSlice{dGuildRole}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testDGuildRolesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dGuildRole := &DGuildRole{}
	if err = randomize.Struct(seed, dGuildRole, dGuildRoleDBTypes, true, dGuildRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DGuildRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dGuildRole.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := DGuildRoles(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	dGuildRoleDBTypes = map[string]string{`Color`: `integer`, `CreatedAt`: `timestamp with time zone`, `DeletedAt`: `timestamp with time zone`, `GuildID`: `bigint`, `Hoist`: `boolean`, `ID`: `bigint`, `Managed`: `boolean`, `Mentionable`: `boolean`, `Name`: `text`, `Permissions`: `integer`, `Position`: `integer`, `Synced`: `boolean`}
	_                 = bytes.MinRead
)

func testDGuildRolesUpdate(t *testing.T) {
	t.Parallel()

	if len(dGuildRoleColumns) == len(dGuildRolePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	dGuildRole := &DGuildRole{}
	if err = randomize.Struct(seed, dGuildRole, dGuildRoleDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DGuildRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dGuildRole.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DGuildRoles(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, dGuildRole, dGuildRoleDBTypes, true, dGuildRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DGuildRole struct: %s", err)
	}

	if err = dGuildRole.Update(tx); err != nil {
		t.Error(err)
	}
}

func testDGuildRolesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(dGuildRoleColumns) == len(dGuildRolePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	dGuildRole := &DGuildRole{}
	if err = randomize.Struct(seed, dGuildRole, dGuildRoleDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DGuildRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dGuildRole.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DGuildRoles(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, dGuildRole, dGuildRoleDBTypes, true, dGuildRolePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DGuildRole struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(dGuildRoleColumns, dGuildRolePrimaryKeyColumns) {
		fields = dGuildRoleColumns
	} else {
		fields = strmangle.SetComplement(
			dGuildRoleColumns,
			dGuildRolePrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(dGuildRole))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := DGuildRoleSlice{dGuildRole}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testDGuildRolesUpsert(t *testing.T) {
	t.Parallel()

	if len(dGuildRoleColumns) == len(dGuildRolePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	dGuildRole := DGuildRole{}
	if err = randomize.Struct(seed, &dGuildRole, dGuildRoleDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DGuildRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dGuildRole.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert DGuildRole: %s", err)
	}

	count, err := DGuildRoles(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &dGuildRole, dGuildRoleDBTypes, false, dGuildRolePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DGuildRole struct: %s", err)
	}

	if err = dGuildRole.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert DGuildRole: %s", err)
	}

	count, err = DGuildRoles(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
