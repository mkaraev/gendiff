package format

import (
	"fmt"

	"gendiff/pkg/format/plain"
	"gendiff/pkg/format/stylish"
	"gendiff/pkg/types"
)

type Formatter func(diff types.Dict) string

var Formatters map[string]Formatter = map[string]Formatter{
	"plain":   plain.Plain,
	"stylish": stylish.Stylish,
	"json":    Json,
}

func Format(diff types.Dict, format string) (string, error) {
	formatter, ok := Formatters[format]
	if !ok {
		return "", fmt.Errorf("unknown format %s", format)
	}
	return formatter(diff), nil
}
