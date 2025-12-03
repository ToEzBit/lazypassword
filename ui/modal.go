package ui

import (
	"strings"

	"github.com/jroimartin/gocui"
	"github.com/toezbit/lazypassword/constants"
	"github.com/toezbit/lazypassword/models"
)

func (uim *UiManagerImpl) openAddWorkspaceModal(g *gocui.Gui, v *gocui.View) error {
	maxX, maxY := g.Size()

	modalAddWorkSpace, _ := g.SetView(constants.ModalAddWorkspace, maxX/2-60, maxY/2-8, maxX/2+60, maxY/2-6)
	modalAddWorkSpace.Title = " WorkSpace Name "
	modalAddWorkSpace.Editable = true

	// g.Cursor = true
	g.SetCurrentView(constants.ModalAddWorkspace)
	uim.ClearKeybindingNavigation()
	uim.ClearKeybindingGlobal()

	return nil

}

func (uim *UiManagerImpl) closeAddWorkspaceModal(g *gocui.Gui, v *gocui.View) error {
	g.DeleteView(constants.ModalAddWorkspace)
	g.SetCurrentView(constants.WorkSpace)
	// g.Cursor = false

	wsView, _ := g.View(constants.WorkSpace)
	wsView.Editable = false

	uim.SetupKeybindingNavigation()
	uim.SetupKeybindingGlobal()

	return nil
}

func (uim *UiManagerImpl) handleAddWorkspace(g *gocui.Gui, view *gocui.View) error {

	modalWorkSpaceInput := strings.TrimSpace(uim.gui.CurrentView().Buffer())

	if modalWorkSpaceInput == "" {
		return nil
	}

	uim.workspaceManager.AddWorkspace(modalWorkSpaceInput)

	uim.gui.DeleteView(constants.ModalAddWorkspace)
	uim.gui.SetCurrentView(constants.WorkSpace)

	uim.SetupKeybindingNavigation()
	uim.SetupKeybindingGlobal()

	return nil
}

func (uim *UiManagerImpl) openAddCredentialModal(g *gocui.Gui, v *gocui.View) error {

	maxX, maxY := g.Size()

	width := 80
	startX := maxX/2 - (width / 2)
	endX := maxX/2 + (width / 2)
	currentY := maxY/2 - 10

	appNameInput, _ := g.SetView(constants.ModalAddCredentialAppNameInput, startX, currentY, endX, currentY+2)
	appNameInput.Title = " App Name "
	appNameInput.Editable = true
	currentY += 4

	emailInput, _ := g.SetView(constants.ModalAddCredentialEmailInput, startX, currentY, endX, currentY+2)
	emailInput.Title = " Email/Username "
	emailInput.Editable = true
	currentY += 4

	passwordInput, _ := g.SetView(constants.ModalAddCredentialPasswordInput, startX, currentY, endX, currentY+2)
	passwordInput.Title = " Password "
	passwordInput.Editable = true
	currentY += 4

	urlInput, _ := g.SetView(constants.ModalAddCredentialUrlInput, startX, currentY, endX, currentY+2)
	urlInput.Title = " Url "
	urlInput.Editable = true
	currentY += 4

	noteInput, _ := g.SetView(constants.ModalAddCredentialNoteInput, startX, currentY, endX, currentY+8)
	noteInput.Title = " Note "
	noteInput.Editable = true

	g.SetCurrentView(constants.ModalAddCredentialAppNameInput)

	uim.ClearKeybindingNavigation()
	uim.ClearKeybindingGlobal()

	return nil

}

func (uim *UiManagerImpl) handleAddCredential(g *gocui.Gui, v *gocui.View) error {

	appNameInputView, _ := g.View(constants.ModalAddCredentialAppNameInput)
	emailInputView, _ := g.View(constants.ModalAddCredentialEmailInput)
	passwordInputView, _ := g.View(constants.ModalAddCredentialPasswordInput)
	urlInputView, _ := g.View(constants.ModalAddCredentialUrlInput)
	noteInputView, _ := g.View(constants.ModalAddCredentialNoteInput)

	currentWorkspace := uim.GetCurrentSelectedWorkspace()

	data := models.Credential{
		Id:       "1",
		AppName:  strings.TrimSpace(appNameInputView.Buffer()),
		Email:    strings.TrimSpace(emailInputView.Buffer()),
		Password: strings.TrimSpace(passwordInputView.Buffer()),
		Url:      strings.TrimSpace(urlInputView.Buffer()),
		Note:     strings.TrimSpace(noteInputView.Buffer()),
	}

	uim.workspaceManager.AddCredential(currentWorkspace.Id, data)

	g.DeleteView(constants.ModalAddCredentialAppNameInput)
	g.DeleteView(constants.ModalAddCredentialEmailInput)
	g.DeleteView(constants.ModalAddCredentialPasswordInput)
	g.DeleteView(constants.ModalAddCredentialUrlInput)
	g.DeleteView(constants.ModalAddCredentialNoteInput)

	g.SetCurrentView(constants.Overview)
	uim.SetupKeybindingNavigation()
	uim.SetupKeybindingGlobal()

	return nil
}

func (uim *UiManagerImpl) closeAddCredentialModal(g *gocui.Gui, v *gocui.View) error {
	g.DeleteView(constants.ModalAddCredentialAppNameInput)
	g.DeleteView(constants.ModalAddCredentialEmailInput)
	g.DeleteView(constants.ModalAddCredentialPasswordInput)
	g.DeleteView(constants.ModalAddCredentialUrlInput)
	g.DeleteView(constants.ModalAddCredentialNoteInput)

	g.SetCurrentView(constants.Credential)

	// wsView, _ := g.View(constants.Credential)
	// wsView.Editable = false

	uim.SetupKeybindingNavigation()
	uim.SetupKeybindingGlobal()

	return nil
}
