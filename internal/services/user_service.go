package services

// UserService is responsible for CRUD operations against a user
type UserService interface{
	// Register adds a user to the system.
	Register(name string) error
}

// UserServiceImpl is the implementation of the UserService interface
type UserServiceImpl struct {

}

// Register adds a user to the system.
func (*UserServiceImpl) Register(name string) error{
	return nil
}