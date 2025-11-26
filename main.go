package main

import (
	"fmt"
	"log"

	"github.com/jroimartin/gocui"
	"github.com/toezbit/lazypassword/app"
	"github.com/toezbit/lazypassword/ui"
	"github.com/toezbit/lazypassword/workspace"
)

func main() {

	g, err := gocui.NewGui(gocui.OutputNormal)

	if err != nil {
		log.Panicln(err)
	}

	defer closefunction(g)
	g.InputEsc = true
	g.Highlight = true
	g.SelFgColor = gocui.ColorCyan

	application := app.New(g)

	workspaceManager := workspace.NewWorkspaceManagerImpl(g)

	uiManager := ui.NewUiManagerImpl(g, workspaceManager)

	application.Initialize(uiManager, workspaceManager)

	if err := application.Gui.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}

}

func closefunction(g *gocui.Gui) {
	g.Close()
	workspace.Save()
	fmt.Println("👋 Cleanup complete. Application terminated.")

}
