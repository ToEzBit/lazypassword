package app

import (
	"github.com/jroimartin/gocui"
)

type UiManager interface {
	WorkSpace()
	AccountList()
	AccountDetail()
	Layout(g *gocui.Gui) error

	SetupKeyblidingNavigation()
	SetupKeyblidingWorkspace()
	SetupKeyblidingGlobal()
	SetupKeyblidingAccountList()
}

type WorkspaceManager interface {
	// GetWorkspace() []models.Workspace
	GetWorkspaceNames() []string
	AddWorkspace(workspaceName string)
	GetCredentialNameList(workspaceId string) []string
}
