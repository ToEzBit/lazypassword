package ui

import (
	"github.com/jroimartin/gocui"
	"github.com/toezbit/lazypassword/constants"
)

var addCredentialInputList = []string{constants.ModalAddCredentialAppNameInput, constants.ModalAddCredentialEmailInput, constants.ModalAddCredentialPasswordInput, constants.ModalAddCredentialUrlInput, constants.ModalAddCredentialNoteInput}

func toggleFocusAddCredentialInput(g *gocui.Gui, v *gocui.View) error {
	current := g.CurrentView().Name()

	for i, view := range addCredentialInputList {
		if view == current {
			next := addCredentialInputList[(i+1)%len(addCredentialInputList)]
			g.SetCurrentView(next)
			break
		}
	}
	return nil

}

func handleSelectCredential(g *gocui.Gui, v *gocui.View) error {
	g.SetCurrentView(constants.Overview)

	return nil
}
