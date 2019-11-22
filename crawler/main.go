package main

func main() {
	engine.Run(engine.Request{
		Url:"http://www.zhenai.com/zhenghun",
		ParserFunc:parser.ParseCityList,
	})

	//content,err := ioutil.ReadFile("README.md")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(string(content[:]))
}
