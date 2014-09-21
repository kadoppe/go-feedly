package feedly

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestCategoriesService_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v3/categories", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `[{"id": "TestId", "label": "TestLabel"}]`)
	})

	categories, _, err := client.Categories.List()

	if err != nil {
		t.Errorf("Categories.List returned error: %v", err)
	}

	want := []Category{{Id: String("TestId"), Label: String("TestLabel")}}
	if !reflect.DeepEqual(categories, want) {
		t.Errorf("Categories.List returned error: $+v, want %+v", categories, want)
	}
}
