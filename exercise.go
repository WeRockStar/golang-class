package golang

import (
	"strings"
)

func enc(xx, rot byte) byte {
	if xx == 32 {
		return xx
	}
	return (xx+rot)%26 + 97
}

func caesar(s string, rot byte) string {
	result := []string{}
	for _, v := range s {
		result = append(result, string(enc(byte(v), rot)))
	}
	return strings.Join(result, "")
}
