package services

import (
	"errors"
	"path/filepath"
	"sort"
	"time"
	"virtual-file-system/internal/models"
)

// FileService is responsible for CRUD operations against a file
type FileService interface {
	Upload(createdBy string, folderID int, filename string, desc string) error
	Delete(deletedBy string, folderID int, filename string) error
	GetAll(username string, folderID int, sortBy string, sortOrder string) ([]models.File, error)
}

// FileServiceImpl is the implementation of the FileService
type FileServiceImpl struct {
	files         map[string]models.File
	userService   UserService
	folderService FolderService
}

// Upload creates the file under the folder with given ID.
// An error will be returned if the folder or the user is not found on the system.
func (service *FileServiceImpl) Upload(createdBy string, folderID int, filename string, desc string) error {
	if !service.userService.Exists(createdBy) {
		return errors.New("authentication failed")
	}

	if !service.folderService.Exists(folderID) {
		return errors.New("folder does not exist")
	}

	if _, exists := service.files[filename]; exists {
		return errors.New("file already exists")
	}

	file := &models.File{
		FolderID:  folderID,
		Name:      filename,
		Ext:       filepath.Ext(filename),
		Desc:      desc,
		CreatedAt: time.Now(),
	}

	service.files[filename] = *file
	return nil
}

// Delete removes the specific file under the given folder.
// An error will be returned if the folder or file or user is not found on the system.
func (service *FileServiceImpl) Delete(deletedBy string, folderID int, filename string) error {
	if !service.userService.Exists(deletedBy) {
		return errors.New("authentication failed")
	}

	if !service.folderService.Exists(folderID) {
		return errors.New("folder does not exist")
	}

	if _, exists := service.files[filename]; !exists {
		return errors.New("file does not exist")
	}

	delete(service.files, filename)
	return nil
}

// GetAll retrieves all files under given folder, applying specific ordering if supplied.
// An error will be returned if the folder or the user is not found on the system.
func (service *FileServiceImpl) GetAll(username string, folderID int, sortBy string, sortOrder string) ([]models.File, error) {
	if !service.userService.Exists(username) {
		return nil, errors.New("authentication failed")
	}

	if !service.folderService.Exists(folderID) {
		return nil, errors.New("folder does not exist")
	}

	files := make([]models.File, 0, len(service.files))
	for _, file := range service.files {
		if file.FolderID == folderID {
			files = append(files, file)
		}
	}

	// sort_name asc
	defaultOrdering := func(i, j int) bool {
		return files[i].Name < files[j].Name
	}

	ordering := defaultOrdering

	if sortBy == "sort_name" && sortOrder == "dsc" {
		ordering = func(i, j int) bool {
			return files[i].Name > files[j].Name
		}
	} else if sortBy == "sort_time" && sortOrder == "asc" {
		ordering = func(i, j int) bool {
			return files[i].CreatedAt.Before(files[j].CreatedAt)
		}
	} else if sortBy == "sort_time" && sortOrder == "dsc" {
		ordering = func(i, j int) bool {
			return files[i].CreatedAt.After(files[j].CreatedAt)
		}
	} else if sortBy == "sort_extension" && sortOrder == "asc" {
		ordering = func(i, j int) bool {
			// if extensions are the same
			// use default ordering, aka order by name asc
			if files[i].Ext == files[j].Ext {
				return defaultOrdering(i, j)
			}

			return files[i].Ext < files[j].Ext
		}
	} else if sortBy == "sort_extension" && sortOrder == "dsc" {
		ordering = func(i, j int) bool {
			// if extensions are the same
			// use default ordering, aka order by name asc
			if files[i].Ext == files[j].Ext {
				return defaultOrdering(i, j)
			}

			return files[i].Ext > files[j].Ext
		}
	}

	sort.SliceStable(files, ordering)

	return files, nil
}
