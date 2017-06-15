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
	"gopkg.in/nullbio/null.v6"
)

// DUser is an object representing the database table.
type DUser struct {
	ID            int64       `boil:"id" json:"id" toml:"id" yaml:"id"`
	CreatedAt     time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	Username      string      `boil:"username" json:"username" toml:"username" yaml:"username"`
	Discriminator string      `boil:"discriminator" json:"discriminator" toml:"discriminator" yaml:"discriminator"`
	Bot           bool        `boil:"bot" json:"bot" toml:"bot" yaml:"bot"`
	Avatar        string      `boil:"avatar" json:"avatar" toml:"avatar" yaml:"avatar"`
	Status        string      `boil:"status" json:"status" toml:"status" yaml:"status"`
	GameName      null.String `boil:"game_name" json:"game_name,omitempty" toml:"game_name" yaml:"game_name,omitempty"`
	GameType      null.Int    `boil:"game_type" json:"game_type,omitempty" toml:"game_type" yaml:"game_type,omitempty"`
	GameURL       null.String `boil:"game_url" json:"game_url,omitempty" toml:"game_url" yaml:"game_url,omitempty"`

	R *dUserR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L dUserL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// dUserR is where relationships are stored.
type dUserR struct {
	UserDMembers DMemberSlice
}

// dUserL is where Load methods for each relationship are stored.
type dUserL struct{}

var (
	dUserColumns               = []string{"id", "created_at", "username", "discriminator", "bot", "avatar", "status", "game_name", "game_type", "game_url"}
	dUserColumnsWithoutDefault = []string{"id", "created_at", "username", "discriminator", "bot", "avatar", "status", "game_name", "game_type", "game_url"}
	dUserColumnsWithDefault    = []string{}
	dUserPrimaryKeyColumns     = []string{"id"}
)

type (
	// DUserSlice is an alias for a slice of pointers to DUser.
	// This should generally be used opposed to []DUser.
	DUserSlice []*DUser

	dUserQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	dUserType                 = reflect.TypeOf(&DUser{})
	dUserMapping              = queries.MakeStructMapping(dUserType)
	dUserPrimaryKeyMapping, _ = queries.BindMapping(dUserType, dUserMapping, dUserPrimaryKeyColumns)
	dUserInsertCacheMut       sync.RWMutex
	dUserInsertCache          = make(map[string]insertCache)
	dUserUpdateCacheMut       sync.RWMutex
	dUserUpdateCache          = make(map[string]updateCache)
	dUserUpsertCacheMut       sync.RWMutex
	dUserUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)

// OneP returns a single dUser record from the query, and panics on error.
func (q dUserQuery) OneP() *DUser {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single dUser record from the query.
func (q dUserQuery) One() (*DUser, error) {
	o := &DUser{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for d_users")
	}

	return o, nil
}

// AllP returns all DUser records from the query, and panics on error.
func (q dUserQuery) AllP() DUserSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all DUser records from the query.
func (q dUserQuery) All() (DUserSlice, error) {
	var o DUserSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to DUser slice")
	}

	return o, nil
}

// CountP returns the count of all DUser records in the query, and panics on error.
func (q dUserQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all DUser records in the query.
func (q dUserQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count d_users rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q dUserQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q dUserQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if d_users exists")
	}

	return count > 0, nil
}

// UserDMembersG retrieves all the d_member's d members via user_id column.
func (o *DUser) UserDMembersG(mods ...qm.QueryMod) dMemberQuery {
	return o.UserDMembers(boil.GetDB(), mods...)
}

// UserDMembers retrieves all the d_member's d members with an executor via user_id column.
func (o *DUser) UserDMembers(exec boil.Executor, mods ...qm.QueryMod) dMemberQuery {
	queryMods := []qm.QueryMod{
		qm.Select("\"a\".*"),
	}

	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"a\".\"user_id\"=?", o.ID),
	)

	query := DMembers(exec, queryMods...)
	queries.SetFrom(query.Query, "\"d_members\" as \"a\"")
	return query
}

// LoadUserDMembers allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (dUserL) LoadUserDMembers(e boil.Executor, singular bool, maybeDUser interface{}) error {
	var slice []*DUser
	var object *DUser

	count := 1
	if singular {
		object = maybeDUser.(*DUser)
	} else {
		slice = *maybeDUser.(*DUserSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		if object.R == nil {
			object.R = &dUserR{}
		}
		args[0] = object.ID
	} else {
		for i, obj := range slice {
			if obj.R == nil {
				obj.R = &dUserR{}
			}
			args[i] = obj.ID
		}
	}

	query := fmt.Sprintf(
		"select * from \"d_members\" where \"user_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)
	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load d_members")
	}
	defer results.Close()

	var resultSlice []*DMember
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice d_members")
	}

	if singular {
		object.R.UserDMembers = resultSlice
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.UserID {
				local.R.UserDMembers = append(local.R.UserDMembers, foreign)
				break
			}
		}
	}

	return nil
}

// AddUserDMembersG adds the given related objects to the existing relationships
// of the d_user, optionally inserting them as new records.
// Appends related to o.R.UserDMembers.
// Sets related.R.User appropriately.
// Uses the global database handle.
func (o *DUser) AddUserDMembersG(insert bool, related ...*DMember) error {
	return o.AddUserDMembers(boil.GetDB(), insert, related...)
}

// AddUserDMembersP adds the given related objects to the existing relationships
// of the d_user, optionally inserting them as new records.
// Appends related to o.R.UserDMembers.
// Sets related.R.User appropriately.
// Panics on error.
func (o *DUser) AddUserDMembersP(exec boil.Executor, insert bool, related ...*DMember) {
	if err := o.AddUserDMembers(exec, insert, related...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// AddUserDMembersGP adds the given related objects to the existing relationships
// of the d_user, optionally inserting them as new records.
// Appends related to o.R.UserDMembers.
// Sets related.R.User appropriately.
// Uses the global database handle and panics on error.
func (o *DUser) AddUserDMembersGP(insert bool, related ...*DMember) {
	if err := o.AddUserDMembers(boil.GetDB(), insert, related...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// AddUserDMembers adds the given related objects to the existing relationships
// of the d_user, optionally inserting them as new records.
// Appends related to o.R.UserDMembers.
// Sets related.R.User appropriately.
func (o *DUser) AddUserDMembers(exec boil.Executor, insert bool, related ...*DMember) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.UserID = o.ID
			if err = rel.Insert(exec); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"d_members\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"user_id"}),
				strmangle.WhereClause("\"", "\"", 2, dMemberPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.UserID, rel.GuildID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}

			if _, err = exec.Exec(updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.UserID = o.ID
		}
	}

	if o.R == nil {
		o.R = &dUserR{
			UserDMembers: related,
		}
	} else {
		o.R.UserDMembers = append(o.R.UserDMembers, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &dMemberR{
				User: o,
			}
		} else {
			rel.R.User = o
		}
	}
	return nil
}

// DUsersG retrieves all records.
func DUsersG(mods ...qm.QueryMod) dUserQuery {
	return DUsers(boil.GetDB(), mods...)
}

// DUsers retrieves all the records using an executor.
func DUsers(exec boil.Executor, mods ...qm.QueryMod) dUserQuery {
	mods = append(mods, qm.From("\"d_users\""))
	return dUserQuery{NewQuery(exec, mods...)}
}

// FindDUserG retrieves a single record by ID.
func FindDUserG(id int64, selectCols ...string) (*DUser, error) {
	return FindDUser(boil.GetDB(), id, selectCols...)
}

// FindDUserGP retrieves a single record by ID, and panics on error.
func FindDUserGP(id int64, selectCols ...string) *DUser {
	retobj, err := FindDUser(boil.GetDB(), id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindDUser retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDUser(exec boil.Executor, id int64, selectCols ...string) (*DUser, error) {
	dUserObj := &DUser{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"d_users\" where \"id\"=$1", sel,
	)

	q := queries.Raw(exec, query, id)

	err := q.Bind(dUserObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from d_users")
	}

	return dUserObj, nil
}

// FindDUserP retrieves a single record by ID with an executor, and panics on error.
func FindDUserP(exec boil.Executor, id int64, selectCols ...string) *DUser {
	retobj, err := FindDUser(exec, id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *DUser) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *DUser) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *DUser) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *DUser) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no d_users provided for insertion")
	}

	var err error
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(dUserColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	dUserInsertCacheMut.RLock()
	cache, cached := dUserInsertCache[key]
	dUserInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			dUserColumns,
			dUserColumnsWithDefault,
			dUserColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(dUserType, dUserMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(dUserType, dUserMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"d_users\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "models: unable to insert into d_users")
	}

	if !cached {
		dUserInsertCacheMut.Lock()
		dUserInsertCache[key] = cache
		dUserInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single DUser record. See Update for
// whitelist behavior description.
func (o *DUser) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single DUser record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *DUser) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the DUser, and panics on error.
// See Update for whitelist behavior description.
func (o *DUser) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the DUser.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *DUser) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	key := makeCacheKey(whitelist, nil)
	dUserUpdateCacheMut.RLock()
	cache, cached := dUserUpdateCache[key]
	dUserUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(dUserColumns, dUserPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update d_users, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"d_users\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, dUserPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(dUserType, dUserMapping, append(wl, dUserPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update d_users row")
	}

	if !cached {
		dUserUpdateCacheMut.Lock()
		dUserUpdateCache[key] = cache
		dUserUpdateCacheMut.Unlock()
	}

	return nil
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q dUserQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q dUserQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for d_users")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o DUserSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o DUserSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o DUserSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o DUserSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dUserPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"d_users\" SET %s WHERE (\"id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(dUserPrimaryKeyColumns), len(colNames)+1, len(dUserPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in dUser slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *DUser) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *DUser) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *DUser) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *DUser) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no d_users provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(dUserColumnsWithDefault, o)

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

	dUserUpsertCacheMut.RLock()
	cache, cached := dUserUpsertCache[key]
	dUserUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			dUserColumns,
			dUserColumnsWithDefault,
			dUserColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			dUserColumns,
			dUserPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert d_users, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(dUserPrimaryKeyColumns))
			copy(conflict, dUserPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"d_users\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(dUserType, dUserMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(dUserType, dUserMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert d_users")
	}

	if !cached {
		dUserUpsertCacheMut.Lock()
		dUserUpsertCache[key] = cache
		dUserUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteP deletes a single DUser record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *DUser) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single DUser record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *DUser) DeleteG() error {
	if o == nil {
		return errors.New("models: no DUser provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single DUser record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *DUser) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single DUser record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *DUser) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no DUser provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), dUserPrimaryKeyMapping)
	sql := "DELETE FROM \"d_users\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from d_users")
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q dUserQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q dUserQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no dUserQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from d_users")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o DUserSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o DUserSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no DUser slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o DUserSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o DUserSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no DUser slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dUserPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"d_users\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, dUserPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(dUserPrimaryKeyColumns), 1, len(dUserPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from dUser slice")
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *DUser) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *DUser) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *DUser) ReloadG() error {
	if o == nil {
		return errors.New("models: no DUser provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *DUser) Reload(exec boil.Executor) error {
	ret, err := FindDUser(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *DUserSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *DUserSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DUserSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty DUserSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DUserSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	dUsers := DUserSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dUserPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"d_users\".* FROM \"d_users\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, dUserPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(dUserPrimaryKeyColumns), 1, len(dUserPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&dUsers)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in DUserSlice")
	}

	*o = dUsers

	return nil
}

// DUserExists checks if the DUser row exists.
func DUserExists(exec boil.Executor, id int64) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"d_users\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, id)
	}

	row := exec.QueryRow(sql, id)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if d_users exists")
	}

	return exists, nil
}

// DUserExistsG checks if the DUser row exists.
func DUserExistsG(id int64) (bool, error) {
	return DUserExists(boil.GetDB(), id)
}

// DUserExistsGP checks if the DUser row exists. Panics on error.
func DUserExistsGP(id int64) bool {
	e, err := DUserExists(boil.GetDB(), id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// DUserExistsP checks if the DUser row exists. Panics on error.
func DUserExistsP(exec boil.Executor, id int64) bool {
	e, err := DUserExists(exec, id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
