package mysqlhelper

import (
	"context"
	"errors"
	"fmt"
	"github.com/dembygenesis/local.tools/internal/global"
	"github.com/dembygenesis/local.tools/internal/persistence/database_helpers/mysql/mysqlutil"
	"github.com/dembygenesis/local.tools/internal/sysconsts"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"os"
	"strings"
	"time"
)

type Tables [][]string

const (
	migrateDirFolder = "/internal/database/migrations"
)

var (
	migrateDir = fmt.Sprintf("%s/%s", os.Getenv(global.OsEnvAppDir), migrateDirFolder)
)

func Migrate(ctx context.Context, cp *mysqlutil.ConnectionSettings, mode CreateMode) (Tables, error) {
	if err := cp.Validate(true); err != nil {
		return nil, fmt.Errorf("validate conn parameters: %v", err)
	}

	if strings.TrimSpace(migrateDir) == "" {
		return nil, errors.New(sysconsts.ErrEmptyDir)
	}

	_, err := NewDbClient(ctx, &ClientOptions{
		ConnString: cp.GetConnectionString(false),
		Close:      true,
	})
	if err != nil {
		return nil, fmt.Errorf("get client: %v", err)
	}

	_, err = NewBuildTeardown(ctx, cp, &NewOpts{
		Mode:  mode,
		Close: true,
	})
	if err != nil {
		return nil, fmt.Errorf("mysql new: %v", err)
	}

	// Execute migration statements
	m, err := migrate.New(
		fmt.Sprintf("file://%s", migrateDir),
		fmt.Sprintf("mysql://%v", cp.GetConnectionString(false)),
	)

	if err != nil {
		return nil, fmt.Errorf("migrate new: %v, %s", err, migrateDir)
	}

	if err = m.Up(); err != nil {
		if !strings.Contains(err.Error(), "no change") {
			return nil, fmt.Errorf("up: %v", err)
		}
	}

	showTablesStmt := fmt.Sprintf(`
			SELECT table_name FROM information_schema.tables
			WHERE table_schema = '%v'
		`, cp.Database)

	conn, err := NewConnection(&Settings{
		ConnectTimeout:       time.Second * 10,
		QueryTimeout:         time.Second * 10,
		ExecTimeout:          time.Second * 10,
		ConnectionParameters: cp,
	})
	if err != nil {
		return nil, fmt.Errorf("connect: %v", err)
	}

	tables, _, err := conn.QueryIntoStructSettings(ctx, &QueryAsArrFilter{
		Stmt:       showTablesStmt,
		Pagination: nil,
	})
	if err != nil {
		return nil, fmt.Errorf("query into struct: %v", err)
	}

	return tables, nil
}
