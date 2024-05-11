package mysql_repository

import (
	"context"
	"fmt"
	"github.com/dembygenesis/local.tools/internal/model"
	"github.com/dembygenesis/local.tools/internal/persistence/database_helpers/mysql/assets/mysqlmodel"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (m *repository) GetOrganizations(ctx context.Context, t Transaction, f *model.OrganizationFilters) (model.Organizations, error) {
	var (
		organizations model.Organizations
		err           error
	)

	h, err := m.getTxHandler(t)
	if err != nil {
		return nil, fmt.Errorf("tx handler: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, m.cfg.QueryTimeout)
	defer cancel()

	queryMods := make([]qm.QueryMod, 0)

	if len(f.IdsIn) > 0 {
		queryMods = append(queryMods, mysqlmodel.OrganizationWhere.ID.IN(f.IdsIn))
	}

	if err = mysqlmodel.Organizations(queryMods...).Bind(ctx, h.getExec(), &organizations); err != nil {
		return nil, fmt.Errorf("get organization: %v", err)
	}

	return organizations, nil
}
