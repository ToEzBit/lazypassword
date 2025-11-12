package main

import (
	"log"

	"github.com/jroimartin/gocui"
	"github.com/toezbit/lazypassword/constants"
	"github.com/toezbit/lazypassword/keybliding"
	"github.com/toezbit/lazypassword/views"
)

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.InputEsc = true
	g.Highlight = true
	g.SelFgColor = gocui.ColorCyan

	g.SetManagerFunc(layout)

	keybliding.SetUpNavigation(g)
	keybliding.SetUpWorkSpaceKeyBligind(g)

	if err := g.SetKeybinding("", 'q', gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}

}

func layout(g *gocui.Gui) error {

	views.WorkSpace(g)
	views.AccountList(g)
	views.AccountDetail(g)

	if g.CurrentView() == nil {
		g.SetCurrentView(constants.WorkSpace)
	}

	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

// func moveRight(g *gocui.Gui, v *gocui.View) error {
// 	rightView, _ := g.View("right-view")
// 	g.SetCurrentView("right-view")
//
// 	views.DrawMenu(g, rightView, "right-view")
// 	return nil
// }
//
// func moveLeft(g *gocui.Gui, v *gocui.View) error {
// 	g.SetCurrentView("left-view")
// 	return nil
//
// }
