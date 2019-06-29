package main

import (
	"fmt"
	"regexp"
)

const text = `My email is zhuguo_wow@163.com
And email is 1204792945@qq.com
And email is zhuguo_wow@126.com
And email is royce_zhu@gmail.com
`

const testAge = `<div class="m-btn purple" data-v-bff6f798>32岁</div>`

var ageRe = regexp.MustCompile(`<div class="m-btn purple" [^>]*>([\d]+)岁</div>`)

func main() {
	compile := regexp.MustCompile(`([a-zA-Z0-9_]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)
	match := compile.FindAllStringSubmatch(text, -1)
	for _, value := range match {
		fmt.Println(value)
	}

	subMatch := ageRe.FindAllStringSubmatch(testAge, -1)
	for _, value := range subMatch {
		fmt.Println(value)
	}
}
