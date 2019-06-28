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

func main() {
	compile := regexp.MustCompile(`([a-zA-Z0-9_]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)
	match := compile.FindAllStringSubmatch(text, -1)
	for _, value := range match {
		fmt.Println(value)
	}
}
