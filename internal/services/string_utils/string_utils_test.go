package string_utils

import (
	"errors"
	"github.com/dembygenesis/local.tools/internal/cli/clifakes"
	"github.com/dembygenesis/local.tools/internal/config"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_New_Success(t *testing.T) {
	conf := config.Config{}
	osLayer := clifakes.FakeStringUtils{}

	osLayer.CopyRootPathToClipboardReturns(nil, nil)

	_, err := New(&conf, &osLayer)
	require.NoError(t, err, "config error")
}

func Test_CopyRootPathToClipboard_Success(t *testing.T) {
	conf := config.Config{}
	conf.CopyToClipboard = config.CopyToClipboard{
		Exclusions: []string{"ab", "cd"},
	}
	osLayer := clifakes.FakeStringUtils{}

	osLayer.CopyRootPathToClipboardReturns(nil, nil)

	fakeStringUtils, err := New(&conf, &osLayer)
	require.NoError(t, err, "config error")

	_, err = fakeStringUtils.CopyRootPathToClipboard("test", nil)
	require.NoError(t, err, "no error expected")
}

func Test_CopyRootPathToClipboard_Fail_Empty_Root(t *testing.T) {
	conf := config.Config{}
	osLayer := clifakes.FakeStringUtils{}

	osLayer.CopyRootPathToClipboardReturns(nil, errors.New("mock error"))

	fakeStringUtils, err := New(&conf, &osLayer)
	require.NoError(t, err, "config error")

	_, err = fakeStringUtils.CopyRootPathToClipboard("test", nil)
	require.Error(t, err, "error expected")
	require.Contains(t, err.Error(), "os:")
}
