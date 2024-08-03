package mysqlstore

import (
	"context"
	"fmt"
	"github.com/dembygenesis/local.tools/internal/model"
	"github.com/dembygenesis/local.tools/internal/persistence"
	"github.com/dembygenesis/local.tools/internal/persistence/database_helpers/mysql/assets/mysqlmodel"
	"github.com/dembygenesis/local.tools/internal/persistence/database_helpers/mysql/mysqltx"
	"github.com/dembygenesis/local.tools/internal/sysconsts"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (m *Repository) DeleteOrganization(
	ctx context.Context,
	tx persistence.TransactionHandler,
	id int,
) error {
	ctxExec, err := mysqltx.GetCtxExecutor(tx)
	if err != nil {
		return fmt.Errorf("get ctx exec: %v", err)
	}

	entry := &mysqlmodel.Organization{
		ID:       id,
		IsActive: false,
	}
	if _, err = entry.Update(ctx, ctxExec, boil.Whitelist("is_active")); err != nil {
		return fmt.Errorf("delete: %v", err)
	}

	return nil
}

func (m *Repository) UpdateOrganization(ctx context.Context, tx persistence.TransactionHandler, params *model.UpdateOrganization) (*model.Organization, error) {
	if params == nil {
		return nil, ErrOrgNil
	}
	ctxExec, err := mysqltx.GetCtxExecutor(tx)
	if err != nil {
		return nil, fmt.Errorf("extract context executor: %v", err)
	}

	entry := &mysqlmodel.Organization{ID: params.Id}
	cols := []string{mysqlmodel.OrganizationColumns.ID}

	//if params.OrganizationTypeRefId.Valid {
	//    entry.OrganizationRefUsers() = params.OrganizationTypeRefId.Int
	//    cols = append(cols, mysqlmodel.OrganizationColumns.IsActive)
	//}

	if params.Name.Valid {
		entry.Name = params.Name.String
		cols = append(cols, mysqlmodel.OrganizationColumns.Name)
	}

	_, err = entry.Update(ctx, ctxExec, boil.Whitelist(cols...))
	if err != nil {
		return nil, fmt.Errorf("update failed: %v", err)
	}

	organization, err := m.GetOrganizationById(ctx, tx, entry.ID)
	if err != nil {
		return nil, fmt.Errorf("get organitzation by id: %v", err)
	}

	return organization, nil
}

func (m *Repository) GetOrganizationById(ctx context.Context, tx persistence.TransactionHandler, id int) (*model.Organization, error) {
	paginated, err := m.GetOrganizations(ctx, tx, &model.OrganizationFilters{
		IdsIn: []int{id},
	})
	if err != nil {
		return nil, fmt.Errorf("organization filtered by id: %v", err)
	}

	if paginated.Pagination.RowCount != 1 {

		return nil, fmt.Errorf(sysconsts.ErrExpectedExactlyOneEntry, id)
	}

	return &paginated.Organizations[0], nil
}

func (m *Repository) GetOrganizations(ctx context.Context, tx persistence.TransactionHandler, filters *model.OrganizationFilters) (*model.PaginatedOrganization, error) {
	ctxExec, err := mysqltx.GetCtxExecutor(tx)
	if err != nil {
		return nil, fmt.Errorf("extract context executor: %v", err)
	}

	res, err := m.getOrganizations(ctx, ctxExec, filters)

	if err != nil {
		return nil, fmt.Errorf("read organizations: %v", err)
	}

	return res, nil
}

func (m *Repository) getOrganizations(ctx context.Context,
	ctxExec boil.ContextExecutor,
	filters *model.OrganizationFilters) (*model.PaginatedOrganization, error) {

	var (
		paginated  model.PaginatedOrganization
		pagination = model.NewPagination()
		res        = make([]model.Organization, 0)
		err        error
	)

	ctx, cancel := context.WithTimeout(ctx, m.cfg.QueryTimeouts.Query)
	defer cancel()

	queryMods := []qm.QueryMod{
		qm.Select(
			fmt.Sprintf("%s.%s AS %s",
				mysqlmodel.TableNames.Organization,
				mysqlmodel.OrganizationColumns.ID,
				mysqlmodel.OrganizationColumns.ID,
			),
			fmt.Sprintf("%s.%s AS %s",
				mysqlmodel.TableNames.Organization,
				mysqlmodel.OrganizationColumns.Name,
				mysqlmodel.OrganizationColumns.Name,
			),
			fmt.Sprintf("%s.%s AS %s",
				mysqlmodel.TableNames.Organization,
				mysqlmodel.OrganizationColumns.IsActive,
				mysqlmodel.OrganizationColumns.IsActive,
			),
		),
	}

	if filters != nil {
		if len(filters.IdsIn) > 0 {
			queryMods = append(queryMods, mysqlmodel.OrganizationWhere.ID.IN(filters.IdsIn))
		}

		if len(filters.OrganizationIsActive) > 0 {
			queryMods = append(queryMods, mysqlmodel.OrganizationWhere.IsActive.EQ(true))
		}

		if len(filters.OrganizationNameIn) > 0 {
			queryMods = append(queryMods, mysqlmodel.OrganizationWhere.Name.IN(filters.OrganizationNameIn))
		}
	}

	q := mysqlmodel.Organizations(queryMods...)
	totalCount, err := q.Count(ctx, ctxExec)
	if err != nil {
		return nil, fmt.Errorf("get organizations count: %v", err)
	}

	page := pagination.Page
	maxRows := pagination.MaxRows
	if filters != nil {
		if filters.Page.Valid {
			page = filters.Page.Int
		}
		if filters.MaxRows.Valid {
			maxRows = filters.MaxRows.Int
		}
	}

	pagination.SetQueryBoundaries(page, maxRows, int(totalCount))

	queryMods = append(queryMods, qm.Limit(pagination.MaxRows), qm.Offset(pagination.Offset))
	q = mysqlmodel.Organizations(queryMods...)

	if err = q.Bind(ctx, ctxExec, &res); err != nil {
		return nil, fmt.Errorf("get organizations: %v", err)
	}

	pagination.RowCount = len(res)
	paginated.Organizations = res
	paginated.Pagination = pagination

	return &paginated, nil
}

func (m *Repository) CreateOrganization(ctx context.Context, tx persistence.TransactionHandler, organization *model.Organization) (*model.Organization, error) {
	ctxExec, err := mysqltx.GetCtxExecutor(tx)
	if err != nil {
		return nil, fmt.Errorf("extract context executor: %v", err)
	}

	entry := mysqlmodel.Organization{
		Name: organization.Name,
	}
	if err = entry.Insert(ctx, ctxExec, boil.Infer()); err != nil {
		return nil, fmt.Errorf("insert organization: %v", err)
	}

	organization, err = m.GetOrganizationById(ctx, tx, entry.ID)
	if err != nil {
		return nil, fmt.Errorf("get organization by id: %v", err)
	}

	return organization, nil
}

func (m *Repository) AddOrganization(ctx context.Context, tx persistence.TransactionHandler, organization *model.Organization) (*model.Organization, error) {
	ctxExec, err := mysqltx.GetCtxExecutor(tx)
	if err != nil {
		return nil, fmt.Errorf("extract context executor: %v", err)
	}

	entry := &mysqlmodel.Organization{
		Name: organization.Name,
	}
	if err = entry.Insert(ctx, ctxExec, boil.Infer()); err != nil {
		return nil, fmt.Errorf("insert organization: %v", err)
	}

	organization, err = m.GetOrganizationById(ctx, tx, entry.ID)
	if err != nil {
		return nil, fmt.Errorf("get organization by id: %v", err)
	}

	return organization, nil
}

func (m *Repository) RestoreOrganization(
	ctx context.Context,
	tx persistence.TransactionHandler,
	id int,
) error {
	ctxExec, err := mysqltx.GetCtxExecutor(tx)
	if err != nil {
		return fmt.Errorf("get ctx exec: %v", err)
	}

	entry := &mysqlmodel.Organization{ID: id, IsActive: true}
	if _, err = entry.Update(ctx, ctxExec, boil.Whitelist("is_active")); err != nil {
		return fmt.Errorf("restore: %w", err)
	}

	return nil
}
