package flags

import (
	"flag"
)

// Array will set CLI flag as an Array and return the results as string slices
func Array(cmdFlag, description string) func() []string {
	var arrFlags ArrFlags
	flag.Var(&arrFlags, cmdFlag, description)

	f := func(arr *ArrFlags) func() []string {
		return func() []string {
			var values []string
			if len(*arr) > 0 {
				for _, val := range *arr {
					if val != "" {
						values = append(values, val)
					}
				}
			}

			return values
		}
	}(&arrFlags)

	return f
}
