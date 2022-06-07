// Code generated by SQLBoiler 4.11.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/v4/types"
	"github.com/volatiletech/strmangle"
)

// LigandEff is an object representing the database table.
type LigandEff struct {
	ActivityID int64             `boil:"activity_id" json:"activity_id" toml:"activity_id" yaml:"activity_id"`
	Bei        types.NullDecimal `boil:"bei" json:"bei,omitempty" toml:"bei" yaml:"bei,omitempty"`
	Sei        types.NullDecimal `boil:"sei" json:"sei,omitempty" toml:"sei" yaml:"sei,omitempty"`
	Le         types.NullDecimal `boil:"le" json:"le,omitempty" toml:"le" yaml:"le,omitempty"`
	Lle        types.NullDecimal `boil:"lle" json:"lle,omitempty" toml:"lle" yaml:"lle,omitempty"`

	R *ligandEffR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L ligandEffL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var LigandEffColumns = struct {
	ActivityID string
	Bei        string
	Sei        string
	Le         string
	Lle        string
}{
	ActivityID: "activity_id",
	Bei:        "bei",
	Sei:        "sei",
	Le:         "le",
	Lle:        "lle",
}

var LigandEffTableColumns = struct {
	ActivityID string
	Bei        string
	Sei        string
	Le         string
	Lle        string
}{
	ActivityID: "ligand_eff.activity_id",
	Bei:        "ligand_eff.bei",
	Sei:        "ligand_eff.sei",
	Le:         "ligand_eff.le",
	Lle:        "ligand_eff.lle",
}

// Generated where

var LigandEffWhere = struct {
	ActivityID whereHelperint64
	Bei        whereHelpertypes_NullDecimal
	Sei        whereHelpertypes_NullDecimal
	Le         whereHelpertypes_NullDecimal
	Lle        whereHelpertypes_NullDecimal
}{
	ActivityID: whereHelperint64{field: "\"ligand_eff\".\"activity_id\""},
	Bei:        whereHelpertypes_NullDecimal{field: "\"ligand_eff\".\"bei\""},
	Sei:        whereHelpertypes_NullDecimal{field: "\"ligand_eff\".\"sei\""},
	Le:         whereHelpertypes_NullDecimal{field: "\"ligand_eff\".\"le\""},
	Lle:        whereHelpertypes_NullDecimal{field: "\"ligand_eff\".\"lle\""},
}

// LigandEffRels is where relationship names are stored.
var LigandEffRels = struct {
	Activity string
}{
	Activity: "Activity",
}

// ligandEffR is where relationships are stored.
type ligandEffR struct {
	Activity *Activity `boil:"Activity" json:"Activity" toml:"Activity" yaml:"Activity"`
}

// NewStruct creates a new relationship struct
func (*ligandEffR) NewStruct() *ligandEffR {
	return &ligandEffR{}
}

func (r *ligandEffR) GetActivity() *Activity {
	if r == nil {
		return nil
	}
	return r.Activity
}

// ligandEffL is where Load methods for each relationship are stored.
type ligandEffL struct{}

var (
	ligandEffAllColumns            = []string{"activity_id", "bei", "sei", "le", "lle"}
	ligandEffColumnsWithoutDefault = []string{"activity_id"}
	ligandEffColumnsWithDefault    = []string{"bei", "sei", "le", "lle"}
	ligandEffPrimaryKeyColumns     = []string{"activity_id"}
	ligandEffGeneratedColumns      = []string{}
)

type (
	// LigandEffSlice is an alias for a slice of pointers to LigandEff.
	// This should almost always be used instead of []LigandEff.
	LigandEffSlice []*LigandEff
	// LigandEffHook is the signature for custom LigandEff hook methods
	LigandEffHook func(context.Context, boil.ContextExecutor, *LigandEff) error

	ligandEffQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	ligandEffType                 = reflect.TypeOf(&LigandEff{})
	ligandEffMapping              = queries.MakeStructMapping(ligandEffType)
	ligandEffPrimaryKeyMapping, _ = queries.BindMapping(ligandEffType, ligandEffMapping, ligandEffPrimaryKeyColumns)
	ligandEffInsertCacheMut       sync.RWMutex
	ligandEffInsertCache          = make(map[string]insertCache)
	ligandEffUpdateCacheMut       sync.RWMutex
	ligandEffUpdateCache          = make(map[string]updateCache)
	ligandEffUpsertCacheMut       sync.RWMutex
	ligandEffUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var ligandEffAfterSelectHooks []LigandEffHook

var ligandEffBeforeInsertHooks []LigandEffHook
var ligandEffAfterInsertHooks []LigandEffHook

var ligandEffBeforeUpdateHooks []LigandEffHook
var ligandEffAfterUpdateHooks []LigandEffHook

var ligandEffBeforeDeleteHooks []LigandEffHook
var ligandEffAfterDeleteHooks []LigandEffHook

var ligandEffBeforeUpsertHooks []LigandEffHook
var ligandEffAfterUpsertHooks []LigandEffHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *LigandEff) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ligandEffAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *LigandEff) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ligandEffBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *LigandEff) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ligandEffAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *LigandEff) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ligandEffBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *LigandEff) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ligandEffAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *LigandEff) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ligandEffBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *LigandEff) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ligandEffAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *LigandEff) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ligandEffBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *LigandEff) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ligandEffAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddLigandEffHook registers your hook function for all future operations.
func AddLigandEffHook(hookPoint boil.HookPoint, ligandEffHook LigandEffHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		ligandEffAfterSelectHooks = append(ligandEffAfterSelectHooks, ligandEffHook)
	case boil.BeforeInsertHook:
		ligandEffBeforeInsertHooks = append(ligandEffBeforeInsertHooks, ligandEffHook)
	case boil.AfterInsertHook:
		ligandEffAfterInsertHooks = append(ligandEffAfterInsertHooks, ligandEffHook)
	case boil.BeforeUpdateHook:
		ligandEffBeforeUpdateHooks = append(ligandEffBeforeUpdateHooks, ligandEffHook)
	case boil.AfterUpdateHook:
		ligandEffAfterUpdateHooks = append(ligandEffAfterUpdateHooks, ligandEffHook)
	case boil.BeforeDeleteHook:
		ligandEffBeforeDeleteHooks = append(ligandEffBeforeDeleteHooks, ligandEffHook)
	case boil.AfterDeleteHook:
		ligandEffAfterDeleteHooks = append(ligandEffAfterDeleteHooks, ligandEffHook)
	case boil.BeforeUpsertHook:
		ligandEffBeforeUpsertHooks = append(ligandEffBeforeUpsertHooks, ligandEffHook)
	case boil.AfterUpsertHook:
		ligandEffAfterUpsertHooks = append(ligandEffAfterUpsertHooks, ligandEffHook)
	}
}

// One returns a single ligandEff record from the query.
func (q ligandEffQuery) One(ctx context.Context, exec boil.ContextExecutor) (*LigandEff, error) {
	o := &LigandEff{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for ligand_eff")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all LigandEff records from the query.
func (q ligandEffQuery) All(ctx context.Context, exec boil.ContextExecutor) (LigandEffSlice, error) {
	var o []*LigandEff

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to LigandEff slice")
	}

	if len(ligandEffAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all LigandEff records in the query.
func (q ligandEffQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count ligand_eff rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q ligandEffQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if ligand_eff exists")
	}

	return count > 0, nil
}

// Activity pointed to by the foreign key.
func (o *LigandEff) Activity(mods ...qm.QueryMod) activityQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"activity_id\" = ?", o.ActivityID),
	}

	queryMods = append(queryMods, mods...)

	return Activities(queryMods...)
}

// LoadActivity allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (ligandEffL) LoadActivity(ctx context.Context, e boil.ContextExecutor, singular bool, maybeLigandEff interface{}, mods queries.Applicator) error {
	var slice []*LigandEff
	var object *LigandEff

	if singular {
		object = maybeLigandEff.(*LigandEff)
	} else {
		slice = *maybeLigandEff.(*[]*LigandEff)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &ligandEffR{}
		}
		args = append(args, object.ActivityID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &ligandEffR{}
			}

			for _, a := range args {
				if a == obj.ActivityID {
					continue Outer
				}
			}

			args = append(args, obj.ActivityID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`activities`),
		qm.WhereIn(`activities.activity_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Activity")
	}

	var resultSlice []*Activity
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Activity")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for activities")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for activities")
	}

	if len(ligandEffAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Activity = foreign
		if foreign.R == nil {
			foreign.R = &activityR{}
		}
		foreign.R.LigandEff = object
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.ActivityID == foreign.ActivityID {
				local.R.Activity = foreign
				if foreign.R == nil {
					foreign.R = &activityR{}
				}
				foreign.R.LigandEff = local
				break
			}
		}
	}

	return nil
}

// SetActivity of the ligandEff to the related item.
// Sets o.R.Activity to related.
// Adds o to related.R.LigandEff.
func (o *LigandEff) SetActivity(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Activity) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"ligand_eff\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, []string{"activity_id"}),
		strmangle.WhereClause("\"", "\"", 0, ligandEffPrimaryKeyColumns),
	)
	values := []interface{}{related.ActivityID, o.ActivityID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.ActivityID = related.ActivityID
	if o.R == nil {
		o.R = &ligandEffR{
			Activity: related,
		}
	} else {
		o.R.Activity = related
	}

	if related.R == nil {
		related.R = &activityR{
			LigandEff: o,
		}
	} else {
		related.R.LigandEff = o
	}

	return nil
}

// LigandEffs retrieves all the records using an executor.
func LigandEffs(mods ...qm.QueryMod) ligandEffQuery {
	mods = append(mods, qm.From("\"ligand_eff\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"ligand_eff\".*"})
	}

	return ligandEffQuery{q}
}

// FindLigandEff retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindLigandEff(ctx context.Context, exec boil.ContextExecutor, activityID int64, selectCols ...string) (*LigandEff, error) {
	ligandEffObj := &LigandEff{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"ligand_eff\" where \"activity_id\"=?", sel,
	)

	q := queries.Raw(query, activityID)

	err := q.Bind(ctx, exec, ligandEffObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from ligand_eff")
	}

	if err = ligandEffObj.doAfterSelectHooks(ctx, exec); err != nil {
		return ligandEffObj, err
	}

	return ligandEffObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *LigandEff) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no ligand_eff provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(ligandEffColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	ligandEffInsertCacheMut.RLock()
	cache, cached := ligandEffInsertCache[key]
	ligandEffInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			ligandEffAllColumns,
			ligandEffColumnsWithDefault,
			ligandEffColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(ligandEffType, ligandEffMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(ligandEffType, ligandEffMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"ligand_eff\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"ligand_eff\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into ligand_eff")
	}

	if !cached {
		ligandEffInsertCacheMut.Lock()
		ligandEffInsertCache[key] = cache
		ligandEffInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the LigandEff.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *LigandEff) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	ligandEffUpdateCacheMut.RLock()
	cache, cached := ligandEffUpdateCache[key]
	ligandEffUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			ligandEffAllColumns,
			ligandEffPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update ligand_eff, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"ligand_eff\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, ligandEffPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(ligandEffType, ligandEffMapping, append(wl, ligandEffPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update ligand_eff row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for ligand_eff")
	}

	if !cached {
		ligandEffUpdateCacheMut.Lock()
		ligandEffUpdateCache[key] = cache
		ligandEffUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q ligandEffQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for ligand_eff")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for ligand_eff")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o LigandEffSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), ligandEffPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"ligand_eff\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, ligandEffPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in ligandEff slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all ligandEff")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *LigandEff) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no ligand_eff provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(ligandEffColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
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
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	ligandEffUpsertCacheMut.RLock()
	cache, cached := ligandEffUpsertCache[key]
	ligandEffUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			ligandEffAllColumns,
			ligandEffColumnsWithDefault,
			ligandEffColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			ligandEffAllColumns,
			ligandEffPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert ligand_eff, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(ligandEffPrimaryKeyColumns))
			copy(conflict, ligandEffPrimaryKeyColumns)
		}
		cache.query = buildUpsertQuerySQLite(dialect, "\"ligand_eff\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(ligandEffType, ligandEffMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(ligandEffType, ligandEffMapping, ret)
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

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert ligand_eff")
	}

	if !cached {
		ligandEffUpsertCacheMut.Lock()
		ligandEffUpsertCache[key] = cache
		ligandEffUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single LigandEff record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *LigandEff) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no LigandEff provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), ligandEffPrimaryKeyMapping)
	sql := "DELETE FROM \"ligand_eff\" WHERE \"activity_id\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from ligand_eff")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for ligand_eff")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q ligandEffQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no ligandEffQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from ligand_eff")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for ligand_eff")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o LigandEffSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(ligandEffBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), ligandEffPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"ligand_eff\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, ligandEffPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from ligandEff slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for ligand_eff")
	}

	if len(ligandEffAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *LigandEff) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindLigandEff(ctx, exec, o.ActivityID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *LigandEffSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := LigandEffSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), ligandEffPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"ligand_eff\".* FROM \"ligand_eff\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, ligandEffPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in LigandEffSlice")
	}

	*o = slice

	return nil
}

// LigandEffExists checks if the LigandEff row exists.
func LigandEffExists(ctx context.Context, exec boil.ContextExecutor, activityID int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"ligand_eff\" where \"activity_id\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, activityID)
	}
	row := exec.QueryRowContext(ctx, sql, activityID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if ligand_eff exists")
	}

	return exists, nil
}