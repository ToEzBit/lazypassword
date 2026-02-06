package ui

import (
	"fmt"
	"strings"

	"github.com/jroimartin/gocui"
	"github.com/toezbit/lazypassword/app"
	"github.com/toezbit/lazypassword/constants"
)

var padding = 1
var keybindingHeight = 3
var isUnlocked = false

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
	workSpaceView, _ := uim.gui.SetView(constants.WorkSpace, 0, 0, maxX/2-padding, (maxY-(keybindingHeight/2))/2)
	workSpaceView.Title = " Work Space "

	DrawMenus(uim.gui, workSpaceView, constants.WorkSpace, uim.workspaceManager.GetWorkspaceNames())
}

func (uim *UiManagerImpl) Credential() {
	maxX, maxY := uim.gui.Size()
	credentialView, _ := uim.gui.SetView(constants.Credential, 0, (maxY-(keybindingHeight/2))/2+padding, maxX/2-padding, maxY-3-padding)
	credentialView.Title = " Credential "

	currentWorkspace := uim.GetCurrentSelectedWorkspace()
	DrawMenus(uim.gui, credentialView, constants.Credential, uim.workspaceManager.GetCredentialNameList(currentWorkspace.Id))

}

func (uim *UiManagerImpl) Overview() {
	maxX, maxY := uim.gui.Size()
	overviewView, _ := uim.gui.SetView(constants.Overview, maxX/2+padding, 0, maxX-padding, maxY-keybindingHeight-padding)
	overviewView.Title = " Overview "

	currentCredential := uim.GetCurrentSelectedCredential()

	isFocusCredential := uim.gui.CurrentView() != nil && uim.gui.CurrentView().Name() == constants.Credential
	isFocusOverview := uim.gui.CurrentView() != nil && uim.gui.CurrentView().Name() == constants.Overview

	isHasCredential := len(uim.GetCurrentSelectedWorkspace().Credentials) > 0

	if (isFocusCredential || isFocusOverview) && isHasCredential {
		DrawOverview(uim.gui, overviewView, currentCredential)
	} else {
		overviewView, _ := uim.gui.View(constants.Overview)
		overviewView.Clear()
	}

}

func (uim *UiManagerImpl) KeybindingList() {
	maxX, maxY := uim.gui.Size()

	keybindingListView, _ := uim.gui.SetView(constants.KeybindingList, 0, maxY-3, maxX-padding, maxY-padding)

	keybindingListView.Clear()

	var leftText string
	rightText := "↓:l | ↑:h | ↓↓:j | ↑↑:k | Quit:q"

	currentView := uim.gui.CurrentView()

	if currentView == nil {
		return
	}

	switch currentView.Name() {
	case constants.WorkSpace:
		leftText = "Add: n | Delete: d"
	case constants.ModalAddWorkspace:
		leftText = "Confirm : <enter> | Exit : <esc>"
	case constants.Credential:
		leftText = "Add: n | Delete: d | Select: <enter>"
	case constants.ModalAddCredentialAppNameInput:
		leftText = "Confirm : <enter> | Exit : <esc> | Toggle-Tab : <tab>"
	case constants.ModalAddCredentialEmailInput, constants.ModalAddCredentialPasswordInput, constants.ModalAddCredentialUrlInput, constants.ModalAddCredentialNoteInput:
		leftText = "Toggle-Tab : <tab>"
	case constants.Overview:
		leftText = "Copy : y | Back-to-Workspace : <esc>"
	case constants.ModalMasterPassword:
		leftText = "Unlock : <enter>"
	default:
		leftText = ""
	}

	viewWidth := maxX - padding - 2

	spacesNeeded := viewWidth - len(leftText) - len(rightText)
	if spacesNeeded < 0 {
		spacesNeeded = 0
	}

	spaces := strings.Repeat(" ", spacesNeeded)

	fmt.Fprintf(keybindingListView, "\033[37m%s%s%s\033[0m\n", leftText, spaces, rightText)

	// fmt.Println(currentView.Name())

}

var modalShown = false

func (uim *UiManagerImpl) Layout(g *gocui.Gui) error {
	uim.WorkSpace()
	uim.Credential()
	uim.Overview()
	uim.KeybindingList()

	if !isUnlocked {
		uim.openMasterPasswordModal(g)
		if !modalShown {
			uim.SetupKeybindingMasterPassword()
			modalShown = true
		}
		return nil
	}

	if uim.gui.CurrentView() == nil {
		uim.gui.SetCurrentView(constants.WorkSpace)
	}

	return nil
}
