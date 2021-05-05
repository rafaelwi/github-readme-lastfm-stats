package generator

import (
	"strings"

	"github.com/rafaelwi/github-readme-lastfm-stats/src/fetcher"
)

func GenerateCard(data fetcher.LastfmData, style string) string {
	return generateCardTop() + generateCardStyle(style) + generateCardBody(data) + generateCardBottom()
}

// TODO: Add an input param for the theme (light, dark, dimmed, etc.)
func generateCardStyle(style string) string {
	const lightTheme = `#card{fill:#fff;stroke:#e1e4e8;stroke-width:5px}#header{font:600 20px 'Segoe UI',Ubuntu,Sans-Serif;fill:#0366d6}#song{font:500 18px 'Segoe UI',Ubuntu,Sans-Serif;fill:#586069}#artist{font:400 18px 'Segoe UI',Ubuntu,Sans-Serif;fill:#586069}#project{font:14px 'Segoe UI',Ubuntu,Sans-Serif;fill:#586069;font-variant:small-caps;font-style:italic}`
	const dimmedTheme = `#card{fill:#22272e;stroke:#22272e;stroke-width:5px}#header{font:600 20px 'Segoe UI',Ubuntu,Sans-Serif;fill:#539bf5}#song{font:500 18px 'Segoe UI',Ubuntu,Sans-Serif;fill:#768390}#artist{font:400 18px 'Segoe UI',Ubuntu,Sans-Serif;fill:#768390}#project{font:14px 'Segoe UI',Ubuntu,Sans-Serif;fill:#768390;font-variant:small-caps;font-style:italic}`
	const darkTheme = `#card{fill:#0d1117;stroke:#0d1117;stroke-width:5px}#header{font:600 20px 'Segoe UI',Ubuntu,Sans-Serif;fill:#58a6ff}#song{font:500 18px 'Segoe UI',Ubuntu,Sans-Serif;fill:#8b949e}#artist{font:400 18px 'Segoe UI',Ubuntu,Sans-Serif;fill:#8b949e}#project{font:14px 'Segoe UI',Ubuntu,Sans-Serif;fill:#8b949e;font-variant:small-caps;font-style:italic}`

	// Set up map to get style css
	styleMap := make(map[string]string)
	styleMap["light"] = lightTheme
	styleMap["dimmed"] = dimmedTheme
	styleMap["dark"] = darkTheme

	// Get theme from map
	theme, ok := styleMap[strings.ToLower(style)]
	if !ok {
		theme = lightTheme
	}

	return `<style>` + theme + `</style>`
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
