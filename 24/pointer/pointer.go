package pointer

type Pointer struct {
	x, y float64
}

func NewPointer(x, y float64) *Pointer {
	return &Pointer{x, y}
}

func (p *Pointer) X() float64 {
	return p.x
}

func (p *Pointer) SetX(val float64) {
	p.x = val
}

func (p *Pointer) SetY(val float64) {
	p.y = val
}

func (p *Pointer) CalcDistance() float64 {
	return p.x - p.y
}