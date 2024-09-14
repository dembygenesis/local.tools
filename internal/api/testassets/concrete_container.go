package testassets

import (
	"github.com/dembygenesis/local.tools/di/ctn/dic"
	"github.com/dembygenesis/local.tools/internal/persistence/database_helpers/mysql/mysqlhelper"
	"github.com/stretchr/testify/require"
	"testing"
)

func GetConcreteContainer(t *testing.T) (*Container, mysqlhelper.CleanFn) {
	db, _, cleanup := mysqlhelper.TestGetMockMariaDB(t)
	ctn, err := dic.NewContainer()
	require.NoError(t, err, "unexpected error instantiating a new container")

	mySQLTxHandler, err := ctn.SafeGetTxHandler()
	require.NoError(t, err, "unexpected error: SafeGetTxHandler")

	err = mySQLTxHandler.Set(db)
	require.Nil(t, err, "unexpected error for setting mysqlTxHandler.Set")

	category, err := ctn.SafeGetLogicCategory()
	require.NoError(t, err, "unexpected error: SafeGetLogicCategory")

	mysqlStore, err := ctn.SafeGetPersistenceMysql()
	require.NoError(t, err, "unexpected error: SafeGetPersistenceMysql")

	mysqlTxProvider, err := ctn.SafeGetTxProvider()
	require.NoError(t, err, "unexpected error: SafeGetTxProvider")

	resourceGetter, err := ctn.SafeGetResourceGetter()
	require.NoError(t, err, "unexpected error: SafeGetResourceGetter")

	return &Container{
		CategoryService: category,
		MySQLStore:      mysqlStore,
		ConnProvider:    mysqlTxProvider,
		ResourceGetter:  resourceGetter,
	}, cleanup
}
