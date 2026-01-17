package triangle

type Kind int

const (
	NaT Kind = iota
	Equ      // equilateral
	Iso      // isosceles
	Sca      // scalene
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	switch {
	case !(a+b >= c && b+c >= a && a+c >= b) || a+b+c <= 0:
		return NaT
	case a == b && b == c:
		return Equ
	case a == b || a == c || b == c:
		return Iso
	default:
		return Sca
	}

}
