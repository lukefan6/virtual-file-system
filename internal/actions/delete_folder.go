package actions

import (
	"fmt"
	"strconv"
	"virtual-file-system/internal/services"
)

type deleteFolder struct {
	folderService services.FolderService
}

// Exec deletes a folder.
func (act *deleteFolder) Exec(args []string) bool {
	if len(args) < 3 {
		fmt.Println("Error - Missing arguments: delete_folder {username} {folder_id}")
		return true
	}

	username := args[1]
	folderID, err := strconv.Atoi(args[2])

	if err != nil {
		fmt.Println("Error - {folder_id} should be integer")
		return true
	}

	err = act.folderService.Delete(folderID, username)
	if err != nil {
		fmt.Println("Error - ", err)
	} else {
		fmt.Println("Success")
	}

	return true
}
