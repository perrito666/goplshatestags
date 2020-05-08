// +build unit

package targetcode

import (
	"testing"

	"github.com/perrito666/goplshatestags/importantcode"
)

func TestUseSomething(t *testing.T) {
	type args struct {
		s *importantcode.SomethingFromOrigin
	}
	tests := []struct {
		name string
		args args
	}{{
		name: "this should trigger gopls not findin CreateSomethingFromOrigin",
		args: importantcode.CreateSomethingFromOrigin(),
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UseSomething(tt.args.s)
		})
	}
}
