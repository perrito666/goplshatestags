package targetcode

import (
	"fmt"

	"github.com/perrito666/goplshatestags/originpack"
)

// UseSomething just makes use of importantcode
func UseSomething(s *originpack.SomethingFromOrigin) {
	if s == nil {
		return
	}
	fmt.Println(s.MeaninglessField)
}
