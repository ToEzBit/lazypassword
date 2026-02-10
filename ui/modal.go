package ui

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/jroimartin/gocui"
	"github.com/toezbit/lazypassword/constants"
	"github.com/toezbit/lazypassword/crypto"
	"github.com/toezbit/lazypassword/models"
	"github.com/toezbit/lazypassword/workspace"
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
		Id:       uuid.New().String(),
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

func (uim *UiManagerImpl) openConfirmDeleteWorkspaceModal(g *gocui.Gui, v *gocui.View) error {

	currentWorkspace := uim.GetCurrentSelectedWorkspace()

	maxX, maxY := g.Size()

	confirmModal, _ := g.SetView(constants.ModalConfirmDeleteWorkspace, maxX/2-60, maxY/2-8, maxX/2+60, maxY/2-6)
	confirmModal.Title = " Confirm Delete Workspace " + "`" + currentWorkspace.Name + "`" + "? "
	g.SetCurrentView(constants.ModalConfirmDeleteWorkspace)

	fmt.Fprintf(confirmModal, "%s\n", " Confirm (y) Cancel (n) ")

	return nil
}

func (uim *UiManagerImpl) closeConfirmDeleteWorkspaceModal(g *gocui.Gui, v *gocui.View) error {
	g.SetCurrentView(constants.WorkSpace)
	g.DeleteView(constants.ModalConfirmDeleteWorkspace)
	return nil
}

func (uim *UiManagerImpl) handleDeleteWorkspace(g *gocui.Gui, v *gocui.View) error {
	currentWorkspace := uim.GetCurrentSelectedWorkspace()

	uim.workspaceManager.DeleteWorkspace(currentWorkspace.Id)

	ClearSelectedMenuIdx()

	uim.closeConfirmDeleteWorkspaceModal(g, v)

	return nil
}

func (uim *UiManagerImpl) openConfirmDeleteCredentialModal(g *gocui.Gui, v *gocui.View) error {

	currentCredential := uim.GetCurrentSelectedCredential()

	maxX, maxY := g.Size()

	confirmModal, _ := g.SetView(constants.ModalConfirmDeleteCredential, maxX/2-60, maxY/2-8, maxX/2+60, maxY/2-6)
	confirmModal.Title = " Confirm Delete App " + "`" + currentCredential.AppName + "`" + "? "
	g.SetCurrentView(constants.ModalConfirmDeleteCredential)

	fmt.Fprintf(confirmModal, "%s\n", " Confirm (y) Cancel (n) ")

	return nil
}

func (uim *UiManagerImpl) closeConfirmDeleteCredentialModal(g *gocui.Gui, v *gocui.View) error {
	g.SetCurrentView(constants.Credential)
	g.DeleteView(constants.ModalConfirmDeleteCredential)
	return nil
}

func (uim *UiManagerImpl) handleDeleteCredential(g *gocui.Gui, v *gocui.View) error {
	currentWorkspace := uim.GetCurrentSelectedWorkspace()
	currentCredential := uim.GetCurrentSelectedCredential()

	uim.workspaceManager.DeleteCredential(currentWorkspace.Id, currentCredential.Id)

	ClearSelectedCredentialIdx()

	uim.closeConfirmDeleteCredentialModal(g, v)

	return nil
}

func (uim *UiManagerImpl) openMasterPasswordModal(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	modalMasterPassword, _ := g.SetView(constants.ModalMasterPassword, maxX/2-30, maxY/2-2, maxX/2+30, maxY/2)
	modalMasterPassword.Title = " 🔐 Enter Master Password "
	modalMasterPassword.Editable = true
	modalMasterPassword.Mask = '*'

	g.SetCurrentView(constants.ModalMasterPassword)

	return nil
}

func (uim *UiManagerImpl) closeMasterPasswordModal(g *gocui.Gui, v *gocui.View) error {
	g.DeleteView(constants.ModalMasterPassword)
	isUnlocked = true
	g.SetCurrentView(constants.WorkSpace)

	return nil
}

func (uim *UiManagerImpl) handleMasterPasswordSubmit(g *gocui.Gui, v *gocui.View) error {
	passwordInput := strings.TrimSpace(v.Buffer())

	if passwordInput == "" {
		return nil
	}

	fileData := workspace.GetFileData()
	key := crypto.DeriveKey(passwordInput, fileData.Salt)

	if fileData.PasswordHash == "" {
		hash := crypto.HashPassword(passwordInput, fileData.Salt)
		workspace.SetPasswordHash(hash)
		workspace.SetEncryptionKey(key)
		return uim.closeMasterPasswordModal(g, v)
	}

	if crypto.VerifyPassword(passwordInput, fileData.Salt, fileData.PasswordHash) {
		workspace.SetEncryptionKey(key)
		return uim.closeMasterPasswordModal(g, v)
	}

	v.Clear()
	v.SetCursor(0, 0)
	v.Title = " ❌ Wrong Password - Try Again "

	return nil
}
