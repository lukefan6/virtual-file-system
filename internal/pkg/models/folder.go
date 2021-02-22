package models

// Folder is the virtual folder that the user can add/delete/rename the files.
type Folder struct{
	// Name should be uniqued and case insensitive.
	Name string

	// Description of a folder is not a necessary field.
	Description string

	// Files that were uploaded to this folder by a user.
	Files []File
}