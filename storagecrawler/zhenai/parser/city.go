package parser

import (
	"learngo/storagecrawler/engine"
	"regexp"
)

var (
	profileRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	cityUrlRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
)

func ParseCity(contents []byte, _ string) engine.ParseResult {
	matches := profileRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	//m[1]是url，m[2]人名
	for _, m := range matches {

		//result.Items = append(result.Items, "User "+name)
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ProfileParser(string(m[2])),
		})
	}

	// matches = cityUrlRe.FindAllSubmatch(contents, -1)
	// for _, m := range matches {
	// 	result.Requests = append(result.Requests, engine.Request{
	// 		Url:        string(m[1]),
	// 		ParserFunc: ParseCity,
	// 	})
	// }

	return result
}
