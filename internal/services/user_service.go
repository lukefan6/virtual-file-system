package services

import (
	"errors"
	"strings"
	"virtual-file-system/internal/models"
)

// UserService is responsible for CRUD operations against a user
type UserService interface {
	// Register adds a user to the system.
	// If user already exists, an error is returned.
	Register(name string) error

	// Exists returns true if the given user name exists in the internal user storage.
	// The user name comparison is case insensitive
	Exists(username string) bool
}

// UserServiceImpl is the implementation of the UserService interface
type UserServiceImpl struct {
	users map[string]models.User
}

// Register adds a user to the system.
// If user already exists, an error is returned.
func (service *UserServiceImpl) Register(name string) error {
	if service.Exists(name) {
		return errors.New("user already exists")
	}

	key := service.makeKey(name)
	service.users[key] = models.User{Name: name}
	return nil
}

// Exists returns true if the given user name exists in the internal user storage.
// The user name comparison is case insensitive
func (service *UserServiceImpl) Exists(username string) bool {
	key := service.makeKey(username)
	_, exists := service.users[key]

	return exists
}

func (service *UserServiceImpl) makeKey(username string) string {
	return strings.ToLower(username)
}
