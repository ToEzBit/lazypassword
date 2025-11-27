package ui

import (
	"fmt"
	"strings"

	"github.com/jroimartin/gocui"
	"github.com/toezbit/lazypassword/constants"
	"github.com/toezbit/lazypassword/models"
	"github.com/toezbit/lazypassword/workspace"
)

var selectedWorkspaceIdx = 0
var selectedCredentialIdx = 0

func DrawMenus(g *gocui.Gui, v *gocui.View, viewName string, menuList []string) {

	maxX, _ := v.Size()
	v.Clear()

	currentView := g.CurrentView()
	isFocus := currentView != nil && g.CurrentView().Name() == viewName

	isNotFocusWorkspace := currentView != nil && g.CurrentView().Name() != constants.WorkSpace && viewName == constants.WorkSpace

	var selectedViewIdx int

	if viewName == constants.WorkSpace {
		selectedViewIdx = selectedWorkspaceIdx
	} else if viewName == constants.Credential {
		selectedViewIdx = selectedCredentialIdx
	}

	for i, item := range menuList {
		if isFocus && i == selectedViewIdx {

			padding := strings.Repeat(" ", maxX-len(item)-2)

			fmt.Fprintf(v, "\033[40;37m%s%s\033[0m\n", item, padding)
		} else if isNotFocusWorkspace && i == selectedViewIdx {

			padding := strings.Repeat(" ", maxX-len(item)-2)

			fmt.Fprintf(v, "\033[37m%s%s\033[0m\n", item, padding)
		} else {
			fmt.Fprintf(v, "%s\n", item)
		}

	}

}

func ClearSelectedMenuIdx() {
	selectedWorkspaceIdx = 0
}

func IncressSelectedWorkspcaeIdx() {
	if selectedWorkspaceIdx+1 >= workspace.CountWorkspace() {
		return
	} else {
		selectedWorkspaceIdx = selectedWorkspaceIdx + 1
	}

}

func DecressSelectedWorkspaceIdx() {
	if selectedWorkspaceIdx <= 0 {
		return
	} else {
		selectedWorkspaceIdx = selectedWorkspaceIdx - 1
	}
}

func IncressSelectedCredentialIdx() {
	workspaces := workspace.GetWorkspaces()
	currentWorkspace := workspaces[selectedWorkspaceIdx]

	countCredential := workspace.CountCredential(currentWorkspace.Id)

	if selectedCredentialIdx+1 >= countCredential {
		return
	} else {

		selectedCredentialIdx = selectedCredentialIdx + 1
	}

}

func DecressSelectedCredentialIdx() {

	if selectedCredentialIdx <= 0 {
		return
	} else {

		selectedCredentialIdx = selectedCredentialIdx - 1
	}

}

func ResetSelectedCredentialIdx() {
	selectedCredentialIdx = 0
}

func (uim *UiManagerImpl) GetCurrentSelectedWorkspace() models.Workspace {
	workspaces := uim.workspaceManager.GetWorkspaces()
	return workspaces[selectedWorkspaceIdx]
}

func (uim *UiManagerImpl) GetCurrentSelectedCredential() models.Credential {
	workspaces := workspace.GetWorkspaces()
	currentWorkspace := workspaces[selectedWorkspaceIdx]
	currentCredential := currentWorkspace.Credentials[selectedCredentialIdx]

	return currentCredential
}
