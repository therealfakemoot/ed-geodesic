package geodesic

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"sort"
)

type Coord struct {
	X int `json:"x"`
	Y int `json:"y"`
	Z int `json:"z"`
}

type System struct {
	Name      string  `json:"name"`
	ID        int     `json:"id"`
	Coords    Coord   `json:"coords"`
	BodyCount int     `json:"bodyCount"`
	Distance  float64 `json:"distance"`
}

func (s System) String() string {
	return fmt.Sprintf("[ %s -> %.2fly | Bodies(%d) ]", s.Name, s.Distance, s.BodyCount)
}

func (s System) Details() (SystemDetails, error) {
	var sd SystemDetails
	u, _ := url.Parse("https://www.edsm.net/api-v1/system")
	q := url.Values{}
	q.Set("systemName", s.Name)
	q.Set("showInformation", "1")
	q.Set("showPermit", "1")

	u.RawQuery = q.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		return sd, fmt.Errorf("unable to generate get system details from EDSM: %s", err)
	}
	enc := json.NewDecoder(resp.Body)
	err = enc.Decode(&sd)
	if err != nil {
		return sd, fmt.Errorf("unable to decode system details: %s", err)
	}
	return sd, err
}

type SystemInformation struct {
	Allegiance    string `json:"allegiance"`
	Government    string `json:"government"`
	Faction       string `json:"faction"`
	Factionstate  string `json:"factionState"`
	Population    int    `json:"population"`
	Security      string `json:"security"`
	Economy       string `json:"economy"`
	Secondeconomy string `json:"secondEconomy"`
	Reserve       string `json:"reserve"`
}

type Primarystar struct {
	Type        string `json:"type"`
	Name        string `json:"name"`
	Isscoopable bool   `json:"isScoopable"`
}

type SystemDetails struct {
	Name          string            `json:"name"`
	Requirepermit bool              `json:"requirePermit"`
	Information   SystemInformation `json:"information"`
	PrimaryStar   Primarystar       `json:"primaryStar"`
}

func Cube(system string, size int) ([]System, error) {
	var systems []System
	q := url.Values{}
	q.Set("systemName", system)
	q.Set("size", fmt.Sprintf("%d", size))

	var edsm, _ = url.Parse("https://www.edsm.net/api-v1/cube-systems")
	edsm.RawQuery = q.Encode()
	resp, err := http.Get(edsm.String())
	if err != nil {
		return systems, fmt.Errorf("unable to request system cube from EDSM: %s", err)

	}

	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&systems)
	if err != nil {
		return systems, fmt.Errorf("unable to decode system cube from EDSM: %s", err)
	}

	sort.Slice(systems, func(i, j int) bool {
		return systems[i].Distance < systems[j].Distance
	})
	return systems, nil
}
