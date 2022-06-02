package parse

import (
	parser "github.com/openvenues/gopostal/parser"
)

func AddrParse(addr string) map[string]string {
	parsed := parser.ParseAddress(addr)
	result := make(map[string]string)

	for i := 0; i < len(parsed); i++ {
		result[parsed[i].Label] = parsed[i].Value
	}

	return result
}
