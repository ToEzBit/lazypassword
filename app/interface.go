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
	SetupKeyblidingAccountList()
}

type ValutManager interface {
	GetValut() models.VaultData
	GetWorkspacesWithoutCredentials() []models.VaultWithoutCredentails
	GetWorkspaceNames() []string
	AddWorkspace(workspaceName string)

	GetAccountList(workspaceId string) []string
}
