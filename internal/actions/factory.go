package actions

// Factory decides which action to execute
type Factory struct{}

// CreateAction decides which action to execute
func (f *Factory) CreateAction(args []string) Action {
	if len(args) == 0 {
		return &Unknown{}
	}

	switch args[0] {
	case "exit":
		return &Exit{}
	default:
		return &Unknown{}
	}
}
