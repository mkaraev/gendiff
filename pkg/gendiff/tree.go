package gendiff

import (
	"reflect"
	"slices"

	"github.com/mkaraev/gendiff/pkg/types"
)

func BuildDiffTree(before, after types.Dict) types.Dict {
	keys := GetKeysUnion(before, after)
	diff := types.Dict{}
	for _, key := range keys {
		if isRemoved(before, after, key) {
			diff[key] = types.Dict{
				"type":  types.Removed,
				"value": before[key],
			}
		} else if isAdded(before, after, key) {
			diff[key] = types.Dict{
				"type":  types.Added,
				"value": after[key],
			}
		} else if reflect.DeepEqual(before[key], after[key]) {
			diff[key] = types.Dict{
				"type":  types.Unchanged,
				"value": before[key],
			}
		} else if isDict(before[key]) && isDict(after[key]) {
			diff[key] = types.Dict{
				"type": types.Nested,
				"value": BuildDiffTree(
					types.ToDict(before[key]),
					types.ToDict(after[key]),
				),
			}
		} else {
			diff[key] = types.Dict{
				"type": types.Changed,
				"old":  before[key],
				"new":  after[key],
			}
		}
	}
	return diff
}

func GetKeysUnion(first, second types.Dict) []string {
	keys := first.Keys()
	keys = append(keys, second.Keys()...)
	slices.Sort(keys)
	return slices.Compact(keys)
}

func isDict(v interface{}) bool {
	value := reflect.ValueOf(v)
	return value.Kind() == reflect.Map
}

func isRemoved(before, after types.Dict, key string) bool {
	return before.HasKey(key) && !after.HasKey(key)
}

func isAdded(before, after types.Dict, key string) bool {
	return !before.HasKey(key) && after.HasKey(key)
}
