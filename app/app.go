package app

import (
	"github.com/jroimartin/gocui"
)

type App struct {
	Gui               *gocui.Gui
	KeyBlidingManager KeyBlidingManager
	ViewManager       ViewManager
	ValutManager      ValutManager
}

func New(g *gocui.Gui) *App {
	return &App{Gui: g}
}

func (a *App) Initialize(keyBlidingManager KeyBlidingManager, viewManager ViewManager, valutManager ValutManager) error {
	a.KeyBlidingManager = keyBlidingManager
	a.ViewManager = viewManager
	a.ValutManager = valutManager

	a.Gui.SetManagerFunc(a.ViewManager.Layout)

	a.KeyBlidingManager.SetupWorkspace()
	a.KeyBlidingManager.SetupNavigation()
	a.KeyBlidingManager.SetupGlobal()

	return nil
}
