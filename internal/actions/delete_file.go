package actions

import (
	"fmt"
	"strconv"
	"virtual-file-system/internal/services"
)

type deleteFile struct {
	fileService services.FileService
}

// Exec deletes a file
func (act *deleteFile) Exec(args []string) bool {
	//delete_file {username} {folder_id} {file_name}
	if len(args) < 4 {
		fmt.Println("Error - Missing arguments: delete_folder {username} {folder_id}")
		return true
	}

	username := args[1]
	fileName := args[3]
	folderID, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("Error - {folder_id} should be integer")
		return true
	}

	err = act.fileService.Delete(username, folderID, fileName)
	if err != nil {
		fmt.Println("Error - ", err)
	} else {
		fmt.Println("Success")
	}

	return true
}
