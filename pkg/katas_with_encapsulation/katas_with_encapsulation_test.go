package kataswithencapsulation

import (
	"reflect"
	"testing"
)

func TestKatas(t *testing.T) {
	f1 := NewIterable()
	f2 := NewRecursive()
	f3 := NewRecursiveReference()
	f4 := NewRecursiveSlice()

	r1 := reflect.TypeOf(f1)
	r2 := reflect.TypeOf(f2)
	r3 := reflect.TypeOf(f3)
	r4 := reflect.TypeOf(f4)

	_, got1 := r1.MethodByName("Chop")
	_, got2 := r2.MethodByName("Chop")
	_, got3 := r3.MethodByName("Chop")
	_, got4 := r4.MethodByName("Chop")

	if !got1 {
		t.Errorf("Got and Expected are not equals. Got: %v, expected: nil", got1)
	}

	if !got2 {
		t.Errorf("Got and Expected are not equals. Got: %v, expected: nil", got2)
	}

	if !got3 {
		t.Errorf("Got and Expected are not equals. Got: %v, expected: nil", got3)
	}

	if !got4 {
		t.Errorf("Got and Expected are not equals. Got: %v, expected: nil", got4)
	}

}
