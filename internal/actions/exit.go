package actions

import "fmt"

// Exit implements Action interface
type Exit struct{}

// Exec of Exit returns false to indicate the program should be terminated
func (e *Exit) Exec(args []string) bool {
	fmt.Println("bye")
	return false
}
