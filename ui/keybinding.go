package ui

import (
	"github.com/jroimartin/gocui"
	"github.com/toezbit/lazypassword/constants"
)

func (uim *UiManagerImpl) SetupKeybindingNavigation() {
	uim.gui.SetKeybinding("", 'h', gocui.ModNone, cycleViewUp)
	uim.gui.SetKeybinding("", 'l', gocui.ModNone, cycleViewDown)
	uim.gui.SetKeybinding(constants.WorkSpace, 'j', gocui.ModNone, moveDownWorkspace)
	uim.gui.SetKeybinding(constants.WorkSpace, 'k', gocui.ModNone, moveUpWorkspace)
}

func (uim *UiManagerImpl) SetupKeybindingWorkspace() {
	uim.gui.SetKeybinding(constants.WorkSpace, 'n', gocui.ModNone, uim.openAddWorkspaceModal)
	uim.gui.SetKeybinding(constants.ModalAddWorkspace, gocui.KeyEsc, gocui.ModNone, uim.closeAddWorkspaceModal)
	uim.gui.SetKeybinding(constants.ModalAddWorkspace, gocui.KeyEnter, gocui.ModNone, uim.handleAddWorkspace)
}

func (uim *UiManagerImpl) SetupKeybindingCredential() {
	uim.gui.SetKeybinding(constants.Credential, 'n', gocui.ModNone, uim.openAddCredentialModal)
	uim.gui.SetKeybinding(constants.ModalAddCredentialAppNameInput, gocui.KeyEsc, gocui.ModNone, uim.closeAddCredentialModal)
	uim.gui.SetKeybinding(constants.ModalAddCredentialAppNameInput, gocui.KeyEnter, gocui.ModNone, uim.handleAddCredential)
	uim.gui.SetKeybinding("", gocui.KeyTab, gocui.ModNone, toggleFocusAddCredentialInput)
	uim.gui.SetKeybinding(constants.Credential, 'j', gocui.ModNone, moveDownCredential)
	uim.gui.SetKeybinding(constants.Credential, 'k', gocui.ModNone, moveUpCredential)
	uim.gui.SetKeybinding(constants.Credential, gocui.KeyEnter, gocui.ModNone, handleSelectCredential)
	uim.gui.SetKeybinding(constants.Credential, 'd', gocui.ModNone, uim.openConfirmDeleteCredentialModal)
	uim.gui.SetKeybinding(constants.ModalConfirmDeleteCredential, 'n', gocui.ModNone, uim.closeConfirmDeleteCredentialModal)
	uim.gui.SetKeybinding(constants.ModalConfirmDeleteCredential, 'y', gocui.ModNone, uim.handleDeleteCredential)
}

func (uim *UiManagerImpl) SetupKeybindingOverview() {
	uim.gui.SetKeybinding(constants.Overview, 'j', gocui.ModNone, moveDownOverview)
	uim.gui.SetKeybinding(constants.Overview, 'k', gocui.ModNone, moveUpOverview)
	uim.gui.SetKeybinding(constants.Overview, gocui.KeyEsc, gocui.ModNone, exitOverview)
	uim.gui.SetKeybinding(constants.Overview, 'y', gocui.ModNone, copyCurrentSelectedOverview)

}

func (uim *UiManagerImpl) SetupKeybindingGlobal() {
	uim.gui.SetKeybinding("", 'q', gocui.ModNone, quit)
}

func (uim *UiManagerImpl) ClearKeybindingNavigation() {
	uim.gui.DeleteKeybinding("", 'h', gocui.ModNone)
	uim.gui.DeleteKeybinding("", 'l', gocui.ModNone)
	uim.gui.DeleteKeybinding(constants.WorkSpace, 'j', gocui.ModNone)
	uim.gui.DeleteKeybinding(constants.WorkSpace, 'k', gocui.ModNone)
}

func (uim *UiManagerImpl) ClearKeybindingGlobal() {
	uim.gui.DeleteKeybinding("", 'q', gocui.ModNone)
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
