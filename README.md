# GitHub Readme Last.fm Stats
Dynamically generated last.fm stats in your profile readme

[![Netlify Status](https://api.netlify.com/api/v1/badges/f13a0213-6b93-4c0f-a3a3-5288199e9f42/deploy-status)](https://app.netlify.com/sites/github-readme-lastfm-stats/deploys)

<hr>

## Contents
- [Usage](#usage)
- [Options](#options)
- [Demo](#demo)
- [Development & Deployment](#development--deployment)
- [Issues, Requests, and Contributing](#issues-requests-and-contributing)

<hr>

## Usage
To embed in your readme:
```
![alt text](https://github-readme-lastfm-stats.netlify.app/.netlify/functions/card?user=USERNAME&OPTIONS)
```
or
```
<img src="https://github-readme-lastfm-stats.netlify.app/.netlify/functions/card?user=USERNAME&OPTIONS"
```
The latter will allow you to better format the card (ex. `align="center"`)

<hr>

## Options
| Option | Example | Description |
| ------ | ------- | ----------- |
| user | `user=st-silver` | **(required)** The user to fetch info for |
| theme | `theme=light` **(default)**<br>`theme=dimmed`<br>`theme=dark` | **(optional)** The theme of the card. See the demo below to see how they look |
| show_scrobbles | `show_scrobbles=false` **(default)**<br>`show_scrobbles=true` | **(optional)** Selects whether or not to show scrobble count


<hr>

## Demo

```
https://github-readme-lastfm-stats.netlify.app/.netlify/functions/card?user=st-silver
```

![card/light-noScrobbles.svg](./docs/card/light-noScrobbles.svg)<hr>

```
https://github-readme-lastfm-stats.netlify.app/.netlify/functions/card?user=st-silver&theme=dimmed
```

![card/dimmed-noScrobbles.svg](./docs/card/dimmed-noScrobbles.svg)<hr>

```
https://github-readme-lastfm-stats.netlify.app/.netlify/functions/card?user=st-silver&theme=dark
```

![card/dark-noScrobbles.svg](./docs/card/dark-noScrobbles.svg)<hr>

```
https://github-readme-lastfm-stats.netlify.app/.netlify/functions/card?user=st-silver&show_scrobbles=true
```

![card/light-scrobbles.svg](./docs/card/light-scrobbles.svg)<hr>

```
https://github-readme-lastfm-stats.netlify.app/.netlify/functions/card?user=st-silver&theme=dimmed&show_scrobbles=true
```

![card/dimmed-scrobbles.svg](./docs/card/dimmed-scrobbles.svg)<hr>

```
https://github-readme-lastfm-stats.netlify.app/.netlify/functions/card?user=st-silver&theme=dark&show_scrobbles=true
```

![card/dark-scrobbles.svg](./docs/card/dark-scrobbles.svg)<hr>

## Development & Deployment
Developing and deploying this project requires setting an environment variable that holds your last.fm API key. [See here to find out how to get an API account.](https://www.last.fm/api#getting-started)

```
export LASTFM_STATS_KEY=your_key_here
```

This also needs to be done on your deployment platform of choice. For Netlify, this can be done under *Site Settings > Build & deploy > Environment*.

<hr>

## Issues, Requests, and Contributing

Found an issue or have a feature request? [Create a new one here](https://github.com/rafaelwi/github-readme-lastfm-stats/issues/new) and I will take a look at it ASAP. Please give as much detail as possible in your commment.

Like the project and want to lend a hand? Just make a pull request and I'll take a look.
