package engine

import (
	"gowork/crawler/fetcher"
	"log"
)

func Run(seeds ...Request) {
	var requests []Request

	for _,v := range seeds{
		requests = append(requests,v)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		log.Printf("Fecthing %s ",r.Url)
		body,err := fetcher.Fetch(r.Url)
		if err != nil{
			log.Printf("Fetcher error : %s %v ",r.Url,err)
			continue
		}

		parseResult := r.ParserFunc(body)
		requests = append(requests,parseResult.Requests...)

		for _,item := range parseResult.Items{
			log.Printf("Got item : %v ",item)
		}
	}
}