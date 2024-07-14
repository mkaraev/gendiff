package format

import (
	"fmt"

	"github.com/mkaraev/gendiff/pkg/format/plain"
	"github.com/mkaraev/gendiff/pkg/format/stylish"
	"github.com/mkaraev/gendiff/pkg/types"
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
