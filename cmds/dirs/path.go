package dirs

import (
	"log"

	"github.com/TeoDev1611/lspm/core/dirs"
)

func PrintConfigDir(opt bool) {
	if opt {
		path := dirs.GetDataPath("")
		log.Printf("PATH: %s", path)
	}
}
