package valut

import (
	"strings"

	"github.com/jroimartin/gocui"
	"github.com/samber/lo"
	"github.com/toezbit/lazypassword/constants"
	"github.com/toezbit/lazypassword/models"
)

type ValutManagerImpl struct {
	gui *gocui.Gui
}

func NewValutManagerImpl(g *gocui.Gui) *ValutManagerImpl {
	return &ValutManagerImpl{
		gui: g,
	}
}

func (v *ValutManagerImpl) GetWorkspaces() []models.VaultWithoutCredentails {
	workSpaceList := lo.Map(store.Valuts, func(el models.Vault, idx int) models.VaultWithoutCredentails {

		return models.VaultWithoutCredentails{
			ID:                   el.ID,
			WorkSpaceName:        el.WorkSpaceName,
			WorkSpaceDescription: el.WorkSpaceDescription,
		}

	})

	return workSpaceList
}

func (v *ValutManagerImpl) GetWorkspaceNames() []string {
	return lo.Map(store.Valuts, func(el models.Vault, idx int) string {
		return el.WorkSpaceName
	})
}

func (v *ValutManagerImpl) AddWorkspace(g *gocui.Gui, vi *gocui.View) error {
	modalWorkSpaceInput := strings.TrimSpace(v.gui.CurrentView().Buffer())

	if modalWorkSpaceInput == "" {
		return nil
	}

	v.gui.DeleteView(constants.ModalAddWorkspace)
	v.gui.SetCurrentView(constants.WorkSpace)

	store.Valuts = append(store.Valuts, models.Vault{
		ID:                   "2",
		WorkSpaceName:        modalWorkSpaceInput,
		WorkSpaceDescription: "",
		Credentials:          []models.Credential{},
	})

	return nil

}
