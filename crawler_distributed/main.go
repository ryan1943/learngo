package main

import (
	"fmt"
	"learngo/crawler_distributed/config"
	"learngo/crawler_distributed/persist/client"
	"learngo/distributedcrawler/engine"
	"learngo/distributedcrawler/scheduler"
	"learngo/distributedcrawler/zhenai/parser"
)

func main() {
	itemChan, err := client.ItemSaver(fmt.Sprintf(":%d", config.ItemSaverPort))
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 50,
		ItemChan:    itemChan,
	}
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
	// e.Run(engine.Request{
	// 	Url:        "http://www.zhenai.com/zhenghun/shanghai",
	// 	ParserFunc: parser.ParseCity,
	// })
}
