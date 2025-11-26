package valut

import (
	"github.com/samber/lo"
	"github.com/toezbit/lazypassword/models"
)

func (v *ValutManagerImpl) GetAccountList(workspaceId string) []string {

	selectedWorkspace, _ := lo.Find(store.Vaults, func(el models.Vault) bool {
		return el.ID == workspaceId
	})

	return lo.Map(selectedWorkspace.Credentials, func(el models.Credential, idx int) string {
		return el.App
	})

}
