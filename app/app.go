package app

import (
	"github.com/jroimartin/gocui"
)

type App struct {
	Gui          *gocui.Gui
	ViewManager  ViewManager
	ValutManager ValutManager
}

func New(g *gocui.Gui) *App {
	return &App{Gui: g}
}

func (a *App) Initialize(viewManager ViewManager, valutManager ValutManager) error {
	a.ViewManager = viewManager
	a.ValutManager = valutManager

	a.Gui.SetManagerFunc(a.ViewManager.Layout)

	a.ViewManager.SetupKeyblidingWorkspace()
	a.ViewManager.SetupKeyblidingNavigation()
	a.ViewManager.SetupKeyblidingGlobal()

	return nil
}
