package fetcher

import (
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
}

func SendTestResponse() string {
	const data = `
    <svg viewBox="0 0 440 120" xmlns="http://www.w3.org/2000/svg">
        <style>#header{font:600 20px 'Segoe UI',Ubuntu,Sans-Serif;fill:#1e2e42}#song{font:500 18px 'Segoe UI',Ubuntu,Sans-Serif}#artist{font:400 18px 'Segoe UI',Ubuntu,Sans-Serif}#project{font:14px 'Segoe UI',Ubuntu,Sans-Serif;font-variant:small-caps;font-style:italic}</style>
        <rect width="440" height="120" rx="5" fill="#fffefe" stroke="#e4e2e2" stroke-width="5"/>
        <image href="https://lastfm.freetls.fastly.net/i/u/174s/478be8d73bdf783c89b709ebe7544180.jpg" width="120" height="120"/>
        <text id="header" x="135" y="35">Currently Listening To:</text>
        <text id="song" x="145" y="60">Primetime</text>
        <text id="artist" x="145" y="85">JAY-Z</text>
        <a href="https://github.com/rafaelwi/github-readme-lastfm-stats" target="_blank">text id="project" x="124" y="115">github.com/rafaelwi/github-readme-lastfm-stats</text>
        </a>
    </svg>
    `
	return data
}

func GetLastfmData(user string, apiKey string) (LastfmData, error) {
	var data LastfmData
	var result map[string]interface{}
	url := fmt.Sprintf("https://ws.audioscrobbler.com/2.0/?method=user.getrecenttracks&user=%s&limit=1&api_key=%s&format=json", user, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	json.Unmarshal(body, &result)
	lfmr, ok := result["recenttracks"].(map[string]interface{})

	// Atypical response or error
	if !ok {
		errorCode := result["error"].(float64)
		errorMsg := result["message"].(string)
		return data, fmt.Errorf("recieved err #%g: %s", errorCode, errorMsg)
	}

	// Set vars and return struct
	attr := lfmr["@attr"].(map[string]interface{})
	track0 := lfmr["track"].([]interface{})[0].(map[string]interface{})
	albumArtArr := track0["image"].([]interface{}) //[0].(map[string]interface{})

	data.User = attr["user"].(string)
	data.Song = track0["name"].(string)
	data.Artist = track0["artist"].(map[string]interface{})["#text"].(string)
	data.Album = track0["album"].(map[string]interface{})["#text"].(string)

	for _, s := range albumArtArr {
		if s.(map[string]interface{})["size"].(string) == "large" {
			data.AlbumArt = s.(map[string]interface{})["#text"].(string)
		}
	}
	return data, nil
}
