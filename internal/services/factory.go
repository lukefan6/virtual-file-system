package services

import "virtual-file-system/internal/models"

// Factory manages service instances across services package.
type Factory struct {
	userService   UserService
	folderService FolderService
}

var instance *Factory

// GetFactory returns singleton instance of service factory
func GetFactory() *Factory {
	if instance == nil {
		instance = &Factory{}
	}

	return instance
}

// GetUserService returns an instance of UserService
func (f *Factory) GetUserService() UserService {
	if f.userService == nil {
		f.userService = &UserServiceImpl{
			make(map[string]models.User),
		}
	}

	return f.userService
}

// GetFolderService returns an instance of FolderService
func (f *Factory) GetFolderService() FolderService {
	if f.folderService == nil {
		f.folderService = &FolderServiceImpl{
			folders:     make(map[int]models.Folder),
			userService: f.GetUserService(),
			initKey:     1001,
		}
	}

	return f.folderService
}
