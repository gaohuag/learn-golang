package main

import (
	"github.com/gaohuag/learn-golang/crawler/engine"
	"github.com/gaohuag/learn-golang/crawler/scheduler"
	"github.com/gaohuag/learn-golang/crawler/types"
	"github.com/gaohuag/learn-golang/crawler/zhenai/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
	}
	e.Run(types.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})
}
