package main

import (
	"fmt"
	"os"
	"ray-tracer/internals/color"
)

func main() {
	var aspect_ratio = 16.0 / 9.0
	var image_width int = 400

	var image_height int = int(image_width / int(aspect_ratio))
	if image_height < 1 {
		image_height = 1
	}

	// Camera
	// Calculating vectors

	fmt.Printf("P3\n%v %v \n255\n", image_width, image_height)

	for i := range image_height {
		fmt.Fprintf(os.Stderr, "\nScanlines remaining: %v", image_height-1-i)
		for j := range image_width {

			r := float64(i) / float64(image_width-1)
			g := float64(j) / float64(image_height-1)
			b := 0.0
			pc := color.Color{X: r, Y: g, Z: b}
			color.WriteColor(os.Stdout, pc)
		}
	}
	fmt.Fprintf(os.Stderr, "\rDone\n")
}
