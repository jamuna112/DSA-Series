package stringss

import "strings"

func DefangIPaddr(address string) string {

	return strings.ReplaceAll(address, ".", "[.]")
}
