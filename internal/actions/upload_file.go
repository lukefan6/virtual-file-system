package actions

import (
	"fmt"
	"strconv"
	"virtual-file-system/internal/services"
)

type uploadFile struct {
	fileService services.FileService
}

// Exec uploads a file
func (act *uploadFile) Exec(args []string) bool {
	//upload_file {username} {folder_id} {file_name} {description}
	if len(args) < 4 {
		fmt.Println("Error - Missing arguments: upload_file {username} {folder_id} {file_name} {description}")
		return true
	}

	username := args[1]
	fileName := args[3]
	folderID, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("Error - {folder_id} should be integer")
		return true
	}

	var description string
	if len(args) == 5 {
		description = args[4]
	}

	err = act.fileService.Upload(username, folderID, fileName, description)
	if err != nil {
		fmt.Println("Error - ", err)
	} else {
		fmt.Println("Success")
	}

	return true
}
