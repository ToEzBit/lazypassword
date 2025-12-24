package ui

import (
	"github.com/jroimartin/gocui"
	"github.com/toezbit/lazypassword/constants"
)

var viewList = []string{constants.WorkSpace, constants.Credential}

func cycleViewDown(g *gocui.Gui, v *gocui.View) error {
	current := g.CurrentView().Name()

	for i, view := range viewList {
		if view == current {
			next := viewList[(i+1)%len(viewList)]
			g.SetCurrentView(next)
			break
		}
	}
	ResetSelectedCredentialIdx()
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

	ResetSelectedCredentialIdx()
	return nil
}

func moveDownWorkspace(g *gocui.Gui, v *gocui.View) error {
	IncreaseSelectedWorkspaceIdx()
	return nil
}

func moveUpWorkspace(g *gocui.Gui, v *gocui.View) error {
	DecreaseSelectedWorkspaceIdx()
	return nil
}

func moveDownCredential(g *gocui.Gui, v *gocui.View) error {
	IncreaseSelectedCredentialIdx()
	return nil
}

func moveUpCredential(g *gocui.Gui, v *gocui.View) error {
	DecreaseSelectedCredentialIdx()
	return nil
}

func moveDownOverview(g *gocui.Gui, v *gocui.View) error {
	IncreaseSelectedOverviewIdx()
	return nil
}

func moveUpOverview(g *gocui.Gui, v *gocui.View) error {
	DecreaseSelectedOverviewIdx()
	return nil
}

func exitOverview(g *gocui.Gui, v *gocui.View) error {
	selectedOverviewIdx = 0
	g.SetCurrentView(constants.WorkSpace)
	return nil
}
