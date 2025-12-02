package ui

import (
	"fmt"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/jroimartin/gocui"
	"github.com/toezbit/lazypassword/constants"
	"github.com/toezbit/lazypassword/models"
	"github.com/toezbit/lazypassword/workspace"
)

var selectedOverviewIdx = 0

const (
	appNameIdx  = 0
	emailIdx    = 1
	passwordIdx = 2
	urlIdx      = 3
	noteIdx     = 4
)

func drawRow(g *gocui.Gui, v *gocui.View, label, value string, isDrawHighlight bool) {
	const (
		reset          = "\033[0m"
		paddingSize    = 15
		highlightColor = "\033[40;37m"
	)
	maxX, _ := v.Size()

	colorCode := reset

	isFocus := g.CurrentView() != nil && g.CurrentView().Name() == constants.Overview

	paddedLabel := fmt.Sprintf("%-*s", paddingSize, label)
	fullText := fmt.Sprintf("%s : %s", paddedLabel, value)

	if isFocus && isDrawHighlight {
		colorCode = highlightColor
	}

	paddingLength := maxX - len(fullText)

	if paddingLength < 0 {
		paddingLength = 0
	}

	padding := strings.Repeat(" ", paddingLength)

	fmt.Fprintf(v, "%s%s%s%s\n",
		colorCode, fullText,
		padding, reset)

}

func DrawOverview(g *gocui.Gui, v *gocui.View, credential models.Credential) {
	v.Clear()

	fmt.Fprintln(v, "")

	drawRow(g, v, "App/Service", credential.AppName, selectedOverviewIdx == appNameIdx)

	fmt.Fprintln(v, "")
	drawRow(g, v, "Username/ID", credential.Email, selectedOverviewIdx == emailIdx)

	fmt.Fprintln(v, "")

	drawRow(g, v, "Password", "********", selectedOverviewIdx == passwordIdx)

	fmt.Fprintln(v, "")

	drawRow(g, v, "Url", credential.Url, selectedOverviewIdx == urlIdx)

	fmt.Fprintln(v, "")

	drawRow(g, v, "Note", credential.Note, selectedOverviewIdx == noteIdx)

}

func IncressSelectedOverviewIdx() {
	if selectedOverviewIdx >= 4 {
		return
	}
	selectedOverviewIdx = selectedOverviewIdx + 1

}

func DecressSelectedOverviewIdx() {
	if selectedOverviewIdx <= 0 {
		return
	}
	selectedOverviewIdx = selectedOverviewIdx - 1

}

func copyCurrentSelectedOverview(g *gocui.Gui, v *gocui.View) error {
	var valueToCopy string

	workspaces := workspace.GetWorkspaces()
	currentWorkspace := workspaces[selectedWorkspaceIdx]
	currentCredential := currentWorkspace.Credentials[selectedCredentialIdx]

	switch selectedOverviewIdx {
	case appNameIdx:
		valueToCopy = currentCredential.AppName
		break
	case emailIdx:
		valueToCopy = currentCredential.Email
		break
	case passwordIdx:
		valueToCopy = currentCredential.Password
		break
	case urlIdx:
		valueToCopy = currentCredential.Url
		break
	case noteIdx:
		valueToCopy = currentCredential.Note
	}
	clipboard.WriteAll(valueToCopy)
	return nil
}
