package app

import (
	"github.com/jroimartin/gocui"
	"github.com/toezbit/lazypassword/models"
)

type ViewManager interface {
	WorkSpace()
	AccountList()
	AccountDetail()
	Layout(g *gocui.Gui) error

	SetupKeyblidingNavigation()
	SetupKeyblidingWorkspace()
	SetupKeyblidingGlobal()
}

type ValutManager interface {
	GetWorkspaces() []models.VaultWithoutCredentails
	GetWorkspaceNames() []string
	AddWorkspace(g *gocui.Gui, v *gocui.View) error
}
