package actions

import "virtual-file-system/internal/services"

// Factory decides which action to execute
type Factory struct{}

// CreateAction decides which action to execute
func (f *Factory) CreateAction(args []string) Action {
	if len(args) == 0 {
		return &unknown{}
	}

	serviceFactory := services.GetFactory()

	switch args[0] {
	case "register":
		return &register{serviceFactory.GetUserService()}
	case "create_folder":
		return &createFolder{serviceFactory.GetFolderService()}
	case "delete_folder":
		return &deleteFolder{serviceFactory.GetFolderService()}
	case "exit":
		return &exit{}
	default:
		return &unknown{}
	}
}
