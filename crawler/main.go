package main

import (
	"gowork/crawler/engine"
	"gowork/crawler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:"http://www.zhenai.com/zhenghun",
		ParserFunc:parser.ParseCityList,
	})
}