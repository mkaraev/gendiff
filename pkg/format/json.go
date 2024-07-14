package format

import (
	"encoding/json"
	"gendiff/pkg/types"
)

func Json(diff types.Dict) string {
	data, _ := json.MarshalIndent(diff, "", "\t")
	return string(data)
}
