package main

import "testing"

func testValidate(t *testing.T) {
	emptyClassResult := validateClass(EntireClass{})
	if emptyClassResult == "" {
		t.Error("For empty class - should return 'false'")
	}
	missingParametersResult := validateClass(EntireClass{Class{Name: "Name", Capacity: 20, Id: 10}, "", ""})
	if missingParametersResult == "" {
		t.Error("For missing parameters - should return 'false'")
	}
	wrongNameResult := validateClass(EntireClass{Class{Name: "N", Capacity: 20, Id: 10}, "2020-09-09", "2020-09-19"})
	if wrongNameResult == "" {
		t.Error("For too short name - should return 'false'")
	}
	wrongDatesResult := validateClass(EntireClass{Class{Name: "Name", Capacity: 20, Id: 10}, "2020-09-19", "2020-09-09"})
	if wrongDatesResult == "" {
		t.Error("For wrong dates - should return 'false'")
	}
	successResult := validateClass(EntireClass{Class{Name: "Name", Capacity: 20, Id: 10}, "2020-09-19", "2020-09-09"})
	if successResult != "" {
		t.Error("For correct request - should return 'true'")
	}

}
