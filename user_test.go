package scouter

import (
	"net/http"
	"testing"
)

func TestCrawler_FetchUser(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	_ = client

	mux.HandleFunc("/users/u/site_admin", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		w.WriteHeader(http.StatusNoContent)
	})

	_, err := GetGithubUser("chechiachang")
	if err != nil {
		t.Errorf("GetUsers returned error: %v", err)
	}
}
