package file

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/toezbit/lazypassword/models"
	"github.com/toezbit/lazypassword/utils"
)

func ReadFile() models.VaultData {
	data, err := os.ReadFile(filepath.Join(utils.GetAppDirectoryPath(), "lazypassword-data.json"))

	fmt.Println(data)

	if err != nil {
		return models.VaultData{}
	}

	var valuts = models.VaultData{}

	err = json.Unmarshal(data, &valuts)

	if err != nil {
		return models.VaultData{}
	}

	return valuts

}
