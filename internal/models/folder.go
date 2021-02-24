package models

import "time"

// Folder is the virtual folder that the user can add/delete/rename the files.
type Folder struct {
	// The unique identifier.
	ID int

	// Name should be uniqued and case insensitive.
	Name string

	// Description of a folder is not a necessary field.
	Description string

	// CreatedBy is the user that created this folder.
	CreatedBy string

	// CreatedAt is the time this folder was created.
	CreatedAt time.Time

	// Files that were uploaded to this folder by a user.
	Files []File
}
