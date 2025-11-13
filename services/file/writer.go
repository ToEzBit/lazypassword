package file

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/toezbit/lazypassword/models"
	"github.com/toezbit/lazypassword/utils"
)

func WriteFile(valuts models.ValutData) {
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
