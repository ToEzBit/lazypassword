package keybliding

import (
	"github.com/jroimartin/gocui"
	"github.com/toezbit/lazypassword/constants"
	"github.com/toezbit/lazypassword/views"
)

func SetUpNavigation(g *gocui.Gui) {
	g.SetKeybinding("", 'h', gocui.ModNone, cycleViewUp)
	g.SetKeybinding("", 'l', gocui.ModNone, cycleViewDown)
}

func SetUpWorkSpaceKeyBligind(g *gocui.Gui) {
	g.SetKeybinding(constants.WorkSpace, 'n', gocui.ModNone, views.OpenAddModalWorkSpace)
	g.SetKeybinding(constants.ModalAddWorkspace, gocui.KeyEsc, gocui.ModNone, views.CloseAddModalWorkSpace)
}
