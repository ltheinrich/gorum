package handlers

import (
	"testing"
)

func TestGetString(t *testing.T) {
	// call GetString
	s := Request{RequestMap: map[string]interface{}{"string": "Hallo Welt"}}.GetString("string")

	// check if value is correct
	if s != "Hallo Welt" {
		// failed
		t.Error("Could not get string, value not correct")
	}
}

func TestGetStringArray(t *testing.T) {
	// call GetStringArray
	sa := Request{RequestMap: map[string]interface{}{"stringArray": []interface{}{"Hallo", "Welt"}}}.GetStringArray("stringArray")

	// check length
	if len(sa) != 2 {
		// failed
		t.Error("Could not get string array, length not correct")
	} else if sa[0] != "Hallo" {
		// failed
		t.Error("Could not get string array, first value not correct")
	} else if sa[1] != "Welt" {
		// failed
		t.Error("Could not get string array, second value not correct")
	}
}

func TestGetInt(t *testing.T) {
	// call GetInt
	i := Request{RequestMap: map[string]interface{}{"int": 10}}.GetInt("int")
	f := Request{RequestMap: map[string]interface{}{"float": 10.0}}.GetInt("float")

	// check if values are correct
	if i != 10 {
		// failed
		t.Error("Could not get int, int value not correct")
	} else if f != 10 {
		// failed
		t.Error("Could not get int, float value not correct")
	}
}

func TestGetBool(t *testing.T) {
	// call GetBool
	b := Request{RequestMap: map[string]interface{}{"bool": true}}.GetBool("bool")

	// check if value correct
	if !b {
		// failed
		t.Error("Could not get bool, value not correct")
	}
}
