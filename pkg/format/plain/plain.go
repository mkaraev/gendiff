package plain

import (
	"fmt"
	"gendiff/pkg/types"
	"reflect"
	"strconv"
	"strings"
)

func stringifyValue(value reflect.Value) string {
	switch value.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(value.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(value.Uint(), 10)
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(value.Float(), 'f', -1, 64)
	case reflect.Bool:
		return strconv.FormatBool(value.Bool())
	case reflect.String:
		return strconv.Quote(value.String())
	case reflect.Array, reflect.Slice, reflect.Map:
		return "[complex value]"
	}
	return ""
}

func stringifyNode(keyPath string, rawNode interface{}) string {
	node := rawNode.(types.Dict)
	nodeType := node["type"]
	switch nodeType {
	case types.Changed:
		oldValue := stringifyValue(reflect.ValueOf(node["old"]))
		newValue := stringifyValue(reflect.ValueOf(node["new"]))
		return fmt.Sprintf("Property %s was updated. From %s to %s", keyPath, oldValue, newValue)
	case types.Added:
		value := stringifyValue(reflect.ValueOf(node["value"]))
		return fmt.Sprintf("Property %s was added with value: %s", keyPath, value)
	case types.Removed:
		return fmt.Sprintf("Property %s was removed", keyPath)
	case types.Unchanged:
		return ""
	case types.Nested:
		nodes := node["value"].(types.Dict)
		lines := make([]string, 0)
		for _, key := range nodes.Keys() {
			keyPath := fmt.Sprintf("%s.%s", keyPath, key)
			stringifiedNode := stringifyNode(keyPath, nodes[key])
			if len(stringifiedNode) > 0 {
				lines = append(lines, stringifiedNode)
			}
		}
		return strings.Join(lines, "\n")
	}
	return ""
}

func Plain(diff types.Dict) string {
	lines := make([]string, 0)
	for _, key := range diff.Keys() {
		stringifiedNode := stringifyNode(key, diff[key])
		if len(stringifiedNode) > 0 {
			lines = append(lines, stringifiedNode)
		}
	}
	return strings.Join(lines, "\n")
}
