package main

import (
	"fmt"
	"os"
)

func main() {
	var image_width int = 256
	var image_height int = 256

	fmt.Printf("P3\n%v %v \n255\n", image_width, image_height)

	for i := range image_height {
		fmt.Fprintf(os.Stderr, "\nScanlines remaining: %v", image_height-1-i)
		for j := range image_width {
			r := float64(i) / float64(image_width-1)
			g := float64(j) / float64(image_height-1)
			b := 0.0

			ir := int(255.999 * r)
			ig := int(255.999 * g)
			ib := int(255.999 * b)

			fmt.Printf("%v %v %v\n", ir, ig, ib)
		}
	}
	fmt.Fprintf(os.Stderr, "\rDone\n")
}
