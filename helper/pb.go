package helper

var (
	mapping = map[string]string{
		"string":  "string",
		"int":     "int32",
		"int64":   "int32", //TODO
		"bool":    "bool",
		"float64": "float64",
	}
)

func GoType2pbType(src string) string {
	return mapping[src]
}
