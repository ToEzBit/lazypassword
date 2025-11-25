package file

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/jroimartin/gocui"
	"github.com/toezbit/lazypassword/models"
	"github.com/toezbit/lazypassword/utils"
)

func WriteFile(valuts models.VaultData) {
	jsonData, err := json.MarshalIndent(valuts, "", "")

	if err != nil {
		fmt.Printf("Error marshaling JSON: %v\n", err)
		return
	}

	appDir := utils.GetAppDirectoryPath()

	err = os.MkdirAll(appDir, 0755)
	if err != nil {
		fmt.Printf("Error creating directory: %v\n", err)
		return
	}

	err = os.WriteFile(filepath.Join(appDir, "lazypassword-data.json"), jsonData, 0644)

	if err != nil {
		fmt.Printf("Error writing file: %v\n", err)
		return
	}

}

func TestWriter(g *gocui.Gui, v *gocui.View) error {

	data := models.VaultData{}

	data.Vaults = append(data.Vaults, models.Vault{
		ID:                   "1",
		WorkSpaceName:        "kuy",
		WorkSpaceDescription: "some desc",
		Credentials: []models.Credential{{
			ID:       "1",
			AcountId: "some",
			Password: "kuy",
			App:      "kako",
			Url:      "kuykuykuy",
		}},
	})

	WriteFile(data)

	return nil
}
