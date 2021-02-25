package actions

import "fmt"

type unknown struct{}

// Exec does nothing and simply returns true to the caller.
func (act *unknown) Exec(args []string) bool {
	fmt.Println("Unknown command: ", args)
	return true
}
