package stylish

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/mkaraev/gendiff/pkg/types"
)

var Marks = map[string]string{
	types.Removed:   "-",
	types.Added:     "+",
	types.Unchanged: " ",
}

const Indent = "    "

func stringifyValue(value reflect.Value, depth int) string {
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
		return value.String()
	case reflect.Array, reflect.Slice:
		elems := make([]string, value.Len())
		for i := 0; i < value.Len(); i++ {
			elems[i] = stringifyValue(value.Index(i), depth)
		}
		return fmt.Sprintf("[%s]", strings.Join(elems, ", "))
	case reflect.Map:
		lines := make([]string, 0)
		lines = append(lines, "{")
		for _, key := range value.MapKeys() {
			lines = append(lines, fmt.Sprintf(
				"%s%s: %s", strings.Repeat(Indent, depth+1),
				key,
				stringifyValue(reflect.ValueOf(value.MapIndex(key).Interface()), depth+1),
			))
		}
		lines = append(lines, fmt.Sprintf("%s}", strings.Repeat(Indent, depth)))
		return strings.Join(lines, "\n")
	}
	return "null"
}

func stringifyNode(key string, rawNode interface{}, depth int) string {
	node := rawNode.(types.Dict)
	nodeType := node["type"]
	switch nodeType {
	case types.Nested:
		nodes := node["value"].(types.Dict)
		lines := make([]string, 0)
		lines = append(lines, fmt.Sprintf("%s%s: {", strings.Repeat(Indent, depth), key))
		for _, key := range nodes.Keys() {
			stringifiedNode := stringifyNode(key, nodes[key], depth+1)
			lines = append(lines, stringifiedNode)
		}
		lines = append(lines, fmt.Sprintf("%s}", strings.Repeat(Indent, depth)))
		return strings.Join(lines, "\n")
	case types.Added, types.Removed, types.Unchanged:
		mark := Marks[nodeType.(string)]
		stringifiedNode := fmt.Sprintf(
			"%s  %s %s: %s",
			strings.Repeat(Indent, depth-1),
			mark,
			key,
			stringifyValue(reflect.ValueOf(node["value"]), depth),
		)
		return stringifiedNode
	case types.Changed:
		removedMark := Marks[types.Removed]
		removed := fmt.Sprintf(
			"%s  %s %s: %s",
			strings.Repeat(Indent, depth-1),
			removedMark,
			key,
			stringifyValue(reflect.ValueOf(node["old"]), depth),
		)
		addedMark := Marks[types.Added]
		added := fmt.Sprintf(
			"%s  %s %s: %s",
			strings.Repeat(Indent, depth-1),
			addedMark,
			key,
			stringifyValue(reflect.ValueOf(node["new"]), depth),
		)
		return fmt.Sprintf("%s\n%s", removed, added)
	}

	return ""
}

func Stylish(diff types.Dict) string {
	lines := make([]string, 0)
	lines = append(lines, "{")
	keys := diff.Keys()
	for _, key := range keys {
		lines = append(lines, stringifyNode(key, diff[key], 1))
	}
	lines = append(lines, "}")
	return strings.Join(lines, "\n")
}
