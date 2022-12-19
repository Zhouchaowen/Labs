package main

import (
	mlog "mosn.io/pkg/log"
	"net/http"
	"time"

	"mosn.io/holmes"
)

// run `curl http://localhost:10003/make1gb` after 15s(warn up)
func init() {
	http.HandleFunc("/make1gb", make1gbslice)
	go http.ListenAndServe(":10003", nil) //nolint:errcheck
}

func main() {
	h, _ := holmes.New(
		holmes.WithCollectInterval("2s"),
		holmes.WithDumpPath("./tmp"),
		holmes.WithLogger(holmes.NewFileLog("./tmp/holmes.log", mlog.DEBUG)),
		holmes.WithTextDump(),
		holmes.WithMemDump(3, 25, 80, time.Minute),
	)
	h.EnableMemDump().Start()
	time.Sleep(time.Hour)
}

func make1gbslice(wr http.ResponseWriter, req *http.Request) {
	var a = make([]byte, 1073741824)
	_ = a
}
