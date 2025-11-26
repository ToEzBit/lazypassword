package ui

import (
	"github.com/jroimartin/gocui"
	"github.com/toezbit/lazypassword/app"
	"github.com/toezbit/lazypassword/constants"
)

var padding = 1

type UiManagerImpl struct {
	gui              *gocui.Gui
	workspaceManager app.WorkspaceManager
}

func NewUiManagerImpl(g *gocui.Gui, workspaceManager app.WorkspaceManager) *UiManagerImpl {
	return &UiManagerImpl{
		gui:              g,
		workspaceManager: workspaceManager,
	}
}

func (uim *UiManagerImpl) WorkSpace() {
	maxX, maxY := uim.gui.Size()
	workSpaceView, _ := uim.gui.SetView(constants.WorkSpace, 0, 0, maxX/2-padding, maxY/2)
	workSpaceView.Title = " Work Space "

	DrawMenus(uim.gui, workSpaceView, constants.WorkSpace, uim.workspaceManager.GetWorkspaceNames())
}

func (uim *UiManagerImpl) Credential() {
	maxX, maxY := uim.gui.Size()
	accountListView, _ := uim.gui.SetView(constants.Credential, 0, maxY/2+padding, maxX/2-padding, maxY-padding)
	accountListView.Title = " Credential "

	currentWorkspace := uim.GetCurrentSelectedWorkspace()
	DrawMenus(uim.gui, accountListView, constants.Credential, uim.workspaceManager.GetCredentialNameList(currentWorkspace.Id))

}

func (uim *UiManagerImpl) Overview() {
	maxX, maxY := uim.gui.Size()
	accountDetailView, _ := uim.gui.SetView(constants.Overview, maxX/2+padding, 0, maxX-padding, maxY-padding)
	accountDetailView.Title = " Overview "
}

func (uim *UiManagerImpl) Layout(g *gocui.Gui) error {
	uim.WorkSpace()
	uim.Credential()
	uim.Overview()

	if uim.gui.CurrentView() == nil {
		uim.gui.SetCurrentView(constants.WorkSpace)
	}

	return nil

}
