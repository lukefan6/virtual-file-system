package actions

import (
	"fmt"
	"strconv"
	"virtual-file-system/internal/services"
)

type getFiles struct {
	fileService services.FileService
}

// Exec get files
func (act *getFiles) Exec(args []string) bool {
	if len(args) < 3 {
		fmt.Println("Error - Missing arguments: get_files {username} {folder_id} {sort_name|sort_time|sort_extension} {asc|dsc}")
		return true
	}
	username := args[1]
	folderID, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("Error - {folder_id} should be integer")
		return true
	}

	var sortBy, ascOrDsc string
	if len(args) == 5 {
		sortBy = args[3]
		ascOrDsc = args[4]
	}

	files, err := act.fileService.GetAll(username, folderID, sortBy, ascOrDsc)
	if err != nil {
		fmt.Println("Error - ", err)
	} else {
		for _, f := range files {
			fmt.Print(f.Name)
			fmt.Print("|")
			fmt.Print(f.Ext)
			fmt.Print("|")
			fmt.Print(f.Desc)
			fmt.Print("|")

			// https://stackoverflow.com/questions/20234104/how-to-format-current-time-using-a-yyyymmddhhmmss-format
			fmt.Print(f.CreatedAt.Format("2006-01-02 15:04:05"))
			fmt.Print("|")
			fmt.Print(f.CreatedBy)
			fmt.Println()
		}
	}

	return true
}
