/* This simple example demonstrates how to display UTF8 */
package main

import (
	"log"
	"time"

	"github.com/famz/SetLocale"
	gc "github.com/gbin/goncurses"
)

func main() {
	SetLocale.SetLocale(SetLocale.LC_ALL, "")
	scr, err := gc.Init()
	if err != nil {
		log.Fatalf("Failed to initialize screen: %w", err)
	}
	defer gc.End()

	scr.MovePrintf(0, 0, "😊")

	_ = scr.AttrSet(gc.A_NORMAL)
	scr.Refresh()

	time.Sleep(5 * time.Second)
}
