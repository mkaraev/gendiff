package format

import (
	"fmt"
	"gendiff/pkg/types"
)

type Formatter func(diff types.AnyDict) string

var Formatters map[string]Formatter = map[string]Formatter{
	"plain":   Plain,
	"stylish": Stylish,
	"json":    Json,
}

func Format(diff types.AnyDict, format string) (string, error) {
	formatter, ok := Formatters[format]
	if !ok {
		return "", fmt.Errorf("unknown format %s", format)
	}
	return formatter(diff), nil
}
