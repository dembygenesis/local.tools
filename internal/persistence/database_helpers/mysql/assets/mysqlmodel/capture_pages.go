// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package mysqlmodel

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
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// CapturePage is an object representing the database table.
type CapturePage struct {
	ID               int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name             string    `boil:"name" json:"name" toml:"name" yaml:"name"`
	HTML             string    `boil:"html" json:"html" toml:"html" yaml:"html"`
	Clicks           int       `boil:"clicks" json:"clicks" toml:"clicks" yaml:"clicks"`
	IsControl        int       `boil:"is_control" json:"is_control" toml:"is_control" yaml:"is_control"`
	CapturePageSetID int       `boil:"capture_page_set_id" json:"capture_page_set_id" toml:"capture_page_set_id" yaml:"capture_page_set_id"`
	CreatedBy        int       `boil:"created_by" json:"created_by" toml:"created_by" yaml:"created_by"`
	UpdatedBy        int       `boil:"updated_by" json:"updated_by" toml:"updated_by" yaml:"updated_by"`
	LastImpressionAt null.Time `boil:"last_impression_at" json:"last_impression_at,omitempty" toml:"last_impression_at" yaml:"last_impression_at,omitempty"`
	Impressions      int       `boil:"impressions" json:"impressions" toml:"impressions" yaml:"impressions"`
	CreatedAt        null.Time `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	UpdatedAt        null.Time `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
	DeletedAt        null.Time `boil:"deleted_at" json:"deleted_at,omitempty" toml:"deleted_at" yaml:"deleted_at,omitempty"`

	R *capturePageR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L capturePageL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CapturePageColumns = struct {
	ID               string
	Name             string
	HTML             string
	Clicks           string
	IsControl        string
	CapturePageSetID string
	CreatedBy        string
	UpdatedBy        string
	LastImpressionAt string
	Impressions      string
	CreatedAt        string
	UpdatedAt        string
	DeletedAt        string
}{
	ID:               "id",
	Name:             "name",
	HTML:             "html",
	Clicks:           "clicks",
	IsControl:        "is_control",
	CapturePageSetID: "capture_page_set_id",
	CreatedBy:        "created_by",
	UpdatedBy:        "updated_by",
	LastImpressionAt: "last_impression_at",
	Impressions:      "impressions",
	CreatedAt:        "created_at",
	UpdatedAt:        "updated_at",
	DeletedAt:        "deleted_at",
}

var CapturePageTableColumns = struct {
	ID               string
	Name             string
	HTML             string
	Clicks           string
	IsControl        string
	CapturePageSetID string
	CreatedBy        string
	UpdatedBy        string
	LastImpressionAt string
	Impressions      string
	CreatedAt        string
	UpdatedAt        string
	DeletedAt        string
}{
	ID:               "capture_pages.id",
	Name:             "capture_pages.name",
	HTML:             "capture_pages.html",
	Clicks:           "capture_pages.clicks",
	IsControl:        "capture_pages.is_control",
	CapturePageSetID: "capture_pages.capture_page_set_id",
	CreatedBy:        "capture_pages.created_by",
	UpdatedBy:        "capture_pages.updated_by",
	LastImpressionAt: "capture_pages.last_impression_at",
	Impressions:      "capture_pages.impressions",
	CreatedAt:        "capture_pages.created_at",
	UpdatedAt:        "capture_pages.updated_at",
	DeletedAt:        "capture_pages.deleted_at",
}

// Generated where

var CapturePageWhere = struct {
	ID               whereHelperint
	Name             whereHelperstring
	HTML             whereHelperstring
	Clicks           whereHelperint
	IsControl        whereHelperint
	CapturePageSetID whereHelperint
	CreatedBy        whereHelperint
	UpdatedBy        whereHelperint
	LastImpressionAt whereHelpernull_Time
	Impressions      whereHelperint
	CreatedAt        whereHelpernull_Time
	UpdatedAt        whereHelpernull_Time
	DeletedAt        whereHelpernull_Time
}{
	ID:               whereHelperint{field: "`capture_pages`.`id`"},
	Name:             whereHelperstring{field: "`capture_pages`.`name`"},
	HTML:             whereHelperstring{field: "`capture_pages`.`html`"},
	Clicks:           whereHelperint{field: "`capture_pages`.`clicks`"},
	IsControl:        whereHelperint{field: "`capture_pages`.`is_control`"},
	CapturePageSetID: whereHelperint{field: "`capture_pages`.`capture_page_set_id`"},
	CreatedBy:        whereHelperint{field: "`capture_pages`.`created_by`"},
	UpdatedBy:        whereHelperint{field: "`capture_pages`.`updated_by`"},
	LastImpressionAt: whereHelpernull_Time{field: "`capture_pages`.`last_impression_at`"},
	Impressions:      whereHelperint{field: "`capture_pages`.`impressions`"},
	CreatedAt:        whereHelpernull_Time{field: "`capture_pages`.`created_at`"},
	UpdatedAt:        whereHelpernull_Time{field: "`capture_pages`.`updated_at`"},
	DeletedAt:        whereHelpernull_Time{field: "`capture_pages`.`deleted_at`"},
}

// CapturePageRels is where relationship names are stored.
var CapturePageRels = struct {
	CapturePageSet string
}{
	CapturePageSet: "CapturePageSet",
}

// capturePageR is where relationships are stored.
type capturePageR struct {
	CapturePageSet *CapturePageSet `boil:"CapturePageSet" json:"CapturePageSet" toml:"CapturePageSet" yaml:"CapturePageSet"`
}

// NewStruct creates a new relationship struct
func (*capturePageR) NewStruct() *capturePageR {
	return &capturePageR{}
}

func (r *capturePageR) GetCapturePageSet() *CapturePageSet {
	if r == nil {
		return nil
	}
	return r.CapturePageSet
}

// capturePageL is where Load methods for each relationship are stored.
type capturePageL struct{}

var (
	capturePageAllColumns            = []string{"id", "name", "html", "clicks", "is_control", "capture_page_set_id", "created_by", "updated_by", "last_impression_at", "impressions", "created_at", "updated_at", "deleted_at"}
	capturePageColumnsWithoutDefault = []string{"name", "html", "capture_page_set_id", "created_by", "updated_by", "last_impression_at", "created_at", "updated_at", "deleted_at"}
	capturePageColumnsWithDefault    = []string{"id", "clicks", "is_control", "impressions"}
	capturePagePrimaryKeyColumns     = []string{"id"}
	capturePageGeneratedColumns      = []string{}
)

type (
	// CapturePageSlice is an alias for a slice of pointers to CapturePage.
	// This should almost always be used instead of []CapturePage.
	CapturePageSlice []*CapturePage

	capturePageQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	capturePageType                 = reflect.TypeOf(&CapturePage{})
	capturePageMapping              = queries.MakeStructMapping(capturePageType)
	capturePagePrimaryKeyMapping, _ = queries.BindMapping(capturePageType, capturePageMapping, capturePagePrimaryKeyColumns)
	capturePageInsertCacheMut       sync.RWMutex
	capturePageInsertCache          = make(map[string]insertCache)
	capturePageUpdateCacheMut       sync.RWMutex
	capturePageUpdateCache          = make(map[string]updateCache)
	capturePageUpsertCacheMut       sync.RWMutex
	capturePageUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single capturePage record from the query.
func (q capturePageQuery) One(ctx context.Context, exec boil.ContextExecutor) (*CapturePage, error) {
	o := &CapturePage{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "mysqlmodel: failed to execute a one query for capture_pages")
	}

	return o, nil
}

// All returns all CapturePage records from the query.
func (q capturePageQuery) All(ctx context.Context, exec boil.ContextExecutor) (CapturePageSlice, error) {
	var o []*CapturePage

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "mysqlmodel: failed to assign all query results to CapturePage slice")
	}

	return o, nil
}

// Count returns the count of all CapturePage records in the query.
func (q capturePageQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "mysqlmodel: failed to count capture_pages rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q capturePageQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "mysqlmodel: failed to check if capture_pages exists")
	}

	return count > 0, nil
}

// CapturePageSet pointed to by the foreign key.
func (o *CapturePage) CapturePageSet(mods ...qm.QueryMod) capturePageSetQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`id` = ?", o.CapturePageSetID),
	}

	queryMods = append(queryMods, mods...)

	return CapturePageSets(queryMods...)
}

// LoadCapturePageSet allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (capturePageL) LoadCapturePageSet(ctx context.Context, e boil.ContextExecutor, singular bool, maybeCapturePage interface{}, mods queries.Applicator) error {
	var slice []*CapturePage
	var object *CapturePage

	if singular {
		var ok bool
		object, ok = maybeCapturePage.(*CapturePage)
		if !ok {
			object = new(CapturePage)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeCapturePage)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeCapturePage))
			}
		}
	} else {
		s, ok := maybeCapturePage.(*[]*CapturePage)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeCapturePage)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeCapturePage))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &capturePageR{}
		}
		args[object.CapturePageSetID] = struct{}{}

	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &capturePageR{}
			}

			args[obj.CapturePageSetID] = struct{}{}

		}
	}

	if len(args) == 0 {
		return nil
	}

	argsSlice := make([]interface{}, len(args))
	i := 0
	for arg := range args {
		argsSlice[i] = arg
		i++
	}

	query := NewQuery(
		qm.From(`capture_page_sets`),
		qm.WhereIn(`capture_page_sets.id in ?`, argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load CapturePageSet")
	}

	var resultSlice []*CapturePageSet
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice CapturePageSet")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for capture_page_sets")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for capture_page_sets")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.CapturePageSet = foreign
		if foreign.R == nil {
			foreign.R = &capturePageSetR{}
		}
		foreign.R.CapturePages = append(foreign.R.CapturePages, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.CapturePageSetID == foreign.ID {
				local.R.CapturePageSet = foreign
				if foreign.R == nil {
					foreign.R = &capturePageSetR{}
				}
				foreign.R.CapturePages = append(foreign.R.CapturePages, local)
				break
			}
		}
	}

	return nil
}

// SetCapturePageSet of the capturePage to the related item.
// Sets o.R.CapturePageSet to related.
// Adds o to related.R.CapturePages.
func (o *CapturePage) SetCapturePageSet(ctx context.Context, exec boil.ContextExecutor, insert bool, related *CapturePageSet) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `capture_pages` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"capture_page_set_id"}),
		strmangle.WhereClause("`", "`", 0, capturePagePrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.CapturePageSetID = related.ID
	if o.R == nil {
		o.R = &capturePageR{
			CapturePageSet: related,
		}
	} else {
		o.R.CapturePageSet = related
	}

	if related.R == nil {
		related.R = &capturePageSetR{
			CapturePages: CapturePageSlice{o},
		}
	} else {
		related.R.CapturePages = append(related.R.CapturePages, o)
	}

	return nil
}

// CapturePages retrieves all the records using an executor.
func CapturePages(mods ...qm.QueryMod) capturePageQuery {
	mods = append(mods, qm.From("`capture_pages`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`capture_pages`.*"})
	}

	return capturePageQuery{q}
}

// FindCapturePage retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCapturePage(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*CapturePage, error) {
	capturePageObj := &CapturePage{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `capture_pages` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, capturePageObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "mysqlmodel: unable to select from capture_pages")
	}

	return capturePageObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *CapturePage) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("mysqlmodel: no capture_pages provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
		if queries.MustTime(o.UpdatedAt).IsZero() {
			queries.SetScanner(&o.UpdatedAt, currTime)
		}
	}

	nzDefaults := queries.NonZeroDefaultSet(capturePageColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	capturePageInsertCacheMut.RLock()
	cache, cached := capturePageInsertCache[key]
	capturePageInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			capturePageAllColumns,
			capturePageColumnsWithDefault,
			capturePageColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(capturePageType, capturePageMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(capturePageType, capturePageMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `capture_pages` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `capture_pages` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `capture_pages` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, capturePagePrimaryKeyColumns))
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
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "mysqlmodel: unable to insert into capture_pages")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == capturePageMapping["id"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "mysqlmodel: unable to populate default values for capture_pages")
	}

CacheNoHooks:
	if !cached {
		capturePageInsertCacheMut.Lock()
		capturePageInsertCache[key] = cache
		capturePageInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the CapturePage.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *CapturePage) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	key := makeCacheKey(columns, nil)
	capturePageUpdateCacheMut.RLock()
	cache, cached := capturePageUpdateCache[key]
	capturePageUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			capturePageAllColumns,
			capturePagePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("mysqlmodel: unable to update capture_pages, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `capture_pages` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, capturePagePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(capturePageType, capturePageMapping, append(wl, capturePagePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "mysqlmodel: unable to update capture_pages row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "mysqlmodel: failed to get rows affected by update for capture_pages")
	}

	if !cached {
		capturePageUpdateCacheMut.Lock()
		capturePageUpdateCache[key] = cache
		capturePageUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q capturePageQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "mysqlmodel: unable to update all for capture_pages")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "mysqlmodel: unable to retrieve rows affected for capture_pages")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CapturePageSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("mysqlmodel: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), capturePagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `capture_pages` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, capturePagePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "mysqlmodel: unable to update all in capturePage slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "mysqlmodel: unable to retrieve rows affected all in update all capturePage")
	}
	return rowsAff, nil
}

var mySQLCapturePageUniqueColumns = []string{
	"id",
	"name",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *CapturePage) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("mysqlmodel: no capture_pages provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	nzDefaults := queries.NonZeroDefaultSet(capturePageColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLCapturePageUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
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
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	capturePageUpsertCacheMut.RLock()
	cache, cached := capturePageUpsertCache[key]
	capturePageUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			capturePageAllColumns,
			capturePageColumnsWithDefault,
			capturePageColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			capturePageAllColumns,
			capturePagePrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("mysqlmodel: unable to upsert capture_pages, could not build update column list")
		}

		ret := strmangle.SetComplement(capturePageAllColumns, strmangle.SetIntersect(insert, update))

		cache.query = buildUpsertQueryMySQL(dialect, "`capture_pages`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `capture_pages` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(capturePageType, capturePageMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(capturePageType, capturePageMapping, ret)
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
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "mysqlmodel: unable to upsert for capture_pages")
	}

	var lastID int64
	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == capturePageMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(capturePageType, capturePageMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "mysqlmodel: unable to retrieve unique values for capture_pages")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "mysqlmodel: unable to populate default values for capture_pages")
	}

CacheNoHooks:
	if !cached {
		capturePageUpsertCacheMut.Lock()
		capturePageUpsertCache[key] = cache
		capturePageUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single CapturePage record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *CapturePage) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("mysqlmodel: no CapturePage provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), capturePagePrimaryKeyMapping)
	sql := "DELETE FROM `capture_pages` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "mysqlmodel: unable to delete from capture_pages")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "mysqlmodel: failed to get rows affected by delete for capture_pages")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q capturePageQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("mysqlmodel: no capturePageQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "mysqlmodel: unable to delete all from capture_pages")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "mysqlmodel: failed to get rows affected by deleteall for capture_pages")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CapturePageSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), capturePagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `capture_pages` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, capturePagePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "mysqlmodel: unable to delete all from capturePage slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "mysqlmodel: failed to get rows affected by deleteall for capture_pages")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *CapturePage) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCapturePage(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CapturePageSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CapturePageSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), capturePagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `capture_pages`.* FROM `capture_pages` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, capturePagePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "mysqlmodel: unable to reload all in CapturePageSlice")
	}

	*o = slice

	return nil
}

// CapturePageExists checks if the CapturePage row exists.
func CapturePageExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `capture_pages` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "mysqlmodel: unable to check if capture_pages exists")
	}

	return exists, nil
}

// Exists checks if the CapturePage row exists.
func (o *CapturePage) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return CapturePageExists(ctx, exec, o.ID)
}
