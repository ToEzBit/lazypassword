package valut

import (
	"github.com/jroimartin/gocui"
	"github.com/toezbit/lazypassword/models"
	"github.com/toezbit/lazypassword/services/file"
)

type ValutManagerImpl struct {
	gui *gocui.Gui
}

func NewValutManagerImpl(g *gocui.Gui) *ValutManagerImpl {
	return &ValutManagerImpl{
		gui: g,
	}
}

var store models.VaultData

func init() {
	store = file.ReadFile()
}

func Save() {
	file.WriteFile(store)
}

func (v *ValutManagerImpl) GetValut() models.VaultData {
	return store
}
