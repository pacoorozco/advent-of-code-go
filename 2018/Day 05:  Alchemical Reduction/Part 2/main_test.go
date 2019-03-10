package main

import (
	"fmt"
	"testing"
)

func TestRemoveUnitType(t *testing.T) {
	tt := []struct {
		in, unit, out string
	}{
		{"dabAcCaCBAcCcaDA", "", "dabAcCaCBAcCcaDA"},
		{"dabAcCaCBAcCcaDA", "a", "dbcCCBcCcD"},
		{"dabAcCaCBAcCcaDA", "A", "dbcCCBcCcD"},
		{"dabAcCaCBAcCcaDA", "b", "daAcCaCAcCcaDA"},
		{"dabAcCaCBAcCcaDA", "C", "dabAaBAaDA"},
		{"dabAcCaCBAcCcaDA", "e", "dabAcCaCBAcCcaDA"},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%s minus '%s'", tc.in, tc.unit), func(t *testing.T) {
			out := removeUnitType(tc.unit, tc.in)
			if tc.out != out {
				t.Errorf("expected %s, got %s", tc.out, out)
			}
		})
	}
}

func TestStep(t *testing.T) {
	tt := []struct {
		in, out string
		ok      bool
	}{
		{"aA", "", true},
		{"aa", "aa", false},
		{"aAbB", "bB", true},
		{"aAa", "a", true},
		{"dabAcCaCBAcCcaDA", "dabAaCBAcCcaDA", true},
	}

	for _, tc := range tt {
		t.Run(tc.in, func(t *testing.T) {
			out, ok := step(tc.in)
			if tc.ok != ok {
				t.Errorf("expected ok to be %v; got %v", tc.ok, ok)
			}
			if tc.out != out {
				t.Errorf("expected out to be %v; got %v", tc.out, out)
			}
		})
	}
}

func TestOpposite(t *testing.T) {
	tt := []struct {
		a, b byte
		res  bool
	}{
		{'a', 'A', true},
		{'A', 'a', true},
		{'a', 'a', false},
		{'a', 'b', false},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%c-%c", tc.a, tc.b), func(t *testing.T) {
			if res := opposite(tc.a, tc.b); res != tc.res {
				t.Fatalf("expected opposite of %c and %c to be %v; got %v", tc.a, tc.b, tc.res, res)
			}
		})
	}
}
