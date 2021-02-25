package actions

import (
	"fmt"
	"virtual-file-system/internal/services"
)

type getFolders struct {
	folderService services.FolderService
}

// Exec gets folders
func (act *getFolders) Exec(args []string) bool {
	if len(args) < 2 {
		fmt.Println("Error - Missing arguments: get_folders {username} {sort_name | sort_time} {asc|dsc}")
		return true
	}

	username := args[1]

	var sortBy, ascOrDsc string
	if len(args) == 4 {
		sortBy = args[2]
		ascOrDsc = args[3]
	}

	folders, err := act.folderService.GetAll(username, sortBy, ascOrDsc)
	if err != nil {
		fmt.Println("Error - ", err)
	} else {
		for _, f := range folders {
			fmt.Print(f.ID)
			fmt.Print("|")
			fmt.Print(f.Name)
			fmt.Print("|")
			fmt.Print(f.Description)
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
