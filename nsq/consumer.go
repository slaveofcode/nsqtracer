package nsq

import (
	nsqio "github.com/nsqio/go-nsq"
	"github.com/slaveofcode/nsqtracer/pansi"
)

// NewConsumer will return new NSQ consumer instance
func NewConsumer(topic, channel string) *nsqio.Consumer {
	defaultCfg := nsqio.NewConfig()
	consumer, err := nsqio.NewConsumer(topic, channel, defaultCfg)

	if err != nil {
		pansi.N(
			pansi.SBlue("➜"),
			" ",
			pansi.SYellow("Error creating new consumer: "),
			pansi.SRedBg(pansi.ForeWhite, err.Error()),
		)
	}

	return consumer
}
