package workspace

import (
	"github.com/samber/lo"
	"github.com/toezbit/lazypassword/models"
)

func (wm *WorkspaceManagerImpl) GetCredentialNameList(workspaceId string) []string {

	selectedWorkspace, _ := lo.Find(workspaces, func(el models.Workspace) bool {
		return el.Id == workspaceId
	})

	return lo.Map(selectedWorkspace.Credentials, func(el models.Credential, idx int) string {
		return el.AppName
	})

}
