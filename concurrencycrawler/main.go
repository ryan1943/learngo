package main

import (
	"learngo/concurrencycrawler/engine"
	"learngo/concurrencycrawler/zhenai/parser"
)

func main() {
	engine.SimpleEngine{}.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
