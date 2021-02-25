package models

import "time"

// File is the virtual file
type File struct {
	// Name should be included as an extension and should be uniqued.
	Name string

	// Ext is the file extension
	Ext string

	// FolderID is the ID of the folder that this file is under.
	FolderID int

	// Desc of a folder is not a necessary field.
	Desc string

	// CreatedAt is the time this file was created.
	CreatedAt time.Time

	// CreatedBy is the user that created this file.
	CreatedBy string
}
