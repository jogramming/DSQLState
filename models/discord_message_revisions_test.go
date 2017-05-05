package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testDiscordMessageRevisions(t *testing.T) {
	t.Parallel()

	query := DiscordMessageRevisions(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testDiscordMessageRevisionsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordMessageRevision := &DiscordMessageRevision{}
	if err = randomize.Struct(seed, discordMessageRevision, discordMessageRevisionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordMessageRevision struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordMessageRevision.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = discordMessageRevision.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := DiscordMessageRevisions(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDiscordMessageRevisionsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordMessageRevision := &DiscordMessageRevision{}
	if err = randomize.Struct(seed, discordMessageRevision, discordMessageRevisionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordMessageRevision struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordMessageRevision.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = DiscordMessageRevisions(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := DiscordMessageRevisions(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDiscordMessageRevisionsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordMessageRevision := &DiscordMessageRevision{}
	if err = randomize.Struct(seed, discordMessageRevision, discordMessageRevisionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordMessageRevision struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordMessageRevision.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := DiscordMessageRevisionSlice{discordMessageRevision}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := DiscordMessageRevisions(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testDiscordMessageRevisionsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordMessageRevision := &DiscordMessageRevision{}
	if err = randomize.Struct(seed, discordMessageRevision, discordMessageRevisionDBTypes, true, discordMessageRevisionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordMessageRevision struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordMessageRevision.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := DiscordMessageRevisionExists(tx, discordMessageRevision.ID)
	if err != nil {
		t.Errorf("Unable to check if DiscordMessageRevision exists: %s", err)
	}
	if !e {
		t.Errorf("Expected DiscordMessageRevisionExistsG to return true, but got false.")
	}
}
func testDiscordMessageRevisionsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordMessageRevision := &DiscordMessageRevision{}
	if err = randomize.Struct(seed, discordMessageRevision, discordMessageRevisionDBTypes, true, discordMessageRevisionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordMessageRevision struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordMessageRevision.Insert(tx); err != nil {
		t.Error(err)
	}

	discordMessageRevisionFound, err := FindDiscordMessageRevision(tx, discordMessageRevision.ID)
	if err != nil {
		t.Error(err)
	}

	if discordMessageRevisionFound == nil {
		t.Error("want a record, got nil")
	}
}
func testDiscordMessageRevisionsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordMessageRevision := &DiscordMessageRevision{}
	if err = randomize.Struct(seed, discordMessageRevision, discordMessageRevisionDBTypes, true, discordMessageRevisionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordMessageRevision struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordMessageRevision.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = DiscordMessageRevisions(tx).Bind(discordMessageRevision); err != nil {
		t.Error(err)
	}
}

func testDiscordMessageRevisionsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordMessageRevision := &DiscordMessageRevision{}
	if err = randomize.Struct(seed, discordMessageRevision, discordMessageRevisionDBTypes, true, discordMessageRevisionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordMessageRevision struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordMessageRevision.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := DiscordMessageRevisions(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testDiscordMessageRevisionsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordMessageRevisionOne := &DiscordMessageRevision{}
	discordMessageRevisionTwo := &DiscordMessageRevision{}
	if err = randomize.Struct(seed, discordMessageRevisionOne, discordMessageRevisionDBTypes, false, discordMessageRevisionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordMessageRevision struct: %s", err)
	}
	if err = randomize.Struct(seed, discordMessageRevisionTwo, discordMessageRevisionDBTypes, false, discordMessageRevisionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordMessageRevision struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordMessageRevisionOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = discordMessageRevisionTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := DiscordMessageRevisions(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testDiscordMessageRevisionsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	discordMessageRevisionOne := &DiscordMessageRevision{}
	discordMessageRevisionTwo := &DiscordMessageRevision{}
	if err = randomize.Struct(seed, discordMessageRevisionOne, discordMessageRevisionDBTypes, false, discordMessageRevisionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordMessageRevision struct: %s", err)
	}
	if err = randomize.Struct(seed, discordMessageRevisionTwo, discordMessageRevisionDBTypes, false, discordMessageRevisionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordMessageRevision struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordMessageRevisionOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = discordMessageRevisionTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DiscordMessageRevisions(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testDiscordMessageRevisionsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordMessageRevision := &DiscordMessageRevision{}
	if err = randomize.Struct(seed, discordMessageRevision, discordMessageRevisionDBTypes, true, discordMessageRevisionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordMessageRevision struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordMessageRevision.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DiscordMessageRevisions(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDiscordMessageRevisionsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordMessageRevision := &DiscordMessageRevision{}
	if err = randomize.Struct(seed, discordMessageRevision, discordMessageRevisionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordMessageRevision struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordMessageRevision.Insert(tx, discordMessageRevisionColumns...); err != nil {
		t.Error(err)
	}

	count, err := DiscordMessageRevisions(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDiscordMessageRevisionToManyRevisionDiscordMessageEmbeds(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a DiscordMessageRevision
	var b, c DiscordMessageEmbed

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, discordMessageRevisionDBTypes, true, discordMessageRevisionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordMessageRevision struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, discordMessageEmbedDBTypes, false, discordMessageEmbedColumnsWithDefault...)
	randomize.Struct(seed, &c, discordMessageEmbedDBTypes, false, discordMessageEmbedColumnsWithDefault...)

	b.RevisionID.Valid = true
	c.RevisionID.Valid = true
	b.RevisionID.Int64 = a.ID
	c.RevisionID.Int64 = a.ID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	discordMessageEmbed, err := a.RevisionDiscordMessageEmbeds(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range discordMessageEmbed {
		if v.RevisionID.Int64 == b.RevisionID.Int64 {
			bFound = true
		}
		if v.RevisionID.Int64 == c.RevisionID.Int64 {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := DiscordMessageRevisionSlice{&a}
	if err = a.L.LoadRevisionDiscordMessageEmbeds(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.RevisionDiscordMessageEmbeds); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.RevisionDiscordMessageEmbeds = nil
	if err = a.L.LoadRevisionDiscordMessageEmbeds(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.RevisionDiscordMessageEmbeds); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", discordMessageEmbed)
	}
}

func testDiscordMessageRevisionToManyAddOpRevisionDiscordMessageEmbeds(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a DiscordMessageRevision
	var b, c, d, e DiscordMessageEmbed

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, discordMessageRevisionDBTypes, false, strmangle.SetComplement(discordMessageRevisionPrimaryKeyColumns, discordMessageRevisionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*DiscordMessageEmbed{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, discordMessageEmbedDBTypes, false, strmangle.SetComplement(discordMessageEmbedPrimaryKeyColumns, discordMessageEmbedColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*DiscordMessageEmbed{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddRevisionDiscordMessageEmbeds(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.RevisionID.Int64 {
			t.Error("foreign key was wrong value", a.ID, first.RevisionID.Int64)
		}
		if a.ID != second.RevisionID.Int64 {
			t.Error("foreign key was wrong value", a.ID, second.RevisionID.Int64)
		}

		if first.R.Revision != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Revision != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.RevisionDiscordMessageEmbeds[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.RevisionDiscordMessageEmbeds[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.RevisionDiscordMessageEmbeds(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testDiscordMessageRevisionToManySetOpRevisionDiscordMessageEmbeds(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a DiscordMessageRevision
	var b, c, d, e DiscordMessageEmbed

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, discordMessageRevisionDBTypes, false, strmangle.SetComplement(discordMessageRevisionPrimaryKeyColumns, discordMessageRevisionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*DiscordMessageEmbed{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, discordMessageEmbedDBTypes, false, strmangle.SetComplement(discordMessageEmbedPrimaryKeyColumns, discordMessageEmbedColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err = a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	err = a.SetRevisionDiscordMessageEmbeds(tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.RevisionDiscordMessageEmbeds(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetRevisionDiscordMessageEmbeds(tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.RevisionDiscordMessageEmbeds(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if b.RevisionID.Valid {
		t.Error("want b's foreign key value to be nil")
	}
	if c.RevisionID.Valid {
		t.Error("want c's foreign key value to be nil")
	}
	if a.ID != d.RevisionID.Int64 {
		t.Error("foreign key was wrong value", a.ID, d.RevisionID.Int64)
	}
	if a.ID != e.RevisionID.Int64 {
		t.Error("foreign key was wrong value", a.ID, e.RevisionID.Int64)
	}

	if b.R.Revision != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Revision != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Revision != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.Revision != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.RevisionDiscordMessageEmbeds[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.RevisionDiscordMessageEmbeds[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testDiscordMessageRevisionToManyRemoveOpRevisionDiscordMessageEmbeds(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a DiscordMessageRevision
	var b, c, d, e DiscordMessageEmbed

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, discordMessageRevisionDBTypes, false, strmangle.SetComplement(discordMessageRevisionPrimaryKeyColumns, discordMessageRevisionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*DiscordMessageEmbed{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, discordMessageEmbedDBTypes, false, strmangle.SetComplement(discordMessageEmbedPrimaryKeyColumns, discordMessageEmbedColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	err = a.AddRevisionDiscordMessageEmbeds(tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.RevisionDiscordMessageEmbeds(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveRevisionDiscordMessageEmbeds(tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.RevisionDiscordMessageEmbeds(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if b.RevisionID.Valid {
		t.Error("want b's foreign key value to be nil")
	}
	if c.RevisionID.Valid {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.Revision != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Revision != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Revision != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.Revision != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.RevisionDiscordMessageEmbeds) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.RevisionDiscordMessageEmbeds[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.RevisionDiscordMessageEmbeds[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testDiscordMessageRevisionToOneDiscordMessageUsingMessage(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local DiscordMessageRevision
	var foreign DiscordMessage

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, discordMessageRevisionDBTypes, true, discordMessageRevisionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordMessageRevision struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, discordMessageDBTypes, true, discordMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordMessage struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.MessageID = foreign.ID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Message(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := DiscordMessageRevisionSlice{&local}
	if err = local.L.LoadMessage(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Message == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Message = nil
	if err = local.L.LoadMessage(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Message == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testDiscordMessageRevisionToOneSetOpDiscordMessageUsingMessage(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a DiscordMessageRevision
	var b, c DiscordMessage

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, discordMessageRevisionDBTypes, false, strmangle.SetComplement(discordMessageRevisionPrimaryKeyColumns, discordMessageRevisionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, discordMessageDBTypes, false, strmangle.SetComplement(discordMessagePrimaryKeyColumns, discordMessageColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, discordMessageDBTypes, false, strmangle.SetComplement(discordMessagePrimaryKeyColumns, discordMessageColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*DiscordMessage{&b, &c} {
		err = a.SetMessage(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Message != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.MessageDiscordMessageRevisions[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.MessageID != x.ID {
			t.Error("foreign key was wrong value", a.MessageID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.MessageID))
		reflect.Indirect(reflect.ValueOf(&a.MessageID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.MessageID != x.ID {
			t.Error("foreign key was wrong value", a.MessageID, x.ID)
		}
	}
}
func testDiscordMessageRevisionsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordMessageRevision := &DiscordMessageRevision{}
	if err = randomize.Struct(seed, discordMessageRevision, discordMessageRevisionDBTypes, true, discordMessageRevisionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordMessageRevision struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordMessageRevision.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = discordMessageRevision.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testDiscordMessageRevisionsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordMessageRevision := &DiscordMessageRevision{}
	if err = randomize.Struct(seed, discordMessageRevision, discordMessageRevisionDBTypes, true, discordMessageRevisionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordMessageRevision struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordMessageRevision.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := DiscordMessageRevisionSlice{discordMessageRevision}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testDiscordMessageRevisionsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordMessageRevision := &DiscordMessageRevision{}
	if err = randomize.Struct(seed, discordMessageRevision, discordMessageRevisionDBTypes, true, discordMessageRevisionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordMessageRevision struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordMessageRevision.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := DiscordMessageRevisions(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	discordMessageRevisionDBTypes = map[string]string{`Content`: `text`, `CreatedAt`: `timestamp with time zone`, `Embeds`: `ARRAYbigint`, `ID`: `bigint`, `MessageID`: `bigint`, `Number`: `integer`}
	_                             = bytes.MinRead
)

func testDiscordMessageRevisionsUpdate(t *testing.T) {
	t.Parallel()

	if len(discordMessageRevisionColumns) == len(discordMessageRevisionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	discordMessageRevision := &DiscordMessageRevision{}
	if err = randomize.Struct(seed, discordMessageRevision, discordMessageRevisionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordMessageRevision struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordMessageRevision.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DiscordMessageRevisions(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, discordMessageRevision, discordMessageRevisionDBTypes, true, discordMessageRevisionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordMessageRevision struct: %s", err)
	}

	if err = discordMessageRevision.Update(tx); err != nil {
		t.Error(err)
	}
}

func testDiscordMessageRevisionsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(discordMessageRevisionColumns) == len(discordMessageRevisionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	discordMessageRevision := &DiscordMessageRevision{}
	if err = randomize.Struct(seed, discordMessageRevision, discordMessageRevisionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordMessageRevision struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordMessageRevision.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DiscordMessageRevisions(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, discordMessageRevision, discordMessageRevisionDBTypes, true, discordMessageRevisionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DiscordMessageRevision struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(discordMessageRevisionColumns, discordMessageRevisionPrimaryKeyColumns) {
		fields = discordMessageRevisionColumns
	} else {
		fields = strmangle.SetComplement(
			discordMessageRevisionColumns,
			discordMessageRevisionPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(discordMessageRevision))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := DiscordMessageRevisionSlice{discordMessageRevision}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testDiscordMessageRevisionsUpsert(t *testing.T) {
	t.Parallel()

	if len(discordMessageRevisionColumns) == len(discordMessageRevisionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	discordMessageRevision := DiscordMessageRevision{}
	if err = randomize.Struct(seed, &discordMessageRevision, discordMessageRevisionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordMessageRevision struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordMessageRevision.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert DiscordMessageRevision: %s", err)
	}

	count, err := DiscordMessageRevisions(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &discordMessageRevision, discordMessageRevisionDBTypes, false, discordMessageRevisionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DiscordMessageRevision struct: %s", err)
	}

	if err = discordMessageRevision.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert DiscordMessageRevision: %s", err)
	}

	count, err = DiscordMessageRevisions(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
