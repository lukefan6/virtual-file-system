package models

// User is anyone who uses the file management system.
type User struct {
	// Name is the identifier, which should be unique throughout the application, but ​case insensitive.
	Name string

	// Folders can be created by each users.
	Folders []Folder
}

// NewUser creates a user with no folders created.
func NewUser(name string) *User {
	user := User{Name: name}
	return &user
}
