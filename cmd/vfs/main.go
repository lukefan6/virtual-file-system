package main

import (
	"fmt"
	"virtual-file-system/internal/models"
	"virtual-file-system/internal/services"
)

func main() {
	user := models.NewUser("Luke")
	fmt.Println(user)

	userService:=services.UserServiceImpl{}
	err:=userService.Register("Luke")
	fmt.Println(err)
}
