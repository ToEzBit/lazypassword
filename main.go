package main

import (
	"fmt"
	"log"

	"github.com/jroimartin/gocui"
	"github.com/toezbit/lazypassword/app"
	"github.com/toezbit/lazypassword/services/valut"
	"github.com/toezbit/lazypassword/views"
)

func main() {

	g, err := gocui.NewGui(gocui.OutputNormal)

	if err != nil {
		log.Panicln(err)
		fmt.Println("kuy i sas")
	}

	defer closefunction(g)
	g.InputEsc = true
	g.Highlight = true
	g.SelFgColor = gocui.ColorCyan

	application := app.New(g)

	valutManager := valut.NewValutManagerImpl(g)

	viewManager := views.NewViewManagerImpl(g, valutManager)

	application.Initialize(viewManager, valutManager)

	if err := application.Gui.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
		fmt.Println("kuy i sas")
	}

}

func closefunction(g *gocui.Gui) {
	g.Close()
	valut.Save()
	fmt.Println("👋 Cleanup complete. Application terminated.")

	valutManager := valut.NewValutManagerImpl(g)

	valutManager.GetAccountList("f5cc33b5-c164-450d-b9b3-704b16503604")
	//
	// fmt.Println(valutManager.GetAccountList('1'))

}
