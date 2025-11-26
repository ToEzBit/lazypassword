package valut

import (
	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/toezbit/lazypassword/models"
)

func (v *ValutManagerImpl) GetWorkspacesWithoutCredentials() []models.VaultWithoutCredentails {
	workSpaceList := lo.Map(store.Vaults, func(el models.Vault, idx int) models.VaultWithoutCredentails {

		return models.VaultWithoutCredentails{
			ID:                   el.ID,
			WorkSpaceName:        el.WorkSpaceName,
			WorkSpaceDescription: el.WorkSpaceDescription,
		}

	})

	return workSpaceList
}

func (v *ValutManagerImpl) GetWorkspaceNames() []string {
	return lo.Map(store.Vaults, func(el models.Vault, idx int) string {
		return el.WorkSpaceName
	})
}

func (v *ValutManagerImpl) AddWorkspace(workspaceName string) {

	store.Vaults = append(store.Vaults, models.Vault{
		ID:                   uuid.NewString(),
		WorkSpaceName:        workspaceName,
		WorkSpaceDescription: "",
		Credentials:          []models.Credential{},
	})

	return
}

func CountWorkspace() int {
	return len(store.Vaults)
}
