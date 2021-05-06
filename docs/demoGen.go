package main

import (
	"fmt"
	"log"
	"os"

	"github.com/rafaelwi/github-readme-lastfm-stats/src/fetcher"
	"github.com/rafaelwi/github-readme-lastfm-stats/src/generator"
)

type key struct {
	style         string
	showScrobbles bool
}

func main() {
	numScrobbles := "102,673 scrobbles"
	m := make(map[key]fetcher.LastfmData)
	m[key{"light", false}] = fetcher.LastfmData{"Midnight in a Perfect World",
		"DJ Shadow", "Endtroducing.....",
		"https://lastfm.freetls.fastly.net/i/u/770x0/8792784aab5c4aeab78ad8525c1c2440.webp#8792784aab5c4aeab78ad8525c1c2440",
		numScrobbles}
	m[key{"dimmed", false}] = fetcher.LastfmData{"All Caps", "Madvillain",
		"Madvillainy", "https://lastfm.freetls.fastly.net/i/u/770x0/2081b8ad6ae8d40db032d0380cb9d2bc.webp#2081b8ad6ae8d40db032d0380cb9d2bc",
		numScrobbles}
	m[key{"dark", false}] = fetcher.LastfmData{"Good Morning", "Kanye West",
		"Graduation", "https://lastfm.freetls.fastly.net/i/u/770x0/3b96418b0b1321fc83a25ce14eea0643.webp#3b96418b0b1321fc83a25ce14eea0643",
		numScrobbles}
	m[key{"light", true}] = m[key{"light", false}]
	m[key{"dimmed", true}] = m[key{"dimmed", false}]
	m[key{"dark", true}] = m[key{"dark", false}]

	scrob := make(map[bool]string)
	scrob[true] = "scrobbles"
	scrob[false] = "noScrobbles"

	for k, v := range m {
		s := generator.GenerateCard(v, k.style, k.showScrobbles)
		filename := fmt.Sprintf("card/%s-%s.svg", k.style, scrob[k.showScrobbles])

		f, err := os.Create(filename)
		if err != nil {
			log.Fatal(err)
		}

		f.WriteString(s)
		err = f.Close()
		if err != nil {
			log.Fatal(err)
		}

		usage := "https://github-readme-lastfm-stats.netlify.app/.netlify/functions/card?user=st-silver"
		if k.style != "light" {
			usage += "&theme=" + k.style
		}

		if k.showScrobbles {
			usage += "&show_scrobbles=true"
		}

		fmt.Printf("```\n%s\n```\n\n", usage)
		fmt.Printf("![%s](%s)<hr>\n\n", filename, "./docs/"+filename)
	}
}
