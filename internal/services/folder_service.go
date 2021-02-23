package services

import (
	"errors"
	"sort"
	"strings"
	"virtual-file-system/internal/models"
)

// FolderService is responsible for CRUD operations against a folder
type FolderService interface {
	// Create adds a folder to the system.
	// If the given `createdBy` does not match existing users in the system, an error is returned.
	// If the given folder name already exists in the system, an error is returned.
	Create(name string, createdBy string, desc string) (*models.Folder, error)

	// Delete removes a folder with given id from the system.
	// If the given `deletedBy` does not match existing users in the system, an error is returned.
	// If the given id does not match existing folders in the system, an error is returned.
	Delete(id int, deletedBy string) error

	// GetAll retrives all folders created by given user.
	// If the sorting conditions were supplied, they will be applied as well.
	// TODO Should empty folders consider an error? Currently it is not.
	// If the given `deletedBy` does not match existing users in the system, an error is returned.
	GetAll(createdBy string, orderBy string, sort string) ([]models.Folder, error)

	// Rename gives the folder with given id a new name.
	// If the given `renamedBy` does not match existing users or the original owner, an error is returned.
	// If the given id does not match existing folders in the system, an error is returned.
	Rename(id int, name string, renamedBy string) error

	// Exists returns true if the given folder id exists in the internal folder storage.
	Exists(id int) bool
}

// FolderServiceImpl is the implementation of the FolderService interface
type FolderServiceImpl struct {
	folders     map[int]models.Folder
	userService UserService
	initKey     int
}

// Create adds a folder to the system.
// If the given `createdBy` does not match existing users in the system, an error is returned.
// If the given folder name already exists in the system, an error is returned.
func (service *FolderServiceImpl) Create(name string, createdBy string, desc string) (*models.Folder, error) {
	if !service.userService.Exists(createdBy) {
		return nil, errors.New("unknown user")
	}

	if service.isNameAlreadyExist(name) {
		return nil, errors.New("folder name already exists")
	}

	key := service.makeNewKey()
	folder := &models.Folder{
		ID:          key,
		Name:        name,
		Description: desc,
		Files:       []models.File{},
	}

	return folder, nil
}

// Delete removes a folder with given id from the system.
// If the given `deletedBy` does not match existing users in the system, an error is returned.
// If the given id does not match existing folders in the system, an error is returned.
func (service *FolderServiceImpl) Delete(id int, deletedBy string) error {
	return nil
}

// GetAll retrives all folders created by given user.
// If the sorting conditions were supplied, they will be applied as well.
// TODO Should empty folders consider an error? Currently it is not.
// If the given `deletedBy` does not match existing users in the system, an error is returned.
func (service *FolderServiceImpl) GetAll(createdBy string, orderBy string, sort string) ([]models.Folder, error) {
	return []models.Folder{}, nil
}

// Rename gives the folder with given id a new name.
// If the given `renamedBy` does not match existing users or the original owner, an error is returned.
// If the given id does not match existing folders in the system, an error is returned.
func (service *FolderServiceImpl) Rename(id int, name string, renamedBy string) error {
	return nil
}

// Exists returns true if the given folder id exists in the internal folder storage.
func (service *FolderServiceImpl) Exists(id int) bool {
	return false
}

func (service *FolderServiceImpl) makeNewKey() int {
	size := len(service.folders)
	if size == 0 {
		return service.initKey
	}

	sortedKeys := service.getSortedKeys()
	return sortedKeys[size-1]
}

func (service *FolderServiceImpl) getSortedKeys() []int {
	keys := make([]int, 0, len(service.folders))
	for k := range service.folders {
		keys = append(keys, k)
	}
	return sort.IntSlice(keys)
}

func (service *FolderServiceImpl) isNameAlreadyExist(name string) bool {
	for _, v := range service.folders {
		if strings.EqualFold(v.Name, name) {
			return true
		}
	}

	return false
}
