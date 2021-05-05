package generator

import (
	"github.com/rafaelwi/github-readme-lastfm-stats/src/fetcher"
)

func GenerateCard(data fetcher.LastfmData) string {
	return generateCardTop() + generateCardStyle() + generateCardBody(data) + generateCardBottom()
}

// TODO: Add an input param for the theme (light, dark, dimmed, etc.)
func generateCardStyle() string {
	const lightTheme = `<style>#card{fill:#fffefe;stroke:#e4e2e2;stroke-width:5px;}#header{font:600 20px 'Segoe UI',Ubuntu,Sans-Serif;fill:#1e2e42}#song{font:500 18px 'Segoe UI',Ubuntu,Sans-Serif}#artist{font:400 18px 'Segoe UI',Ubuntu,Sans-Serif}#project{font:14px 'Segoe UI',Ubuntu,Sans-Serif;font-variant:small-caps;font-style:italic}</style>`
	return lightTheme
}

// TODO: In the future I may need to add various flags for setting specific aspects of the cards
func generateCardBody(data fetcher.LastfmData) string {
	body := "<rect id=\"card\" width=\"440\" height=\"120\" rx=\"5\"/>" +
		"<image href=\"" + data.AlbumArt + "\" width=\"120\" height=\"120\"/>" +
		"<text id=\"header\" x=\"135\" y=\"35\">Currently Listening To:</text>" +
		"<text id=\"song\" x=\"145\" y=\"60\">" + data.Song + "</text>" +
		"<text id=\"artist\" x=\"145\" y=\"85\">" + data.Artist + "</text>" +
		"<a href=\"https://github.com/rafaelwi/github-readme-lastfm-stats\" target=\"_blank\"><text id=\"project\" x=\"124\" y=\"115\">github.com/rafaelwi/github-readme-lastfm-stats</text></a>"
	return body
}

func generateCardTop() string {
	return `<svg width="440" height="120" viewBox="0 0 440 120" xmlns="http://www.w3.org/2000/svg">`
}

func generateCardBottom() string {
	return `</svg>`
}
