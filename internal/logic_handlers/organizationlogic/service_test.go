package organizationlogic

import (
	"context"
	"github.com/dembygenesis/local.tools/internal/lib/logger"
	"github.com/dembygenesis/local.tools/internal/logic_handlers/organizationlogic/organizationlogicfakes"
	"github.com/dembygenesis/local.tools/internal/model"
	"github.com/dembygenesis/local.tools/internal/persistence"
	"github.com/dembygenesis/local.tools/internal/persistence/database_helpers/mysql/assets/mysqlmodel"
	"github.com/dembygenesis/local.tools/internal/persistence/database_helpers/mysql/mysqlconn"
	"github.com/dembygenesis/local.tools/internal/persistence/database_helpers/mysql/mysqlhelper"
	"github.com/dembygenesis/local.tools/internal/persistence/database_helpers/mysql/mysqltx"
	"github.com/dembygenesis/local.tools/internal/persistence/persistencefakes"
	"github.com/dembygenesis/local.tools/internal/persistence/persistors/mysqlstore"
	"github.com/dembygenesis/local.tools/internal/persistence/persistors/mysqlstore/testhelper"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"strconv"
	"testing"
	"time"
)

var (
	mockTimeout      = 5 * time.Second
	mockLogger       = logger.New(context.TODO())
	mockDbReturnsErr = "error getting db"
)

type dependencies struct {
	Persistor  persistor
	Logger     *logrus.Entry
	TxProvider persistence.TransactionProvider
	Db         *sqlx.DB
	Cleanup    func(ignoreErrors ...bool)
}

func getConcreteDependencies(t *testing.T) (*dependencies, func(ignoreErrors ...bool)) {
	db, cp, cleanup := mysqlhelper.TestGetMockMariaDB(t)

	store, err := mysqlstore.New(&mysqlstore.Config{
		Logger: mockLogger,
		QueryTimeouts: &persistence.QueryTimeouts{
			Query: mockTimeout,
			Exec:  mockTimeout,
		},
	})
	require.NoError(t, err, "unexpected new mysqlstore error")

	tx, err := mysqltx.New(&mysqltx.Config{
		Logger:       mockLogger,
		Db:           db,
		DatabaseName: cp.Database,
	})
	require.NoError(t, err, "unexpected new mysqltx error")

	prov, err := mysqlconn.New(&mysqlconn.Config{
		Logger:    mockLogger,
		TxHandler: tx,
	})
	require.NoError(t, err, "unexpected new mysqlconn error")

	return &dependencies{
		Persistor:  store,
		TxProvider: prov,
		Logger:     mockLogger,
		Cleanup:    cleanup,
		Db:         db,
	}, cleanup
}

// get orgs
type argsGetOrganizations struct {
	ctx    context.Context
	filter *model.OrganizationFilters
}

type testCaseGetOrganizations struct {
	name            string
	getDependencies func(t *testing.T) (*dependencies, func(ignoreErrors ...bool))
	args            argsGetOrganizations
	mutations       func(t *testing.T, db *sqlx.DB) (org *model.CreateOrganization)
	assertions      func(t *testing.T, organization *model.PaginatedOrganizations, err error)
}

func getGetOrganizationTestCases() []testCaseGetOrganizations {
	return []testCaseGetOrganizations{
		{
			name:            "success",
			getDependencies: getConcreteDependencies,
			args: argsGetOrganizations{
				ctx:    context.Background(),
				filter: &model.OrganizationFilters{},
			},
			mutations: func(t *testing.T, db *sqlx.DB) (org *model.CreateOrganization) {
				entry := mysqlmodel.Organization{
					Name:          "Demby",
					CreatedBy:     null.IntFrom(1),
					LastUpdatedBy: null.IntFrom(1),
				}

				err := entry.Insert(context.Background(), db, boil.Infer())
				require.NoError(t, err, "error inserting in the user db")

				createOrg := model.CreateOrganization{
					Name:   "Demby",
					UserId: 1,
				}

				return &createOrg
			},
			assertions: func(t *testing.T, paginated *model.PaginatedOrganizations, err error) {
				require.NoError(t, err, "unexpected get organizations error")
				require.NotNil(t, paginated, "unexpected nil organizations")
				require.NotNil(t, paginated.Pagination, "unexpected nil pagination")
				require.NotNil(t, paginated.Organizations, "unexpected nil organizations")
				assert.NotEqual(t, 0, paginated.Pagination.RowCount, "unexpected total count")
			},
		},
		{
			name:            "fail-get-organizations",
			getDependencies: getConcreteDependencies,
			args: argsGetOrganizations{
				ctx:    context.Background(),
				filter: &model.OrganizationFilters{},
			},
			mutations: func(t *testing.T, db *sqlx.DB) (org *model.CreateOrganization) {
				testhelper.DropTable(t, db, mysqlmodel.TableNames.Organization)
				return nil
			},
			assertions: func(t *testing.T, paginated *model.PaginatedOrganizations, err error) {
				require.Error(t, err, "unexpected get organizations error")
				require.Contains(t, err.Error(), "get organizations:")
			},
		},
		{
			name: "fail-mock-get-db",
			getDependencies: func(t *testing.T) (*dependencies, func(ignoreErrors ...bool)) {
				cleanup := func(ignoreErrors ...bool) {

				}

				mockTxProvider := persistencefakes.FakeTransactionProvider{}
				mockTxProvider.DbReturns(nil, errors.New(mockDbReturnsErr))

				return &dependencies{
					Persistor:  &organizationlogicfakes.FakePersistor{},
					TxProvider: &mockTxProvider,
					Logger:     mockLogger,
					Cleanup:    cleanup,
				}, cleanup
			},
			args: argsGetOrganizations{
				ctx:    context.Background(),
				filter: &model.OrganizationFilters{},
			},
			mutations: func(t *testing.T, db *sqlx.DB) (org *model.CreateOrganization) {
				return nil
			},
			assertions: func(t *testing.T, paginated *model.PaginatedOrganizations, err error) {
				require.Error(t, err, "unexpected get organizations error")
				require.Contains(t, err.Error(), "get db:")
			},
		},
	}
}

func TestService_GetOrganizations(t *testing.T) {
	for _, tt := range getGetOrganizationTestCases() {
		t.Run(tt.name, func(t *testing.T) {
			_dependencies, cleanup := tt.getDependencies(t)
			defer cleanup()

			svc, err := New(&Config{
				TxProvider: _dependencies.TxProvider,
				Logger:     _dependencies.Logger,
				Persistor:  _dependencies.Persistor,
			})
			require.NoError(t, err, "unexpected new error")

			tt.mutations(t, _dependencies.Db)

			paginatedOrganizations, err := svc.ListOrganizations(tt.args.ctx, tt.args.filter)
			tt.assertions(t, paginatedOrganizations, err)
		})
	}
}

// update org
type argsUpdateOrganizations struct {
	ctx    context.Context
	params *model.UpdateOrganization
}

type testCaseUpdateOrganization struct {
	name            string
	getDependencies func(t *testing.T) (*dependencies, func(ignoreErrors ...bool))
	args            argsUpdateOrganizations
	mutations       func(t *testing.T, db *sqlx.DB) (updateData model.UpdateOrganization)
	assertions      func(t *testing.T, params *model.UpdateOrganization, organization *model.Organization, err error)
}

func getTestCasesUpdateOrganizations() []testCaseUpdateOrganization {
	return []testCaseUpdateOrganization{
		{
			name: "success",
			args: argsUpdateOrganizations{
				ctx: context.TODO(),
				params: &model.UpdateOrganization{
					Id: 1,
					Name: null.String{
						String: "Demby",
						Valid:  true,
					},
					UserId: null.Int{
						Int:   3,
						Valid: true,
					},
				},
			},
			getDependencies: getConcreteDependencies,
			mutations: func(t *testing.T, db *sqlx.DB) (updateData model.UpdateOrganization) {
				entry := mysqlmodel.Organization{
					ID:            1,
					Name:          "Demby",
					CreatedBy:     null.IntFrom(3),
					LastUpdatedBy: null.IntFrom(3),
					CreatedAt:     time.Now(),
					LastUpdatedAt: null.TimeFrom(time.Now()),
					IsActive:      true,
				}

				err := entry.Insert(context.Background(), db, boil.Infer())
				require.NoError(t, err, "error inserting sample data")

				updatedOrganization := model.UpdateOrganization{
					Id:   1,
					Name: null.StringFrom("Hello iam test"),
				}

				return updatedOrganization
			},
			assertions: func(t *testing.T, params *model.UpdateOrganization, organization *model.Organization, err error) {
				createdByConvToInt, _ := strconv.Atoi(organization.CreatedBy)
				require.NoError(t, err, "unexpected error")
				require.NotNil(t, organization, "unexpected nil organization")
				assert.Equal(t, params.Id, organization.Id, "expected id to be equal")
				assert.Equal(t, params.UserId.Int, createdByConvToInt, "expected created_by id to be equal")
				assert.Equal(t, params.Name.String, organization.Name, "expected name to be equal")
			},
		},
		{
			name: "success-name-only",
			args: argsUpdateOrganizations{
				ctx: context.TODO(),
				params: &model.UpdateOrganization{
					Id: 1,
					Name: null.String{
						String: "Demby",
						Valid:  true,
					},
				},
			},
			getDependencies: getConcreteDependencies,
			mutations: func(t *testing.T, db *sqlx.DB) (updateData model.UpdateOrganization) {
				entry := mysqlmodel.Organization{
					ID:            1,
					Name:          "Demby",
					CreatedBy:     null.IntFrom(3),
					LastUpdatedBy: null.IntFrom(3),
					CreatedAt:     time.Now(),
					LastUpdatedAt: null.TimeFrom(time.Now()),
					IsActive:      true,
				}

				err := entry.Insert(context.Background(), db, boil.Infer())
				require.NoError(t, err, "error inserting sample data")

				updatedOrganization := model.UpdateOrganization{
					Id:   1,
					Name: null.StringFrom("Hello iam test"),
				}

				return updatedOrganization
			},
			assertions: func(t *testing.T, params *model.UpdateOrganization, organization *model.Organization, err error) {
				require.NoError(t, err, "unexpected error")
				require.NotNil(t, organization, "unexpected nil organization")
				assert.Equal(t, params.Name.String, organization.Name, "expected name to be equal")
			},
		},
		{
			name: "fail-mock",
			args: argsUpdateOrganizations{
				ctx: context.TODO(),
				params: &model.UpdateOrganization{
					Id: 1,
					UserId: null.Int{
						Int:   3,
						Valid: true,
					},
				},
			},
			getDependencies: func(t *testing.T) (*dependencies, func(ignoreErrors ...bool)) {
				cleanup := func(ignoreErrors ...bool) {}

				mockTxProvider := persistencefakes.FakeTransactionProvider{}
				mockTxProvider.TxReturns(nil, errors.New(mockDbReturnsErr))

				return &dependencies{
					Persistor:  &organizationlogicfakes.FakePersistor{},
					TxProvider: &mockTxProvider,
					Logger:     mockLogger,
					Cleanup:    cleanup,
				}, cleanup
			},
			mutations: func(t *testing.T, db *sqlx.DB) (updateData model.UpdateOrganization) {
				return updateData
			},
			assertions: func(t *testing.T, params *model.UpdateOrganization, organization *model.Organization, err error) {
				assert.Error(t, err, "unexpected error")
				assert.Nil(t, organization, "unexpected nil organization")
			},
		},
		{
			name: "fail-internal-server-error",
			args: argsUpdateOrganizations{
				ctx: context.TODO(),
				params: &model.UpdateOrganization{
					Id: 1,
					UserId: null.Int{
						Int:   3,
						Valid: true,
					},
				},
			},
			getDependencies: getConcreteDependencies,
			mutations: func(t *testing.T, db *sqlx.DB) (updateData model.UpdateOrganization) {
				testhelper.DropTable(t, db, mysqlmodel.TableNames.Organization)
				return updateData
			},
			assertions: func(t *testing.T, params *model.UpdateOrganization, organization *model.Organization, err error) {
				assert.Error(t, err, "unexpected error")
				assert.Nil(t, organization, "unexpected nil organization")
			},
		},
	}
}

func TestService_UpdateOrganization(t *testing.T) {
	for _, tt := range getTestCasesUpdateOrganizations() {
		t.Run(tt.name, func(t *testing.T) {
			_dependencies, cleanup := tt.getDependencies(t)
			defer cleanup()

			svc, err := New(&Config{
				TxProvider: _dependencies.TxProvider,
				Logger:     _dependencies.Logger,
				Persistor:  _dependencies.Persistor,
			})
			require.NoError(t, err, "unexpected new error")

			tt.mutations(t, _dependencies.Db)

			organization, err := svc.UpdateOrganization(tt.args.ctx, tt.args.params)
			tt.assertions(t, tt.args.params, organization, err)
		})
	}
}

// create org
type argsCreateOrganization struct {
	ctx          context.Context
	organization *model.CreateOrganization
}

type testCaseCreateOrganization struct {
	name            string
	getDependencies func(t *testing.T) (*dependencies, func(ignoreErrors ...bool))
	ctx             context.Context
	mutations       func(t *testing.T, db *sqlx.DB) (organization *model.CreateOrganization)
	assertions      func(t *testing.T, organization *model.Organization, err error)
}

func getTestCasesCreateOrganization() []testCaseCreateOrganization {
	return []testCaseCreateOrganization{
		{
			name:            "success",
			getDependencies: getConcreteDependencies,
			ctx:             context.TODO(),
			assertions: func(t *testing.T, organization *model.Organization, err error) {
				require.NoError(t, err, "unexpected error")
				require.NotNil(t, organization, "unexpected nil orga	nization")

				require.NotEqual(t, 0, organization.Id, "unexpected nil organization")
				require.NotEmpty(t, organization.Name, "unexpected empty organization name")
				//require.NotEqual(t, 0, organization.CreatedBy, "unexpected empty Created_by")
			},
			mutations: func(t *testing.T, db *sqlx.DB) (organization *model.CreateOrganization) {
				entryUser := mysqlmodel.User{
					Firstname:         "Demby",
					Lastname:          "Abella",
					Email:             "demby@test.com",
					Password:          "password",
					CategoryTypeRefID: 1,
				}
				err := entryUser.Insert(context.Background(), db, boil.Infer())
				require.NoError(t, err, "error inserting in the user db")

				entryOrganization := mysqlmodel.Organization{
					ID:            4,
					Name:          "TEST",
					CreatedBy:     null.IntFrom(entryUser.ID),
					LastUpdatedBy: null.IntFrom(entryUser.ID),
					CreatedAt:     time.Now(),
					LastUpdatedAt: null.TimeFrom(time.Now()),
					IsActive:      false,
				}
				err = entryOrganization.Insert(context.Background(), db, boil.Infer())
				require.NoError(t, err, "error inserting sample data")

				createdOrganization := &model.CreateOrganization{
					Name:   "Demby",
					UserId: 4,
				}

				return createdOrganization
			},
		},
	}
}

func TestService_CreateOrganization(t *testing.T) {
	for _, tt := range getTestCasesCreateOrganization() {
		t.Run(tt.name, func(t *testing.T) {
			_dependencies, cleanup := tt.getDependencies(t)
			defer cleanup()

			svc, err := New(&Config{
				TxProvider: _dependencies.TxProvider,
				Logger:     _dependencies.Logger,
				Persistor:  _dependencies.Persistor,
			})
			require.NoError(t, err, "unexpected new error")

			toBeCreatedOrganization := tt.mutations(t, _dependencies.Db)

			organization, err := svc.AddOrganization(tt.ctx, toBeCreatedOrganization)
			tt.assertions(t, organization, err)
		})
	}
}

// delete test
type testCaseDeleteOrganizations struct {
	name            string
	getDependencies func(t *testing.T) (*dependencies, func(ignoreErrors ...bool))
	ctx             context.Context
	mutations       func(t *testing.T, db *sqlx.DB) *model.DeleteOrganization
	assertions      func(t *testing.T, db *sqlx.DB, id int)
}

func getTestCasesDeleteOrganizations() []testCaseDeleteOrganizations {
	return []testCaseDeleteOrganizations{
		{
			name:            "success",
			getDependencies: getConcreteDependencies,
			ctx:             context.TODO(),
			mutations: func(t *testing.T, db *sqlx.DB) *model.DeleteOrganization {
				entryUser := mysqlmodel.User{
					Firstname:         "Demby",
					Lastname:          "Abella",
					Email:             "demby@test.com",
					Password:          "password",
					CategoryTypeRefID: 1,
				}
				err := entryUser.Insert(context.Background(), db, boil.Infer())
				require.NoError(t, err, "error inserting in the user db")

				entryOrganization := mysqlmodel.Organization{
					ID:            4,
					Name:          "TEST",
					CreatedBy:     null.IntFrom(entryUser.ID),
					LastUpdatedBy: null.IntFrom(entryUser.ID),
					CreatedAt:     time.Now(),
					LastUpdatedAt: null.TimeFrom(time.Now()),
					IsActive:      true,
				}
				err = entryOrganization.Insert(context.Background(), db, boil.Infer())
				require.NoError(t, err, "error inserting sample data")

				toBeDeletedOrganization := &model.DeleteOrganization{
					ID: entryOrganization.ID,
				}

				return toBeDeletedOrganization
			},
			assertions: func(t *testing.T, db *sqlx.DB, id int) {
				returnOrganization, err := mysqlmodel.FindOrganization(context.TODO(), db, id)
				require.Nil(t, returnOrganization, "expected to be nil")
				require.Error(t, err, "error fetching organization from db")
			},
		},
		{
			name:            "fail-organization-not-found",
			getDependencies: getConcreteDependencies,
			ctx:             context.TODO(),
			mutations: func(t *testing.T, db *sqlx.DB) *model.DeleteOrganization {
				entryUser := mysqlmodel.User{
					Firstname:         "Demby",
					Lastname:          "Abella",
					Email:             "demby@test.com",
					Password:          "password",
					CategoryTypeRefID: 1,
				}
				err := entryUser.Insert(context.Background(), db, boil.Infer())
				require.NoError(t, err, "error inserting in the user db")

				entryOrganization := mysqlmodel.Organization{
					ID:            4,
					Name:          "TEST",
					CreatedBy:     null.IntFrom(entryUser.ID),
					LastUpdatedBy: null.IntFrom(entryUser.ID),
					CreatedAt:     time.Now(),
					LastUpdatedAt: null.TimeFrom(time.Now()),
					IsActive:      true,
				}
				err = entryOrganization.Insert(context.Background(), db, boil.Infer())
				require.NoError(t, err, "error inserting sample data")

				toBeDeletedOrganization := &model.DeleteOrganization{
					ID: entryOrganization.ID,
				}

				return toBeDeletedOrganization
			},
			assertions: func(t *testing.T, db *sqlx.DB, id int) {
				returnOrganization, err := mysqlmodel.FindOrganization(context.TODO(), db, id)
				require.Nil(t, returnOrganization, "expected to be nil")
				require.Error(t, err, "expected an error when deleting a non-existent organization")
			},
		},
	}
}

func TestDeleteOrganization(t *testing.T) {
	for _, testCase := range getTestCasesDeleteOrganizations() {
		t.Run(testCase.name, func(t *testing.T) {
			db, _, cleanup := mysqlhelper.TestGetMockMariaDB(t)
			defer cleanup()
			require.NotNil(t, testCase.mutations, "unexpected nil mutations")
			require.NotNil(t, testCase.assertions, "unexpected nil assertions")

			_dependencies, cleanup := testCase.getDependencies(t)
			defer cleanup()

			svc, err := New(&Config{
				TxProvider: _dependencies.TxProvider,
				Logger:     _dependencies.Logger,
				Persistor:  _dependencies.Persistor,
			})
			require.NoError(t, err, "unexpected new error")

			toBeRestoredOrganization := testCase.mutations(t, _dependencies.Db)
			require.NoError(t, err, "unexpected new service error")

			err = svc.DeleteOrganization(testCase.ctx, toBeRestoredOrganization)
			require.NoError(t, err, "unexpected error deleting organization.")
			testCase.assertions(t, db, toBeRestoredOrganization.ID)
		})
	}
}

type testCaseRestoreOrganizations struct {
	name            string
	getDependencies func(t *testing.T) (*dependencies, func(ignoreErrors ...bool))
	ctx             context.Context
	mutations       func(t *testing.T, db *sqlx.DB) *model.RestoreOrganization
	assertions      func(t *testing.T, db *sqlx.DB, id int, err error)
}

func getTestCasesRestoreOrganizations() []testCaseRestoreOrganizations {
	return []testCaseRestoreOrganizations{
		{
			name:            "success",
			getDependencies: getConcreteDependencies,
			ctx:             context.TODO(),
			mutations: func(t *testing.T, db *sqlx.DB) *model.RestoreOrganization {
				entryUser := mysqlmodel.User{
					Firstname:         "Demby",
					Lastname:          "Abella",
					Email:             "demby@test.com",
					Password:          "password",
					CategoryTypeRefID: 1,
				}
				err := entryUser.Insert(context.Background(), db, boil.Infer())
				require.NoError(t, err, "error inserting in the user db")

				entryOrganization := mysqlmodel.Organization{
					ID:            4,
					Name:          "TEST",
					CreatedBy:     null.IntFrom(entryUser.ID),
					LastUpdatedBy: null.IntFrom(entryUser.ID),
					CreatedAt:     time.Now(),
					LastUpdatedAt: null.TimeFrom(time.Now()),
					IsActive:      false,
				}
				err = entryOrganization.Insert(context.Background(), db, boil.Infer())
				require.NoError(t, err, "error inserting sample data")

				toBeRestoredOrganization := &model.RestoreOrganization{
					ID: entryOrganization.ID,
				}

				return toBeRestoredOrganization
			},
			assertions: func(t *testing.T, db *sqlx.DB, id int, err error) {
				require.NoError(t, err, "error restoring organization from db")

				returnOrganization, err := mysqlmodel.FindOrganization(context.TODO(), db, id)
				require.NoError(t, err, "error fetching organization from db")
				require.NotNil(t, returnOrganization, "expected organization to be not nil")

				require.True(t, returnOrganization.IsActive, "expected organization to be active")
				assert.Equal(t, returnOrganization.IsActive, true)
			},
		},
		{
			name:            "fail-organization-not-found",
			getDependencies: getConcreteDependencies,
			ctx:             context.TODO(),
			mutations: func(t *testing.T, db *sqlx.DB) *model.RestoreOrganization {
				entryUser := mysqlmodel.User{
					Firstname:         "Demby",
					Lastname:          "Abella",
					Email:             "demby@test.com",
					Password:          "password",
					CategoryTypeRefID: 1,
				}
				err := entryUser.Insert(context.Background(), db, boil.Infer())
				require.NoError(t, err, "error inserting in the user db")

				entryOrganization := mysqlmodel.Organization{
					ID:            4,
					Name:          "TEST",
					CreatedBy:     null.IntFrom(entryUser.ID),
					LastUpdatedBy: null.IntFrom(entryUser.ID),
					CreatedAt:     time.Now(),
					LastUpdatedAt: null.TimeFrom(time.Now()),
					IsActive:      false,
				}
				err = entryOrganization.Insert(context.Background(), db, boil.Infer())
				require.NoError(t, err, "error inserting sample data")

				toBeRestoredOrganization := &model.RestoreOrganization{
					ID: entryOrganization.ID + 12321,
				}

				return toBeRestoredOrganization
			},
			assertions: func(t *testing.T, db *sqlx.DB, id int, err error) {
				returnOrganization, fetchErr := mysqlmodel.FindOrganization(context.TODO(), db, id)
				require.Error(t, fetchErr, "expected an error when fetching organization from db")
				require.Nil(t, returnOrganization, "expected organization to be nil")
			},
		},
	}
}

func TestRestoreOrganization(t *testing.T) {
	for _, testCase := range getTestCasesRestoreOrganizations() {
		t.Run(testCase.name, func(t *testing.T) {
			_dependencies, cleanup := testCase.getDependencies(t)
			defer cleanup()

			svc, err := New(&Config{
				TxProvider: _dependencies.TxProvider,
				Logger:     _dependencies.Logger,
				Persistor:  _dependencies.Persistor,
			})
			require.NoError(t, err, "unexpected new error")

			toBeRestoredOrganization := testCase.mutations(t, _dependencies.Db)
			require.NoError(t, err, "unexpected error in mutations")

			err = svc.RestoreOrganization(testCase.ctx, toBeRestoredOrganization)
			testCase.assertions(t, _dependencies.Db, toBeRestoredOrganization.ID, err)
		})
	}
}
