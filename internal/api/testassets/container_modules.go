package testassets

import (
	apirsrc "github.com/dembygenesis/local.tools/internal/api/resource"
	"github.com/dembygenesis/local.tools/internal/logic_handlers/categorylogic"
	"github.com/dembygenesis/local.tools/internal/persistence/database_helpers/mysql/mysqlconn"
	"github.com/dembygenesis/local.tools/internal/persistence/persistors/mysqlstore"
)

type Container struct {
	CategoryService *categorylogic.Service
	MySQLStore      *mysqlstore.Repository
	ConnProvider    *mysqlconn.Provider
	ResourceGetter  *apirsrc.Provider
}
