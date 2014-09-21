package feedly

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestMarkerService_UnreadCounts(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v3/markers/counts", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{
			"unreadcounts": [
			{
				"updated": 1367539068016,
				"count": 605,
				"id": "user/c805fcbf-3acf-4302-a97e-d82f9d7c897f/category/global.all"

			},
			{
				"updated": 1367539068016,
				"count": 601,
				"id": "user/c805fcbf-3acf-4302-a97e-d82f9d7c897f/category/design"
			},
			{
				"updated": 1367539068016,
				"count": 508,
				"id": "feed/http://www.autoblog.com/rss.xml"
			},
			{
				"updated": 1367539068016,
				"count": 3,
				"id": "feed/http://feeds.feedburner.com/BakingObsession"
			},
			{
				"updated": 1367539068016,
				"count": 2,
				"id": "feed/http://vimeo.com/mattrunks/videos/rss"
			},
			{
				"updated": 1367539068016,
				"count": 1,
				"id": "feed/http://feeds.feedburner.com/DorieGreenspan"
			},
			{
				"updated": 1367539068016,
				"count": 3,
				"id": "feed/http://chasingdelicious.com/feed/"
			}
			]}`)
	})

	unreadCounts, _, err := client.Markers.ListUnreadCounts()

	if err != nil {
		t.Errorf("Markers.ListUnreadCounts returned error: %v", err)
	}

	want := []UnreadCount{
		{
			Updated: Int(1367539068016),
			Count:   Int(605),
			Id:      String("user/c805fcbf-3acf-4302-a97e-d82f9d7c897f/category/global.all"),
		},
		{
			Updated: Int(1367539068016),
			Count:   Int(601),
			Id:      String("user/c805fcbf-3acf-4302-a97e-d82f9d7c897f/category/design"),
		},
		{
			Updated: Int(1367539068016),
			Count:   Int(508),
			Id:      String("feed/http://www.autoblog.com/rss.xml"),
		},
		{
			Updated: Int(1367539068016),
			Count:   Int(3),
			Id:      String("feed/http://feeds.feedburner.com/BakingObsession"),
		},
		{
			Updated: Int(1367539068016),
			Count:   Int(2),
			Id:      String("feed/http://vimeo.com/mattrunks/videos/rss"),
		},
		{
			Updated: Int(1367539068016),
			Count:   Int(1),
			Id:      String("feed/http://feeds.feedburner.com/DorieGreenspan"),
		},
		{
			Updated: Int(1367539068016),
			Count:   Int(3),
			Id:      String("feed/http://chasingdelicious.com/feed/"),
		},
	}

	if !reflect.DeepEqual(unreadCounts, want) {
		t.Errorf("Marker.ListUnreadCounts returned error: $+v, want %+v", unreadCounts, want)
	}
}
