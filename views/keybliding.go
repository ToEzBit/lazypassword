package views

import (
	"github.com/jroimartin/gocui"
	"github.com/toezbit/lazypassword/constants"
)

func (v *ViewManagerImpl) SetupKeyblidingNavigation() {
	v.gui.SetKeybinding("", 'h', gocui.ModNone, cycleViewUp)
	v.gui.SetKeybinding("", 'l', gocui.ModNone, cycleViewDown)
	v.gui.SetKeybinding(constants.WorkSpace, 'j', gocui.ModNone, moveDownMenu)
	v.gui.SetKeybinding(constants.WorkSpace, 'k', gocui.ModNone, moveUpMenu)
}

func (v *ViewManagerImpl) SetupKeyblidingWorkspace() {
	v.gui.SetKeybinding(constants.WorkSpace, 'n', gocui.ModNone, v.openAddWorkspaceModal)
	v.gui.SetKeybinding(constants.ModalAddWorkspace, gocui.KeyEsc, gocui.ModNone, v.closeAddWorkspaceModal)
	v.gui.SetKeybinding(constants.ModalAddWorkspace, gocui.KeyEnter, gocui.ModNone, v.valutManager.AddWorkspace)
}

func (v *ViewManagerImpl) SetupKeyblidingGlobal() {
	v.gui.SetKeybinding("", 'q', gocui.ModNone, quit)

}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
