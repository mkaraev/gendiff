package gendiff

import (
	"encoding/json"
	"io"
	"os"
	"path"

	output_format "gendiff/pkg/format"
	"gendiff/pkg/types"

	"gopkg.in/yaml.v3"
)

func GenerateDiff(firstFilePath string, secondFilePath string, format string) (string, error) {
	firstFile, err := os.Open(firstFilePath)
	if err != nil {
		return "", err
	}
	secondFile, err := os.Open(secondFilePath)
	if err != nil {
		return "", err
	}

	before, err := Parse(firstFile, path.Ext(firstFilePath))
	if err != nil {
		return "", err
	}
	after, err := Parse(secondFile, path.Ext(secondFilePath))
	if err != nil {
		return "", err
	}

	diff := BuildDiffTree(before, after)

	return output_format.Format(diff, format)
}

func Parse(r io.Reader, format string) (types.Dict, error) {
	type Decoder interface {
		Decode(v interface{}) error
	}

	data := map[string]interface{}{}
	var decoder Decoder

	switch format {
	case ".json":
		decoder = json.NewDecoder(r)
	case ".yml", ".yaml":
		decoder = yaml.NewDecoder(r)
	}

	if err := decoder.Decode(&data); err != nil {
		return nil, err
	}

	return types.Dict(data), nil
}
