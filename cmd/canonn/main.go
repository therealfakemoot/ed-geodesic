package main

import (
	"log"

	"github.com/therealfakemoot/ed-geodesic"
)

func main() {
	p, err := geodesic.Patrols(false)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Cannon Patrol POI count: %d", len(p))
}
