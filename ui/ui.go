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
	credentialView, _ := uim.gui.SetView(constants.Credential, 0, maxY/2+padding, maxX/2-padding, maxY-padding)
	credentialView.Title = " Credential "

	currentWorkspace := uim.GetCurrentSelectedWorkspace()
	DrawMenus(uim.gui, credentialView, constants.Credential, uim.workspaceManager.GetCredentialNameList(currentWorkspace.Id))

}

func (uim *UiManagerImpl) Overview() {
	maxX, maxY := uim.gui.Size()
	overviewView, _ := uim.gui.SetView(constants.Overview, maxX/2+padding, 0, maxX-padding, maxY-padding)
	overviewView.Title = " Overview "

	currentCredential := uim.GetCurrentSelectedCredential()

	isFocusCredential := uim.gui.CurrentView() != nil && uim.gui.CurrentView().Name() == constants.Credential
	isFocusOverview := uim.gui.CurrentView() != nil && uim.gui.CurrentView().Name() == constants.Overview

	if isFocusCredential || isFocusOverview {
		DrawOverview(uim.gui, overviewView, currentCredential)
	} else {
		overviewView, _ := uim.gui.View(constants.Overview)
		overviewView.Clear()
	}

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
