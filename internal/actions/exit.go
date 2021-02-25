package actions

import "fmt"

type exit struct{}

// Exec returns false to indicate the program should be terminated
func (act *exit) Exec(args []string) bool {
	fmt.Println("bye")
	return false
}
