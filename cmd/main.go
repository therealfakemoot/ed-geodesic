package main

import (
	"flag"
	"fmt"
	"sort"

	"github.com/therealfakemoot/ed-geodesic"
)

func main() {
	var (
		system string
		radius int
	)
	flag.StringVar(&system, "system", "", "e.g. Ceos, Syneuefe NL-N C24-4")
	flag.IntVar(&radius, "radius", 25, "radius of the cube in lightyears. maximum: 200ly")

	flag.Parse()

	systems := geodesic.Cube(system, radius)
	sort.Slice(systems, func(i, j int) bool {
		return systems[i].Distance < systems[j].Distance
	})
	fmt.Printf("found %d systems", len(systems))
	fmt.Println(systems[0])

}
