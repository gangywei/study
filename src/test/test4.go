package test

import (
	"strconv"
	"strings"
)

func CompressString(S string) string {
	var str strings.Builder
	strLen := len(S)
	for i := 0; i < strLen; i++ {
		c := S[i]
		iCount := 1
		for ;i < strLen && i+1 < strLen && c == S[i+1]; i++ {
			iCount++
		}
		str.WriteByte(c)
		str.WriteString(strconv.Itoa(iCount))
	}

	if str.Len() >= strLen {
		return S
	}

	return str.String()
}
