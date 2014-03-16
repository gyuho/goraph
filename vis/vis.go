// Package vis provides graph visualizing tools
package vis

import (
	"github.com/gyuho/dgo/twd"
)

type Node struct {
	Canvas    *twd.Canvas
	Position  *twd.Vector
	Neighbors []*Node
	Ch        chan *Node
	Power     uint8
}
