package parser

import (
	"fmt"
	"strconv"
	"strings"
)

// Unsigned is any unsigned integer type.
type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

// ParseUints turns a comma-separated string such as "2,3,4" into [2 3 4] of
// type T, e.g. ParseUints[uint8]("1,2"). Spaces around an item are ignored. An
// empty (or blank) string gives a nil slice and no error; an item that isn't
// an unsigned integer, or doesn't fit in T, gives an error.
func ParseUints[T Unsigned](raw string) ([]T, error) {
	if strings.TrimSpace(raw) == "" {
		return nil, nil
	}

	parts := strings.Split(raw, ",")
	list := make([]T, 0, len(parts))
	for _, part := range parts {
		n, err := strconv.ParseUint(strings.TrimSpace(part), 10, 64)
		value := T(n)
		if err != nil || uint64(value) != n {
			return nil, fmt.Errorf("invalid number %q in list %q", part, raw)
		}
		list = append(list, value)
	}
	return list, nil
}
