package comparator

type ModeEnum int

const (
	Year ModeEnum = iota
	Size
	Rate
)

func (m ModeEnum) String() string {
	switch m {
	case Year:
		return "year"
	case Size:
		return "size"
	case Rate:
		return "rate"
	default:
		return "unknown"
	}
}
