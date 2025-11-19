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
	OpenAddWorkspaceModal(g *gocui.Gui, v *gocui.View) error
	CloseAddWorkspaceModal(g *gocui.Gui, v *gocui.View) error
}

type KeyBlidingManager interface {
	SetupNavigation()
	SetupWorkspace()
	SetupGlobal()
}

type ValutManager interface {
	GetWorkspaces() []models.VaultWithoutCredentails
	GetWorkspaceNames() []string
	AddWorkspace(g *gocui.Gui, v *gocui.View) error
}
