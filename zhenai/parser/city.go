package parser

import (
	"regexp"
	"reptileProject/engine"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func ParseCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		name := string(m[2])
		result.Items = append(result.Items, "User "+name)
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParseFunc: func(bytes []byte) engine.ParseResult {
				// 利用闭包概念，将外部作用域的name变量，引入到该函数中
				return ParseProfile(bytes, name)
			},
		})
	}
	return result
}
