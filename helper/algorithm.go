package helper

func Contact(src []byte, more ...[]byte) []byte {
	for _, x := range more {
		src = append(src, x...)
	}
	return src
}
