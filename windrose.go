package main

import (
	"bytes"
	"fmt"
	"math"
	"os"
	"text/template"
)

// Coordinates represents coordinates on a cartesian plane
type Coordinates struct {
	X, Y int
}

// Arrow represents an SVG line object with marker
type Arrow struct {
	Start, Finish Coordinates
	Width         float64
}

// GenWindrose takes an input angle and fills a byte buffer with SVG data
func GenWindrose(angleDeg float64, svgWindroseBuf *bytes.Buffer) error {
	svgBaseTmplData, err := os.ReadFile("windrose_base.svg.tmpl")
	if err != nil {
		return err
	}
	svgArrowTmplData, err := os.ReadFile("windrose_arrow.svg.tmpl")
	if err != nil {
		return err
	}
	svgArrowTmpl, err := template.New("arrow").Parse(string(svgArrowTmplData))
	if err != nil {
		return err
	}
	svgBaseTmpl, err := template.New("arrow").Parse(string(svgBaseTmplData))
	if err != nil {
		return err
	}

	var (
		arrow         Arrow
		start, finish Coordinates
	)
	// starting point is static
	start.X = 307 // not sure why this is offset by 7 and Y is not
	start.Y = 300

	arrowLength := 200 // 200 = 2/3 of 1 quadrant of canvas
	angleRad := (angleDeg - 90) * (math.Pi / 180)
	fmt.Fprintf(os.Stderr, "angle: %f degrees, %f radians\n", angleDeg, angleRad)
	finish.X = start.X + int(math.Cos(angleRad)*float64(arrowLength))
	finish.Y = start.Y + int(math.Sin(angleRad)*float64(arrowLength))
	arrow.Start = start
	arrow.Finish = finish
	arrow.Width = 3.5

	svgArrowBuf := &bytes.Buffer{}
	err = svgArrowTmpl.Execute(svgArrowBuf, arrow)
	if err != nil {
		return err
	}
	err = svgBaseTmpl.Execute(svgWindroseBuf, svgArrowBuf.String())
	if err != nil {
		return err
	}
	return nil
}

/*
func main() {
	var (
		angleDeg float64
		err      error
	)
	if len(os.Args) > 1 {
		angleDeg, err = strconv.ParseFloat(os.Args[1], 64)
		if err != nil {
			panic(err)
		}
	}

	svgWindroseBuf := &bytes.Buffer{}
	err = GenWindrose(angleDeg, svgWindroseBuf)
	// err = svgBaseTmpl.Execute(os.Stdout, svgArrowBuf.String())
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", svgWindroseBuf.String())
}
*/
