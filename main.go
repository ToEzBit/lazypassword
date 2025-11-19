package main

import (
	"log"

	"github.com/jroimartin/gocui"
	"github.com/toezbit/lazypassword/app"
	"github.com/toezbit/lazypassword/keybliding"
	"github.com/toezbit/lazypassword/services/valut"
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

	application := app.New(g)

	valutManager := valut.NewValutManagerImpl(g)

	viewManager := views.NewViewManagerImpl(g, valutManager)

	keyBlidingManger := keybliding.NewKeyBlidingManagerImpl(g, viewManager, valutManager)

	application.Initialize(keyBlidingManger, viewManager, valutManager)

	if err := application.Gui.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}

}
