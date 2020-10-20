package nsq

import (
	"github.com/slaveofcode/nsqtracer/pansi"
)

// DefaultLogger just noob logger
type DefaultLogger struct {
	Name string
	Kind string
}

// Output will executed on logging
func (l *DefaultLogger) Output(calldepth int, s string) error {
	if l.Name == "" {
		l.Name = "Default"
	}

	pansi.N(
		pansi.SGray("#"),
		pansi.SGreen("["+l.Name+"]"),
		pansi.SYellow("["+l.Kind+"]"),
		" ",
		s,
	)

	return nil
}
