package release

type stringDiff []string

func (s stringDiff) valueA() string {
	return s[0]
}

func (s stringDiff) valueB() string {
	return s[1]
}

func newStringDiff(a string, b string) *stringDiff {
	return &stringDiff{a, b}
}

type commandDiff struct {
	path    string
	missing bool
	example *stringDiff
	usage   *stringDiff
	short   *stringDiff
	flags   []flagDiff
}

func newCommandDiff(path string) commandDiff {
	return commandDiff{path: path, missing: false}
}

func newCommandMissing(path string) commandDiff {
	return commandDiff{path: path, missing: true}
}

type flagDiff struct {
	name         string
	usage        *stringDiff
	defaultValue *stringDiff
}

func newFlagDiff(name string) flagDiff {
	return flagDiff{name: name}
}
