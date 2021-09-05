package main

import (
	"github.com/TeoDev1611/lspm/cmds/dirs"
	"github.com/integrii/flaggy"
)

var (
	pathData bool
)

func main() {

	// Write a subcommand if is necesary
	dirSubCmd := flaggy.NewSubcommand("dir")
	dirSubCmd.Description = `Here show all dir utils like config path cache path servers path etc`
	dirSubCmd.Bool(&pathData, "p", "path", "Prints the Url to the path")

	// Set the Options for the parser 
	flaggy.DefaultParser.Name = "lspm"
	flaggy.DefaultParser.Version = "0.0.1"
	flaggy.DefaultParser.ShortName = "lsm"
	flaggy.DefaultParser.ShowHelpOnUnexpected = true
	flaggy.DefaultParser.Description = `Language Server Protocol Manager

- Why this ? You don't need any plugin for install the language servers for your editor
- What's this ? The best and fastest way to install any language server :)
- Help with this ? Check the documentation on github :p

*Credits:*
Made with love in Ecuador By Teo
`
	flaggy.AttachSubcommand(dirSubCmd, 1)
	//Parse the cli
	flaggy.Parse()

	// Run the cmds
	dirs.PrintConfigDir(pathData)

}
