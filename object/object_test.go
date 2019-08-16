package object

import (
	"testing"
)

func TestStringHashKey(t *testing.T) {
	hello1 := &String{Value: "Hello world"}
	hello2 := &String{Value: "Hello world"} //identical value, but different object, so dif memory address
	diff1 := &String{Value: "goodbye"}
	diff2 := &String{Value: "goodbye"}

	if hello1.HashKey() != hello2.HashKey() {
		t.Errorf("Strings w/ same content have different hashkeys")
	}

	if diff1.HashKey() != diff2.HashKey() {
		t.Errorf("Strings w/ same content have different hashkeys")
	}

	if hello1.HashKey() == diff1.HashKey() {
		t.Errorf("Strings w/ different content have same hashkeys")
	}
}
