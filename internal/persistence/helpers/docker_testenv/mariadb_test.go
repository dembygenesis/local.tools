package docker_testenv

import (
	"context"
	"github.com/dembygenesis/local.tools/internal/utils_common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_New_MariaDB_Success(t *testing.T) {

	var (
		mappedPort    = 3340
		containerName = utils_common.GetUuidUnderscore()
		host          = "localhost"
		pass          = "secret"
	)

	cfg := MariaDBConfig{
		ContainerName:       containerName,
		Host:                host,
		Pass:                pass,
		ExposedInternalPort: mappedPort,
		ExposedExternalPort: mappedPort,
		RecreateContainer:   true,
	}

	mariaDbCtn, err := NewMariaDB(&cfg)
	require.NoError(t, err, "error starting new maria db container")

	err = mariaDbCtn.Cleanup(context.Background())
	require.NoError(t, err, "error cleaning up database")
}

func Test_New_MariaDB_Fail_Validate(t *testing.T) {
	var (
		mappedPort    = 0
		containerName = utils_common.GetUuidUnderscore()
		host          = "localhost"
		pass          = "secret"
	)

	cfg := MariaDBConfig{
		ContainerName:       containerName,
		Host:                host,
		Pass:                pass,
		ExposedInternalPort: mappedPort,
		RecreateContainer:   true,
	}

	_, err := NewMariaDB(&cfg)
	require.Error(t, err, "there should be a validation error")
	assert.Contains(t, err.Error(), "validate:")
}

func Test_New_MariaDB_Fail_Invalid_Setting(t *testing.T) {
	var (
		mappedPort    = 999999999999
		containerName = utils_common.GetUuidUnderscore()
		host          = "localhost"
		pass          = "secret"
	)

	cfg := MariaDBConfig{
		ContainerName:       containerName,
		Host:                host,
		Pass:                pass,
		ExposedInternalPort: mappedPort,
		ExposedExternalPort: mappedPort,
		RecreateContainer:   true,
	}

	_, err := NewMariaDB(&cfg)
	require.Error(t, err, "there should be a validation error")
	assert.Contains(t, err.Error(), "upsert container:")
}
