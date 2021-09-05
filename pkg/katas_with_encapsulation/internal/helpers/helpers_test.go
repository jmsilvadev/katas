package helpers

import "testing"

func TestHelperScenario(t *testing.T) {
	got := HelperScenario()
	if len(got) <= 0 {
		t.Errorf("Got and Expected are not equals. Got: %v, expected: >= 0", got)
	}

}
