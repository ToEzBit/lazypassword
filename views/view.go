package views

import (
	"fmt"
	"strings"

	"github.com/jroimartin/gocui"
	"github.com/toezbit/lazypassword/constants"
)

var padding = 1

func WorkSpace(g *gocui.Gui) {
	maxX, maxY := g.Size()
	v, _ := g.SetView(constants.WorkSpace, 0, 0, maxX/2-padding, maxY/2)
	v.Title = " Work Space "

}

func AccountList(g *gocui.Gui) {
	maxX, maxY := g.Size()
	v, _ := g.SetView(constants.AccountList, 0, maxY/2+padding, maxX/2-padding, maxY-padding)
	v.Title = " Account List "
}

func AccountDetail(g *gocui.Gui) {
	maxX, maxY := g.Size()
	v, _ := g.SetView(constants.AccountdDetail, maxX/2+padding, 0, maxX-padding, maxY-padding)
	v.Title = " Account Detail "
}

var menuItems = []string{"Item 1", "Item 2", "Item 3"}
var selectedIdx = 0

func LeftView(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	v, err := g.SetView("left-view", 0, 0, maxX/2-4, maxY/5)

	if err != gocui.ErrUnknownView {
		return err
	}

	fmt.Fprintln(v, "left view")

	return nil
}

func RightView(g *gocui.Gui) error {

	maxX, maxY := g.Size()
	if v, err := g.SetView("right-view", maxX/2, 0, maxX-4, maxY/5); err != nil {
		v.Title = "Right View"

		if err != gocui.ErrUnknownView {
			return err
		}

		fmt.Fprintln(v, "Right view")
		DrawMenu(g, v, "right-view")

		g.SetKeybinding("right-view", 'j', gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
			selectedIdx++
			DrawMenu(g, v, "right-view")
			return nil
		})

	}
	return nil
}

func DrawMenu(g *gocui.Gui, v *gocui.View, view string) {
	maxX, _ := v.Size()
	v.Clear()

	currentView := g.CurrentView()
	isFocus := currentView != nil && g.CurrentView().Name() == view

	for i, item := range menuItems {
		if isFocus && i == selectedIdx {
			padding := strings.Repeat(" ", maxX-len(item)-1)
			fmt.Fprintf(v, "\033[42;37m%s%s\033[0m\n", item, padding)

		} else {
			fmt.Fprintf(v, "  %s\n", item)
		}

	}

}
