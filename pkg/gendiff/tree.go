package gendiff

import (
	"reflect"
	"slices"

	"gendiff/pkg/types"
)

func BuildDiffTree(before, after types.AnyDict) types.AnyDict {
	keys := GetKeysUnion(before, after)
	diff := types.AnyDict{}
	for _, key := range keys {
		if isRemoved(before, after, key) {
			diff[key] = map[string]interface{}{
				"type":  types.Removed,
				"value": before[key],
			}
		} else if isAdded(before, after, key) {
			diff[key] = map[string]interface{}{
				"type":  types.Added,
				"value": after[key],
			}
		} else if reflect.DeepEqual(before[key], after[key]) {
			diff[key] = map[string]interface{}{
				"type":  types.Unchanged,
				"value": before[key],
			}
		} else if isAnyDict(before[key]) && isAnyDict(after[key]) {
			diff[key] = map[string]interface{}{
				"type": types.Nested,
				"value": BuildDiffTree(
					before[key].(types.AnyDict),
					after[key].(types.AnyDict),
				),
			}
		} else {
			diff[key] = map[string]interface{}{
				"type": types.Changed,
				"old":  before[key],
				"new":  after[key],
			}
		}
	}
	return diff
}

func GetKeysUnion(first, second types.AnyDict) []string {
	keys := first.Keys()
	keys = append(keys, second.Keys()...)
	slices.Sort(keys)
	return slices.Compact(keys)
}

func isAnyDict(v interface{}) bool {
	_, ok := v.(types.AnyDict)
	return ok
}

func isRemoved(before, after types.AnyDict, key string) bool {
	return before.HasKey(key) && !after.HasKey(key)
}

func isAdded(before, after types.AnyDict, key string) bool {
	return !before.HasKey(key) && after.HasKey(key)
}
