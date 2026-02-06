package file

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/toezbit/lazypassword/models"
	"github.com/toezbit/lazypassword/utils"
)

func ReadFile() (models.FileData, bool) {
	var fileData models.FileData

	data, err := os.ReadFile(filepath.Join(utils.GetAppDirectoryPath(), "lazypassword-data.json"))

	if err != nil {
		if os.IsNotExist(err) {
			return models.FileData{}, true
		}
		return models.FileData{}, false
	}

	err = json.Unmarshal(data, &fileData)
	if err != nil {
		return models.FileData{}, false
	}

	return fileData, true
}
