package ui

import (
	"fmt"
	"strings"

	"github.com/jroimartin/gocui"
	"github.com/toezbit/lazypassword/workspace"
)

var selectedWorkspaceIdx = 0

func DrawMenus(g *gocui.Gui, v *gocui.View, viewName string, menuList []string) {

	maxX, _ := v.Size()
	v.Clear()

	currentView := g.CurrentView()
	isFocus := currentView != nil && g.CurrentView().Name() == viewName

	for i, item := range menuList {
		if isFocus && i == selectedWorkspaceIdx {

			padding := strings.Repeat(" ", maxX-len(item)-2)

			fmt.Fprintf(v, "\033[40;37m%s%s\033[0m\n", item, padding)
		} else {
			fmt.Fprintf(v, "%s\n", item)
		}

	}

}

func ClearSelectedMenuIdx() {
	selectedWorkspaceIdx = 0
}

func IncressSelectedIdx() {
	if selectedWorkspaceIdx+1 >= workspace.CountWorkspace() {
		return
	} else {
		selectedWorkspaceIdx = selectedWorkspaceIdx + 1
	}

}

func DecressSelectedIdx() {
	if selectedWorkspaceIdx <= 0 {
		return
	} else {
		selectedWorkspaceIdx = selectedWorkspaceIdx - 1
	}
}

// func (vm *ViewManagerImpl) GetCurrentSelectedWorkspace() models.Vault {
// 	valuts := vm.valutManager.GetValut()
// 	return valuts[selectedWorkspaceIdx]
// }
