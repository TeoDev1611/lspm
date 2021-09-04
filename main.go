package main

import (
	"github.com/TeoDev1611/lspm/cmds"
	"github.com/integrii/flaggy"
)

var (
	Version bool
)

func main() {
	flaggy.Bool(&Version, "V", "version", "Prints the version of lspm")
	flaggy.DefaultParser.ShowVersionWithVersionFlag = false
	flaggy.Parse()
	cmds.VersionParsing(Version)
}
