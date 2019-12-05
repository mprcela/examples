package main

import (
	"fmt"
	"time"

	"github.com/minus5/svckit/nsq"
)

func main() {
	p := nsq.RrPub("marin2")
	nsq.Set(nsq.Concurrency(1024), nsq.MaxInFlight(1024), nsq.OutputBufferTimeout(250*time.Millisecond))

	for { //range time.NewTicker(time.Second / 6).C {
		// go func() {
		start := time.Now()
		var r rsp
		p.ReqRsp("marin", "prijava", rsp{start.UnixNano()}, &r, nil, 0, nil)
		now := time.Now()
		total := int(now.Sub(start).Nanoseconds())
		fmt.Println("TOTAL:", total/1000000)
		// }()
	}
}

type rsp struct {
	Ts int64
}
