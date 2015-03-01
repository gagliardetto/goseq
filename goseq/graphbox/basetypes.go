// Graphics display model
//

package graphbox

import (
    "github.com/ajstarks/svgo"
)


type GraphboxItem interface {
    // Defines a constraint.  It is provided with the coordinates
    // of the item.
    Constraint(r, c int) Constraint

    // Call to draw this box 
    Draw(ctx DrawContext, point Point)
}

type Constraint interface {
    Apply(cm ConstraintChanger)
}

type ConstraintChanger interface {
    // Calculate the current size between the two grid points
    GridPointRect(fr, fc, tr, tc int) (int, int)

    // Ensure that the left side of this column has this much space.
    // Provide space if needed.
    EnsureLeftIsAtleast(col, newLeft int)

    // Ensure that the top side of this row has this much space.
    // Provide space if needed.
    EnsureTopIsAtLeast(row, newTop int) 

    AddLeftToCol(col, newLeft int)

    AddTopToRow(row, newTop int)
}


// A drawing context
type DrawContext struct {
    Canvas          *svg.SVG
    Graphic         *Graphic
    R, C            int
}

// Returns the outer rectangle of a particular cell
func (dc *DrawContext) PointAt(r, c int) (Point, bool) {
    return dc.Graphic.PointAt(r, c)
}


// An anchor point located in a rectangle at 0, 0 with the w, h passed in
type Gravity         func(w, h int) (int, int)

var NorthWestGravity Gravity = func(w, h int) (int, int) { return 0, 0 }
var EastGravity Gravity = func(w, h int) (int, int) { return w, h / 2 }
var WestGravity Gravity = func(w, h int) (int, int) { return 0, h / 2 }
var CenterGravity Gravity = func(w, h int) (int, int) { return w / 2, h / 2 }
var SouthGravity Gravity = func(w, h int) (int, int) { return w / 2, h }
var SouthWestGravity Gravity = func(w, h int) (int, int) { return 0, h }


// A specific gravity
func AtSpecificGravity(fx, fy float64) Gravity {
    return func(w, h int) (int, int) {
        return int(fx * float64(w)), int(fy * float64(h))
    }
}


// A rectangle
type Rect struct {
    X, Y            int
    W, H            int
}

// Returns a point located at a specific gravity within the rectangle
func (r Rect) PointAt(gravity Gravity) (int, int) {
    lx, ly := gravity(r.W, r.H)
    return r.X + lx, r.Y + ly
}

// Returns a rectangle position at a specific point and a gravity relative
func (r Rect) PositionAt(x, y int, gravity Gravity) Rect {
    lx, ly := gravity(r.W, r.H)
    nx := x - lx
    ny := y - ly
    return Rect{nx, ny, r.W, r.H}
}

// Returns a rectangle blown out by a given size
func (r Rect) BlowOut(dims Point) Rect {
    return Rect{r.X - dims.X, r.Y - dims.Y, r.W + dims.X * 2, r.H + dims.Y * 2}
}



// A point
type Point struct {
    X, Y            int
}

// A box frame
type BoxFrame struct {
    // The outer rectangle.  This encompasses margins, etc.
    OuterRect       Rect

    // The inner rectangle.
    InnerRect       Rect
}