package tests

import "testing"

func TestCategories(t *testing.T) {
	categories, _, err := client.Categories.List()
	if err != nil {
		t.Fatalf("Categories.List returned error: %v", err)
	}

	if len(categories) == 0 {
		t.Errorf("List returned no categories")
	}
}
