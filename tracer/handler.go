package tracer

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/TylerBrock/colorjson"
	"github.com/fatih/color"
	"github.com/nsqio/go-nsq"
	"github.com/slaveofcode/nsqtracer/pansi"
)

type Handler struct {
	TraceName    string
	IsAutoFinish bool
}

func (h *Handler) HandleMessage(msg *nsq.Message) error {
	timeStamp := time.Now().Local().Format(time.RFC822)
	var tData map[string]interface{}

	f := colorjson.NewFormatter()
	f.KeyColor = color.New(color.FgHiYellow)
	f.Indent = 2

	var beautified []byte
	if err := json.Unmarshal(msg.Body, &tData); err != nil {
		// not a JSON data
		pansi.N(
			pansi.SGreen("➜"),
			pansi.SBlue("["+h.TraceName+"]"),
			" ",
			pansi.SMagenta("["+timeStamp+"]"),
			" ",
			pansi.SYellow("[Published Msg]"),
			" ",
			string(msg.Body),
		)

		goto finish
	}

	beautified, _ = f.Marshal(tData)
	// beautified, _ = json.MarshalIndent(colorized, "", " ")

	pansi.N(
		pansi.SGreen("➜"),
		pansi.SBlue("["+h.TraceName+"]"),
		" ",
		pansi.SMagenta("["+timeStamp+"]"),
		" ",
		pansi.SYellow("[Published Msg]"),
		" ",
	)

	fmt.Println(string(beautified))
	fmt.Println()

finish:

	if h.IsAutoFinish {
		msg.Finish()
	}

	return nil
}
