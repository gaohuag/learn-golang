package engine

import (
	"github.com/gaohuag/learn-golang/crawler/types"
	"github.com/gaohuag/learn-golang/crawler/zhenai/parser"
	"testing"
)

func TestSimpleEngine_Run(t *testing.T) {
	SimpleEngine{}.Run(types.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})
}
