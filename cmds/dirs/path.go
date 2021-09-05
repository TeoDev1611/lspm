package dirs

import (
	"log"

	"github.com/TeoDev1611/lspm/core/dirs"
)

func PrintServersPath(opt bool) {
	if opt {
		path := dirs.GetServersPath("")
		log.Printf("PATH: %s", path)
	}
}

func PrintCachePath(opt bool) {
	if opt {
		path := dirs.GetCachePath("")
		log.Printf("PATH: %s", path)
	}
}

func PrintConfigFile(opt bool) {
	if opt {
		path := dirs.GetConfigPath()
		log.Printf("PATH: %s", path)
	}
}
