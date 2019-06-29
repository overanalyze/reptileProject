package main

import (
	"reptileProject/engine"
	"reptileProject/scheduler"
	"reptileProject/zhenai/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 100,
	}
	e.Run(engine.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})
	// engine.SimpleEngine{}.Run(engine.Request{
	// 	Url:       "http://www.zhenai.com/zhenghun",
	// 	ParseFunc: parser.ParseCityList,
	// })
}
