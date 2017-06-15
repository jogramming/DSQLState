package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testDMessages(t *testing.T) {
	t.Parallel()

	query := DMessages(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testDMessagesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dMessage := &DMessage{}
	if err = randomize.Struct(seed, dMessage, dMessageDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DMessage struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dMessage.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = dMessage.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := DMessages(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDMessagesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dMessage := &DMessage{}
	if err = randomize.Struct(seed, dMessage, dMessageDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DMessage struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dMessage.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = DMessages(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := DMessages(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDMessagesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dMessage := &DMessage{}
	if err = randomize.Struct(seed, dMessage, dMessageDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DMessage struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dMessage.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := DMessageSlice{dMessage}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := DMessages(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testDMessagesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dMessage := &DMessage{}
	if err = randomize.Struct(seed, dMessage, dMessageDBTypes, true, dMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DMessage struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dMessage.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := DMessageExists(tx, dMessage.ID)
	if err != nil {
		t.Errorf("Unable to check if DMessage exists: %s", err)
	}
	if !e {
		t.Errorf("Expected DMessageExistsG to return true, but got false.")
	}
}
func testDMessagesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dMessage := &DMessage{}
	if err = randomize.Struct(seed, dMessage, dMessageDBTypes, true, dMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DMessage struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dMessage.Insert(tx); err != nil {
		t.Error(err)
	}

	dMessageFound, err := FindDMessage(tx, dMessage.ID)
	if err != nil {
		t.Error(err)
	}

	if dMessageFound == nil {
		t.Error("want a record, got nil")
	}
}
func testDMessagesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dMessage := &DMessage{}
	if err = randomize.Struct(seed, dMessage, dMessageDBTypes, true, dMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DMessage struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dMessage.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = DMessages(tx).Bind(dMessage); err != nil {
		t.Error(err)
	}
}

func testDMessagesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dMessage := &DMessage{}
	if err = randomize.Struct(seed, dMessage, dMessageDBTypes, true, dMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DMessage struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dMessage.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := DMessages(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testDMessagesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dMessageOne := &DMessage{}
	dMessageTwo := &DMessage{}
	if err = randomize.Struct(seed, dMessageOne, dMessageDBTypes, false, dMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DMessage struct: %s", err)
	}
	if err = randomize.Struct(seed, dMessageTwo, dMessageDBTypes, false, dMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DMessage struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dMessageOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = dMessageTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := DMessages(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testDMessagesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	dMessageOne := &DMessage{}
	dMessageTwo := &DMessage{}
	if err = randomize.Struct(seed, dMessageOne, dMessageDBTypes, false, dMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DMessage struct: %s", err)
	}
	if err = randomize.Struct(seed, dMessageTwo, dMessageDBTypes, false, dMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DMessage struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dMessageOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = dMessageTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DMessages(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testDMessagesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dMessage := &DMessage{}
	if err = randomize.Struct(seed, dMessage, dMessageDBTypes, true, dMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DMessage struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dMessage.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DMessages(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDMessagesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dMessage := &DMessage{}
	if err = randomize.Struct(seed, dMessage, dMessageDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DMessage struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dMessage.Insert(tx, dMessageColumns...); err != nil {
		t.Error(err)
	}

	count, err := DMessages(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDMessageToManyMessageDMessageRevisions(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a DMessage
	var b, c DMessageRevision

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, dMessageDBTypes, true, dMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DMessage struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, dMessageRevisionDBTypes, false, dMessageRevisionColumnsWithDefault...)
	randomize.Struct(seed, &c, dMessageRevisionDBTypes, false, dMessageRevisionColumnsWithDefault...)

	b.MessageID = a.ID
	c.MessageID = a.ID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	dMessageRevision, err := a.MessageDMessageRevisions(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range dMessageRevision {
		if v.MessageID == b.MessageID {
			bFound = true
		}
		if v.MessageID == c.MessageID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := DMessageSlice{&a}
	if err = a.L.LoadMessageDMessageRevisions(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.MessageDMessageRevisions); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.MessageDMessageRevisions = nil
	if err = a.L.LoadMessageDMessageRevisions(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.MessageDMessageRevisions); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", dMessageRevision)
	}
}

func testDMessageToManyMessageDMessageEmbeds(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a DMessage
	var b, c DMessageEmbed

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, dMessageDBTypes, true, dMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DMessage struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, dMessageEmbedDBTypes, false, dMessageEmbedColumnsWithDefault...)
	randomize.Struct(seed, &c, dMessageEmbedDBTypes, false, dMessageEmbedColumnsWithDefault...)

	b.MessageID = a.ID
	c.MessageID = a.ID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	dMessageEmbed, err := a.MessageDMessageEmbeds(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range dMessageEmbed {
		if v.MessageID == b.MessageID {
			bFound = true
		}
		if v.MessageID == c.MessageID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := DMessageSlice{&a}
	if err = a.L.LoadMessageDMessageEmbeds(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.MessageDMessageEmbeds); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.MessageDMessageEmbeds = nil
	if err = a.L.LoadMessageDMessageEmbeds(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.MessageDMessageEmbeds); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", dMessageEmbed)
	}
}

func testDMessageToManyAddOpMessageDMessageRevisions(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a DMessage
	var b, c, d, e DMessageRevision

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, dMessageDBTypes, false, strmangle.SetComplement(dMessagePrimaryKeyColumns, dMessageColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*DMessageRevision{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, dMessageRevisionDBTypes, false, strmangle.SetComplement(dMessageRevisionPrimaryKeyColumns, dMessageRevisionColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*DMessageRevision{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddMessageDMessageRevisions(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.MessageID {
			t.Error("foreign key was wrong value", a.ID, first.MessageID)
		}
		if a.ID != second.MessageID {
			t.Error("foreign key was wrong value", a.ID, second.MessageID)
		}

		if first.R.Message != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Message != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.MessageDMessageRevisions[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.MessageDMessageRevisions[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.MessageDMessageRevisions(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testDMessageToManyAddOpMessageDMessageEmbeds(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a DMessage
	var b, c, d, e DMessageEmbed

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, dMessageDBTypes, false, strmangle.SetComplement(dMessagePrimaryKeyColumns, dMessageColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*DMessageEmbed{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, dMessageEmbedDBTypes, false, strmangle.SetComplement(dMessageEmbedPrimaryKeyColumns, dMessageEmbedColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*DMessageEmbed{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddMessageDMessageEmbeds(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.MessageID {
			t.Error("foreign key was wrong value", a.ID, first.MessageID)
		}
		if a.ID != second.MessageID {
			t.Error("foreign key was wrong value", a.ID, second.MessageID)
		}

		if first.R.Message != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Message != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.MessageDMessageEmbeds[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.MessageDMessageEmbeds[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.MessageDMessageEmbeds(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testDMessagesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dMessage := &DMessage{}
	if err = randomize.Struct(seed, dMessage, dMessageDBTypes, true, dMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DMessage struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dMessage.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = dMessage.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testDMessagesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dMessage := &DMessage{}
	if err = randomize.Struct(seed, dMessage, dMessageDBTypes, true, dMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DMessage struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dMessage.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := DMessageSlice{dMessage}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testDMessagesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dMessage := &DMessage{}
	if err = randomize.Struct(seed, dMessage, dMessageDBTypes, true, dMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DMessage struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dMessage.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := DMessages(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	dMessageDBTypes = map[string]string{`AuthorAvatar`: `text`, `AuthorBot`: `boolean`, `AuthorDiscrim`: `integer`, `AuthorID`: `bigint`, `AuthorUsername`: `character varying`, `ChannelID`: `bigint`, `Content`: `text`, `DeletedAt`: `timestamp with time zone`, `EditedTimestamp`: `timestamp with time zone`, `Embeds`: `ARRAYbigint`, `ID`: `bigint`, `MentionEveryone`: `boolean`, `MentionRoles`: `ARRAYbigint`, `Mentions`: `ARRAYbigint`, `Timestamp`: `timestamp with time zone`}
	_               = bytes.MinRead
)

func testDMessagesUpdate(t *testing.T) {
	t.Parallel()

	if len(dMessageColumns) == len(dMessagePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	dMessage := &DMessage{}
	if err = randomize.Struct(seed, dMessage, dMessageDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DMessage struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dMessage.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DMessages(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, dMessage, dMessageDBTypes, true, dMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DMessage struct: %s", err)
	}

	if err = dMessage.Update(tx); err != nil {
		t.Error(err)
	}
}

func testDMessagesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(dMessageColumns) == len(dMessagePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	dMessage := &DMessage{}
	if err = randomize.Struct(seed, dMessage, dMessageDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DMessage struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dMessage.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DMessages(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, dMessage, dMessageDBTypes, true, dMessagePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DMessage struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(dMessageColumns, dMessagePrimaryKeyColumns) {
		fields = dMessageColumns
	} else {
		fields = strmangle.SetComplement(
			dMessageColumns,
			dMessagePrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(dMessage))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := DMessageSlice{dMessage}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testDMessagesUpsert(t *testing.T) {
	t.Parallel()

	if len(dMessageColumns) == len(dMessagePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	dMessage := DMessage{}
	if err = randomize.Struct(seed, &dMessage, dMessageDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DMessage struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = dMessage.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert DMessage: %s", err)
	}

	count, err := DMessages(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &dMessage, dMessageDBTypes, false, dMessagePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DMessage struct: %s", err)
	}

	if err = dMessage.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert DMessage: %s", err)
	}

	count, err = DMessages(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
