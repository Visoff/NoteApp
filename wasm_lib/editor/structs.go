package editor

type Pos struct {
    X int
    Y int
}

func NewPos(x int, y int) Pos {
    return Pos{X: x, Y: y}
}

type Size struct {
    W int
    H int
}

func NewSize(w int, h int) Size {
    return Size{W: w, H: h}
}

type CursorMode int
const (
    NORMAL CursorMode = iota
    INSERT
    VISUAL
)

type Cursor struct {
    pos Pos
    selectionStart Pos
    mode CursorMode
}

func NewCursor() *Cursor {
    return &Cursor{pos: NewPos(0, 0), selectionStart: NewPos(0, 0)}
}

func (c *Cursor) Pos() Pos {
    return c.pos
}

func (c *Cursor) SelectionStart() Pos {
    if c.mode != VISUAL {
        return c.pos
    }
    return c.selectionStart
}

func (c *Cursor) Mode() CursorMode {
    return c.mode
}

func (c *Cursor) Move(x, y int) {
    c.pos.X += x
    c.pos.Y += y
}

type TextField struct {
    text []string
    cursors []*Cursor
}

func NewTextField() TextField {
    return TextField{
        text: []string{""},
        cursors: []*Cursor{NewCursor()},
    }
}
