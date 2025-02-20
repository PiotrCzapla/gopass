//go:build !windows
// +build !windows

package appdir

import (
	"os"
	"testing"

	"github.com/gopasspw/gopass/tests/gptest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserConfig(t *testing.T) { //nolint:paralleltest
	ov := gptest.UnsetVars("GOPASS_HOMEDIR", "XDG_CONFIG_HOME", "HOME")
	defer ov()

	t.Run("gopass homedir", func(t *testing.T) { //nolint:paralleltest
		require.NoError(t, os.Setenv("GOPASS_HOMEDIR", "/foo/bar"))
		assert.Equal(t, "/foo/bar/.config/gopass", UserConfig())
		require.NoError(t, os.Unsetenv("GOPASS_HOMEDIR"))
	})

	t.Run("xdg_config_home", func(t *testing.T) { //nolint:paralleltest
		require.NoError(t, os.Setenv("XDG_CONFIG_HOME", "/foo/baz/myconfig"))
		assert.Equal(t, "/foo/baz/myconfig/gopass", UserConfig())
		require.NoError(t, os.Unsetenv("XDG_CONFIG_HOME"))
	})

	t.Run("default", func(t *testing.T) { //nolint:paralleltest
		require.NoError(t, os.Setenv("HOME", "/home/gopass"))
		assert.Equal(t, "/home/gopass/.config/gopass", UserConfig())
		require.NoError(t, os.Unsetenv("HOME"))
	})
}

func TestUserCache(t *testing.T) { //nolint:paralleltest
	ov := gptest.UnsetVars("GOPASS_HOMEDIR", "XDG_CACHE_HOME", "HOME")
	defer ov()

	t.Run("gopass homedir", func(t *testing.T) { //nolint:paralleltest
		require.NoError(t, os.Setenv("GOPASS_HOMEDIR", "/foo/bar"))
		assert.Equal(t, "/foo/bar/.cache/gopass", UserCache())
		require.NoError(t, os.Unsetenv("GOPASS_HOMEDIR"))
	})

	t.Run("xdg_cache_home", func(t *testing.T) { //nolint:paralleltest
		require.NoError(t, os.Setenv("XDG_CACHE_HOME", "/foo/baz/mycache"))
		assert.Equal(t, "/foo/baz/mycache/gopass", UserCache())
		require.NoError(t, os.Unsetenv("XDG_CACHE_HOME"))
	})

	t.Run("default", func(t *testing.T) { //nolint:paralleltest
		require.NoError(t, os.Setenv("HOME", "/home/gopass"))
		assert.Equal(t, "/home/gopass/.cache/gopass", UserCache())
		require.NoError(t, os.Unsetenv("HOME"))
	})
}

func TestUserData(t *testing.T) { //nolint:paralleltest
	ov := gptest.UnsetVars("GOPASS_HOMEDIR", "XDG_DATA_HOME", "HOME")
	defer ov()

	t.Run("gopass homedir", func(t *testing.T) { //nolint:paralleltest
		require.NoError(t, os.Setenv("GOPASS_HOMEDIR", "/foo/bar"))
		assert.Equal(t, "/foo/bar/.local/share/gopass", UserData())
		require.NoError(t, os.Unsetenv("GOPASS_HOMEDIR"))
	})

	t.Run("xdg_data_home", func(t *testing.T) { //nolint:paralleltest
		require.NoError(t, os.Setenv("XDG_DATA_HOME", "/foo/baz/mydata"))
		assert.Equal(t, "/foo/baz/mydata/gopass", UserData())
		require.NoError(t, os.Unsetenv("XDG_DATA_HOME"))
	})

	t.Run("default", func(t *testing.T) { //nolint:paralleltest
		require.NoError(t, os.Setenv("HOME", "/home/gopass"))
		assert.Equal(t, "/home/gopass/.local/share/gopass", UserData())
		require.NoError(t, os.Unsetenv("HOME"))
	})
}
