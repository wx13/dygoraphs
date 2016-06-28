package main

import (
	"fmt"
	"github.com/wx13/dygoraphs"
)

func linePlot() {

	fmt.Println("<h2>Line Plot</h2>")

	g := dygoraphs.NewGraph()
	g.SetStyle("width", "500px")
	g.SetStyle("height", "300px")
	data := "time,height,weight\n" +
		"0,100,150\n" +
		"1,102,140\n" +
		"2,104,145\n" +
		"3,101,155\n" +
		"4,110,160\n" +
		"5,105,170\n"
	fmt.Println(g.PlotCSV(data))

}

func symErrBars() {

	fmt.Println("<h2>Symmetric Error Bars</h2>")

	g := dygoraphs.NewGraph()
	g.SetStyle("width", "500px")
	g.SetStyle("height", "300px")
	g.SetOption("errorBars", "true")
	data := "time,height,weight\n" +
		"0,100,10,150,20\n" +
		"1,102,10,140,20\n" +
		"2,104,10,145,20\n" +
		"3,101,10,155,20\n" +
		"4,110,10,160,20\n" +
		"5,105,10,170,20\n"
	fmt.Println(g.PlotCSV(data))

}

func asymErrBars() {

	fmt.Println("<h2>Asymmetric Error Bars</h2>")

	g := dygoraphs.NewGraph()
	g.SetStyle("width", "500px")
	g.SetStyle("height", "300px")
	g.SetOption("customBars", "true")
	data := "time,height,weight\n" +
		"0,95;100;110,140;150;160\n" +
		"1,100;102;112,140;140;160\n" +
		"2,100;104;114,140;145;160\n" +
		"3,100;101;111,140;155;160\n" +
		"4,100;110;129,140;160;160\n" +
		"5,100;105;115,140;170;160\n"
	fmt.Println(g.PlotCSV(data))

}

func errBarTrans() {

	fmt.Println("<h2>Symmetric Error Bars</h2>")

	g := dygoraphs.NewGraph()
	g.SetStyle("width", "500px")
	g.SetStyle("height", "300px")
	g.SetOption("errorBars", "true")
	g.SetOption(`'height'`, "{color: \"blue\", fillAlpha: 0.5}")
	g.SetOption(`'weight'`, "{color: \"blue\", fillAlpha: 0.1}")
	g.SetOption(`'depth'`, "{color: \"blue\", fillAlpha: 0.3}")
	data := "time,height,weight,depth\n" +
		"0,100,10,150,20,80,10\n" +
		"1,102,10,140,20,80,10\n" +
		"2,104,10,145,20,80,10\n" +
		"3,101,10,155,20,80,10\n" +
		"4,110,10,160,20,80,10\n" +
		"5,105,10,170,20,80,10\n"
	fmt.Println(g.PlotCSV(data))

}

func main() {

	fmt.Println("<html><head><title></title>")
	fmt.Println(dygoraphs.Script)
	fmt.Println("</head><body>")

	linePlot()
	symErrBars()
	asymErrBars()
	errBarTrans()

	fmt.Println("</body></html>")

}
