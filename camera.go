package ebitencam

import "github.com/hajimehoshi/ebiten/v2"

// Cam is a camera, it's X and Y position are it's position relative to the world.
type Cam struct {
	X, Y float64
}

// Cam constructor.
func NewCam(x, y float64) *Cam {
	return &Cam{x, y}
}

// Set's the camera position.
func (c *Cam) Set(x, y float64) {
	c.X = x
	c.Y = y
}

// Translates the given coordinates.
func (c *Cam) Apply(x, y float64) (float64, float64) {
	return x - c.X, y - c.Y
}

// Follow's the given coordinates smoothly with the factor delta.
func (c *Cam) Follow(x, y, delta float64) {
	c.X += (x - c.X) * delta
	c.Y += (y - c.Y) * delta
}

// Like Follow but it follow's ahead of the given coordinates, vx and vy are the
// velocity of the object, maxDist is the maximum distance in pixels that the camera
// is allowed to reach from the object.
func (c *Cam) FollowAhead(x, y, vx, vy, delta, maxDist float64) {
	tx := x + (vx * maxDist)
	ty := y + (vy * maxDist)

	c.X += (tx - c.X) * delta
	c.Y += (ty - c.Y) * delta
}

// Like Follow but it adds an offset based of the cursor's position on screen,
// sw and sh are the screen's dimensions.
func (c *Cam) FollowCursor(x, y, delta, maxDist, sw, sh float64) {
	cx, cy := ebiten.CursorPosition()

	tx := float64(cx) - sw/2
	ty := float64(cy) - sh/2

	if tx > maxDist {
		tx = maxDist
	}
	if tx < maxDist {
		tx = -maxDist
	}
	if ty > maxDist {
		ty = maxDist
	}
	if ty < maxDist {
		ty = -maxDist
	}

	c.X += (x + tx - c.X) * delta
	c.Y += (y + ty - c.Y) * delta
}
