package main

import (
	"flag"
	"log"
	"os"

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

	r, err := geodesic.NewReport(system, radius)
	if err != nil {
		log.Fatal("error generating report: %s", err)
	}
	err = r.Render(os.Stdout)
	if err != nil {
		log.Fatal("error rendering report: %s", err)
	}

}
