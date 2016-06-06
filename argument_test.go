package cape

import (
	"testing"
)

func TestArgumentDefault(t *testing.T) {
	a := newArgument()
	a.Default("123")

	if *a.intValue != 123 {
		t.Error("IntValue is wrong. Expected 123 but got", *a.intValue)
	}
	if *a.stringValue != "123" {
		t.Error("StringValue is wrong. Expected 123 but got", *a.stringValue)
	}
	if *a.boolValue != false {
		t.Error("BoolValue is wrong. Expected false but got", *a.boolValue, ". I have no idea what happened -.-")
	}

	a = newArgument()
	a.Default("hallo")

	if *a.intValue != 0 {
		t.Error("IntValue is wrong. Expected 123 but got", *a.intValue)
	}
	if *a.stringValue != "hallo" {
		t.Error("StringValue is wrong. Expected 123 but got", *a.stringValue)
	}
	if *a.boolValue != false {
		t.Error("BoolValue is wrong. Expected false but got", *a.boolValue, ". I have no idea what happened -.-")
	}

	a = newArgument()
	a.Default("true")

	if *a.intValue != 0 {
		t.Error("IntValue is wrong. Expected 123 but got", *a.intValue)
	}
	if *a.stringValue != "true" {
		t.Error("StringValue is wrong. Expected 123 but got", *a.stringValue)
	}
	if *a.boolValue != true {
		t.Error("BoolValue is wrong. Expected false but got", *a.boolValue, ". I have no idea what happened -.-")
	}
}

func TestArgumentDefault_PointerGetter(t *testing.T) {
	a := newArgument()
	a.Default("123")
	if *a.intValue != *a.Int() {
		t.Error("Int getter failed")
	}

	a = newArgument()
	a.Default("hallo")
	if *a.stringValue != *a.String() {
		t.Error("String getter failed")
	}

	a = newArgument()
	a.Default("true")
	if *a.boolValue != *a.Bool() {
		t.Error("Bool getter failed")
	}
}

func TestArgumentSetter(t *testing.T) {
	a := newArgument()
	a.set("123")

	if *a.intValue != 123 {
		t.Error("IntValue is wrong. Expected 123 but got", *a.intValue)
	}
	if *a.stringValue != "123" {
		t.Error("StringValue is wrong. Expected 123 but got", *a.stringValue)
	}
	if *a.boolValue != false {
		t.Error("BoolValue is wrong. Expected false but got", *a.boolValue, ". I have no idea what happened -.-")
	}

	a = newArgument()
	a.set("hallo")

	if *a.intValue != 0 {
		t.Error("IntValue is wrong. Expected 123 but got", *a.intValue)
	}
	if *a.stringValue != "hallo" {
		t.Error("StringValue is wrong. Expected 123 but got", *a.stringValue)
	}
	if *a.boolValue != false {
		t.Error("BoolValue is wrong. Expected false but got", *a.boolValue, ". I have no idea what happened -.-")
	}

	a = newArgument()
	a.set("true")

	if *a.intValue != 0 {
		t.Error("IntValue is wrong. Expected 123 but got", *a.intValue)
	}
	if *a.stringValue != "true" {
		t.Error("StringValue is wrong. Expected 123 but got", *a.stringValue)
	}
	if *a.boolValue != true {
		t.Error("BoolValue is wrong. Expected false but got", *a.boolValue, ". I have no idea what happened -.-")
	}
}

func TestArgumentSetter_PointerGetter(t *testing.T) {
	a := newArgument()
	a.set("123")
	if *a.intValue != *a.Int() {
		t.Error("Int getter failed")
	}

	a = newArgument()
	a.set("hallo")
	if *a.stringValue != *a.String() {
		t.Error("String getter failed")
	}

	a = newArgument()
	a.set("true")
	if *a.boolValue != *a.Bool() {
		t.Error("Bool getter failed")
	}
}
