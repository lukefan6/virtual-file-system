package actions

import (
	"fmt"
	"strconv"
	"virtual-file-system/internal/services"
)

type renameFolder struct {
	folderService services.FolderService
}

// Exec renames a folder
func (act *renameFolder) Exec(args []string) bool {
	if len(args) < 4 {
		fmt.Println("Error - Missing arguments: rename_folders {username} {folder_id} {new_folder_name}")
		return true
	}

	username := args[1]
	newFolderName := args[3]

	folderID, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("Error - {folder_id} should be integer")
		return true
	}

	err = act.folderService.Rename(folderID, newFolderName, username)
	if err != nil {
		fmt.Println("Error - ", err)
	} else {
		fmt.Println("Success")
	}

	return true
}
