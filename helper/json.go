package helper

import "encoding/json"

// 转json，用于调试
func ToPrettyJson(v interface{}) string {
	data, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		return err.Error()
	}
	return string(data)
}
