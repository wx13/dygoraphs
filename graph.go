package dygoraphs

import (
	"fmt"
	"math/rand"
	"strings"
)

// Graph defines a single graph object.
type Graph struct {
	ID      string
	Style   map[string]string
	Options map[string]string
}

// NewGraph creates a new Graph object.
func NewGraph() *Graph {
	g := Graph{}
	g.ID = fmt.Sprintf("%d", rand.Int())
	g.Style = make(map[string]string)
	g.Options = make(map[string]string)
	return &g
}

// SetOption sets a graph option as a name/value pair.
func (g *Graph) SetOption(name, value string) {
	g.Options[name] = value
}

// SetStyle specifies a CSS style attribute to set.
func (g *Graph) SetStyle(name, value string) {
	g.Style[name] = value
}

// DumpOptions dumps the graph options as JSON.
func (g *Graph) DumpOptions() string {
	optStrs := []string{}
	for name, val := range g.Options {
		s := fmt.Sprintf("%s: %s", name, val)
		optStrs = append(optStrs, s)
	}
	return "{" + strings.Join(optStrs, ", ") + "}"
}

// DumpStyle dumps the graph style options as JSON.
func (g *Graph) DumpStyle() string {
	strs := []string{}
	for name, val := range g.Style {
		s := fmt.Sprintf("%s:%s", name, val)
		strs = append(strs, s)
	}
	return strings.Join(strs, "; ") + ";"
}

// Div creates a graph <div> element to insert into a web page.
func (g *Graph) Div() string {
	return fmt.Sprintf(`<div id="%s" style="%s"></div>`, g.ID, g.DumpStyle())
}

// EscCSV escapes a csv string for inclusion into HTML.
func (g *Graph) EscCSV(csv string) string {
	return fmt.Sprintf("%q", csv)
}

// PlotCSV creates a dygoraphs plot from CSV data.
func (g *Graph) PlotCSV(csv string) string {
	html := g.Div() + "\n"
	html += "<script>\n"
	html += "g = new Dygraph(\n"
	html += fmt.Sprintf(`  document.getElementById("%s"),`, g.ID) + "\n"
	html += g.EscCSV(csv) + ",\n"
	html += g.DumpOptions()
	html += ");\n"
	html += "</script>\n"
	return html
}
