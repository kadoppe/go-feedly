package feedly

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestProfileService_Get(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v3/profile", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{
			"email": "jim.smith@gmail.com",
			"reader": "9080770707070700",
			"gender": "male",
			"wave": "2013.7",
			"id": "c805fcbf-3acf-4302-a97e-d82f9d7c897f",
			"google": "115562565652656565656",
			"familyName": "Smith",
			"picture": "https://www.google.com/profile_images/1771656873/bigger.jpg"
		}`)
	})

	profile, _, err := client.Profile.Get()

	if err != nil {
		t.Errorf("Profile.Get returned error: %v", err)
	}

	want := &Profile{
		Email:      String("jim.smith@gmail.com"),
		Reader:     String("9080770707070700"),
		Gender:     String("male"),
		Wave:       String("2013.7"),
		Id:         String("c805fcbf-3acf-4302-a97e-d82f9d7c897f"),
		Google:     String("115562565652656565656"),
		FamilyName: String("Smith"),
		Picture:    String("https://www.google.com/profile_images/1771656873/bigger.jpg"),
	}

	if !reflect.DeepEqual(profile, want) {
		t.Errorf("Profile.Get returned error: $+v, want %+v", profile, want)
	}
}
