// Maktabkhooneh
// The project of the Fourth Season - the first part
package main

import (
	"fmt"
	"strings"
)

func Search(m map[string]string) string {
	for k, v := range m {
		if strings.Contains(v, k) {
			return "{" + k + "} was found."
		}
	}
	return "Not Found!"
}

func main() {
	m := map[string]string{
		"maktabkhooneh": "maktabkhooneh is the best learning platform",
	}

	fmt.Println(Search(m))
}
