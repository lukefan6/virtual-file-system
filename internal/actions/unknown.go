package actions

import "fmt"

// Unknown implements Actions interface.
type Unknown struct{}

// Exec of Unknown does nothing and simply returns true to the caller.
func (e *Unknown) Exec(args []string) bool {
	fmt.Println("Unknown command: ", args)
	return true
}
