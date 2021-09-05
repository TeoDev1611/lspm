package main

import (
	"github.com/integrii/flaggy"
)

func main() {

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
	//Parse the cli
	flaggy.Parse()
}
