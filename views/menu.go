package views

import (
	"fmt"
	"strings"

	"github.com/jroimartin/gocui"
)

var selectedMenuIdx = 0

func DrawMenus(g *gocui.Gui, v *gocui.View, viewName string, menuList []string) {

	maxX, _ := v.Size()
	v.Clear()

	currentView := g.CurrentView()
	isFocus := currentView != nil && g.CurrentView().Name() == viewName

	for i, item := range menuList {
		if isFocus && i == selectedMenuIdx {

			padding := strings.Repeat(" ", maxX-len(item)-2)

			// fmt.Fprintf(v, "\033[42;37m%s%s\033[0m\n", item, padding)
			fmt.Fprintf(v, "\033[40;37m%s%s\033[0m\n", item, padding)
		} else {
			fmt.Fprintf(v, "%s\n", item)
		}

	}

}

func ClearSelectedMenuIdx() {
	selectedMenuIdx = 0
}

func IncressSelectedIdx() {
	selectedMenuIdx = selectedMenuIdx + 1
}

func DecressSelectedIdx() {
	if selectedMenuIdx <= 0 {
		return
	} else {
		selectedMenuIdx = selectedMenuIdx - 1
	}
}
