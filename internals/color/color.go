package color

import (
	"fmt"
	"io"
	"ray-tracer/internals/vec3"
)

type Color = vec3.Vec3

func WriteColor(out io.Writer, pc Color) error {
	r := pc.X
	g := pc.Y
	b := pc.Z

	ir := int(255.999 * r)
	ig := int(255.999 * g)
	ib := int(255.999 * b)

	_, err := fmt.Fprintf(out, "%d %d %d\n", ir, ig, ib)
	return err
}
