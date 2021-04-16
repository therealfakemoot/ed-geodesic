package geodesic

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
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

// func Graph(s string)

func Cube(system string, size int) []System {
	var systems []System
	q := url.Values{"systemName": []string{system}, "size": []string{fmt.Sprintf("%d", size)}}

	var edsm, _ = url.Parse("https://www.edsm.net/api-v1/cube-systems")
	edsm.RawQuery = q.Encode()
	resp, err := http.Get(edsm.String())
	if err != nil {
		log.Fatalf("unable to request system cube from EDSM: %s", err)
		return systems
	}

	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&systems)
	if err != nil {
		log.Fatalf("unable to decode system cube from EDSM: %s", err)
	}

	return systems
}
