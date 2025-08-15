package main

import (
	"fmt"
	"os"
	"ray-tracer/internals/color"
	"ray-tracer/internals/ray"
	"ray-tracer/internals/vec3"
)

func rayColor(r ray.Ray) color.Color {
	unitDirection := r.Direction.Normalize()
	a := 0.5*unitDirection.Y + 1.0
	return color.Color{
		X: (1.0-a)*1.0 + a*0.5,
		Y: (1.0-a)*1.0 + a*0.7,
		Z: (1.0-a)*1.0 + a*1.0,
	}
}

func main() {

	// Image
	aspectRatio := 16.0 / 9.0
	imageWidth := 400

	imageHeight := int(float64(imageWidth) / aspectRatio)
	if imageHeight < 1 {
		imageHeight = 1
	}

	// Camera
	focalLength := 1.0
	viewportHeight := 2.0
	viewportWidth := viewportHeight * (float64(imageWidth) / float64(imageHeight))
	cameraCenter := vec3.Vec3{X: 0, Y: 0, Z: 0}

	// Calculating vectors
	viewportU := vec3.Vec3{X: viewportWidth, Y: 0, Z: 0}
	viewportV := vec3.Vec3{X: 0, Y: -viewportHeight, Z: 0}

	// Deltas from pixel to pixel
	pixelDeltaU := viewportU.Div(float64(imageWidth))
	pixelDeltaV := viewportV.Div(float64(imageHeight))

	// Calculate location of upper left pixel
	viewportUpperLeft := cameraCenter.
		Sub(vec3.Vec3{X: 0, Y: 0, Z: focalLength}).
		Sub(viewportU.Div(2)).
		Sub(viewportV.Div(2))

	pixel00Loc := viewportUpperLeft.Add(
		pixelDeltaU.Add(pixelDeltaV).Scale(0.5),
	)

	fmt.Printf("P3\n%v %v \n255\n", imageWidth, imageHeight)

	for i := range imageHeight {
		fmt.Fprintf(os.Stderr, "\nScanlines remaining: %v", imageHeight-1-i)
		for j := range imageWidth {
			pixelCenter := pixel00Loc.
				Add(pixelDeltaU.Scale(float64(i))).
				Add(pixelDeltaV.Scale(float64(j)))
			rayDirection := pixelCenter.Sub(cameraCenter)
			r := ray.NewRay(cameraCenter, rayDirection)
			pc := rayColor(r)
			color.WriteColor(os.Stdout, pc)
		}
	}
	fmt.Fprintf(os.Stderr, "\rDone\n")
}
