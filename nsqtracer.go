package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/slaveofcode/nsqtracer/flags"
	"github.com/slaveofcode/nsqtracer/pansi"
	"github.com/slaveofcode/nsqtracer/tracer"
)

var channelName = "NSQTracer"
var topicNames []string
var nsqLookups []string
var nsqds []string

func main() {
	isAutoFinish := flag.Bool("auto-finish", false, "Auto finish message after traced, by default \"false\"")
	enableDebug := flag.Bool("debug-log", false, "Enabling debug log message, by default \"false\"")
	enableWarn := flag.Bool("warning-log", false, "Enabling warning log message, by default \"false\"")
	enableInfo := flag.Bool("info-log", false, "Enabling info log message, by default \"false\"")
	channelName = *flag.String("channel-name", "NSQTracer", "Set custom name for this NSQTracer channel subscription, by default \"NSQTracer\"")
	ftopicNames := flags.Array("topic", "Topic name to trace, can be used multiple times")
	fnsqds := flags.Array("nsqd-tcp", "NSQd TCP address (e.g. localhost:4150), can be used multiple times")
	fnsqLookups := flags.Array("nsqlookup-http", "NSQ Lookup HTTP address (e.g. localhost:4161), can be used multiple times")
	flag.Parse()

	topicNames = ftopicNames()
	nsqds = fnsqds()
	nsqLookups = fnsqLookups()

	if len(topicNames) == 0 {
		pansi.PrintWarn("No topic name is provided, please set by adding \"--topic\" option")
	}

	pansi.PrintInfo("Will tracing for topic(s):")
	for _, t := range topicNames {
		pansi.N(
			"  ",
			pansi.SBlue("➜"),
			" ",
			pansi.SYellow(t),
		)
	}

	fmt.Println()

	var tracers []*tracer.Tracer
	for _, t := range topicNames {
		t := tracer.NewDefaultTracer(&tracer.TracerConfig{
			Topic:          t,
			Channel:        channelName,
			NSQdAddrs:      nsqds,
			NSQLookupAddrs: nsqLookups,
			EnableDebug:    *enableDebug,
			EnableWarn:     *enableWarn,
			EnableInfo:     *enableInfo,
			AutoFinish:     *isAutoFinish,
		})

		if err := t.StartTrace(); err != nil {
			os.Exit(0)
		}

		tracers = append(tracers, t)
	}

	fmt.Println()

	shutdown := make(chan os.Signal, 2)
	signal.Notify(shutdown, syscall.SIGINT)
	for {
		select {
		case <-shutdown:
			fmt.Println()
			pansi.PrintWarn("Stopping tracers...")

			for _, t := range tracers {
				t.StopTrace()
			}

			pansi.PrintWarn("Shutting down...")
			return
		}
	}
}
