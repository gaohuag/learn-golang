package engine

import (
	"github.com/gaohuag/learn-golang/crawler/scheduler"
	"github.com/gaohuag/learn-golang/crawler/types"
	"github.com/gaohuag/learn-golang/crawler/zhenai/parser"
	"testing"
)

func TestConcurrentEngine_simpleScheduler(t *testing.T) {
	e := ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 100,
	}
	e.Run(types.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})
}

func TestConcurrentEngine_queuedScheduler(t *testing.T) {
	e := ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
	}
	e.Run(types.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})
}
