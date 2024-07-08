package format

import (
	"encoding/json"
	"gendiff/pkg/types"
)

func Json(diff types.AnyDict) string {
	data, _ := json.Marshal(diff)
	return string(data)
}
