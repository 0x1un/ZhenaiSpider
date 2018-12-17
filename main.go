package main

import (
	"learning/crawler/engine"

	"learning/crawler/zhenai/parser"
)
const url = "http://www.zhenai.com/zhenghun/"
func main() {
	engine.Run(engine.Request{url, parser.ParseCityList})
}