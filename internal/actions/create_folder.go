package actions

import (
	"fmt"
	"virtual-file-system/internal/services"
)

type createFolder struct {
	folderService services.FolderService
}

// Exec creates a folder.
func (act *createFolder) Exec(args []string) bool {
	if len(args) < 3 {
		fmt.Println("Error - Missing arguments: create_folder {username} {folder_name} {description}")
		return true
	}

	username := args[1]
	folderName := args[2]

	var description string
	if len(args) == 4 {
		description = args[3]
	}

	f, err := act.folderService.Create(folderName, username, description)
	if err != nil {
		fmt.Println("Error - ", err)
	} else {
		fmt.Println(f.ID)
	}

	return true
}
