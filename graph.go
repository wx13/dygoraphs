package dygoraphs

import (
	"fmt"
	"math/rand"
	"strings"
)

type Graph struct {
	Id      string
	Style   map[string]string
	Options map[string]string
}

func NewGraph() *Graph {
	g := Graph{}
	g.Id = fmt.Sprintf("%d", rand.Int())
	g.Style = make(map[string]string)
	g.Options = make(map[string]string)
	return &g
}

func (g *Graph) SetOption(name, value string) {
	g.Options[name] = value
}

func (g *Graph) SetStyle(name, value string) {
	g.Style[name] = value
}

func (g *Graph) DumpOptions() string {
	optStrs := []string{}
	for name, val := range g.Options {
		s := fmt.Sprintf("%s: %s", name, val)
		optStrs = append(optStrs, s)
	}
	return "{" + strings.Join(optStrs, ", ") + "}"
}

func (g *Graph) DumpStyle() string {
	strs := []string{}
	for name, val := range g.Style {
		s := fmt.Sprintf("%s:%s", name, val)
		strs = append(strs, s)
	}
	return strings.Join(strs, "; ")
}

//func (g *Graph) CSVdata() string {
//	csv := `"` + strings.Join(g.Columns, ",") + `\n"` + " +\n"
//	for row := range g.Data[0] {
//		csv += `"`
//		for col := range g.Data {
//			csv += fmt.Sprintf("%f", g.Data[col][row])
//			if col < len(g.Data)-1 {
//				csv += ","
//			}
//		}
//		csv += `\n"`
//		if row < len(g.Data[0])-1 {
//			csv += " +\n"
//		}
//	}
//	return csv
//}

func (g *Graph) Div() string {
	return fmt.Sprintf(`<div id="%s" style="%s"></div>`, g.Id, g.DumpStyle())
}

func (g *Graph) EscCSV(csv string) string {
	return fmt.Sprintf("%q", csv)
}

func (g *Graph) PlotCSV(csv string) string {
	html := g.Div() + "\n"
	html += "<script>\n"
	html += "g = new Dygraph(\n"
	html += fmt.Sprintf(`  document.getElementById("%s"),`, g.Id) + "\n"
	html += g.EscCSV(csv) + ",\n"
	html += g.DumpOptions()
	html += ");\n"
	html += "</script>\n"
	return html
}
