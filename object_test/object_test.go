package object_test

import (
	"monkey/object/boolean"
	objectInteger "monkey/object/integer"
	objectString "monkey/object/string"
	"testing"
)

func TestStringHashKey(t *testing.T) {
	hello1 := objectString.New("Hello World")
	hello2 := objectString.New("Hello World")
	diff1 := objectString.New("My name is johnny")
	diff2 := objectString.New("My name is johnny")

	if hello1.HashKey() != hello2.HashKey() {
		t.Errorf("strings with same content have different hash keys")
	}

	if diff1.HashKey() != diff2.HashKey() {
		t.Errorf("strings with same content have different hash keys")
	}

	if hello1.HashKey() == diff1.HashKey() {
		t.Errorf("strings with different content have same hash keys")
	}
}

func TestBooleanHashKey(t *testing.T) {
	true1 := boolean.GetBoolean(true)
	true2 := boolean.GetBoolean(true)
	false1 := boolean.GetBoolean(false)
	false2 := boolean.GetBoolean(false)

	if true1.HashKey() != true2.HashKey() {
		t.Errorf("trues do not have same hash key")
	}

	if false1.HashKey() != false2.HashKey() {
		t.Errorf("falses do not have same hash key")
	}

	if true1.HashKey() == false1.HashKey() {
		t.Errorf("true has same hash key as false")
	}
}

func TestIntegerHashKey(t *testing.T) {
	one1 := objectInteger.New(1)
	one2 := objectInteger.New(1)
	two1 := objectInteger.New(2)
	two2 := objectInteger.New(2)

	if one1.HashKey() != one2.HashKey() {
		t.Errorf("integers with same content have twoerent hash keys")
	}

	if two1.HashKey() != two2.HashKey() {
		t.Errorf("integers with same content have twoerent hash keys")
	}

	if one1.HashKey() == two1.HashKey() {
		t.Errorf("integers with twoerent content have same hash keys")
	}
}
