package main

import (
	"encoding/json"
	"os"
	"time"

	"github.com/mprcela/svckit/nsq"
)

func main() {
	nsq.Set(nsq.Concurrency(1024), nsq.MaxInFlight(1024))
	p := nsq.Pub("rsp")
	handler := func(m *nsq.Message) error {
		r := times{}
		json.Unmarshal(m.Body, &r)
		r.ReqReceive = time.Now().UnixNano()
		buf, _ := json.Marshal(&r)
		p.Publish(buf)
		return nil
	}
	nsq.Sub("req", handler)
	<-make(chan os.Signal, 1)
}

type times struct {
	ReqSend    int64
	ReqReceive int64
	RspReceive int64
}
