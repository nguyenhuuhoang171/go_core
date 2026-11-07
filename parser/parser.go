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
	return parseList(raw, func(item string) (uint64, error) {
		return strconv.ParseUint(item, 10, 64)
	})
}

// ParseInts is ParseUints for signed ints: "1,2,3" becomes [1 2 3].
func ParseInts(raw string) ([]int, error) {
	return parseList(raw, strconv.Atoi)
}

func parseList[T any](raw string, parse func(string) (T, error)) ([]T, error) {
	if strings.TrimSpace(raw) == "" {
		return nil, nil
	}

	parts := strings.Split(raw, ",")
	list := make([]T, 0, len(parts))
	for _, part := range parts {
		n, err := parse(strings.TrimSpace(part))
		if err != nil {
			return nil, fmt.Errorf("invalid number %q in list %q", part, raw)
		}
		list = append(list, n)
	}
	return list, nil
}
