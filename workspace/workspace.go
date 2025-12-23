package workspace

import (
	"slices"

	"github.com/google/uuid"
	"github.com/jroimartin/gocui"
	"github.com/samber/lo"
	"github.com/toezbit/lazypassword/file"
	"github.com/toezbit/lazypassword/models"
)

type WorkspaceManagerImpl struct {
	gui *gocui.Gui
}

func NewWorkspaceManagerImpl(g *gocui.Gui) *WorkspaceManagerImpl {
	return &WorkspaceManagerImpl{
		gui: g,
	}
}

var workspaces []models.Workspace

func init() {
	workspaces = file.ReadFile()
}

func Save() {
	file.WriteFile(workspaces)
}

func (vm *WorkspaceManagerImpl) GetWorkspaces() []models.Workspace {
	return workspaces
}

func GetWorkspaces() []models.Workspace {
	return workspaces
}

func (wm *WorkspaceManagerImpl) GetWorkspaceNames() []string {
	return lo.Map(workspaces, func(el models.Workspace, idx int) string {
		return el.Name
	})
}

func (wm *WorkspaceManagerImpl) AddWorkspace(workspaceName string) {

	workspaces = append(workspaces, models.Workspace{
		Id:          uuid.NewString(),
		Name:        workspaceName,
		Credentials: []models.Credential{},
	})

	return
}

func (wm *WorkspaceManagerImpl) DeleteWorkspace(workspaceId string) {

	_, targetWorkspaceIdx, _ := lo.FindIndexOf(workspaces, func(el models.Workspace) bool {
		return el.Id == workspaceId
	})

	updatedWorkspace := slices.Delete(workspaces, targetWorkspaceIdx, targetWorkspaceIdx+1)

	workspaces = updatedWorkspace

}

func CountWorkspace() int {
	return len(workspaces)
}
