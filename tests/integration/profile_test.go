package tests

import "testing"

func TestProfile(t *testing.T) {
	profile, _, err := client.Profile.Get()
	if err != nil {
		t.Fatalf("Profile.Get returned error: %v", err)
	}

	if profile == nil {
		t.Errorf("Get returned no profile")
	}
}
