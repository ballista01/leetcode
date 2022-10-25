package main

import (
	"bytes"
)

func convertToTitle(columnNumber int) string {
	buf := new(bytes.Buffer)
	var b byte
	const a = byte('A')
	for num := columnNumber; num > 0; num /= 26 {
		if num%26 == 0 {
			buf.WriteByte('Z')
			num--
		}
		b = a + byte(num%26) - 1
		// fmt.Printf("remainder = %d, byte = %c\n", num%26, b)
		buf.WriteByte(b)
	}

	length := buf.Len()
	resArr := make([]byte, length)
	str := buf.String()
	for i := 0; i < length; i++ {
		resArr[i] = str[length-i-1]
	}
	return string(resArr[:])
}
