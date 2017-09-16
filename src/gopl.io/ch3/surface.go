package main

import (
	"io"
	"fmt"
	"math"
)

const (
	cells = 100
	xyrange = 30.0
	angle = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func writeSurface(out io.Writer, width, height int) {
	w := float64(width)
	h := float64(height)
	xyscale := w / 2 / xyrange
	zscale := h * 0.4
	
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' " +
		"style='stroke: grey; fill: white; stroke-width: 0.7' " +
		"width='%d' height='%d'>", width, height)
	var r, g, b uint8
	g = 0xEF
    for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az:= corner(i+1, j, w, h, xyscale, zscale)
			bx, by, bz:= corner(i, j, w, h, xyscale, zscale)
			cx, cy, cz:= corner(i, j+1, w, h, xyscale, zscale)
			dx, dy, dz:= corner(i+1, j+1, w, h, xyscale, zscale)
			hz := az + bz + cz + dz
			if hz > 0.0 {
				if hz < 1 / 256 {
					hz *= 256
				}
				r = uint8(0xFF * hz)
			} else {
				hz = math.Abs(hz)
				if hz < 1 / 256 {
					hz *= 256
				}
				b = uint8(0xFF * hz)
			}
			fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g' style='fill: #%02X%02X%02X' />\n",
				ax, ay, bx, by, cx, cy, dx, dy, r, g, b)
		}
	}
	fmt.Fprintf(out, "</svg>")
}

func corner(i, j int, w, h, xyscale, zscale float64) (float64, float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z := f(x, y)
	sx := w/2 + (x-y)*cos30*xyscale
	sy := h/2 + (x+y)*sin30*xyscale - z*zscale
	if math.IsNaN(sx) {
		sx = 0.0
	}
	if math.IsNaN(sy) {
		sy = 0.0
	}
	return sx, sy, z
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	// r := math.Nextafter(x, y)
	return math.Sin(r) / r
}

