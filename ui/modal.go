package ui

import (
	"strings"

	"github.com/jroimartin/gocui"
	"github.com/toezbit/lazypassword/constants"
)

func (uim *UiManagerImpl) openAddWorkspaceModal(g *gocui.Gui, v *gocui.View) error {
	maxX, maxY := g.Size()

	modalAddWorkSpace, _ := g.SetView(constants.ModalAddWorkspace, maxX/2-60, maxY/2-8, maxX/2+60, maxY/2-6)
	modalAddWorkSpace.Title = " WorkSpace Name "
	modalAddWorkSpace.Editable = true

	// g.Cursor = true
	g.SetCurrentView(constants.ModalAddWorkspace)
	uim.ClearKeyblidingNavigation()
	uim.ClearKeyblidingGlobal()

	return nil

}

func (uim *UiManagerImpl) closeAddWorkspaceModal(g *gocui.Gui, v *gocui.View) error {
	g.DeleteView(constants.ModalAddWorkspace)
	g.SetCurrentView(constants.WorkSpace)
	// g.Cursor = false

	wsView, _ := g.View(constants.WorkSpace)
	wsView.Editable = false

	uim.SetupKeyblidingNavigation()
	uim.SetupKeyblidingGlobal()

	return nil
}

func (uim *UiManagerImpl) hanldeAddWorkspace(g *gocui.Gui, view *gocui.View) error {

	modalWorkSpaceInput := strings.TrimSpace(uim.gui.CurrentView().Buffer())

	if modalWorkSpaceInput == "" {
		return nil
	}

	uim.workspaceManager.AddWorkspace(modalWorkSpaceInput)

	uim.gui.DeleteView(constants.ModalAddWorkspace)
	uim.gui.SetCurrentView(constants.WorkSpace)

	uim.SetupKeyblidingNavigation()
	uim.SetupKeyblidingGlobal()

	return nil
}

func (uim *UiManagerImpl) openAddAccountModal(g *gocui.Gui, v *gocui.View) error {

	maxX, maxY := g.Size()

	width := 80
	startX := maxX/2 - (width / 2)
	endX := maxX/2 + (width / 2)
	currentY := maxY/2 - 10

	appNameInput, _ := g.SetView(constants.ModalAddAccountAppNameInput, startX, currentY, endX, currentY+2)
	appNameInput.Title = " App Name "
	appNameInput.Editable = true
	currentY += 4

	accountIdInput, _ := g.SetView(constants.ModalAddAccountIdInput, startX, currentY, endX, currentY+2)
	accountIdInput.Title = " Account Id or Email "
	accountIdInput.Editable = true
	currentY += 4

	passwordInput, _ := g.SetView(constants.ModalAddAccountPasswordInput, startX, currentY, endX, currentY+2)
	passwordInput.Title = " Password "
	passwordInput.Editable = true
	currentY += 4

	urlInput, _ := g.SetView(constants.ModalAddAccountUrlInput, startX, currentY, endX, currentY+2)
	urlInput.Title = " Url "
	urlInput.Editable = true
	currentY += 4

	noteInput, _ := g.SetView(constants.ModalAddAccountNoteInput, startX, currentY, endX, currentY+8)
	noteInput.Title = " Note "
	noteInput.Editable = true

	g.SetCurrentView(constants.ModalAddAccountAppNameInput)

	return nil

}
