package views

import (
	"github.com/jroimartin/gocui"
	"github.com/toezbit/lazypassword/app"
	"github.com/toezbit/lazypassword/constants"
)

var padding = 1

type ViewManagerImpl struct {
	gui          *gocui.Gui
	valutManager app.ValutManager
}

func NewViewManagerImpl(g *gocui.Gui, valm app.ValutManager) *ViewManagerImpl {
	return &ViewManagerImpl{
		gui:          g,
		valutManager: valm,
	}
}

func (v *ViewManagerImpl) WorkSpace() {
	maxX, maxY := v.gui.Size()
	workSpaceView, _ := v.gui.SetView(constants.WorkSpace, 0, 0, maxX/2-padding, maxY/2)
	workSpaceView.Title = " Work Space "

	DrawMenus(v.gui, workSpaceView, constants.WorkSpace, v.valutManager.GetWorkspaceNames())
}

func (v *ViewManagerImpl) AccountList() {
	maxX, maxY := v.gui.Size()
	accountListView, _ := v.gui.SetView(constants.AccountList, 0, maxY/2+padding, maxX/2-padding, maxY-padding)
	accountListView.Title = " Account List "
}

func (v *ViewManagerImpl) AccountDetail() {
	maxX, maxY := v.gui.Size()
	accountDetailView, _ := v.gui.SetView(constants.AccountdDetail, maxX/2+padding, 0, maxX-padding, maxY-padding)
	accountDetailView.Title = " Account Detail "
}

func (v *ViewManagerImpl) Layout(g *gocui.Gui) error {
	v.WorkSpace()
	v.AccountDetail()
	v.AccountList()

	if v.gui.CurrentView() == nil {
		v.gui.SetCurrentView(constants.WorkSpace)
	}

	return nil

}
