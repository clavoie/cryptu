package cryptu_test

import (
	"testing"

	"github.com/clavoie/cryptu"
)

func TestNewDiDefs(t *testing.T) {
	defs := cryptu.NewDiDefs()

	if defs == nil {
		t.Fatal("Expecting non-nil defs")
	}

	defs2 := cryptu.NewDiDefs()
	if defs[0] == defs2[0] {
		t.Fatal("Not expecting defs to match")
	}
}
