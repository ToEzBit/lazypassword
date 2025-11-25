package valut

import (
	"github.com/toezbit/lazypassword/models"
	"github.com/toezbit/lazypassword/services/file"
)

var store models.VaultData

func init() {
	store = file.ReadFile()
}
