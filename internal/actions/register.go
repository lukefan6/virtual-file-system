package actions

import (
	"fmt"
	"virtual-file-system/internal/services"
)

type register struct {
	userService services.UserService
}

// Exec registers the user and returns true regardless of errors
func (act *register) Exec(args []string) bool {
	if len(args) != 2 {
		fmt.Println("Error - Missing arguments: register {username}.")
		return true
	}

	username := args[1]

	if err := act.userService.Register(username); err != nil {
		fmt.Println("Error - ", err)
	} else {
		fmt.Println("Success")
	}

	return true
}
