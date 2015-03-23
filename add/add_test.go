package add

import (
	"testing"
)

func TestAdd(t *testing.T) {
	if Add(5, 6) != 11 {
		t.Errorf("Wrong result")
	}
}
