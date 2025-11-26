package ui

import (
	"github.com/jroimartin/gocui"
	"github.com/toezbit/lazypassword/constants"
)

func (uim *UiManagerImpl) SetupKeyblidingNavigation() {
	uim.gui.SetKeybinding("", 'h', gocui.ModNone, cycleViewUp)
	uim.gui.SetKeybinding("", 'l', gocui.ModNone, cycleViewDown)
	uim.gui.SetKeybinding(constants.WorkSpace, 'j', gocui.ModNone, moveDownWorkspace)
	uim.gui.SetKeybinding(constants.WorkSpace, 'k', gocui.ModNone, moveUpWorkspace)
}

func (uim *UiManagerImpl) SetupKeyblidingWorkspace() {
	uim.gui.SetKeybinding(constants.WorkSpace, 'n', gocui.ModNone, uim.openAddWorkspaceModal)
	uim.gui.SetKeybinding(constants.ModalAddWorkspace, gocui.KeyEsc, gocui.ModNone, uim.closeAddWorkspaceModal)
	uim.gui.SetKeybinding(constants.ModalAddWorkspace, gocui.KeyEnter, gocui.ModNone, uim.hanldeAddWorkspace)
}

func (uim *UiManagerImpl) SetupKeyblidingCredential() {
	uim.gui.SetKeybinding(constants.Credential, 'n', gocui.ModNone, uim.openAddAccountModal)
	uim.gui.SetKeybinding("", gocui.KeyTab, gocui.ModNone, toggleFocusAddAccountInput)
	uim.gui.SetKeybinding(constants.Credential, 'j', gocui.ModNone, moveDownCredential)
	uim.gui.SetKeybinding(constants.Credential, 'k', gocui.ModNone, moveUpCredential)
}

func (uim *UiManagerImpl) SetupKeyblidingGlobal() {
	uim.gui.SetKeybinding("", 'q', gocui.ModNone, quit)
}

func (uim *UiManagerImpl) ClearKeyblidingNavigation() {
	uim.gui.DeleteKeybinding("", 'h', gocui.ModNone)
	uim.gui.DeleteKeybinding("", 'l', gocui.ModNone)
	uim.gui.DeleteKeybinding(constants.WorkSpace, 'j', gocui.ModNone)
	uim.gui.DeleteKeybinding(constants.WorkSpace, 'k', gocui.ModNone)
}

func (uim *UiManagerImpl) ClearKeyblidingGlobal() {
	uim.gui.DeleteKeybinding("", 'q', gocui.ModNone)

}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
