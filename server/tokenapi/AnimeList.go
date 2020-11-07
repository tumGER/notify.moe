package tokenapi

import (
	"errors"

	"github.com/animenotifier/notify.moe/arn"
)

// AnimeUpdate parses animeListEntry.out of the JSON and then tries to integrate them into the database
func AnimeUpdate(request *TokenRequest) error {
	animeListEntry := &arn.AnimeListItem{}
	animeJSON := request.JSON.Get("anime")

	animeListEntry.Status = animeJSON.Get("status").String()
	animeListEntry.AnimeID = animeJSON.Get("id").String()
	animeListEntry.Episodes = int(animeJSON.Get("episode").Int())
	animeListEntry.Notes = animeJSON.Get("notes").String()
	animeListEntry.RewatchCount = int(animeJSON.Get("rewatchCount").Int())
	animeListEntry.Private = animeJSON.Get("isPrivate").Bool()

	if animeListEntry.AnimeID == "" {
		return errors.New("No anime ID has been supplied")
	}

	rating := animeJSON.Get("ratings")
	animeListEntry.Rating.Overall = float64(rating.Get("overall").Float())
	animeListEntry.Rating.Story = float64(rating.Get("story").Float())
	animeListEntry.Rating.Visuals = float64(rating.Get("visuals").Float())
	animeListEntry.Rating.Soundtrack = float64(rating.Get("soundtrack").Float())

	animeList := request.User.AnimeList()
	animeList.Import(animeListEntry)

	return nil
}
