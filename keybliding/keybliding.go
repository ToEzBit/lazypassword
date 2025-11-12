package keybliding

import "github.com/jroimartin/gocui"

func SetUpNavigation(g *gocui.Gui) {
	g.SetKeybinding("", 'h', gocui.ModNone, cycleViewUp)
	g.SetKeybinding("", 'l', gocui.ModNone, cycleViewDown)
}
