package file

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/toezbit/lazypassword/models"
	"github.com/toezbit/lazypassword/utils"
)

func WriteFile(data []models.Workspace) {

	fileData := models.FileData{
		Version: "1.0",
		Data:    data,
	}

	jsonData, err := json.MarshalIndent(fileData, "", "")

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

func TestWriter() {
	var data []models.Workspace

	res := models.Workspace{
		Id:   "1",
		Name: "ai-soft",
		Credentials: []models.Credential{{
			Id:       "1",
			AppName:  "facebook",
			Url:      "facebook.com",
			Email:    "john@email.com",
			Password: "1234567890",
		}},
	}

	data = append(data, res)

	WriteFile(data)

}
