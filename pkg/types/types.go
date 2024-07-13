package types

import "slices"

const (
	Removed   = "Removed"
	Added     = "Added"
	Changed   = "Changed"
	Unchanged = "Unchanged"
	Nested    = "Nested"
)

type Dict map[string]interface{}

func (d Dict) Keys() []string {
	keys := make([]string, 0)
	for key := range d {
		keys = append(keys, key)
	}
	slices.Sort(keys)
	return keys
}

func (d Dict) HasKey(key string) bool {
	_, ok := d[key]
	return ok
}

func ToDict(v interface{}) Dict {
	rawValue := v.(map[string]interface{})
	return Dict(rawValue)
}
