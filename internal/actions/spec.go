package actions

// Action defines the interface to execute an action
type Action interface {
	// Exec executes the action, returns true if next action should be executed.
	Exec(args []string) bool
}
