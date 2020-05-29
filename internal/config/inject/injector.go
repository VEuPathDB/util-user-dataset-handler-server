package inject

// VariableInjector defines an object that can be used to replace command config
// template variables with concrete values.
type VariableInjector interface {

	// Inject replaces all instances of the specific VariableInjector's handled
	// template variable with a concrete value.
	//
	// The resulting string slice will be the same or greater in size than the
	// input string slice depending on the concrete values and whether or not they
	// are quoted.
	Inject(target []string) ([]string, error)
}
