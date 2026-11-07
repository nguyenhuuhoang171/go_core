package parser

import (
	"fmt"
	"strconv"
	"strings"
)

// ParseUints turns a comma-separated string such as "2,3,4" into [2 3 4].
// Spaces around an item are ignored. An empty (or blank) string gives a nil
// slice and no error; an item that isn't an unsigned integer gives an error.
func ParseUints(raw string) ([]uint64, error) {
	if strings.TrimSpace(raw) == "" {
		return nil, nil
	}

	parts := strings.Split(raw, ",")
	list := make([]uint64, 0, len(parts))
	for _, part := range parts {
		n, err := strconv.ParseUint(strings.TrimSpace(part), 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid number %q in list %q", part, raw)
		}
		list = append(list, n)
	}
	return list, nil
}
