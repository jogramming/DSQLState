package models

import (
	"bytes"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/queries"
	"github.com/vattle/sqlboiler/queries/qm"
	"github.com/vattle/sqlboiler/strmangle"
	"github.com/vattle/sqlboiler/types"
	"gopkg.in/nullbio/null.v6"
)

// DMember is an object representing the database table.
type DMember struct {
	UserID    int64            `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	GuildID   int64            `boil:"guild_id" json:"guild_id" toml:"guild_id" yaml:"guild_id"`
	CreatedAt time.Time        `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	Synced    bool             `boil:"synced" json:"synced" toml:"synced" yaml:"synced"`
	LeftAt    null.Time        `boil:"left_at" json:"left_at,omitempty" toml:"left_at" yaml:"left_at,omitempty"`
	JoinedAt  time.Time        `boil:"joined_at" json:"joined_at" toml:"joined_at" yaml:"joined_at"`
	Nick      string           `boil:"nick" json:"nick" toml:"nick" yaml:"nick"`
	Deaf      bool             `boil:"deaf" json:"deaf" toml:"deaf" yaml:"deaf"`
	Mute      bool             `boil:"mute" json:"mute" toml:"mute" yaml:"mute"`
	Roles     types.Int64Array `boil:"roles" json:"roles" toml:"roles" yaml:"roles"`

	R *dMemberR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L dMemberL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// dMemberR is where relationships are stored.
type dMemberR struct {
	User *DUser
}

// dMemberL is where Load methods for each relationship are stored.
type dMemberL struct{}

var (
	dMemberColumns               = []string{"user_id", "guild_id", "created_at", "synced", "left_at", "joined_at", "nick", "deaf", "mute", "roles"}
	dMemberColumnsWithoutDefault = []string{"user_id", "guild_id", "created_at", "synced", "left_at", "joined_at", "nick", "deaf", "mute", "roles"}
	dMemberColumnsWithDefault    = []string{}
	dMemberPrimaryKeyColumns     = []string{"user_id", "guild_id"}
)

type (
	// DMemberSlice is an alias for a slice of pointers to DMember.
	// This should generally be used opposed to []DMember.
	DMemberSlice []*DMember

	dMemberQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	dMemberType                 = reflect.TypeOf(&DMember{})
	dMemberMapping              = queries.MakeStructMapping(dMemberType)
	dMemberPrimaryKeyMapping, _ = queries.BindMapping(dMemberType, dMemberMapping, dMemberPrimaryKeyColumns)
	dMemberInsertCacheMut       sync.RWMutex
	dMemberInsertCache          = make(map[string]insertCache)
	dMemberUpdateCacheMut       sync.RWMutex
	dMemberUpdateCache          = make(map[string]updateCache)
	dMemberUpsertCacheMut       sync.RWMutex
	dMemberUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)

// OneP returns a single dMember record from the query, and panics on error.
func (q dMemberQuery) OneP() *DMember {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single dMember record from the query.
func (q dMemberQuery) One() (*DMember, error) {
	o := &DMember{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for d_members")
	}

	return o, nil
}

// AllP returns all DMember records from the query, and panics on error.
func (q dMemberQuery) AllP() DMemberSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all DMember records from the query.
func (q dMemberQuery) All() (DMemberSlice, error) {
	var o DMemberSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to DMember slice")
	}

	return o, nil
}

// CountP returns the count of all DMember records in the query, and panics on error.
func (q dMemberQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all DMember records in the query.
func (q dMemberQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count d_members rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q dMemberQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q dMemberQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if d_members exists")
	}

	return count > 0, nil
}

// UserG pointed to by the foreign key.
func (o *DMember) UserG(mods ...qm.QueryMod) dUserQuery {
	return o.User(boil.GetDB(), mods...)
}

// User pointed to by the foreign key.
func (o *DMember) User(exec boil.Executor, mods ...qm.QueryMod) dUserQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.UserID),
	}

	queryMods = append(queryMods, mods...)

	query := DUsers(exec, queryMods...)
	queries.SetFrom(query.Query, "\"d_users\"")

	return query
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (dMemberL) LoadUser(e boil.Executor, singular bool, maybeDMember interface{}) error {
	var slice []*DMember
	var object *DMember

	count := 1
	if singular {
		object = maybeDMember.(*DMember)
	} else {
		slice = *maybeDMember.(*DMemberSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		if object.R == nil {
			object.R = &dMemberR{}
		}
		args[0] = object.UserID
	} else {
		for i, obj := range slice {
			if obj.R == nil {
				obj.R = &dMemberR{}
			}
			args[i] = obj.UserID
		}
	}

	query := fmt.Sprintf(
		"select * from \"d_users\" where \"id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load DUser")
	}
	defer results.Close()

	var resultSlice []*DUser
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice DUser")
	}

	if singular && len(resultSlice) != 0 {
		object.R.User = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.UserID == foreign.ID {
				local.R.User = foreign
				break
			}
		}
	}

	return nil
}

// SetUserG of the d_member to the related item.
// Sets o.R.User to related.
// Adds o to related.R.UserDMembers.
// Uses the global database handle.
func (o *DMember) SetUserG(insert bool, related *DUser) error {
	return o.SetUser(boil.GetDB(), insert, related)
}

// SetUserP of the d_member to the related item.
// Sets o.R.User to related.
// Adds o to related.R.UserDMembers.
// Panics on error.
func (o *DMember) SetUserP(exec boil.Executor, insert bool, related *DUser) {
	if err := o.SetUser(exec, insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetUserGP of the d_member to the related item.
// Sets o.R.User to related.
// Adds o to related.R.UserDMembers.
// Uses the global database handle and panics on error.
func (o *DMember) SetUserGP(insert bool, related *DUser) {
	if err := o.SetUser(boil.GetDB(), insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetUser of the d_member to the related item.
// Sets o.R.User to related.
// Adds o to related.R.UserDMembers.
func (o *DMember) SetUser(exec boil.Executor, insert bool, related *DUser) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"d_members\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"user_id"}),
		strmangle.WhereClause("\"", "\"", 2, dMemberPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.UserID, o.GuildID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.UserID = related.ID

	if o.R == nil {
		o.R = &dMemberR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &dUserR{
			UserDMembers: DMemberSlice{o},
		}
	} else {
		related.R.UserDMembers = append(related.R.UserDMembers, o)
	}

	return nil
}

// DMembersG retrieves all records.
func DMembersG(mods ...qm.QueryMod) dMemberQuery {
	return DMembers(boil.GetDB(), mods...)
}

// DMembers retrieves all the records using an executor.
func DMembers(exec boil.Executor, mods ...qm.QueryMod) dMemberQuery {
	mods = append(mods, qm.From("\"d_members\""))
	return dMemberQuery{NewQuery(exec, mods...)}
}

// FindDMemberG retrieves a single record by ID.
func FindDMemberG(userID int64, guildID int64, selectCols ...string) (*DMember, error) {
	return FindDMember(boil.GetDB(), userID, guildID, selectCols...)
}

// FindDMemberGP retrieves a single record by ID, and panics on error.
func FindDMemberGP(userID int64, guildID int64, selectCols ...string) *DMember {
	retobj, err := FindDMember(boil.GetDB(), userID, guildID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindDMember retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDMember(exec boil.Executor, userID int64, guildID int64, selectCols ...string) (*DMember, error) {
	dMemberObj := &DMember{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"d_members\" where \"user_id\"=$1 AND \"guild_id\"=$2", sel,
	)

	q := queries.Raw(exec, query, userID, guildID)

	err := q.Bind(dMemberObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from d_members")
	}

	return dMemberObj, nil
}

// FindDMemberP retrieves a single record by ID with an executor, and panics on error.
func FindDMemberP(exec boil.Executor, userID int64, guildID int64, selectCols ...string) *DMember {
	retobj, err := FindDMember(exec, userID, guildID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *DMember) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *DMember) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *DMember) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *DMember) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no d_members provided for insertion")
	}

	var err error
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(dMemberColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	dMemberInsertCacheMut.RLock()
	cache, cached := dMemberInsertCache[key]
	dMemberInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			dMemberColumns,
			dMemberColumnsWithDefault,
			dMemberColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(dMemberType, dMemberMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(dMemberType, dMemberMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"d_members\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

		if len(cache.retMapping) != 0 {
			cache.query += fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into d_members")
	}

	if !cached {
		dMemberInsertCacheMut.Lock()
		dMemberInsertCache[key] = cache
		dMemberInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single DMember record. See Update for
// whitelist behavior description.
func (o *DMember) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single DMember record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *DMember) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the DMember, and panics on error.
// See Update for whitelist behavior description.
func (o *DMember) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the DMember.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *DMember) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	key := makeCacheKey(whitelist, nil)
	dMemberUpdateCacheMut.RLock()
	cache, cached := dMemberUpdateCache[key]
	dMemberUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(dMemberColumns, dMemberPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update d_members, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"d_members\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, dMemberPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(dMemberType, dMemberMapping, append(wl, dMemberPrimaryKeyColumns...))
		if err != nil {
			return err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err = exec.Exec(cache.query, values...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update d_members row")
	}

	if !cached {
		dMemberUpdateCacheMut.Lock()
		dMemberUpdateCache[key] = cache
		dMemberUpdateCacheMut.Unlock()
	}

	return nil
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q dMemberQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q dMemberQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for d_members")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o DMemberSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o DMemberSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o DMemberSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o DMemberSlice) UpdateAll(exec boil.Executor, cols M) error {
	ln := int64(len(o))
	if ln == 0 {
		return nil
	}

	if len(cols) == 0 {
		return errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dMemberPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"d_members\" SET %s WHERE (\"user_id\",\"guild_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(dMemberPrimaryKeyColumns), len(colNames)+1, len(dMemberPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in dMember slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *DMember) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *DMember) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *DMember) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *DMember) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no d_members provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(dMemberColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs postgres problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range updateColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range whitelist {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	dMemberUpsertCacheMut.RLock()
	cache, cached := dMemberUpsertCache[key]
	dMemberUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			dMemberColumns,
			dMemberColumnsWithDefault,
			dMemberColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			dMemberColumns,
			dMemberPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert d_members, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(dMemberPrimaryKeyColumns))
			copy(conflict, dMemberPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"d_members\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(dMemberType, dMemberMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(dMemberType, dMemberMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert d_members")
	}

	if !cached {
		dMemberUpsertCacheMut.Lock()
		dMemberUpsertCache[key] = cache
		dMemberUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteP deletes a single DMember record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *DMember) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single DMember record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *DMember) DeleteG() error {
	if o == nil {
		return errors.New("models: no DMember provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single DMember record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *DMember) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single DMember record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *DMember) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no DMember provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), dMemberPrimaryKeyMapping)
	sql := "DELETE FROM \"d_members\" WHERE \"user_id\"=$1 AND \"guild_id\"=$2"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from d_members")
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q dMemberQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q dMemberQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no dMemberQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from d_members")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o DMemberSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o DMemberSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no DMember slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o DMemberSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o DMemberSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no DMember slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dMemberPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"d_members\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, dMemberPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(dMemberPrimaryKeyColumns), 1, len(dMemberPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from dMember slice")
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *DMember) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *DMember) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *DMember) ReloadG() error {
	if o == nil {
		return errors.New("models: no DMember provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *DMember) Reload(exec boil.Executor) error {
	ret, err := FindDMember(exec, o.UserID, o.GuildID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *DMemberSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *DMemberSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DMemberSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty DMemberSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DMemberSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	dMembers := DMemberSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dMemberPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"d_members\".* FROM \"d_members\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, dMemberPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(dMemberPrimaryKeyColumns), 1, len(dMemberPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&dMembers)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in DMemberSlice")
	}

	*o = dMembers

	return nil
}

// DMemberExists checks if the DMember row exists.
func DMemberExists(exec boil.Executor, userID int64, guildID int64) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"d_members\" where \"user_id\"=$1 AND \"guild_id\"=$2 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, userID, guildID)
	}

	row := exec.QueryRow(sql, userID, guildID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if d_members exists")
	}

	return exists, nil
}

// DMemberExistsG checks if the DMember row exists.
func DMemberExistsG(userID int64, guildID int64) (bool, error) {
	return DMemberExists(boil.GetDB(), userID, guildID)
}

// DMemberExistsGP checks if the DMember row exists. Panics on error.
func DMemberExistsGP(userID int64, guildID int64) bool {
	e, err := DMemberExists(boil.GetDB(), userID, guildID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// DMemberExistsP checks if the DMember row exists. Panics on error.
func DMemberExistsP(exec boil.Executor, userID int64, guildID int64) bool {
	e, err := DMemberExists(exec, userID, guildID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
