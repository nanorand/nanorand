package nanorand

import (
	"fmt"
	"testing"
)

func TestGenerate(t *testing.T) {
	for i := 1; i <= 50; i++ {
		long, err := Generate(i)
		if err != nil {
			t.Error(err)
		}

		if len(long) != i {
			t.Error(fmt.Sprintf("invalid length of generating (%d) with result: %s", i, long))
		}
	}
}

func TestGenerateShort(t *testing.T) {
	for i := 1; i <= 7; i++ {
		short, err := GenerateShort(i)
		if err != nil {
			t.Error(err)
		}

		g := len(short)
		if g != i {
			t.Error(fmt.Sprintf("invalid result length (%d) for short (%s)", g, short))
		}
	}
}

func TestGenerateArray(t *testing.T) {
	for l := 1; l <= 7; l++ {
		for a := 1; a <= 10; a++ {
			arr, err := GenerateArray(l, a, 23)
			if err != nil {
				fmt.Println(err)
				t.Error(err)
			}

			for _, e := range arr {
				m := len(fmt.Sprint(e))
				if m != l {
					t.Error(fmt.Sprintf("invalid length (%d) of array element: %d", m, e))
				}
			}
		}
	}
}
