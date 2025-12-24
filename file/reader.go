package file

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/toezbit/lazypassword/models"
	"github.com/toezbit/lazypassword/utils"
)

func ReadFile() ([]models.Workspace, bool) {
	var fileData models.FileData

	data, err := os.ReadFile(filepath.Join(utils.GetAppDirectoryPath(), "lazypassword-data.json"))

	if err != nil {
		if os.IsNotExist(err) {
			return []models.Workspace{}, true
		}
		return []models.Workspace{}, false
	}

	err = json.Unmarshal(data, &fileData)
	if err != nil {
		return []models.Workspace{}, false
	}

	return fileData.Data, true
}
