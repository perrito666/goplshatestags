package targetcode

import (
	"fmt"

	"github.com/perrito666/goplshatestags/importantcode"
)

// UseSomething just makes use of importantcode
func UseSomething(s *importantcode.SomethingFromOrigin) {
	if s == nil {
		return
	}
	fmt.Println(s.MeaninglessField)
}
