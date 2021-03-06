// ncurses_example.go
package goncurses_test

import (
	"github.com/gbin/goncurses"
	"fmt"
	"os"
)

func ExampleInit() {
	// You should always test to make sure ncurses has initialized properly.
	// In order for your error messages to be visible on the terminal you will
	// need to either log error messages or output them to to stderr.
	stdscr, err := goncurses.Init()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer goncurses.End()
	stdscr.Print("Press enter to continue...")
	stdscr.Refresh()
}
