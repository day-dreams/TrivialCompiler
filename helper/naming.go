package helper

import (
	"errors"
	"fmt"
	"strings"
)

const (
	a = 'a'
	z = 'z'

	A = 'A'
	Z = 'Z'
)

// x和y是否同时小写，或者同时大写
func bothLowerOrBothLarger(x, y byte) bool {
	ok := (a <= x && x <= z && a <= y && y <= z) ||
		(A <= x && x <= Z && A <= y && y <= Z)
	fmt.Printf("cmping.. %c %c %v\n", x, y, ok)
	return ok
}

const (
	s0 = iota // a -> s1, A -> s2
	s1        // a -> s1, A -> s2, eof -> s3
	s2        // a -> s1, A -> s2, eof -> s3
	s3
)

func take(src string, i int) (byte, error) {
	if len(src) == 0 || i >= len(src) || i < 0 {
		return ' ', errors.New("eof.")
	}
	return src[i], nil
}

func Underlined(src string) string {

	i := 0
	state := s0
	blocks := []string{}
	last := 0
	for {
		x, err := take(src, i)
		if err != nil {
			blocks = append(blocks, strings.ToLower(src[last:i]))
			break
		}

		isLower := a <= x && x <= z

		switch state {
		case s0:
			if isLower {
				state = s1
			} else {
				state = s2
			}
		case s1:
			if isLower {
				state = s1
			} else {
				state = s2
				blocks = append(blocks, strings.ToLower(src[last:i]))
				last = i
			}
		case s2:
			if isLower {
				state = s1
				//blocks = append(blocks, strings.ToLower(src[last:i]))
				//last = i
			} else {
				state = s2
			}
		}

		i += 1
	}

	return strings.Join(blocks, "_")
}
