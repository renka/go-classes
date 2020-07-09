package main

import "testing"

func testValidate(t *testing.T) {
	emptyClassResult := validateBooking(Booking{})
	if emptyClassResult == "" {
		t.Error("For empty booking - should return 'false'")
	}
	missingParametersResult := validateBooking(Booking{Name: "Name", ClassId: 0, Date: "2020-09-09"})
	if missingParametersResult == "" {
		t.Error("For missing parameters - should return 'false'")
	}
	wrongNameResult := validateBooking(Booking{Name: "", ClassId: 1, Date: "2020-09-09"})
	if wrongNameResult == "" {
		t.Error("For empty name - should return 'false'")
	}
	wrongDateResult := validateBooking(Booking{Name: "Name", ClassId: 1, Date: "2020-Sept-09"})
	if wrongDateResult == "" {
		t.Error("For wrong date - should return 'false'")
	}
	successResult := validateBooking(Booking{Name: "Name", ClassId: 1, Date: "2020-09-09"})
	if successResult != "" {
		t.Error("For correct request - should return 'true'")
	}

}
