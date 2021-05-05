package fetcher_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/rafaelwi/github-readme-lastfm-stats/src/fetcher"
)

func TestGetLastfmData_Good(t *testing.T) {
	user := "st-silver"
	apiKey := os.Getenv("LASTFM_STATS_KEY")
	res, err := fetcher.GetLastfmData(user, apiKey)
	fmt.Printf("Resulting struct: %+v\n", res)

	if err != nil {
		t.Fatalf("Error %s", err)
	}
}

func TestGetLastfmData_Bad(t *testing.T) {
	user := "NOT-A-USER"
	apiKey := "NOT-A-KEY"
	_, err := fetcher.GetLastfmData(user, apiKey)

	if err == nil {
		t.Fatalf("Error, no error!")
	}

	t.Log(err)
}
