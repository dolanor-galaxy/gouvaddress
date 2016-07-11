package gouvaddress

import (
	"testing"
)

func TestSearch(t *testing.T) {

	testing := map[string]string{
		"q":   "8 bd du port",
		"lat": "48.357",
		"lon": "2.37",
	}

	feedback_search := Search(&testing)
	feedback_reverse := Reverse(&testing)

	if feedback_search == nil || feedback_reverse == nil {
		t.Errorf("%s", "test")
	}
}
