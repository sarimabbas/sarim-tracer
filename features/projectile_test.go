package features

import (
	"testing"
)

type projectile struct {
	position Tuple
	velocity Tuple
}

type environment struct {
	gravity Tuple
	wind    Tuple
}

func tick(env environment, proj projectile) projectile {
	position := proj.position.Add(proj.velocity)
	velocity := proj.velocity.Add(env.gravity).Add(env.wind)
	return projectile{position, velocity}
}

func TestProjectileEnvironment(t *testing.T) {
	// projectile starts one unit above the origin
	// velocity is normalized to 1 unit/tick
	p := projectile{PointNew(0, 1, 0), VectorNew(1, 1.8, 0).Normalize().ScalarMultiply(11.25)}
	// gravity -0.1 unit/tick, and wind is -0.01 unit/tick
	e := environment{VectorNew(0, -0.1, 0), VectorNew(-0.01, 0, 0)}
	// canvas
	c := CanvasNew(900, 550)

	for i := 0; i < 100; i++ {
		p = tick(e, p)
		c.SetPixel(int(p.position.x), int(p.position.y), ColorNew(1, 0, 0))
	}

	c.ToPPM("projectile_test.ppm", false, true)
}
