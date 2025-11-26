package app

import (
	"github.com/jroimartin/gocui"
)

type App struct {
	Gui              *gocui.Gui
	UiManager        UiManager
	WorkspaceManager WorkspaceManager
}

func New(g *gocui.Gui) *App {
	return &App{Gui: g}
}

func (app *App) Initialize(uiManager UiManager, workspaceManager WorkspaceManager) error {
	app.UiManager = uiManager
	app.WorkspaceManager = workspaceManager

	app.Gui.SetManagerFunc(app.UiManager.Layout)

	app.UiManager.SetupKeyblidingWorkspace()
	app.UiManager.SetupKeyblidingNavigation()
	app.UiManager.SetupKeyblidingGlobal()
	app.UiManager.SetupKeyblidingAccountList()

	return nil
}
