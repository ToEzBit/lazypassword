package file

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/toezbit/lazypassword/models"
	"github.com/toezbit/lazypassword/utils"
)

func ReadFile() []models.Workspace {
	var fileData models.FileData

	data, err := os.ReadFile(filepath.Join(utils.GetAppDirectoryPath(), "lazypassword-data.json"))

	if err != nil {
		return []models.Workspace{}

	}

	json.Unmarshal(data, &fileData)

	return fileData.Data

}
