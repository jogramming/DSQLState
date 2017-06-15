package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testDMembers(t *testing.T) {
	t.Parallel()

	query := DMembers(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testDMembersDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dMember := &DMember{}
	if err = randomize.Struct(seed, dMember, dMemberDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DMember struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dMember.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = dMember.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := DMembers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDMembersQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dMember := &DMember{}
	if err = randomize.Struct(seed, dMember, dMemberDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DMember struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dMember.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = DMembers(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := DMembers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDMembersSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dMember := &DMember{}
	if err = randomize.Struct(seed, dMember, dMemberDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DMember struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dMember.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := DMemberSlice{dMember}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := DMembers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testDMembersExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dMember := &DMember{}
	if err = randomize.Struct(seed, dMember, dMemberDBTypes, true, dMemberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DMember struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dMember.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := DMemberExists(tx, dMember.UserID, dMember.GuildID)
	if err != nil {
		t.Errorf("Unable to check if DMember exists: %s", err)
	}
	if !e {
		t.Errorf("Expected DMemberExistsG to return true, but got false.")
	}
}
func testDMembersFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dMember := &DMember{}
	if err = randomize.Struct(seed, dMember, dMemberDBTypes, true, dMemberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DMember struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dMember.Insert(tx); err != nil {
		t.Error(err)
	}

	dMemberFound, err := FindDMember(tx, dMember.UserID, dMember.GuildID)
	if err != nil {
		t.Error(err)
	}

	if dMemberFound == nil {
		t.Error("want a record, got nil")
	}
}
func testDMembersBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dMember := &DMember{}
	if err = randomize.Struct(seed, dMember, dMemberDBTypes, true, dMemberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DMember struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dMember.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = DMembers(tx).Bind(dMember); err != nil {
		t.Error(err)
	}
}

func testDMembersOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dMember := &DMember{}
	if err = randomize.Struct(seed, dMember, dMemberDBTypes, true, dMemberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DMember struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dMember.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := DMembers(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testDMembersAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dMemberOne := &DMember{}
	dMemberTwo := &DMember{}
	if err = randomize.Struct(seed, dMemberOne, dMemberDBTypes, false, dMemberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DMember struct: %s", err)
	}
	if err = randomize.Struct(seed, dMemberTwo, dMemberDBTypes, false, dMemberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DMember struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dMemberOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = dMemberTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := DMembers(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testDMembersCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	dMemberOne := &DMember{}
	dMemberTwo := &DMember{}
	if err = randomize.Struct(seed, dMemberOne, dMemberDBTypes, false, dMemberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DMember struct: %s", err)
	}
	if err = randomize.Struct(seed, dMemberTwo, dMemberDBTypes, false, dMemberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DMember struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dMemberOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = dMemberTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DMembers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testDMembersInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dMember := &DMember{}
	if err = randomize.Struct(seed, dMember, dMemberDBTypes, true, dMemberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DMember struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dMember.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DMembers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDMembersInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dMember := &DMember{}
	if err = randomize.Struct(seed, dMember, dMemberDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DMember struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dMember.Insert(tx, dMemberColumns...); err != nil {
		t.Error(err)
	}

	count, err := DMembers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDMemberToOneDUserUsingUser(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local DMember
	var foreign DUser

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, dMemberDBTypes, true, dMemberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DMember struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, dUserDBTypes, true, dUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DUser struct: %s", err)
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

	slice := DMemberSlice{&local}
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

func testDMemberToOneSetOpDUserUsingUser(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a DMember
	var b, c DUser

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, dMemberDBTypes, false, strmangle.SetComplement(dMemberPrimaryKeyColumns, dMemberColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, dUserDBTypes, false, strmangle.SetComplement(dUserPrimaryKeyColumns, dUserColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, dUserDBTypes, false, strmangle.SetComplement(dUserPrimaryKeyColumns, dUserColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*DUser{&b, &c} {
		err = a.SetUser(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.User != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.UserDMembers[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.UserID != x.ID {
			t.Error("foreign key was wrong value", a.UserID)
		}

		if exists, err := DMemberExists(tx, a.UserID, a.GuildID); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'a' to exist")
		}

	}
}
func testDMembersReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dMember := &DMember{}
	if err = randomize.Struct(seed, dMember, dMemberDBTypes, true, dMemberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DMember struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dMember.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = dMember.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testDMembersReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dMember := &DMember{}
	if err = randomize.Struct(seed, dMember, dMemberDBTypes, true, dMemberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DMember struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dMember.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := DMemberSlice{dMember}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testDMembersSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dMember := &DMember{}
	if err = randomize.Struct(seed, dMember, dMemberDBTypes, true, dMemberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DMember struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dMember.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := DMembers(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	dMemberDBTypes = map[string]string{`CreatedAt`: `timestamp with time zone`, `Deaf`: `boolean`, `GuildID`: `bigint`, `JoinedAt`: `timestamp with time zone`, `LeftAt`: `timestamp with time zone`, `Mute`: `boolean`, `Nick`: `character varying`, `Roles`: `ARRAYbigint`, `Synced`: `boolean`, `UserID`: `bigint`}
	_              = bytes.MinRead
)

func testDMembersUpdate(t *testing.T) {
	t.Parallel()

	if len(dMemberColumns) == len(dMemberPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	dMember := &DMember{}
	if err = randomize.Struct(seed, dMember, dMemberDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DMember struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dMember.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DMembers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, dMember, dMemberDBTypes, true, dMemberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DMember struct: %s", err)
	}

	if err = dMember.Update(tx); err != nil {
		t.Error(err)
	}
}

func testDMembersSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(dMemberColumns) == len(dMemberPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	dMember := &DMember{}
	if err = randomize.Struct(seed, dMember, dMemberDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DMember struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dMember.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DMembers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, dMember, dMemberDBTypes, true, dMemberPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DMember struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(dMemberColumns, dMemberPrimaryKeyColumns) {
		fields = dMemberColumns
	} else {
		fields = strmangle.SetComplement(
			dMemberColumns,
			dMemberPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(dMember))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := DMemberSlice{dMember}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testDMembersUpsert(t *testing.T) {
	t.Parallel()

	if len(dMemberColumns) == len(dMemberPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	dMember := DMember{}
	if err = randomize.Struct(seed, &dMember, dMemberDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DMember struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dMember.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert DMember: %s", err)
	}

	count, err := DMembers(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &dMember, dMemberDBTypes, false, dMemberPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DMember struct: %s", err)
	}

	if err = dMember.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert DMember: %s", err)
	}

	count, err = DMembers(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
