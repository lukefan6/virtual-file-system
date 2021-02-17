package models

// File is the virtual file
type File struct{
	// Name should be included as an extension and should be uniqued.
	Name string
	
	// Description of a folder is not a necessary field.
	Description string
}