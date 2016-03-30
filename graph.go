package dygoraphs

import (
	"fmt"
	"math/rand"
	"strings"
)

type Graph struct {
	Id      string
	Columns []string
	Data    [][]float64
	Style   string
}

func NewGraph() *Graph {
	g := Graph{}
	g.Id = fmt.Sprintf("%d", rand.Int())
	g.Columns = []string{}
	g.Data = [][]float64{}
	g.Style = ""
	return &g
}

func (g *Graph) Add(columns []string, data [][]float64) {
	g.Columns = append(g.Columns, columns...)
	g.Data = append(g.Data, data...)
}

func (g *Graph) CSVdata() string {
	csv := `"` + strings.Join(g.Columns, ",") + `\n"` + " +\n"
	for row := range g.Data[0] {
		csv += `"`
		for col := range g.Data {
			csv += fmt.Sprintf("%f", g.Data[col][row])
			if col < len(g.Data)-1 {
				csv += ","
			}
		}
		csv += `\n"`
		if row < len(g.Data[0])-1 {
			csv += " +\n"
		}
	}
	return csv
}

func (g *Graph) Div() string {
	return fmt.Sprintf(`<div id="%s" style="%s"></div>`, g.Id, g.Style)
}

func (g *Graph) Plot() string {
	csv := g.CSVdata()
	html := g.Div() + "\n"
	html += "<script>\n"
	html += "g = new Dygraph(\n"
	html += fmt.Sprintf(`  document.getElementById("%s"),`, g.Id) + "\n"
	html += csv
	html += ");\n"
	html += "</script>\n"
	return html
}
