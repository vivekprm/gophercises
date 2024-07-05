package main

import (
	"os"

	svg "github.com/ajstarks/svgo"
)

func main() {
	f, err := os.Create("demo.svg")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	canvas := svg.New(f)
	data := []struct {
		Month string
		Usage int
	}{
		{Month: "Jan", Usage: 171},
		{Month: "Feb", Usage: 180},
		{Month: "Mar", Usage: 100},
		{Month: "Apr", Usage: 87},
		{Month: "May", Usage: 66},
		{Month: "June", Usage: 40},
		{Month: "July", Usage: 32},
		{Month: "Aug", Usage: 55},
		{Month: "Sept", Usage: 0},
		{Month: "Oct", Usage: 0},
		{Month: "Nov", Usage: 0},
		{Month: "Dec", Usage: 0},
	}
	width := len(data)*60 + 10
	height := 300
	threshold := 160
	max := 0
	for _, item := range data {
		if item.Usage > max {
			max = item.Usage
		}
	}
	canvas.Start(width, height)
	canvas.Rect(0, 0, width, height)
	for i, val := range data {
		percent := val.Usage * (height - 50) / max
		canvas.Rect(i*60+10, (height-50)-percent, 50, percent, "fill:rgb(77,200,232)")
		canvas.Text(i*60+35, height-20, val.Month, "font-size:14pt;fill:rgb(150,150,150);text-anchor:middle")
	}
	thresPercent := threshold * (height - 50) / max
	canvas.Line(0, height-thresPercent, width, height-thresPercent, "stroke: rgb(250, 150, 150); opacity: 0.8; stroke-width:2px")
	canvas.Rect(0, 0, width, height-thresPercent, "fill:rgb(255, 100, 100); opacity:0.1")
	canvas.Line(0, height-50, width, height-50, "stroke: rgb(150,150,150);stroke-width:2px")
	canvas.End()
}
