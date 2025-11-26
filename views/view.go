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

func (vm *ViewManagerImpl) WorkSpace() {
	maxX, maxY := vm.gui.Size()
	workSpaceView, _ := vm.gui.SetView(constants.WorkSpace, 0, 0, maxX/2-padding, maxY/2)
	workSpaceView.Title = " Work Space "

	DrawMenus(vm.gui, workSpaceView, constants.WorkSpace, vm.valutManager.GetWorkspaceNames())
}

func (vm *ViewManagerImpl) AccountList() {
	maxX, maxY := vm.gui.Size()
	accountListView, _ := vm.gui.SetView(constants.AccountList, 0, maxY/2+padding, maxX/2-padding, maxY-padding)
	accountListView.Title = " Account List "

	DrawMenus(vm.gui, accountListView, constants.AccountList, vm.valutManager.GetAccountList("1"))

}

func (vm *ViewManagerImpl) AccountDetail() {
	maxX, maxY := vm.gui.Size()
	accountDetailView, _ := vm.gui.SetView(constants.AccountdDetail, maxX/2+padding, 0, maxX-padding, maxY-padding)
	accountDetailView.Title = " Account Detail "
}

func (vm *ViewManagerImpl) Layout(g *gocui.Gui) error {
	vm.WorkSpace()
	vm.AccountDetail()
	vm.AccountList()

	if vm.gui.CurrentView() == nil {
		vm.gui.SetCurrentView(constants.WorkSpace)
	}

	return nil

}
