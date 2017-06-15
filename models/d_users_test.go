package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testDUsers(t *testing.T) {
	t.Parallel()

	query := DUsers(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testDUsersDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dUser := &DUser{}
	if err = randomize.Struct(seed, dUser, dUserDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dUser.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = dUser.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := DUsers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDUsersQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dUser := &DUser{}
	if err = randomize.Struct(seed, dUser, dUserDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dUser.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = DUsers(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := DUsers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDUsersSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dUser := &DUser{}
	if err = randomize.Struct(seed, dUser, dUserDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dUser.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := DUserSlice{dUser}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := DUsers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testDUsersExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dUser := &DUser{}
	if err = randomize.Struct(seed, dUser, dUserDBTypes, true, dUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dUser.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := DUserExists(tx, dUser.ID)
	if err != nil {
		t.Errorf("Unable to check if DUser exists: %s", err)
	}
	if !e {
		t.Errorf("Expected DUserExistsG to return true, but got false.")
	}
}
func testDUsersFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dUser := &DUser{}
	if err = randomize.Struct(seed, dUser, dUserDBTypes, true, dUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dUser.Insert(tx); err != nil {
		t.Error(err)
	}

	dUserFound, err := FindDUser(tx, dUser.ID)
	if err != nil {
		t.Error(err)
	}

	if dUserFound == nil {
		t.Error("want a record, got nil")
	}
}
func testDUsersBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dUser := &DUser{}
	if err = randomize.Struct(seed, dUser, dUserDBTypes, true, dUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dUser.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = DUsers(tx).Bind(dUser); err != nil {
		t.Error(err)
	}
}

func testDUsersOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dUser := &DUser{}
	if err = randomize.Struct(seed, dUser, dUserDBTypes, true, dUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dUser.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := DUsers(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testDUsersAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dUserOne := &DUser{}
	dUserTwo := &DUser{}
	if err = randomize.Struct(seed, dUserOne, dUserDBTypes, false, dUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DUser struct: %s", err)
	}
	if err = randomize.Struct(seed, dUserTwo, dUserDBTypes, false, dUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dUserOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = dUserTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := DUsers(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testDUsersCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	dUserOne := &DUser{}
	dUserTwo := &DUser{}
	if err = randomize.Struct(seed, dUserOne, dUserDBTypes, false, dUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DUser struct: %s", err)
	}
	if err = randomize.Struct(seed, dUserTwo, dUserDBTypes, false, dUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dUserOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = dUserTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DUsers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testDUsersInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dUser := &DUser{}
	if err = randomize.Struct(seed, dUser, dUserDBTypes, true, dUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dUser.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DUsers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDUsersInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dUser := &DUser{}
	if err = randomize.Struct(seed, dUser, dUserDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dUser.Insert(tx, dUserColumns...); err != nil {
		t.Error(err)
	}

	count, err := DUsers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDUserToManyUserDMembers(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a DUser
	var b, c DMember

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, dUserDBTypes, true, dUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DUser struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, dMemberDBTypes, false, dMemberColumnsWithDefault...)
	randomize.Struct(seed, &c, dMemberDBTypes, false, dMemberColumnsWithDefault...)

	b.UserID = a.ID
	c.UserID = a.ID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	dMember, err := a.UserDMembers(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range dMember {
		if v.UserID == b.UserID {
			bFound = true
		}
		if v.UserID == c.UserID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := DUserSlice{&a}
	if err = a.L.LoadUserDMembers(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.UserDMembers); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.UserDMembers = nil
	if err = a.L.LoadUserDMembers(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.UserDMembers); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", dMember)
	}
}

func testDUserToManyAddOpUserDMembers(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a DUser
	var b, c, d, e DMember

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, dUserDBTypes, false, strmangle.SetComplement(dUserPrimaryKeyColumns, dUserColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*DMember{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, dMemberDBTypes, false, strmangle.SetComplement(dMemberPrimaryKeyColumns, dMemberColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*DMember{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddUserDMembers(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.UserID {
			t.Error("foreign key was wrong value", a.ID, first.UserID)
		}
		if a.ID != second.UserID {
			t.Error("foreign key was wrong value", a.ID, second.UserID)
		}

		if first.R.User != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.User != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.UserDMembers[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.UserDMembers[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.UserDMembers(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testDUsersReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dUser := &DUser{}
	if err = randomize.Struct(seed, dUser, dUserDBTypes, true, dUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dUser.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = dUser.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testDUsersReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dUser := &DUser{}
	if err = randomize.Struct(seed, dUser, dUserDBTypes, true, dUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dUser.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := DUserSlice{dUser}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testDUsersSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dUser := &DUser{}
	if err = randomize.Struct(seed, dUser, dUserDBTypes, true, dUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dUser.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := DUsers(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	dUserDBTypes = map[string]string{`Avatar`: `text`, `Bot`: `boolean`, `CreatedAt`: `timestamp with time zone`, `Discriminator`: `character varying`, `GameName`: `text`, `GameType`: `integer`, `GameURL`: `text`, `ID`: `bigint`, `Status`: `text`, `Username`: `character varying`}
	_            = bytes.MinRead
)

func testDUsersUpdate(t *testing.T) {
	t.Parallel()

	if len(dUserColumns) == len(dUserPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	dUser := &DUser{}
	if err = randomize.Struct(seed, dUser, dUserDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dUser.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DUsers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, dUser, dUserDBTypes, true, dUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DUser struct: %s", err)
	}

	if err = dUser.Update(tx); err != nil {
		t.Error(err)
	}
}

func testDUsersSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(dUserColumns) == len(dUserPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	dUser := &DUser{}
	if err = randomize.Struct(seed, dUser, dUserDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dUser.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DUsers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, dUser, dUserDBTypes, true, dUserPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DUser struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(dUserColumns, dUserPrimaryKeyColumns) {
		fields = dUserColumns
	} else {
		fields = strmangle.SetComplement(
			dUserColumns,
			dUserPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(dUser))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := DUserSlice{dUser}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testDUsersUpsert(t *testing.T) {
	t.Parallel()

	if len(dUserColumns) == len(dUserPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	dUser := DUser{}
	if err = randomize.Struct(seed, &dUser, dUserDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dUser.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert DUser: %s", err)
	}

	count, err := DUsers(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &dUser, dUserDBTypes, false, dUserPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DUser struct: %s", err)
	}

	if err = dUser.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert DUser: %s", err)
	}

	count, err = DUsers(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
