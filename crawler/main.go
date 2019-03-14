package main

import (
	"engine"
	"persist"
	"scheduler"
	"zhenai/parser"
)

//var url = "http://www.zhenai.com/zhenghun"
var url = "http://www.zhenai.com/zhenghun/shanghai"

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 50,
		ItemChan:    persist.ItemSaver(),
	}
	e.Run(engine.Request{
		Url:        url,
		ParserFunc: parser.ParseCityList,
	})

}
