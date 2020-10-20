package flags

import (
	"strings"
)

// ArrFlags is type used to store multi-flag CLI values as string slices
type ArrFlags []string

func (i *ArrFlags) String() string {
	return strings.Join(*i, ",")
}

// Set will append CLI flag value to an Array
func (i *ArrFlags) Set(val string) error {
	*i = append(*i, val)
	return nil
}
