package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testDiscordUsers(t *testing.T) {
	t.Parallel()

	query := DiscordUsers(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testDiscordUsersDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordUser := &DiscordUser{}
	if err = randomize.Struct(seed, discordUser, discordUserDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordUser.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = discordUser.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := DiscordUsers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDiscordUsersQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordUser := &DiscordUser{}
	if err = randomize.Struct(seed, discordUser, discordUserDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordUser.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = DiscordUsers(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := DiscordUsers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDiscordUsersSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordUser := &DiscordUser{}
	if err = randomize.Struct(seed, discordUser, discordUserDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordUser.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := DiscordUserSlice{discordUser}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := DiscordUsers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testDiscordUsersExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordUser := &DiscordUser{}
	if err = randomize.Struct(seed, discordUser, discordUserDBTypes, true, discordUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordUser.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := DiscordUserExists(tx, discordUser.ID)
	if err != nil {
		t.Errorf("Unable to check if DiscordUser exists: %s", err)
	}
	if !e {
		t.Errorf("Expected DiscordUserExistsG to return true, but got false.")
	}
}
func testDiscordUsersFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordUser := &DiscordUser{}
	if err = randomize.Struct(seed, discordUser, discordUserDBTypes, true, discordUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordUser.Insert(tx); err != nil {
		t.Error(err)
	}

	discordUserFound, err := FindDiscordUser(tx, discordUser.ID)
	if err != nil {
		t.Error(err)
	}

	if discordUserFound == nil {
		t.Error("want a record, got nil")
	}
}
func testDiscordUsersBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordUser := &DiscordUser{}
	if err = randomize.Struct(seed, discordUser, discordUserDBTypes, true, discordUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordUser.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = DiscordUsers(tx).Bind(discordUser); err != nil {
		t.Error(err)
	}
}

func testDiscordUsersOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordUser := &DiscordUser{}
	if err = randomize.Struct(seed, discordUser, discordUserDBTypes, true, discordUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordUser.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := DiscordUsers(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testDiscordUsersAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordUserOne := &DiscordUser{}
	discordUserTwo := &DiscordUser{}
	if err = randomize.Struct(seed, discordUserOne, discordUserDBTypes, false, discordUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordUser struct: %s", err)
	}
	if err = randomize.Struct(seed, discordUserTwo, discordUserDBTypes, false, discordUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordUserOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = discordUserTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := DiscordUsers(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testDiscordUsersCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	discordUserOne := &DiscordUser{}
	discordUserTwo := &DiscordUser{}
	if err = randomize.Struct(seed, discordUserOne, discordUserDBTypes, false, discordUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordUser struct: %s", err)
	}
	if err = randomize.Struct(seed, discordUserTwo, discordUserDBTypes, false, discordUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordUserOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = discordUserTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DiscordUsers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testDiscordUsersInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordUser := &DiscordUser{}
	if err = randomize.Struct(seed, discordUser, discordUserDBTypes, true, discordUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordUser.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DiscordUsers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDiscordUsersInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordUser := &DiscordUser{}
	if err = randomize.Struct(seed, discordUser, discordUserDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordUser.Insert(tx, discordUserColumns...); err != nil {
		t.Error(err)
	}

	count, err := DiscordUsers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDiscordUserToManyRecipientDiscordPrivateChannels(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a DiscordUser
	var b, c DiscordPrivateChannel

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, discordUserDBTypes, true, discordUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordUser struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, discordPrivateChannelDBTypes, false, discordPrivateChannelColumnsWithDefault...)
	randomize.Struct(seed, &c, discordPrivateChannelDBTypes, false, discordPrivateChannelColumnsWithDefault...)

	b.RecipientID = a.ID
	c.RecipientID = a.ID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	discordPrivateChannel, err := a.RecipientDiscordPrivateChannels(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range discordPrivateChannel {
		if v.RecipientID == b.RecipientID {
			bFound = true
		}
		if v.RecipientID == c.RecipientID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := DiscordUserSlice{&a}
	if err = a.L.LoadRecipientDiscordPrivateChannels(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.RecipientDiscordPrivateChannels); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.RecipientDiscordPrivateChannels = nil
	if err = a.L.LoadRecipientDiscordPrivateChannels(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.RecipientDiscordPrivateChannels); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", discordPrivateChannel)
	}
}

func testDiscordUserToManyUserDiscordMembers(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a DiscordUser
	var b, c DiscordMember

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, discordUserDBTypes, true, discordUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordUser struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, discordMemberDBTypes, false, discordMemberColumnsWithDefault...)
	randomize.Struct(seed, &c, discordMemberDBTypes, false, discordMemberColumnsWithDefault...)

	b.UserID = a.ID
	c.UserID = a.ID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	discordMember, err := a.UserDiscordMembers(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range discordMember {
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

	slice := DiscordUserSlice{&a}
	if err = a.L.LoadUserDiscordMembers(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.UserDiscordMembers); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.UserDiscordMembers = nil
	if err = a.L.LoadUserDiscordMembers(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.UserDiscordMembers); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", discordMember)
	}
}

func testDiscordUserToManyUserDiscordMemberRoles(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a DiscordUser
	var b, c DiscordMemberRole

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, discordUserDBTypes, true, discordUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordUser struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, discordMemberRoleDBTypes, false, discordMemberRoleColumnsWithDefault...)
	randomize.Struct(seed, &c, discordMemberRoleDBTypes, false, discordMemberRoleColumnsWithDefault...)

	b.UserID = a.ID
	c.UserID = a.ID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	discordMemberRole, err := a.UserDiscordMemberRoles(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range discordMemberRole {
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

	slice := DiscordUserSlice{&a}
	if err = a.L.LoadUserDiscordMemberRoles(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.UserDiscordMemberRoles); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.UserDiscordMemberRoles = nil
	if err = a.L.LoadUserDiscordMemberRoles(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.UserDiscordMemberRoles); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", discordMemberRole)
	}
}

func testDiscordUserToManyAddOpRecipientDiscordPrivateChannels(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a DiscordUser
	var b, c, d, e DiscordPrivateChannel

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, discordUserDBTypes, false, strmangle.SetComplement(discordUserPrimaryKeyColumns, discordUserColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*DiscordPrivateChannel{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, discordPrivateChannelDBTypes, false, strmangle.SetComplement(discordPrivateChannelPrimaryKeyColumns, discordPrivateChannelColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*DiscordPrivateChannel{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddRecipientDiscordPrivateChannels(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.RecipientID {
			t.Error("foreign key was wrong value", a.ID, first.RecipientID)
		}
		if a.ID != second.RecipientID {
			t.Error("foreign key was wrong value", a.ID, second.RecipientID)
		}

		if first.R.Recipient != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Recipient != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.RecipientDiscordPrivateChannels[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.RecipientDiscordPrivateChannels[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.RecipientDiscordPrivateChannels(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testDiscordUserToManyAddOpUserDiscordMembers(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a DiscordUser
	var b, c, d, e DiscordMember

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, discordUserDBTypes, false, strmangle.SetComplement(discordUserPrimaryKeyColumns, discordUserColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*DiscordMember{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, discordMemberDBTypes, false, strmangle.SetComplement(discordMemberPrimaryKeyColumns, discordMemberColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*DiscordMember{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddUserDiscordMembers(tx, i != 0, x...)
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

		if a.R.UserDiscordMembers[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.UserDiscordMembers[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.UserDiscordMembers(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testDiscordUserToManyAddOpUserDiscordMemberRoles(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a DiscordUser
	var b, c, d, e DiscordMemberRole

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, discordUserDBTypes, false, strmangle.SetComplement(discordUserPrimaryKeyColumns, discordUserColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*DiscordMemberRole{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, discordMemberRoleDBTypes, false, strmangle.SetComplement(discordMemberRolePrimaryKeyColumns, discordMemberRoleColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*DiscordMemberRole{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddUserDiscordMemberRoles(tx, i != 0, x...)
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

		if a.R.UserDiscordMemberRoles[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.UserDiscordMemberRoles[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.UserDiscordMemberRoles(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testDiscordUsersReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordUser := &DiscordUser{}
	if err = randomize.Struct(seed, discordUser, discordUserDBTypes, true, discordUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordUser.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = discordUser.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testDiscordUsersReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordUser := &DiscordUser{}
	if err = randomize.Struct(seed, discordUser, discordUserDBTypes, true, discordUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordUser.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := DiscordUserSlice{discordUser}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testDiscordUsersSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordUser := &DiscordUser{}
	if err = randomize.Struct(seed, discordUser, discordUserDBTypes, true, discordUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordUser.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := DiscordUsers(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	discordUserDBTypes = map[string]string{`Avatar`: `text`, `Bot`: `boolean`, `CreatedAt`: `timestamp with time zone`, `Discriminator`: `character varying`, `GameName`: `text`, `GameType`: `integer`, `GameURL`: `text`, `ID`: `bigint`, `Status`: `text`, `Username`: `character varying`}
	_                  = bytes.MinRead
)

func testDiscordUsersUpdate(t *testing.T) {
	t.Parallel()

	if len(discordUserColumns) == len(discordUserPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	discordUser := &DiscordUser{}
	if err = randomize.Struct(seed, discordUser, discordUserDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordUser.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DiscordUsers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, discordUser, discordUserDBTypes, true, discordUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordUser struct: %s", err)
	}

	if err = discordUser.Update(tx); err != nil {
		t.Error(err)
	}
}

func testDiscordUsersSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(discordUserColumns) == len(discordUserPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	discordUser := &DiscordUser{}
	if err = randomize.Struct(seed, discordUser, discordUserDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordUser.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DiscordUsers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, discordUser, discordUserDBTypes, true, discordUserPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DiscordUser struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(discordUserColumns, discordUserPrimaryKeyColumns) {
		fields = discordUserColumns
	} else {
		fields = strmangle.SetComplement(
			discordUserColumns,
			discordUserPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(discordUser))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := DiscordUserSlice{discordUser}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testDiscordUsersUpsert(t *testing.T) {
	t.Parallel()

	if len(discordUserColumns) == len(discordUserPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	discordUser := DiscordUser{}
	if err = randomize.Struct(seed, &discordUser, discordUserDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordUser.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert DiscordUser: %s", err)
	}

	count, err := DiscordUsers(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &discordUser, discordUserDBTypes, false, discordUserPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DiscordUser struct: %s", err)
	}

	if err = discordUser.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert DiscordUser: %s", err)
	}

	count, err = DiscordUsers(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
