package comparator

import "errors"

type ModeEnum int

const (
	ComparatorModeYear ModeEnum = iota
	ComparatorModeSize
	ComparatorModeRate
)

var ErrUnknownModeEnum = errors.New("unknown mode")

func (m ModeEnum) String() string {
	switch m {
	case ComparatorModeYear:
		return "year"
	case ComparatorModeSize:
		return "size"
	case ComparatorModeRate:
		return "rate"
	default:
		return "unknown"
	}
}
