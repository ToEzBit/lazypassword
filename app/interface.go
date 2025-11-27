package app

import (
	"github.com/jroimartin/gocui"
	"github.com/toezbit/lazypassword/models"
)

type UiManager interface {
	WorkSpace()
	Credential()
	Overview()
	Layout(g *gocui.Gui) error

	SetupKeyblidingNavigation()
	SetupKeyblidingWorkspace()
	SetupKeyblidingGlobal()
	SetupKeyblidingCredential()
	SetupKeybindingOverview()
}

type WorkspaceManager interface {
	GetWorkspaces() []models.Workspace
	GetWorkspaceNames() []string
	AddWorkspace(workspaceName string)
	GetCredentialNameList(workspaceId string) []string
}
