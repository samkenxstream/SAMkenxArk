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
	"github.com/volatiletech/strmangle"
)

// ReactionsideReaction is an object representing the database table.
type ReactionsideReaction struct {
	Reaction                 string `boil:"reaction" json:"reaction" toml:"reaction" yaml:"reaction"`
	Reactionside             string `boil:"reactionside" json:"reactionside" toml:"reactionside" yaml:"reactionside"`
	ReactionsideReactionType string `boil:"reactionside_reaction_type" json:"reactionside_reaction_type" toml:"reactionside_reaction_type" yaml:"reactionside_reaction_type"`

	R *reactionsideReactionR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L reactionsideReactionL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ReactionsideReactionColumns = struct {
	Reaction                 string
	Reactionside             string
	ReactionsideReactionType string
}{
	Reaction:                 "reaction",
	Reactionside:             "reactionside",
	ReactionsideReactionType: "reactionside_reaction_type",
}

var ReactionsideReactionTableColumns = struct {
	Reaction                 string
	Reactionside             string
	ReactionsideReactionType string
}{
	Reaction:                 "reactionside_reaction.reaction",
	Reactionside:             "reactionside_reaction.reactionside",
	ReactionsideReactionType: "reactionside_reaction.reactionside_reaction_type",
}

// Generated where

var ReactionsideReactionWhere = struct {
	Reaction                 whereHelperstring
	Reactionside             whereHelperstring
	ReactionsideReactionType whereHelperstring
}{
	Reaction:                 whereHelperstring{field: "\"reactionside_reaction\".\"reaction\""},
	Reactionside:             whereHelperstring{field: "\"reactionside_reaction\".\"reactionside\""},
	ReactionsideReactionType: whereHelperstring{field: "\"reactionside_reaction\".\"reactionside_reaction_type\""},
}

// ReactionsideReactionRels is where relationship names are stored.
var ReactionsideReactionRels = struct {
	ReactionsideReactionReactionside string
	ReactionsideReactionReaction     string
}{
	ReactionsideReactionReactionside: "ReactionsideReactionReactionside",
	ReactionsideReactionReaction:     "ReactionsideReactionReaction",
}

// reactionsideReactionR is where relationships are stored.
type reactionsideReactionR struct {
	ReactionsideReactionReactionside *Reactionside `boil:"ReactionsideReactionReactionside" json:"ReactionsideReactionReactionside" toml:"ReactionsideReactionReactionside" yaml:"ReactionsideReactionReactionside"`
	ReactionsideReactionReaction     *Reaction     `boil:"ReactionsideReactionReaction" json:"ReactionsideReactionReaction" toml:"ReactionsideReactionReaction" yaml:"ReactionsideReactionReaction"`
}

// NewStruct creates a new relationship struct
func (*reactionsideReactionR) NewStruct() *reactionsideReactionR {
	return &reactionsideReactionR{}
}

func (r *reactionsideReactionR) GetReactionsideReactionReactionside() *Reactionside {
	if r == nil {
		return nil
	}
	return r.ReactionsideReactionReactionside
}

func (r *reactionsideReactionR) GetReactionsideReactionReaction() *Reaction {
	if r == nil {
		return nil
	}
	return r.ReactionsideReactionReaction
}

// reactionsideReactionL is where Load methods for each relationship are stored.
type reactionsideReactionL struct{}

var (
	reactionsideReactionAllColumns            = []string{"reaction", "reactionside", "reactionside_reaction_type"}
	reactionsideReactionColumnsWithoutDefault = []string{"reaction", "reactionside", "reactionside_reaction_type"}
	reactionsideReactionColumnsWithDefault    = []string{}
	reactionsideReactionPrimaryKeyColumns     = []string{"reaction", "reactionside"}
	reactionsideReactionGeneratedColumns      = []string{}
)

type (
	// ReactionsideReactionSlice is an alias for a slice of pointers to ReactionsideReaction.
	// This should almost always be used instead of []ReactionsideReaction.
	ReactionsideReactionSlice []*ReactionsideReaction
	// ReactionsideReactionHook is the signature for custom ReactionsideReaction hook methods
	ReactionsideReactionHook func(context.Context, boil.ContextExecutor, *ReactionsideReaction) error

	reactionsideReactionQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	reactionsideReactionType                 = reflect.TypeOf(&ReactionsideReaction{})
	reactionsideReactionMapping              = queries.MakeStructMapping(reactionsideReactionType)
	reactionsideReactionPrimaryKeyMapping, _ = queries.BindMapping(reactionsideReactionType, reactionsideReactionMapping, reactionsideReactionPrimaryKeyColumns)
	reactionsideReactionInsertCacheMut       sync.RWMutex
	reactionsideReactionInsertCache          = make(map[string]insertCache)
	reactionsideReactionUpdateCacheMut       sync.RWMutex
	reactionsideReactionUpdateCache          = make(map[string]updateCache)
	reactionsideReactionUpsertCacheMut       sync.RWMutex
	reactionsideReactionUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var reactionsideReactionAfterSelectHooks []ReactionsideReactionHook

var reactionsideReactionBeforeInsertHooks []ReactionsideReactionHook
var reactionsideReactionAfterInsertHooks []ReactionsideReactionHook

var reactionsideReactionBeforeUpdateHooks []ReactionsideReactionHook
var reactionsideReactionAfterUpdateHooks []ReactionsideReactionHook

var reactionsideReactionBeforeDeleteHooks []ReactionsideReactionHook
var reactionsideReactionAfterDeleteHooks []ReactionsideReactionHook

var reactionsideReactionBeforeUpsertHooks []ReactionsideReactionHook
var reactionsideReactionAfterUpsertHooks []ReactionsideReactionHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *ReactionsideReaction) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range reactionsideReactionAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *ReactionsideReaction) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range reactionsideReactionBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *ReactionsideReaction) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range reactionsideReactionAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *ReactionsideReaction) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range reactionsideReactionBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *ReactionsideReaction) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range reactionsideReactionAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *ReactionsideReaction) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range reactionsideReactionBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *ReactionsideReaction) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range reactionsideReactionAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *ReactionsideReaction) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range reactionsideReactionBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *ReactionsideReaction) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range reactionsideReactionAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddReactionsideReactionHook registers your hook function for all future operations.
func AddReactionsideReactionHook(hookPoint boil.HookPoint, reactionsideReactionHook ReactionsideReactionHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		reactionsideReactionAfterSelectHooks = append(reactionsideReactionAfterSelectHooks, reactionsideReactionHook)
	case boil.BeforeInsertHook:
		reactionsideReactionBeforeInsertHooks = append(reactionsideReactionBeforeInsertHooks, reactionsideReactionHook)
	case boil.AfterInsertHook:
		reactionsideReactionAfterInsertHooks = append(reactionsideReactionAfterInsertHooks, reactionsideReactionHook)
	case boil.BeforeUpdateHook:
		reactionsideReactionBeforeUpdateHooks = append(reactionsideReactionBeforeUpdateHooks, reactionsideReactionHook)
	case boil.AfterUpdateHook:
		reactionsideReactionAfterUpdateHooks = append(reactionsideReactionAfterUpdateHooks, reactionsideReactionHook)
	case boil.BeforeDeleteHook:
		reactionsideReactionBeforeDeleteHooks = append(reactionsideReactionBeforeDeleteHooks, reactionsideReactionHook)
	case boil.AfterDeleteHook:
		reactionsideReactionAfterDeleteHooks = append(reactionsideReactionAfterDeleteHooks, reactionsideReactionHook)
	case boil.BeforeUpsertHook:
		reactionsideReactionBeforeUpsertHooks = append(reactionsideReactionBeforeUpsertHooks, reactionsideReactionHook)
	case boil.AfterUpsertHook:
		reactionsideReactionAfterUpsertHooks = append(reactionsideReactionAfterUpsertHooks, reactionsideReactionHook)
	}
}

// One returns a single reactionsideReaction record from the query.
func (q reactionsideReactionQuery) One(ctx context.Context, exec boil.ContextExecutor) (*ReactionsideReaction, error) {
	o := &ReactionsideReaction{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for reactionside_reaction")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all ReactionsideReaction records from the query.
func (q reactionsideReactionQuery) All(ctx context.Context, exec boil.ContextExecutor) (ReactionsideReactionSlice, error) {
	var o []*ReactionsideReaction

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to ReactionsideReaction slice")
	}

	if len(reactionsideReactionAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all ReactionsideReaction records in the query.
func (q reactionsideReactionQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count reactionside_reaction rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q reactionsideReactionQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if reactionside_reaction exists")
	}

	return count > 0, nil
}

// ReactionsideReactionReactionside pointed to by the foreign key.
func (o *ReactionsideReaction) ReactionsideReactionReactionside(mods ...qm.QueryMod) reactionsideQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"accession\" = ?", o.Reactionside),
	}

	queryMods = append(queryMods, mods...)

	return Reactionsides(queryMods...)
}

// ReactionsideReactionReaction pointed to by the foreign key.
func (o *ReactionsideReaction) ReactionsideReactionReaction(mods ...qm.QueryMod) reactionQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"accession\" = ?", o.Reaction),
	}

	queryMods = append(queryMods, mods...)

	return Reactions(queryMods...)
}

// LoadReactionsideReactionReactionside allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (reactionsideReactionL) LoadReactionsideReactionReactionside(ctx context.Context, e boil.ContextExecutor, singular bool, maybeReactionsideReaction interface{}, mods queries.Applicator) error {
	var slice []*ReactionsideReaction
	var object *ReactionsideReaction

	if singular {
		object = maybeReactionsideReaction.(*ReactionsideReaction)
	} else {
		slice = *maybeReactionsideReaction.(*[]*ReactionsideReaction)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &reactionsideReactionR{}
		}
		if !queries.IsNil(object.Reactionside) {
			args = append(args, object.Reactionside)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &reactionsideReactionR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.Reactionside) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.Reactionside) {
				args = append(args, obj.Reactionside)
			}

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`reactionside`),
		qm.WhereIn(`reactionside.accession in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Reactionside")
	}

	var resultSlice []*Reactionside
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Reactionside")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for reactionside")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for reactionside")
	}

	if len(reactionsideReactionAfterSelectHooks) != 0 {
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
		object.R.ReactionsideReactionReactionside = foreign
		if foreign.R == nil {
			foreign.R = &reactionsideR{}
		}
		foreign.R.ReactionsideReactions = append(foreign.R.ReactionsideReactions, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.Reactionside, foreign.Accession) {
				local.R.ReactionsideReactionReactionside = foreign
				if foreign.R == nil {
					foreign.R = &reactionsideR{}
				}
				foreign.R.ReactionsideReactions = append(foreign.R.ReactionsideReactions, local)
				break
			}
		}
	}

	return nil
}

// LoadReactionsideReactionReaction allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (reactionsideReactionL) LoadReactionsideReactionReaction(ctx context.Context, e boil.ContextExecutor, singular bool, maybeReactionsideReaction interface{}, mods queries.Applicator) error {
	var slice []*ReactionsideReaction
	var object *ReactionsideReaction

	if singular {
		object = maybeReactionsideReaction.(*ReactionsideReaction)
	} else {
		slice = *maybeReactionsideReaction.(*[]*ReactionsideReaction)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &reactionsideReactionR{}
		}
		if !queries.IsNil(object.Reaction) {
			args = append(args, object.Reaction)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &reactionsideReactionR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.Reaction) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.Reaction) {
				args = append(args, obj.Reaction)
			}

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`reaction`),
		qm.WhereIn(`reaction.accession in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Reaction")
	}

	var resultSlice []*Reaction
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Reaction")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for reaction")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for reaction")
	}

	if len(reactionsideReactionAfterSelectHooks) != 0 {
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
		object.R.ReactionsideReactionReaction = foreign
		if foreign.R == nil {
			foreign.R = &reactionR{}
		}
		foreign.R.ReactionsideReactions = append(foreign.R.ReactionsideReactions, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.Reaction, foreign.Accession) {
				local.R.ReactionsideReactionReaction = foreign
				if foreign.R == nil {
					foreign.R = &reactionR{}
				}
				foreign.R.ReactionsideReactions = append(foreign.R.ReactionsideReactions, local)
				break
			}
		}
	}

	return nil
}

// SetReactionsideReactionReactionside of the reactionsideReaction to the related item.
// Sets o.R.ReactionsideReactionReactionside to related.
// Adds o to related.R.ReactionsideReactions.
func (o *ReactionsideReaction) SetReactionsideReactionReactionside(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Reactionside) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"reactionside_reaction\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, []string{"reactionside"}),
		strmangle.WhereClause("\"", "\"", 0, reactionsideReactionPrimaryKeyColumns),
	)
	values := []interface{}{related.Accession, o.Reaction, o.Reactionside}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	queries.Assign(&o.Reactionside, related.Accession)
	if o.R == nil {
		o.R = &reactionsideReactionR{
			ReactionsideReactionReactionside: related,
		}
	} else {
		o.R.ReactionsideReactionReactionside = related
	}

	if related.R == nil {
		related.R = &reactionsideR{
			ReactionsideReactions: ReactionsideReactionSlice{o},
		}
	} else {
		related.R.ReactionsideReactions = append(related.R.ReactionsideReactions, o)
	}

	return nil
}

// SetReactionsideReactionReaction of the reactionsideReaction to the related item.
// Sets o.R.ReactionsideReactionReaction to related.
// Adds o to related.R.ReactionsideReactions.
func (o *ReactionsideReaction) SetReactionsideReactionReaction(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Reaction) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"reactionside_reaction\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, []string{"reaction"}),
		strmangle.WhereClause("\"", "\"", 0, reactionsideReactionPrimaryKeyColumns),
	)
	values := []interface{}{related.Accession, o.Reaction, o.Reactionside}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	queries.Assign(&o.Reaction, related.Accession)
	if o.R == nil {
		o.R = &reactionsideReactionR{
			ReactionsideReactionReaction: related,
		}
	} else {
		o.R.ReactionsideReactionReaction = related
	}

	if related.R == nil {
		related.R = &reactionR{
			ReactionsideReactions: ReactionsideReactionSlice{o},
		}
	} else {
		related.R.ReactionsideReactions = append(related.R.ReactionsideReactions, o)
	}

	return nil
}

// ReactionsideReactions retrieves all the records using an executor.
func ReactionsideReactions(mods ...qm.QueryMod) reactionsideReactionQuery {
	mods = append(mods, qm.From("\"reactionside_reaction\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"reactionside_reaction\".*"})
	}

	return reactionsideReactionQuery{q}
}

// FindReactionsideReaction retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindReactionsideReaction(ctx context.Context, exec boil.ContextExecutor, reaction string, reactionside string, selectCols ...string) (*ReactionsideReaction, error) {
	reactionsideReactionObj := &ReactionsideReaction{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"reactionside_reaction\" where \"reaction\"=? AND \"reactionside\"=?", sel,
	)

	q := queries.Raw(query, reaction, reactionside)

	err := q.Bind(ctx, exec, reactionsideReactionObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from reactionside_reaction")
	}

	if err = reactionsideReactionObj.doAfterSelectHooks(ctx, exec); err != nil {
		return reactionsideReactionObj, err
	}

	return reactionsideReactionObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *ReactionsideReaction) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no reactionside_reaction provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(reactionsideReactionColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	reactionsideReactionInsertCacheMut.RLock()
	cache, cached := reactionsideReactionInsertCache[key]
	reactionsideReactionInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			reactionsideReactionAllColumns,
			reactionsideReactionColumnsWithDefault,
			reactionsideReactionColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(reactionsideReactionType, reactionsideReactionMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(reactionsideReactionType, reactionsideReactionMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"reactionside_reaction\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"reactionside_reaction\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into reactionside_reaction")
	}

	if !cached {
		reactionsideReactionInsertCacheMut.Lock()
		reactionsideReactionInsertCache[key] = cache
		reactionsideReactionInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the ReactionsideReaction.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *ReactionsideReaction) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	reactionsideReactionUpdateCacheMut.RLock()
	cache, cached := reactionsideReactionUpdateCache[key]
	reactionsideReactionUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			reactionsideReactionAllColumns,
			reactionsideReactionPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update reactionside_reaction, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"reactionside_reaction\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, reactionsideReactionPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(reactionsideReactionType, reactionsideReactionMapping, append(wl, reactionsideReactionPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update reactionside_reaction row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for reactionside_reaction")
	}

	if !cached {
		reactionsideReactionUpdateCacheMut.Lock()
		reactionsideReactionUpdateCache[key] = cache
		reactionsideReactionUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q reactionsideReactionQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for reactionside_reaction")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for reactionside_reaction")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ReactionsideReactionSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), reactionsideReactionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"reactionside_reaction\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, reactionsideReactionPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in reactionsideReaction slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all reactionsideReaction")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *ReactionsideReaction) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no reactionside_reaction provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(reactionsideReactionColumnsWithDefault, o)

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

	reactionsideReactionUpsertCacheMut.RLock()
	cache, cached := reactionsideReactionUpsertCache[key]
	reactionsideReactionUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			reactionsideReactionAllColumns,
			reactionsideReactionColumnsWithDefault,
			reactionsideReactionColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			reactionsideReactionAllColumns,
			reactionsideReactionPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert reactionside_reaction, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(reactionsideReactionPrimaryKeyColumns))
			copy(conflict, reactionsideReactionPrimaryKeyColumns)
		}
		cache.query = buildUpsertQuerySQLite(dialect, "\"reactionside_reaction\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(reactionsideReactionType, reactionsideReactionMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(reactionsideReactionType, reactionsideReactionMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert reactionside_reaction")
	}

	if !cached {
		reactionsideReactionUpsertCacheMut.Lock()
		reactionsideReactionUpsertCache[key] = cache
		reactionsideReactionUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single ReactionsideReaction record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *ReactionsideReaction) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no ReactionsideReaction provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), reactionsideReactionPrimaryKeyMapping)
	sql := "DELETE FROM \"reactionside_reaction\" WHERE \"reaction\"=? AND \"reactionside\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from reactionside_reaction")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for reactionside_reaction")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q reactionsideReactionQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no reactionsideReactionQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from reactionside_reaction")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for reactionside_reaction")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ReactionsideReactionSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(reactionsideReactionBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), reactionsideReactionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"reactionside_reaction\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, reactionsideReactionPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from reactionsideReaction slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for reactionside_reaction")
	}

	if len(reactionsideReactionAfterDeleteHooks) != 0 {
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
func (o *ReactionsideReaction) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindReactionsideReaction(ctx, exec, o.Reaction, o.Reactionside)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ReactionsideReactionSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ReactionsideReactionSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), reactionsideReactionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"reactionside_reaction\".* FROM \"reactionside_reaction\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, reactionsideReactionPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ReactionsideReactionSlice")
	}

	*o = slice

	return nil
}

// ReactionsideReactionExists checks if the ReactionsideReaction row exists.
func ReactionsideReactionExists(ctx context.Context, exec boil.ContextExecutor, reaction string, reactionside string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"reactionside_reaction\" where \"reaction\"=? AND \"reactionside\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, reaction, reactionside)
	}
	row := exec.QueryRowContext(ctx, sql, reaction, reactionside)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if reactionside_reaction exists")
	}

	return exists, nil
}
