package organizationlogic

import (
	"context"
	"fmt"
	"github.com/dembygenesis/local.tools/internal/model"
	"github.com/dembygenesis/local.tools/internal/persistence"
	"github.com/dembygenesis/local.tools/internal/sysconsts"
	"github.com/dembygenesis/local.tools/internal/utilities/errs"
	"github.com/dembygenesis/local.tools/internal/utilities/strutil"
	"github.com/dembygenesis/local.tools/internal/utilities/validationutils"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

type Config struct {
	TxProvider persistence.TransactionProvider `json:"tx_provider" validate:"required"`
	Logger     *logrus.Entry                   `json:"Logger" validate:"required"`
	Persistor  persistor                       `json:"Persistor" validate:"required"`
}

func (i *Config) Validate() error {
	return validationutils.Validate(i)
}

type Service struct {
	cfg *Config
}

func New(cfg *Config) (*Service, error) {
	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("validate: %w", err)
	}
	return &Service{cfg}, nil
}

func (i *Service) CreateOrganization(ctx context.Context, params *model.CreateOrganization) (*model.Organization, error) {
	if err := params.Validate(); err != nil {
		return nil, errs.New(&errs.Cfg{
			StatusCode: http.StatusBadRequest,
			Err:        fmt.Errorf("validate: %w", err),
		})
	}

	tx, err := i.cfg.TxProvider.Tx(ctx)
	if err != nil {
		return nil, errs.New(&errs.Cfg{
			StatusCode: http.StatusInternalServerError,
			Err:        fmt.Errorf("get db: %w", err),
		})
	}
	defer tx.Rollback(ctx)

	organization := params.ToOrganization()

	exists, err := i.cfg.Persistor.GetOrganizationByName(ctx, tx, organization.Name)
	if err != nil {
		if !strings.Contains(err.Error(), sysconsts.ErrExpectedExactlyOneEntry) {
			return nil, errs.New(&errs.Cfg{
				StatusCode: http.StatusBadRequest,
				Err:        fmt.Errorf("check category unique: %w", err),
			})
		}
	}
	if exists != nil {
		return nil, errs.New(&errs.Cfg{
			StatusCode: http.StatusBadRequest,
			Err:        fmt.Errorf("category already exists"),
		})
	}

	organizationCreated, err := i.cfg.Persistor.CreateOrganization(ctx, tx, organization)
	if err != nil {
		return nil, errs.New(&errs.Cfg{
			StatusCode: http.StatusInternalServerError,
			Err:        fmt.Errorf("create: %w", err),
		})
	}

	if err = tx.Commit(ctx); err != nil {
		return nil, errs.New(&errs.Cfg{
			StatusCode: http.StatusInternalServerError,
			Err:        fmt.Errorf("commit: %v", err),
		})
	}

	return organizationCreated, nil
}

func (i *Service) AddOrganization(ctx context.Context, params *model.CreateOrganization) (*model.Organization, error) {
	if err := params.Validate(); err != nil {
		return nil, errs.New(&errs.Cfg{
			StatusCode: http.StatusBadRequest,
			Err:        fmt.Errorf("validate: %w", err),
		})
	}

	tx, err := i.cfg.TxProvider.Tx(ctx)
	if err != nil {
		return nil, errs.New(&errs.Cfg{
			StatusCode: http.StatusInternalServerError,
			Err:        fmt.Errorf("get db: %w", err),
		})
	}
	defer tx.Rollback(ctx)

	// Might still be useful
	//organization := params.ToOrganization()
	//userId, err := strconv.Atoi(organization.CreatedBy)
	//
	//toBeCreatedOrganization := &model.CreateOrganization{
	//	Name:   organization.Name,
	//	UserId: userId,
	//}

	exists, err := i.cfg.Persistor.GetOrganizationByName(ctx, tx, params.Name)
	if err != nil {
		if !strings.Contains(err.Error(), sysconsts.ErrExpectedExactlyOneEntry) {
			return nil, errs.New(&errs.Cfg{
				StatusCode: http.StatusBadRequest,
				Err:        fmt.Errorf("check category unique: %v", err),
			})
		}
	}
	if exists != nil {
		return nil, errs.New(&errs.Cfg{
			StatusCode: http.StatusBadRequest,
			Err:        fmt.Errorf("category already exists"),
		})
	}

	organization, err := i.cfg.Persistor.AddOrganization(ctx, tx, params)
	if err != nil {
		return nil, errs.New(&errs.Cfg{
			StatusCode: http.StatusInternalServerError,
			Err:        fmt.Errorf("create: %w", err),
		})
	}

	if err = tx.Commit(ctx); err != nil {
		return nil, errs.New(&errs.Cfg{
			StatusCode: http.StatusInternalServerError,
			Err:        fmt.Errorf("commit: %w", err),
		})
	}

	return organization, nil
}

func (i *Service) ListOrganizations(
	ctx context.Context,
	filter *model.OrganizationFilters,
) (*model.PaginatedOrganizations, error) {
	db, err := i.cfg.TxProvider.Db(ctx)
	if err != nil {
		return nil, errs.New(&errs.Cfg{
			StatusCode: http.StatusInternalServerError,
			Err:        fmt.Errorf("get db: %v", err),
		})
	}

	paginated, err := i.cfg.Persistor.GetOrganizations(ctx, db, filter)
	if err != nil {
		return nil, errs.New(&errs.Cfg{
			StatusCode: http.StatusInternalServerError,
			Err:        fmt.Errorf("get organizations: %v", err),
		})
	}

	return paginated, nil
}

func (i *Service) UpdateOrganization(ctx context.Context, params *model.UpdateOrganization) (*model.Organization, error) {
	tx, err := i.cfg.TxProvider.Tx(ctx)
	if err != nil {
		return nil, errs.New(&errs.Cfg{
			StatusCode: http.StatusInternalServerError,
			Err:        fmt.Errorf("get db: %w", err),
		})
	}
	defer tx.Rollback(ctx)

	if err := params.Validate(); err != nil {
		return nil, fmt.Errorf("validate: %w", err)
	}

	organization, err := i.cfg.Persistor.UpdateOrganization(ctx, tx, params)
	if err != nil {
		return nil, fmt.Errorf("update organization: %w", err)
	}

	return organization, nil
}

func (s *Service) DeleteOrganization(ctx context.Context, params *model.DeleteOrganization) error {
	tx, err := s.cfg.TxProvider.Tx(ctx)
	if err != nil {
		return errs.New(&errs.Cfg{
			StatusCode: http.StatusInternalServerError,
			Err:        fmt.Errorf("get db: %w", err),
		})
	}
	defer tx.Rollback(ctx)

	err = s.cfg.Persistor.DeleteOrganization(ctx, tx, params.ID)
	if err != nil {
		return errs.New(&errs.Cfg{
			StatusCode: http.StatusInternalServerError,
			Err:        fmt.Errorf("delete organization: %w", err),
		})
	}

	if err := tx.Commit(ctx); err != nil {
		return errs.New(&errs.Cfg{
			StatusCode: http.StatusInternalServerError,
			Err:        fmt.Errorf("commit transaction: %w", err),
		})
	}

	return nil
}

func (s *Service) RestoreOrganization(ctx context.Context, params *model.RestoreOrganization) error {
	tx, err := s.cfg.TxProvider.Tx(ctx)
	if err != nil {
		return errs.New(&errs.Cfg{
			StatusCode: http.StatusInternalServerError,
			Err:        fmt.Errorf("get db: %w", err),
		})
	}
	defer tx.Rollback(ctx)

	fmt.Println("the params --- ", strutil.GetAsJson(params))

	err = s.cfg.Persistor.RestoreOrganization(ctx, tx, params.ID)
	if err != nil {
		return errs.New(&errs.Cfg{
			StatusCode: http.StatusInternalServerError,
			Err:        fmt.Errorf("restore organization: %w", err),
		})
	}

	if err := tx.Commit(ctx); err != nil {
		return errs.New(&errs.Cfg{
			StatusCode: http.StatusInternalServerError,
			Err:        fmt.Errorf("commit transaction: %w", err),
		})
	}

	return nil
}
