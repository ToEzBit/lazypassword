package views

import (
	"github.com/jroimartin/gocui"
	"github.com/toezbit/lazypassword/constants"
)

func OpenAddModalWorkSpace(g *gocui.Gui, v *gocui.View) error {
	maxX, maxY := g.Size()

	modalAddWorkSpace, _ := g.SetView(constants.ModalAddWorkspace, maxX/2-60, maxY/2-8, maxX/2+60, maxY/2-6)
	modalAddWorkSpace.Title = " WorkSpace Name "
	modalAddWorkSpace.Editable = true
	g.Cursor = true
	g.SetCurrentView(constants.ModalAddWorkspace)
	//INFO : get input value
	// fmt.Fprintf(os.Stderr, "Username: %s\n", modalAddWorkSpace.Buffer())

	return nil

}

func CloseAddModalWorkSpace(g *gocui.Gui, v *gocui.View) error {
	g.DeleteView(constants.ModalAddWorkspace)
	g.SetCurrentView(constants.WorkSpace)
	g.Cursor = false

	return nil
}
