package nanorand

import (
	"fmt"
	"testing"
)

func TestGenerate(t *testing.T) {
	for i := 1; i <= 50; i++ {
		long, err := Gen(i)
		if err != nil {
			t.Error(err)
		}

		if len(long) != i {
			t.Error(fmt.Sprintf("invalid length of generating (%d) with result: %s", i, long))
		}
	}
}
