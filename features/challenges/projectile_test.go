package challenges

import (
	"sarim-tracer/features/canvas"
	"sarim-tracer/features/tuples"
	"testing"
)

type projectile struct {
	position tuples.Tuple
	velocity tuples.Tuple
}

type environment struct {
	gravity tuples.Tuple
	wind    tuples.Tuple
}

func tick(env environment, proj projectile) projectile {
	position := proj.position.Add(proj.velocity)
	velocity := proj.velocity.Add(env.gravity).Add(env.wind)
	return projectile{position, velocity}
}

func TestProjectileEnvironment(t *testing.T) {
	// projectile starts one unit above the origin
	// velocity is normalized to 1 unit/tick
	p := projectile{tuples.PointNew(0, 1, 0), tuples.VectorNew(1, 1.8, 0).Normalize().ScalarMultiply(11.25)}
	// gravity -0.1 unit/tick, and wind is -0.01 unit/tick
	e := environment{tuples.VectorNew(0, -0.1, 0), tuples.VectorNew(-0.01, 0, 0)}
	// canvas
	c := canvas.CanvasNew(900, 550)

	for i := 0; i < 100; i++ {
		p = tick(e, p)
		c.SetPixel(int(p.position.X), int(p.position.Y), tuples.ColorNew(1, 0, 0))
	}

	c.ToPPM("projectile_test.ppm", false, true)
}
