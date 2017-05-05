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

func testDiscordGuildToManyGuildDiscordGuildChannels(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a DiscordGuild
	var b, c DiscordGuildChannel

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, discordGuildDBTypes, true, discordGuildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordGuild struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, discordGuildChannelDBTypes, false, discordGuildChannelColumnsWithDefault...)
	randomize.Struct(seed, &c, discordGuildChannelDBTypes, false, discordGuildChannelColumnsWithDefault...)

	b.GuildID = a.ID
	c.GuildID = a.ID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	discordGuildChannel, err := a.GuildDiscordGuildChannels(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range discordGuildChannel {
		if v.GuildID == b.GuildID {
			bFound = true
		}
		if v.GuildID == c.GuildID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := DiscordGuildSlice{&a}
	if err = a.L.LoadGuildDiscordGuildChannels(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.GuildDiscordGuildChannels); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.GuildDiscordGuildChannels = nil
	if err = a.L.LoadGuildDiscordGuildChannels(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.GuildDiscordGuildChannels); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", discordGuildChannel)
	}
}

func testDiscordGuildToManyGuildDiscordMembers(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a DiscordGuild
	var b, c DiscordMember

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, discordGuildDBTypes, true, discordGuildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordGuild struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, discordMemberDBTypes, false, discordMemberColumnsWithDefault...)
	randomize.Struct(seed, &c, discordMemberDBTypes, false, discordMemberColumnsWithDefault...)

	b.GuildID = a.ID
	c.GuildID = a.ID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	discordMember, err := a.GuildDiscordMembers(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range discordMember {
		if v.GuildID == b.GuildID {
			bFound = true
		}
		if v.GuildID == c.GuildID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := DiscordGuildSlice{&a}
	if err = a.L.LoadGuildDiscordMembers(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.GuildDiscordMembers); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.GuildDiscordMembers = nil
	if err = a.L.LoadGuildDiscordMembers(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.GuildDiscordMembers); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", discordMember)
	}
}

func testDiscordGuildToManyGuildDiscordGuildRoles(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a DiscordGuild
	var b, c DiscordGuildRole

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, discordGuildDBTypes, true, discordGuildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordGuild struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, discordGuildRoleDBTypes, false, discordGuildRoleColumnsWithDefault...)
	randomize.Struct(seed, &c, discordGuildRoleDBTypes, false, discordGuildRoleColumnsWithDefault...)

	b.GuildID = a.ID
	c.GuildID = a.ID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	discordGuildRole, err := a.GuildDiscordGuildRoles(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range discordGuildRole {
		if v.GuildID == b.GuildID {
			bFound = true
		}
		if v.GuildID == c.GuildID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := DiscordGuildSlice{&a}
	if err = a.L.LoadGuildDiscordGuildRoles(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.GuildDiscordGuildRoles); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.GuildDiscordGuildRoles = nil
	if err = a.L.LoadGuildDiscordGuildRoles(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.GuildDiscordGuildRoles); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", discordGuildRole)
	}
}

func testDiscordGuildToManyGuildDiscordMemberRoles(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a DiscordGuild
	var b, c DiscordMemberRole

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, discordGuildDBTypes, true, discordGuildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordGuild struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, discordMemberRoleDBTypes, false, discordMemberRoleColumnsWithDefault...)
	randomize.Struct(seed, &c, discordMemberRoleDBTypes, false, discordMemberRoleColumnsWithDefault...)

	b.GuildID = a.ID
	c.GuildID = a.ID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	discordMemberRole, err := a.GuildDiscordMemberRoles(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range discordMemberRole {
		if v.GuildID == b.GuildID {
			bFound = true
		}
		if v.GuildID == c.GuildID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := DiscordGuildSlice{&a}
	if err = a.L.LoadGuildDiscordMemberRoles(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.GuildDiscordMemberRoles); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.GuildDiscordMemberRoles = nil
	if err = a.L.LoadGuildDiscordMemberRoles(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.GuildDiscordMemberRoles); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", discordMemberRole)
	}
}

func testDiscordGuildToManyAddOpGuildDiscordGuildChannels(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a DiscordGuild
	var b, c, d, e DiscordGuildChannel

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, discordGuildDBTypes, false, strmangle.SetComplement(discordGuildPrimaryKeyColumns, discordGuildColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*DiscordGuildChannel{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, discordGuildChannelDBTypes, false, strmangle.SetComplement(discordGuildChannelPrimaryKeyColumns, discordGuildChannelColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*DiscordGuildChannel{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddGuildDiscordGuildChannels(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.GuildID {
			t.Error("foreign key was wrong value", a.ID, first.GuildID)
		}
		if a.ID != second.GuildID {
			t.Error("foreign key was wrong value", a.ID, second.GuildID)
		}

		if first.R.Guild != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Guild != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.GuildDiscordGuildChannels[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.GuildDiscordGuildChannels[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.GuildDiscordGuildChannels(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testDiscordGuildToManyAddOpGuildDiscordMembers(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a DiscordGuild
	var b, c, d, e DiscordMember

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, discordGuildDBTypes, false, strmangle.SetComplement(discordGuildPrimaryKeyColumns, discordGuildColumnsWithoutDefault)...); err != nil {
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
		err = a.AddGuildDiscordMembers(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.GuildID {
			t.Error("foreign key was wrong value", a.ID, first.GuildID)
		}
		if a.ID != second.GuildID {
			t.Error("foreign key was wrong value", a.ID, second.GuildID)
		}

		if first.R.Guild != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Guild != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.GuildDiscordMembers[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.GuildDiscordMembers[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.GuildDiscordMembers(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testDiscordGuildToManyAddOpGuildDiscordGuildRoles(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a DiscordGuild
	var b, c, d, e DiscordGuildRole

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, discordGuildDBTypes, false, strmangle.SetComplement(discordGuildPrimaryKeyColumns, discordGuildColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*DiscordGuildRole{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, discordGuildRoleDBTypes, false, strmangle.SetComplement(discordGuildRolePrimaryKeyColumns, discordGuildRoleColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*DiscordGuildRole{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddGuildDiscordGuildRoles(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.GuildID {
			t.Error("foreign key was wrong value", a.ID, first.GuildID)
		}
		if a.ID != second.GuildID {
			t.Error("foreign key was wrong value", a.ID, second.GuildID)
		}

		if first.R.Guild != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Guild != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.GuildDiscordGuildRoles[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.GuildDiscordGuildRoles[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.GuildDiscordGuildRoles(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testDiscordGuildToManyAddOpGuildDiscordMemberRoles(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a DiscordGuild
	var b, c, d, e DiscordMemberRole

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, discordGuildDBTypes, false, strmangle.SetComplement(discordGuildPrimaryKeyColumns, discordGuildColumnsWithoutDefault)...); err != nil {
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
		err = a.AddGuildDiscordMemberRoles(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.GuildID {
			t.Error("foreign key was wrong value", a.ID, first.GuildID)
		}
		if a.ID != second.GuildID {
			t.Error("foreign key was wrong value", a.ID, second.GuildID)
		}

		if first.R.Guild != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Guild != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.GuildDiscordMemberRoles[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.GuildDiscordMemberRoles[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.GuildDiscordMemberRoles(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
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
