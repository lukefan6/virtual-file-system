package main

import (
	"fmt"
	"virtual-file-system/internal/models"
)

func main() {
	user := models.NewUser("Luke")
	fmt.Println(user)
}
