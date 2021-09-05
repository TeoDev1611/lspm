package dirs

import (
	"os"
	"path"

	"github.com/TeoDev1611/lspm/core/errors"
)

func GetServersPath(file string) string {
	configPath, err := os.UserConfigDir()
	errors.ErrorChecker(err)
	return path.Join(configPath, "lspm", "servers", file)
}

func GetConfigPath() string {
	configPath, err := os.UserConfigDir()
	errors.ErrorChecker(err)
	return path.Join(configPath, "lspm", "config", "lspm.toml")
}

func GetCachePath(file string) string {
	cachePath, err := os.UserCacheDir()
	errors.ErrorChecker(err)
	return path.Join(cachePath, "lspm", "cache", file)
}
