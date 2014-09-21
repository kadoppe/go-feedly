package tests

import (
	"fmt"
	"os"

	"github.com/kadoppe/go-feedly/feedly"

	"code.google.com/p/goauth2/oauth"
)

var (
	client *feedly.Client
	auth   bool
)

func init() {
	token := os.Getenv("FEEDLY_AUTH_TOKEN")

	if token == "" {
		println("!!! No OAuth token. Some tests won't run. !!!\n")
		client = feedly.NewClient(nil)
	} else {
		t := &oauth.Transport{
			Token: &oauth.Token{AccessToken: token},
		}
		client = feedly.NewClient(t.Client())
		auth = true
	}
}

func checkAuth(name string) bool {
	if !auth {
		fmt.Printf("No auth - skipping portions of %v\n", name)
	}
	return auth
}
