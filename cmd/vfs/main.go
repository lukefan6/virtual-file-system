package main

import (
	"fmt"
	"virtual-file-system/internal/services"
)

func main() {
	userService := services.UserServiceImpl{}
	err := userService.Register("Luke")
	fmt.Println(err)
	fmt.Println(userService)

	err2 := userService.Register("Luke")
	fmt.Println(err2)
	fmt.Println(userService)
}
