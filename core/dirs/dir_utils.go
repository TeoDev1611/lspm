package dirs

import (
	"os"
	"path"

	"github.com/TeoDev1611/lspm/core/errors"
)

func GetDataPath(file string) string {
	configPath, err := os.UserConfigDir()
	errors.ErrorChecker(err)
	return path.Join(configPath, "lspm", file)
}
