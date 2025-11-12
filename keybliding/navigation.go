package keybliding

import (
	"github.com/jroimartin/gocui"
	"github.com/toezbit/lazypassword/constants"
)

var views = []string{constants.WorkSpace, constants.AccountList}

func cycleViewDown(g *gocui.Gui, v *gocui.View) error {
	current := g.CurrentView().Name()

	for i, view := range views {
		if view == current {
			next := views[(i+1)%len(views)]
			g.SetCurrentView(next)
			break
		}
	}
	return nil
}

func cycleViewUp(g *gocui.Gui, v *gocui.View) error {
	current := g.CurrentView().Name()

	for i, view := range views {
		if view == current {
			next := views[(i+1)%len(views)]
			g.SetCurrentView(next)
			break
		}
	}
	return nil
}
