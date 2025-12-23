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

	SetupKeybindingNavigation()
	SetupKeybindingWorkspace()
	SetupKeybindingGlobal()
	SetupKeybindingCredential()
	SetupKeybindingOverview()
}

type WorkspaceManager interface {
	GetWorkspaces() []models.Workspace
	GetWorkspaceNames() []string
	AddWorkspace(workspaceName string)
	AddCredential(worksapceId string, data models.Credential)
	GetCredentialNameList(workspaceId string) []string
	DeleteCredential(workspaceId string, credentialId string)
}
