package workspace

import (
	"slices"

	"github.com/google/uuid"
	"github.com/jroimartin/gocui"
	"github.com/samber/lo"
	"github.com/toezbit/lazypassword/crypto"
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

var fileData models.FileData
var workspaces []models.Workspace
var loadedSuccessfully bool

func init() {
	fileData, loadedSuccessfully = file.ReadFile()
	workspaces = fileData.Data

	if loadedSuccessfully && fileData.Salt == "" {
		salt, err := crypto.GenerateSalt()
		if err == nil {
			fileData.Salt = salt
			fileData.Version = "1.0"
		}
	}
}

func Save() {
	if !loadedSuccessfully {
		return
	}
	fileData.Data = workspaces
	file.WriteFile(fileData)
}

func GetFileData() models.FileData {
	return fileData
}

func SetPasswordHash(hash string) {
	fileData.PasswordHash = hash
	Save()
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
