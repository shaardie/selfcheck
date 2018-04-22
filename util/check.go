package util

type Check interface {
	Name() string
	Description() string
	Run() (string, int, error)
}

type SimpleCheck struct {
	name        string
	description string
	run         func() (string, int, error)
}

func (sc SimpleCheck) Name() string {
	return sc.name
}

func (sc SimpleCheck) Description() string {
	return sc.description
}

func (sc SimpleCheck) Run() (string, int, error) {
	return sc.run()
}

func SimpleChecker(name string, description string, run func() (string, int, error)) Check {
	return SimpleCheck{name, description, run}
}
