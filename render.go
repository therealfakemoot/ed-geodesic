package geodesic

import (
	"io"
	"text/template"
)

type Report struct {
	Systems map[string]SystemDetails
}

func (r Report) Render(w io.Writer) error {
	t := template.Must(template.New("base").Parse(`Name\tDistance\tScoopable
	{{range .Systems}}
	{{.Name}}\t{{.Distance}}\t{{.Scoopable}}
	{{end}}
	`))

	return t.Execute(w, r)

}

func NewReport(system string, radius int) (Report, error) {
	var r Report
	r.Systems = make(map[string]SystemDetails)
	systems, err := Cube(system, radius)
	if err != nil {
		return r, err
	}
	for _, s := range systems {
		d, err := s.Details()
		if err != nil {
			return r, err
		}
		r.Systems[s.Name] = d
	}

	return r, nil
}
