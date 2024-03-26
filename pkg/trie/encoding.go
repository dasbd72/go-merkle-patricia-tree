package trie

import "fmt"

// func keyToNibbles(b []byte) []byte {
// 	l := len(b)*2 + 1
// 	var nibbles = make([]byte, l)
// 	for i, b := range b {
// 		nibbles[i*2] = b / 16
// 		nibbles[i*2+1] = b % 16
// 	}
// 	nibbles[l-1] = 16
// 	return nibbles
// }

// keyToNibbles converts a byte slice to a nibble slice
//
// Use this for matching homework spec
func keyToNibbles(b []byte) []byte {
	l := len(b)
	nibbles := make([]byte, l)
	for i, b := range b {
		if b >= '0' && b <= '9' {
			nibbles[i] = b - '0'
		} else {
			nibbles[i] = b - 'a' + 10
		}
	}
	return nibbles
}

func nibblesToString(b []byte) string {
	var str string
	for _, b := range b {
		str += fmt.Sprintf("%01x", b)
	}
	return str
}
