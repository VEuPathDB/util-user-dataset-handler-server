package inject

type VariableInjector interface {
	Inject(target []string) ([]string, error)
}
