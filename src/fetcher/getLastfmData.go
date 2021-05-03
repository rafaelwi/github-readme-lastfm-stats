package fetcher

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type LastfmData struct {
	User     string
	Song     string
	Artist   string
	Album    string
	AlbumArt string
	IsError  bool
}

func SendTestResponse() {
	const res = `
	<style>
	.header {
		font: 600 20px 'Segoe UI', Ubuntu, Sans-Serif;
		fill: #1e2e42;
	}

	.song {
		font: 500 18px 'Segoe UI', Ubuntu, Sans-Serif;
	}

	.artist {
		font: 400 18px 'Segoe UI', Ubuntu, Sans-Serif;
	}

	.project {
		font: 14px 'Segoe UI', Ubuntu, Sans-Serif;
		font-variant: small-caps;
		font-style: italic;
	}
	</style>
	<svg width="440" height="120" viewBox="0 0 440 120">
		<rect x="0" y="0" width="440" height="120" rx="5" fill="#fffefe" stroke="#e4e2e2" stroke-width="5"/>
		<image href="https://lastfm.freetls.fastly.net/i/u/174s/478be8d73bdf783c89b709ebe7544180.jpg" width="120" height="120"/>
		<g id="card-text">
			<text class="header"  x="135" y="35">Currently Listening To:</text>
			<text class="song"    x="145" y="60">Primetime</text>
			<text class="artist"  x="145" y="85">JAY-Z</text>
			<a href="https://github.com/rafaelwi/github-readme-lastfm-stats" target="_blank">
				<text class="project" x="124" y="115">github.com/rafaelwi/github-readme-lastfm-stats</text>
			</a>
		</g>
	</svg>
	`
}

func GetLastfmData() {
	user := "st-silver"
	apiKey := "ed6f93aff07a748849612cdea67dbc81"
	url := fmt.Sprintf("https://ws.audioscrobbler.com/2.0/?method=user.getrecenttracks&user=%s&limit=1&api_key=%s&format=json", user, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	dstBody := &bytes.Buffer{}
	if err := json.Indent(dstBody, body, "", "    "); err != nil {
		panic(err)
	}
	//fmt.Println(dstBody.String())

	fmt.Println("=== Unmarshaled content ===")
	var result map[string]interface{}
	json.Unmarshal(body, &result)
	lfmr := result["recenttracks"].(map[string]interface{})
	for key, val := range lfmr {
		fmt.Println(key, val)
	}

	fmt.Println("\n\n=== Individual values ===")
	var attr = lfmr["@attr"].(map[string]interface{})
	var track0 = lfmr["track"].([]interface{})[0].(map[string]interface{})
	var trackName = track0["name"].(string)
	var artist = track0["artist"].(map[string]interface{})["#text"].(string)
	var album = track0["album"].(map[string]interface{})["#text"].(string)

	var albumArtArr = track0["image"].([]interface{}) //[0].(map[string]interface{})
	var albumArt string
	for _, s := range albumArtArr {
		if s.(map[string]interface{})["size"].(string) == "large" {
			albumArt = s.(map[string]interface{})["#text"].(string)
		}
	}
	fmt.Printf("     User: %s\n", attr["user"].(string))
	fmt.Printf("     Song: %s\n", trackName)
	fmt.Printf("   Artist: %s\n", artist)
	fmt.Printf("    Album: %s\n", album)
	fmt.Printf("Album Art: %s\n", albumArt)
}
