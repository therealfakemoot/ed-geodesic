package geodesic

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	// "log"
	"net/http"
	"net/url"
)

type Patrol struct {
	Type         string   `json:"type"`
	System       string   `json:"system"`
	X            float64  `json:"x"`
	Y            float64  `json:"y"`
	Z            float64  `json:"z"`
	Regionid     int      `json:"regionId"`
	Region       string   `json:"region"`
	BodyCount    int      `json:"body_count"`
	Balance      int      `json:"balance"`
	SignalCount  int      `json:"signal_count"`
	CodexCount   int      `json:"codex_count"`
	Types        []string `json:"types"`
	Guess        []string `json:"guess"`
	URL          string   `json:"url"`
	Instructions string   `json:"instructions"`
}

func CanonnChallenge() []Patrol {
	var p []Patrol

	return p
}

func JsonPatrols() ([]*url.URL, error) {
	var s []*url.URL
	// this is a csv file
	resp, err := http.Get("https://docs.google.com/spreadsheets/d/e/2PACX-1vQsi1Vbfx4Sk2msNYiqo0PVnW3VHSrvvtIRkjT-JvH_oG9fP67TARWX2jIjehFHKLwh4VXdSh0atk3J/pub?gid=0&single=true&output=csv")
	if err != nil {
		return s, err
	}
	enc := csv.NewReader(resp.Body)
	for {
		r, err := enc.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return s, nil
		}

		// this ignores disabled data sources
		if r[1] != "Y" && r[3] == "json" {
			u, err := url.Parse(r[4])
			if err != nil {
				return s, fmt.Errorf("poorly formatted data source URL for %s patrol: %w\n", r[2], err)
			}
			s = append(s, u)
		}
	}
	return s, nil
}

func Patrols(challenge bool) ([]Patrol, error) {
	var patrols []Patrol

	sources, err := JsonPatrols()
	if err != nil {
		return patrols, fmt.Errorf("error fetching json patrol sources: %w\n", err)
	}

	for _, source := range sources {
		resp, err := http.Get(source.String())
		if err != nil {
			return patrols, fmt.Errorf("error fetching json patrol sources: %w\nsource url: %s", err, source.String())
		}

		var raw []Patrol
		enc := json.NewDecoder(resp.Body)
		err = enc.Decode(&raw)
		for _, p := range raw {
			patrols = append(patrols, p)
		}

	}

	return patrols, nil
}
