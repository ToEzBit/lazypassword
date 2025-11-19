package keybliding

import (
	"github.com/jroimartin/gocui"
	"github.com/toezbit/lazypassword/app"
	"github.com/toezbit/lazypassword/constants"
)

type KeyBlidingManagerImpl struct {
	gui          *gocui.Gui
	viewManager  app.ViewManager
	valutManager app.ValutManager
}

func NewKeyBlidingManagerImpl(g *gocui.Gui, vm app.ViewManager, valm app.ValutManager) *KeyBlidingManagerImpl {
	return &KeyBlidingManagerImpl{
		gui:          g,
		viewManager:  vm,
		valutManager: valm,
	}
}

func (k *KeyBlidingManagerImpl) SetupNavigation() {
	k.gui.SetKeybinding("", 'h', gocui.ModNone, cycleViewUp)
	k.gui.SetKeybinding("", 'l', gocui.ModNone, cycleViewDown)
	k.gui.SetKeybinding(constants.WorkSpace, 'j', gocui.ModNone, moveDownMenu)
	k.gui.SetKeybinding(constants.WorkSpace, 'k', gocui.ModNone, moveUpMenu)
}

func (k *KeyBlidingManagerImpl) SetupWorkspace() {
	k.gui.SetKeybinding(constants.WorkSpace, 'n', gocui.ModNone, k.viewManager.OpenAddWorkspaceModal)
	k.gui.SetKeybinding(constants.ModalAddWorkspace, gocui.KeyEsc, gocui.ModNone, k.viewManager.CloseAddWorkspaceModal)
	k.gui.SetKeybinding(constants.ModalAddWorkspace, gocui.KeyEnter, gocui.ModNone, k.valutManager.AddWorkspace)
}

func (k *KeyBlidingManagerImpl) SetupGlobal() {
	k.gui.SetKeybinding("", 'q', gocui.ModNone, quit)

}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
