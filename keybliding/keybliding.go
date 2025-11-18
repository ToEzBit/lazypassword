package keybliding

import (
	"github.com/jroimartin/gocui"
	"github.com/toezbit/lazypassword/constants"
	"github.com/toezbit/lazypassword/services/valut"
	"github.com/toezbit/lazypassword/views"
)

func SetUpNavigation(g *gocui.Gui) {
	g.SetKeybinding("", 'h', gocui.ModNone, cycleViewUp)
	g.SetKeybinding("", 'l', gocui.ModNone, cycleViewDown)
	g.SetKeybinding(constants.WorkSpace, 'j', gocui.ModNone, moveDownMenu)
	g.SetKeybinding(constants.WorkSpace, 'k', gocui.ModNone, moveUpMenu)
}

func SetUpWorkSpaceKeyBligind(g *gocui.Gui) {
	g.SetKeybinding(constants.WorkSpace, 'n', gocui.ModNone, views.OpenAddModalWorkSpace)
	g.SetKeybinding(constants.ModalAddWorkspace, gocui.KeyEsc, gocui.ModNone, views.CloseAddModalWorkSpace)
	g.SetKeybinding(constants.ModalAddWorkspace, gocui.KeyEnter, gocui.ModNone, valut.AddWorkSpace)
}
