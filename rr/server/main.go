package main

import (
	"encoding/json"
	"os"
	"time"

	"github.com/minus5/svckit/nsq"
)

func main() {
	sigs := make(chan os.Signal, 1)
	nsq.Set(nsq.Concurrency(1024), nsq.MaxInFlight(1024), nsq.OutputBufferTimeout(250*time.Millisecond))
	server := nsq.RrSub("marin", handleNSQ)
	<-sigs
	server.Close()
}

type rsp struct {
	Ts int64
}

func handleNSQ(typ string, body []byte) (interface{}, error) {
	r := rsp{}
	json.Unmarshal(body, &r)
	now := time.Now()
	ts := now.UnixNano()
	// ns := ts - r.Ts
	// fmt.Println("since req ns:", ns)
	// fmt.Println("rsp ts:", ts)
	return rsp{ts}, nil
}
