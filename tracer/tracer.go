package tracer

import (
	"fmt"
	"time"

	nsqio "github.com/nsqio/go-nsq"
	"github.com/slaveofcode/nsqtracer/nsq"
	"github.com/slaveofcode/nsqtracer/pansi"
)

type TracerConfig struct {
	Topic          string
	Channel        string
	NSQdAddrs      []string
	NSQLookupAddrs []string
	Concurrency    int
	MaxInFlighht   int
	EnableDebug    bool
	EnableInfo     bool
	EnableWarn     bool
	AutoFinish     bool
}

type Tracer struct {
	topic          string
	channel        string
	consumer       *nsqio.Consumer
	nsqdAddrs      []string
	nsqLookupAddrs []string
	concurrency    int
	maxInFlight    int
	debug          bool
	info           bool
	warn           bool
	autoFinish     bool
}

const defaultLocalNSQD = "localhost:4150"

func NewDefaultTracer(cfg *TracerConfig) *Tracer {
	c := nsq.NewConsumer(cfg.Topic, cfg.Channel)

	concur := 5
	if cfg.Concurrency > 0 {
		concur = cfg.Concurrency
	}

	maxInf := 10
	if cfg.MaxInFlighht > 0 {
		maxInf = cfg.MaxInFlighht
	}

	return &Tracer{
		topic:          cfg.Topic,
		channel:        cfg.Channel,
		consumer:       c,
		nsqdAddrs:      cfg.NSQdAddrs,
		nsqLookupAddrs: cfg.NSQLookupAddrs,
		concurrency:    concur,
		maxInFlight:    maxInf,
		debug:          cfg.EnableDebug,
		warn:           cfg.EnableWarn,
		info:           cfg.EnableInfo,
		autoFinish:     cfg.AutoFinish,
	}
}

func (t *Tracer) AddHandler(handler nsqio.Handler) {
	t.consumer.AddConcurrentHandlers(handler, t.concurrency)
}

func (t *Tracer) SetConcurrency(v int) {
	t.concurrency = v
}

func (t *Tracer) SetMaxInFlight(v int) {
	t.maxInFlight = v
}

func (t *Tracer) connect() error {
	// prioritize nsqlookups first
	if len(t.nsqLookupAddrs) > 0 {
		err := t.consumer.ConnectToNSQLookupds(t.nsqLookupAddrs)
		if err != nil {
			pansi.PrintErr("Error to connect to NSQ Lookups: ", err)
			return err
		}
		return nil
	}

	if len(t.nsqdAddrs) > 0 {
		err := t.consumer.ConnectToNSQDs(t.nsqdAddrs)
		if err != nil {
			pansi.PrintErr("Error to connect to NSQds: ", err)
			return err
		}
		return nil
	}

	// connect to default address nsq on local
	err := t.consumer.ConnectToNSQD(defaultLocalNSQD)
	if err != nil {
		pansi.PrintErr("Error to connect to NSQd: ", err)
		pansi.PrintWarn("You should set at least one \"nsqd-tcp\" or \"nsqlookup-http\" address to make it work!")
		return err
	}

	return nil
}

func (t *Tracer) setLoggers() {
	if t.debug {
		pansi.PrintIcon(
			pansi.SCyan("["+t.topic+"]"),
			" ",
			pansi.SWhite("Debug Msg Enabled"),
		)
		t.consumer.SetLogger(&nsq.DefaultLogger{
			Name: t.topic,
			Kind: "Debug",
		}, nsqio.LogLevelDebug)
	}

	if t.info {
		pansi.PrintIcon(
			pansi.SCyan("["+t.topic+"]"),
			" ",
			pansi.SWhite("Info Msg Enabled"),
		)
		t.consumer.SetLogger(&nsq.DefaultLogger{
			Name: t.topic,
			Kind: "Info",
		}, nsqio.LogLevelInfo)
	}

	if t.warn {
		pansi.PrintIcon(
			pansi.SCyan("["+t.topic+"]"),
			" ",
			pansi.SWhite("Warning Msg Enabled"),
		)
		t.consumer.SetLogger(&nsq.DefaultLogger{
			Name: t.topic,
			Kind: "Warning",
		}, nsqio.LogLevelWarning)
	}

	t.consumer.SetLogger(&nsq.DefaultLogger{
		Name: t.topic,
		Kind: "Error",
	}, nsqio.LogLevelError)
}

func (t *Tracer) StartTrace() error {
	pansi.PrintInfo("Spawning tracer for " + pansi.SYellow("\""+t.topic+"\""))
	t.AddHandler(&Handler{
		TraceName:    t.topic,
		IsAutoFinish: t.autoFinish,
	})
	t.setLoggers()

	fmt.Println()

	time.Sleep(time.Second)

	if err := t.connect(); err != nil {
		return err
	}

	return nil
}

func (t *Tracer) StopTrace() {
	t.consumer.Stop()
}
