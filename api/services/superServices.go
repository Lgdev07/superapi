package services

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"

	"github.com/Lgdev07/superapi/api/models"
	"github.com/tidwall/gjson"
)

func ApiGetSuperByName(name string) (*models.Super, int, error) {
	apiToken := os.Getenv("SUPERHERO_API_TOKEN")
	url := fmt.Sprintf("https://superheroapi.com/api/%s/search/%s", apiToken, name)

	resp, err := http.Get(url)
	if err != nil {
		return &models.Super{}, http.StatusInternalServerError, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &models.Super{}, http.StatusInternalServerError, err
	}

	path := fmt.Sprintf(`results.#(name==%v)`, name)
	results := gjson.Get(string(body), path)

	if results.String() == "" {
		return &models.Super{}, http.StatusNotFound, err
	}

	relatives := results.Get("connections.relatives").String()

	re := regexp.MustCompile(`[),]+|[);]`)
	relativesSplited := re.FindAllString(relatives, -1)

	super := &models.Super{
		Name:            name,
		FullName:        results.Get("biography.full-name").String(),
		Intelligence:    int(results.Get("powerstats.intelligence").Int()),
		Power:           int(results.Get("powerstats.power").Int()),
		Occupation:      results.Get("work.occupation").String(),
		Image:           results.Get("image.url").String(),
		Alignment:       results.Get("biography.alignment").String(),
		Groups:          results.Get("connections.group-affiliation").String(),
		NumberOfParents: len(relativesSplited),
	}

	return super, 0, nil
}
