package Area

import (
	"fmt"
)

func PrintArea(a Shape) {
	fmt.Printf("\nThis shape has an area of %.2f square meters\n.", a.Area())
}
