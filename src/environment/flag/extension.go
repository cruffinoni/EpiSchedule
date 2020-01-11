package flag

type ArgType []string

func (i *ArgType) String() string {
	return "my string representation"
}

func (i *ArgType) Set(value string) error {
	*i = append(*i, value)
	return nil
}
