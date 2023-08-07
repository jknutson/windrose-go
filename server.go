package main

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"
)

func windrose(w http.ResponseWriter, req *http.Request) {
	// angle := req.URL.Query()["angle"][0]
	params := req.URL.Query()
	angle := params.Get("angle")
	direction := params.Get("direction")
	if direction != "" {
		directions := [8]string{
			"N", "NE",
			"E", "SE",
			"S", "SW",
			"W", "NW"}
		multiplier := 360 / len(directions)
		for i := 0; i < len(directions); i++ {
			if direction == directions[i] {
				angle = strconv.Itoa(int(i) * multiplier)
			}
		}
	}
	if angle == "" {
		angle = "0" // set a default
	}
	angleDeg, err := strconv.ParseFloat(angle, 64)
	if err != nil {
		angle = "0"
		err = nil
		// TODO: handle errors better
		// panic(err)
	}
	svgWindroseBuf := &bytes.Buffer{}
	err = GenWindrose(angleDeg, svgWindroseBuf)
	if err != nil {
		// TODO: what here?
		panic(err)
	}
	w.Header().Set("Content-Type", "image/svg+xml")
	// w.Header().Set("Windrose-Angle-Deg", string(angleDeg))
	fmt.Fprintf(w, svgWindroseBuf.String())
}

func main() {
	http.HandleFunc("/windrose", windrose)

	http.ListenAndServe(":8080", nil)
}
