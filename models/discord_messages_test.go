package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testDiscordMessages(t *testing.T) {
	t.Parallel()

	query := DiscordMessages(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testDiscordMessagesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordMessage := &DiscordMessage{}
	if err = randomize.Struct(seed, discordMessage, discordMessageDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordMessage struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordMessage.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = discordMessage.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := DiscordMessages(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDiscordMessagesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordMessage := &DiscordMessage{}
	if err = randomize.Struct(seed, discordMessage, discordMessageDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordMessage struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordMessage.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = DiscordMessages(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := DiscordMessages(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDiscordMessagesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordMessage := &DiscordMessage{}
	if err = randomize.Struct(seed, discordMessage, discordMessageDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordMessage struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordMessage.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := DiscordMessageSlice{discordMessage}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := DiscordMessages(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testDiscordMessagesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordMessage := &DiscordMessage{}
	if err = randomize.Struct(seed, discordMessage, discordMessageDBTypes, true, discordMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordMessage struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordMessage.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := DiscordMessageExists(tx, discordMessage.ID)
	if err != nil {
		t.Errorf("Unable to check if DiscordMessage exists: %s", err)
	}
	if !e {
		t.Errorf("Expected DiscordMessageExistsG to return true, but got false.")
	}
}
func testDiscordMessagesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordMessage := &DiscordMessage{}
	if err = randomize.Struct(seed, discordMessage, discordMessageDBTypes, true, discordMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordMessage struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordMessage.Insert(tx); err != nil {
		t.Error(err)
	}

	discordMessageFound, err := FindDiscordMessage(tx, discordMessage.ID)
	if err != nil {
		t.Error(err)
	}

	if discordMessageFound == nil {
		t.Error("want a record, got nil")
	}
}
func testDiscordMessagesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordMessage := &DiscordMessage{}
	if err = randomize.Struct(seed, discordMessage, discordMessageDBTypes, true, discordMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordMessage struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordMessage.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = DiscordMessages(tx).Bind(discordMessage); err != nil {
		t.Error(err)
	}
}

func testDiscordMessagesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordMessage := &DiscordMessage{}
	if err = randomize.Struct(seed, discordMessage, discordMessageDBTypes, true, discordMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordMessage struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordMessage.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := DiscordMessages(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testDiscordMessagesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordMessageOne := &DiscordMessage{}
	discordMessageTwo := &DiscordMessage{}
	if err = randomize.Struct(seed, discordMessageOne, discordMessageDBTypes, false, discordMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordMessage struct: %s", err)
	}
	if err = randomize.Struct(seed, discordMessageTwo, discordMessageDBTypes, false, discordMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordMessage struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordMessageOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = discordMessageTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := DiscordMessages(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testDiscordMessagesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	discordMessageOne := &DiscordMessage{}
	discordMessageTwo := &DiscordMessage{}
	if err = randomize.Struct(seed, discordMessageOne, discordMessageDBTypes, false, discordMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordMessage struct: %s", err)
	}
	if err = randomize.Struct(seed, discordMessageTwo, discordMessageDBTypes, false, discordMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordMessage struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordMessageOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = discordMessageTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DiscordMessages(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testDiscordMessagesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordMessage := &DiscordMessage{}
	if err = randomize.Struct(seed, discordMessage, discordMessageDBTypes, true, discordMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordMessage struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordMessage.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DiscordMessages(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDiscordMessagesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordMessage := &DiscordMessage{}
	if err = randomize.Struct(seed, discordMessage, discordMessageDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordMessage struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordMessage.Insert(tx, discordMessageColumns...); err != nil {
		t.Error(err)
	}

	count, err := DiscordMessages(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDiscordMessageToManyMessageDiscordMessageRevisions(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a DiscordMessage
	var b, c DiscordMessageRevision

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, discordMessageDBTypes, true, discordMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordMessage struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, discordMessageRevisionDBTypes, false, discordMessageRevisionColumnsWithDefault...)
	randomize.Struct(seed, &c, discordMessageRevisionDBTypes, false, discordMessageRevisionColumnsWithDefault...)

	b.MessageID = a.ID
	c.MessageID = a.ID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	discordMessageRevision, err := a.MessageDiscordMessageRevisions(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range discordMessageRevision {
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

	slice := DiscordMessageSlice{&a}
	if err = a.L.LoadMessageDiscordMessageRevisions(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.MessageDiscordMessageRevisions); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.MessageDiscordMessageRevisions = nil
	if err = a.L.LoadMessageDiscordMessageRevisions(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.MessageDiscordMessageRevisions); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", discordMessageRevision)
	}
}

func testDiscordMessageToManyMessageDiscordMessageEmbeds(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a DiscordMessage
	var b, c DiscordMessageEmbed

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, discordMessageDBTypes, true, discordMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordMessage struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, discordMessageEmbedDBTypes, false, discordMessageEmbedColumnsWithDefault...)
	randomize.Struct(seed, &c, discordMessageEmbedDBTypes, false, discordMessageEmbedColumnsWithDefault...)

	b.MessageID = a.ID
	c.MessageID = a.ID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	discordMessageEmbed, err := a.MessageDiscordMessageEmbeds(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range discordMessageEmbed {
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

	slice := DiscordMessageSlice{&a}
	if err = a.L.LoadMessageDiscordMessageEmbeds(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.MessageDiscordMessageEmbeds); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.MessageDiscordMessageEmbeds = nil
	if err = a.L.LoadMessageDiscordMessageEmbeds(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.MessageDiscordMessageEmbeds); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", discordMessageEmbed)
	}
}

func testDiscordMessageToManyAddOpMessageDiscordMessageRevisions(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a DiscordMessage
	var b, c, d, e DiscordMessageRevision

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, discordMessageDBTypes, false, strmangle.SetComplement(discordMessagePrimaryKeyColumns, discordMessageColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*DiscordMessageRevision{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, discordMessageRevisionDBTypes, false, strmangle.SetComplement(discordMessageRevisionPrimaryKeyColumns, discordMessageRevisionColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*DiscordMessageRevision{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddMessageDiscordMessageRevisions(tx, i != 0, x...)
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

		if a.R.MessageDiscordMessageRevisions[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.MessageDiscordMessageRevisions[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.MessageDiscordMessageRevisions(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testDiscordMessageToManyAddOpMessageDiscordMessageEmbeds(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a DiscordMessage
	var b, c, d, e DiscordMessageEmbed

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, discordMessageDBTypes, false, strmangle.SetComplement(discordMessagePrimaryKeyColumns, discordMessageColumnsWithoutDefault)...); err != nil {
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
		err = a.AddMessageDiscordMessageEmbeds(tx, i != 0, x...)
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

		if a.R.MessageDiscordMessageEmbeds[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.MessageDiscordMessageEmbeds[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.MessageDiscordMessageEmbeds(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testDiscordMessagesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordMessage := &DiscordMessage{}
	if err = randomize.Struct(seed, discordMessage, discordMessageDBTypes, true, discordMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordMessage struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordMessage.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = discordMessage.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testDiscordMessagesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordMessage := &DiscordMessage{}
	if err = randomize.Struct(seed, discordMessage, discordMessageDBTypes, true, discordMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordMessage struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordMessage.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := DiscordMessageSlice{discordMessage}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testDiscordMessagesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	discordMessage := &DiscordMessage{}
	if err = randomize.Struct(seed, discordMessage, discordMessageDBTypes, true, discordMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordMessage struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordMessage.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := DiscordMessages(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	discordMessageDBTypes = map[string]string{`AuthorAvatar`: `text`, `AuthorBot`: `boolean`, `AuthorDiscrim`: `integer`, `AuthorID`: `bigint`, `AuthorUsername`: `character varying`, `ChannelID`: `bigint`, `Content`: `text`, `DeletedAt`: `timestamp with time zone`, `EditedTimestamp`: `timestamp with time zone`, `Embeds`: `ARRAYbigint`, `ID`: `bigint`, `MentionEveryone`: `boolean`, `MentionRoles`: `ARRAYbigint`, `Mentions`: `ARRAYbigint`, `Timestamp`: `timestamp with time zone`}
	_                     = bytes.MinRead
)

func testDiscordMessagesUpdate(t *testing.T) {
	t.Parallel()

	if len(discordMessageColumns) == len(discordMessagePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	discordMessage := &DiscordMessage{}
	if err = randomize.Struct(seed, discordMessage, discordMessageDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordMessage struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordMessage.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DiscordMessages(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, discordMessage, discordMessageDBTypes, true, discordMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DiscordMessage struct: %s", err)
	}

	if err = discordMessage.Update(tx); err != nil {
		t.Error(err)
	}
}

func testDiscordMessagesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(discordMessageColumns) == len(discordMessagePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	discordMessage := &DiscordMessage{}
	if err = randomize.Struct(seed, discordMessage, discordMessageDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordMessage struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordMessage.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := DiscordMessages(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, discordMessage, discordMessageDBTypes, true, discordMessagePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DiscordMessage struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(discordMessageColumns, discordMessagePrimaryKeyColumns) {
		fields = discordMessageColumns
	} else {
		fields = strmangle.SetComplement(
			discordMessageColumns,
			discordMessagePrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(discordMessage))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := DiscordMessageSlice{discordMessage}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testDiscordMessagesUpsert(t *testing.T) {
	t.Parallel()

	if len(discordMessageColumns) == len(discordMessagePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	discordMessage := DiscordMessage{}
	if err = randomize.Struct(seed, &discordMessage, discordMessageDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DiscordMessage struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = discordMessage.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert DiscordMessage: %s", err)
	}

	count, err := DiscordMessages(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &discordMessage, discordMessageDBTypes, false, discordMessagePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DiscordMessage struct: %s", err)
	}

	if err = discordMessage.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert DiscordMessage: %s", err)
	}

	count, err = DiscordMessages(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
