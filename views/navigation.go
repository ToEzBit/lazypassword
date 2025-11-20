package views

import (
	"github.com/jroimartin/gocui"
	"github.com/toezbit/lazypassword/constants"
)

var viewList = []string{constants.WorkSpace, constants.AccountList}

func cycleViewDown(g *gocui.Gui, v *gocui.View) error {
	current := g.CurrentView().Name()

	for i, view := range viewList {
		if view == current {
			next := viewList[(i+1)%len(viewList)]
			g.SetCurrentView(next)
			break
		}
	}
	return nil
}

func cycleViewUp(g *gocui.Gui, v *gocui.View) error {
	current := g.CurrentView().Name()

	for i, view := range viewList {
		if view == current {
			next := viewList[(i+1)%len(viewList)]
			g.SetCurrentView(next)
			break
		}
	}
	return nil
}

func moveDownMenu(g *gocui.Gui, v *gocui.View) error {
	IncressSelectedIdx()
	return nil
}

func moveUpMenu(g *gocui.Gui, v *gocui.View) error {
	DecressSelectedIdx()
	return nil
}
