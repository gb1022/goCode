package parser

import (
	"engine"
	//	"fmt"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	limit := 10
	for _, m := range matches {
		if limit < 0 {
			break
		}
		result.Items = append(result.Items, "City "+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCity,
		})
		limit--
		//		fmt.Printf("city:%s,  url:%s\n ", m[2], m[1])
	}
	//	fmt.Printf("Matches found:%d\n", len(matches))
	return result
}
