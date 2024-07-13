package format

import (
	"encoding/json"
	"gendiff/pkg/types"
)

func Json(diff types.Dict) string {
	data, _ := json.Marshal(diff)
	return string(data)
}
