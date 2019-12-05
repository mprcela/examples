package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/mprcela/svckit/nsq"
)

func main() {
	nsq.Set(nsq.Concurrency(1024), nsq.MaxInFlight(1024))

	handler := func(m *nsq.Message) error {
		r := times{}
		json.Unmarshal(m.Body, &r)
		r.RspReceive = time.Now().UnixNano()
		fmt.Println((r.ReqReceive-r.ReqSend)/1000000, "\t", (r.RspReceive-r.ReqReceive)/1000000, "\t", (r.RspReceive-r.ReqSend)/1000000)
		return nil
	}
	nsq.Sub("rsp", handler)
	p := nsq.Pub("req")

	for range time.NewTicker(time.Second / 5).C {
		r := times{
			ReqSend: time.Now().UnixNano(),
		}
		buf, _ := json.Marshal(&r)
		p.Publish(buf)
	}
}

type times struct {
	ReqSend    int64
	ReqReceive int64
	RspReceive int64
}
