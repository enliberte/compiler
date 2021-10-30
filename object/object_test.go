package object

import "testing"

func TestStringHashKey(t *testing.T)  {
	a1 := &String{Value: "a"}
	a2 := &String{Value: "a"}
	b1 := &String{Value: "b"}
	b2 := &String{Value: "b"}

	if a1.HashKey() != a2.HashKey() {
		t.Errorf("strings with same content have different hash keys")
	}

	if b1.HashKey() != b2.HashKey() {
		t.Errorf("strings with same content have different hash keys")
	}

	if a1.HashKey() == b1.HashKey() {
		t.Errorf("strings with different content have same hash keys")
	}
}

func TestBooleanHashKey(t *testing.T)  {
	a1 := &Boolean{Value: true}
	a2 := &Boolean{Value: true}
	b1 := &Boolean{Value: false}
	b2 := &Boolean{Value: false}

	if a1.HashKey() != a2.HashKey() {
		t.Errorf("booleans with same content have different hash keys")
	}

	if b1.HashKey() != b2.HashKey() {
		t.Errorf("booleans with same content have different hash keys")
	}

	if a1.HashKey() == b1.HashKey() {
		t.Errorf("booleans with different content have same hash keys")
	}
}

func TestIntegerHashKey(t *testing.T)  {
	a1 := &Integer{Value: 1}
	a2 := &Integer{Value: 1}
	b1 := &Integer{Value: 2}
	b2 := &Integer{Value: 2}

	if a1.HashKey() != a2.HashKey() {
		t.Errorf("integers with same content have different hash keys")
	}

	if b1.HashKey() != b2.HashKey() {
		t.Errorf("integers with same content have different hash keys")
	}

	if a1.HashKey() == b1.HashKey() {
		t.Errorf("integers with different content have same hash keys")
	}
}
