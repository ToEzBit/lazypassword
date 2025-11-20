package views

import (
	"github.com/jroimartin/gocui"
	"github.com/toezbit/lazypassword/constants"
)

func (vm *ViewManagerImpl) openAddWorkspaceModal(g *gocui.Gui, v *gocui.View) error {
	maxX, maxY := g.Size()

	modalAddWorkSpace, _ := g.SetView(constants.ModalAddWorkspace, maxX/2-60, maxY/2-8, maxX/2+60, maxY/2-6)
	modalAddWorkSpace.Title = " WorkSpace Name "
	modalAddWorkSpace.Editable = true
	// g.Cursor = true
	g.SetCurrentView(constants.ModalAddWorkspace)

	return nil

}

func (vm *ViewManagerImpl) closeAddWorkspaceModal(g *gocui.Gui, v *gocui.View) error {
	g.DeleteView(constants.ModalAddWorkspace)
	g.SetCurrentView(constants.WorkSpace)
	// g.Cursor = false

	wsView, _ := g.View(constants.WorkSpace)
	wsView.Editable = false

	return nil
}
