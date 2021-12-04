package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	wizardFlag  = flag.Bool("wizard", false, "instead of the usual micropesce instance, start a cli configuration wizard. if you pass <db_file>, the wizard will modify the configuration in that file. else a new db file will be created.")
	versionFlag = flag.Bool("version", false, "print micropesce version and quit.")
)

// Parse flags and notify the user that none of them are implemented actually.
func parseFlags() {
	flag.Parse()
	if *wizardFlag {
		fmt.Println("sorry, the wizard is not implemented yet.")
		os.Exit(1)
	}
	if *versionFlag {
		fmt.Println("micropesce v?")
		os.Exit(1)
	}
}
