package views

import (
	"github.com/jroimartin/gocui"
	"github.com/toezbit/lazypassword/constants"
)

func (vm *ViewManagerImpl) SetupKeyblidingNavigation() {
	vm.gui.SetKeybinding("", 'h', gocui.ModNone, cycleViewUp)
	vm.gui.SetKeybinding("", 'l', gocui.ModNone, cycleViewDown)
	vm.gui.SetKeybinding(constants.WorkSpace, 'j', gocui.ModNone, moveDownMenu)
	vm.gui.SetKeybinding(constants.WorkSpace, 'k', gocui.ModNone, moveUpMenu)
}

func (vm *ViewManagerImpl) SetupKeyblidingWorkspace() {
	vm.gui.SetKeybinding(constants.WorkSpace, 'n', gocui.ModNone, vm.openAddWorkspaceModal)
	vm.gui.SetKeybinding(constants.ModalAddWorkspace, gocui.KeyEsc, gocui.ModNone, vm.closeAddWorkspaceModal)
	vm.gui.SetKeybinding(constants.ModalAddWorkspace, gocui.KeyEnter, gocui.ModNone, vm.hanldeAddWorkspace)
}

func (vm *ViewManagerImpl) SetupKeyblidingGlobal() {
	vm.gui.SetKeybinding("", 'q', gocui.ModNone, quit)
}

func (vm *ViewManagerImpl) ClearKeyblidingNavigation() {
	vm.gui.DeleteKeybinding("", 'h', gocui.ModNone)
	vm.gui.DeleteKeybinding("", 'l', gocui.ModNone)
	vm.gui.DeleteKeybinding(constants.WorkSpace, 'j', gocui.ModNone)
	vm.gui.DeleteKeybinding(constants.WorkSpace, 'k', gocui.ModNone)

}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
