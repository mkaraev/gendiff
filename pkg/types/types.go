package types

const (
	Removed = iota
	Added
	Changed
	Unchanged
	Nested
)

type AnyDict map[string]interface{}

func (d AnyDict) Keys() []string {
	keys := make([]string, 0)
	for key := range d {
		keys = append(keys, key)
	}
	return keys
}

func (d AnyDict) HasKey(key string) bool {
	_, ok := d[key]
	return ok
}
