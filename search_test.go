package gouvaddress

import (
	"testing"
)

func TestSearch(t *testing.T) {

	// We're taking examples from API doc.
	testing := map[string]string{
		"q":   "8 bd du port",
		"lat": "48.357",
		"lon": "2.37",
	}

	feedbackSearch := Search(&testing)
	feedbackReverse := Reverse(&testing)

	if feedbackSearch == nil || feedbackReverse == nil {
		t.Errorf("%s", "Value's return cannot be nil.")
	}
}
