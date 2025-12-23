package workspace

import (
	"slices"

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

func CountCredential(workspaceId string) int {
	selectedWorkspace, _ := lo.Find(workspaces, func(el models.Workspace) bool {
		return el.Id == workspaceId
	})
	return len(selectedWorkspace.Credentials)
}

func (wm *WorkspaceManagerImpl) AddCredential(workspaceId string, data models.Credential) {

	_, targetWorkspaceIdx, _ := lo.FindIndexOf(workspaces, func(el models.Workspace) bool {
		return el.Id == workspaceId
	})

	updatedCredentials := append(workspaces[targetWorkspaceIdx].Credentials, data)

	workspaces[targetWorkspaceIdx].Credentials = updatedCredentials

}

func (wm *WorkspaceManagerImpl) DeleteCredential(workspaceId string, credentialId string) {

	_, targetWorkspaceIdx, _ := lo.FindIndexOf(workspaces, func(el models.Workspace) bool {
		return el.Id == workspaceId
	})

	_, targetCredentialIdx, _ := lo.FindIndexOf(workspaces[targetWorkspaceIdx].Credentials, func(el models.Credential) bool {
		return el.Id == credentialId
	})

	updatedCredentials := slices.Delete(workspaces[targetWorkspaceIdx].Credentials, targetCredentialIdx, targetCredentialIdx+1)

	workspaces[targetWorkspaceIdx].Credentials = updatedCredentials

}
