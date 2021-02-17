package main

import (
	"fmt"
	"virtual-file-system/pkg/models"
)

func main() {
    user := models.NewUser("Luke")
    fmt.Println(user)
}