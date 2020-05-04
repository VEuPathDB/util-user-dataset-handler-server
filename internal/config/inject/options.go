package inject

type Options struct {
	Commands []Command `yaml:"commands"`
}

type Command struct {
	Command string   `yaml:"command"`
	Args    []string `yaml:"args"`
}
