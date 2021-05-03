package main

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

func test() {
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
