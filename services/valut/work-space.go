package valut

import (
	"strings"

	"github.com/jroimartin/gocui"
	"github.com/samber/lo"
	"github.com/toezbit/lazypassword/constants"
	"github.com/toezbit/lazypassword/models"
)

func GetWorkSpaces() []models.VaultWithoutCredentails {

	workSpaceList := lo.Map(store.Valuts, func(el models.Vault, idx int) models.VaultWithoutCredentails {

		return models.VaultWithoutCredentails{
			ID:                   el.ID,
			WorkSpaceName:        el.WorkSpaceName,
			WorkSpaceDescription: el.WorkSpaceDescription,
		}

	})

	return workSpaceList
}

func GetWorkSpaceNames() []string {
	return lo.Map(store.Valuts, func(el models.Vault, idx int) string {
		return el.WorkSpaceName
	})
}

func AddWorkSpace(g *gocui.Gui, v *gocui.View) error {

	modalWorkSpaceInput := strings.TrimSpace(g.CurrentView().Buffer())

	if modalWorkSpaceInput == "" {
		return nil
	}

	g.DeleteView(constants.ModalAddWorkspace)
	g.SetCurrentView(constants.WorkSpace)

	store.Valuts = append(store.Valuts, models.Vault{
		ID:                   "2",
		WorkSpaceName:        modalWorkSpaceInput,
		WorkSpaceDescription: "",
		Credentials:          []models.Credential{},
	})

	return nil
}
